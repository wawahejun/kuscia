// Copyright 2023 Ant Group Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package handler

import (
	"context"
	"fmt"
	"reflect"

	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/secretflow/kuscia/pkg/common"
	kusciaapisv1alpha1 "github.com/secretflow/kuscia/pkg/crd/apis/kuscia/v1alpha1"
	"github.com/secretflow/kuscia/pkg/utils/nlog"
	utilsres "github.com/secretflow/kuscia/pkg/utils/resources"
)

// RunningHandler will handle kuscia job in Running phase.
type RunningHandler struct {
	*JobScheduler
}

// NewRunningHandler return RunningHandler to handle Running kuscia job.
func NewRunningHandler(deps *Dependencies) *RunningHandler {
	return &RunningHandler{
		JobScheduler: NewJobScheduler(deps),
	}
}

// HandlePhase implements the KusciaJobPhaseHandler interface.
// It will do scheduling when the job phase is Running.
func (h *RunningHandler) HandlePhase(job *kusciaapisv1alpha1.KusciaJob) (needUpdate bool, retErr error) {
	return h.handleRunning(job)
}

// handleRunning
// pending --> running
// Pending --> Failed
func (h *RunningHandler) handleRunning(job *kusciaapisv1alpha1.KusciaJob) (needUpdateStatus bool, err error) {
	now := metav1.Now().Rfc3339Copy()
	defer updateJobTime(now, job)
	// handle stage command, check if the stage command matches the phase of job
	if hasReconciled, handleErr := h.handleStageCommand(now, job); handleErr != nil || hasReconciled {
		return hasReconciled, handleErr
	}
	// set task id
	if utilsres.SelfClusterAsInitiator(h.namespaceLister, job.Spec.Initiator, job.Annotations) {
		if hasSet, setErr := h.setJobTaskID(job); hasSet {
			return false, setErr
		}
	}
	// selector all tasks of job
	selector, err := jobTaskSelector(string(job.UID))
	if err != nil {
		nlog.Errorf("Create job sub-tasks selector failed: %s", err)
		return false, err
	}
	// list all sub-tasks
	subTasks, err := h.kusciaTaskLister.List(selector)
	if err != nil {
		nlog.Errorf("List job sub-tasks selector: %s", err)
		return false, err
	}
	// compute current status.
	// NOTE: We don't believe kusciaJob.TaskStatus, we rebuild it from current sub-task status.
	// MayBe some tasks have been created, but updateStatus failed Or first task creation has been happened,
	// but updateStatus is delayed.
	currentSubTasksStatusWithAlias, currentSubTasksStatusWithID := buildJobSubTaskStatus(subTasks, job)
	currentJobPhase := jobStatusPhaseFrom(job, currentSubTasksStatusWithAlias)
	if updateJobSubTaskStatus(&job.Status, currentSubTasksStatusWithID) {
		needUpdateStatus = true
	}

	// compute ready task and push job when needed.
	readyTask := readyTasksOf(job, currentSubTasksStatusWithAlias)
	if currentJobPhase != kusciaapisv1alpha1.KusciaJobFailed && currentJobPhase != kusciaapisv1alpha1.KusciaJobSucceeded {
		willStartTask := willStartTasksOf(job, readyTask, currentSubTasksStatusWithAlias)
		willStartKusciaTasks, err := h.buildWillStartKusciaTask(job, willStartTask)
		if err != nil {
			return needUpdateStatus, err
		}
		// then we will start KusciaTask
		for _, t := range willStartKusciaTasks {
			nlog.Infof("Create kuscia tasks: %s", t.ObjectMeta.Name)
			_, err = h.kusciaClient.KusciaV1alpha1().KusciaTasks(common.KusciaCrossDomain).Create(context.Background(), t, metav1.CreateOptions{})
			if err != nil {
				if k8serrors.IsAlreadyExists(err) {
					existTask, getErr := h.kusciaTaskLister.KusciaTasks(common.KusciaCrossDomain).Get(t.Name)
					if getErr != nil {
						if k8serrors.IsNotFound(getErr) {
							existTask, getErr = h.kusciaClient.KusciaV1alpha1().KusciaTasks(common.KusciaCrossDomain).Get(context.Background(), t.Name, metav1.GetOptions{})
						}
						if getErr != nil {
							nlog.Errorf("Get exist task %v failed: %v", t.Name, getErr)
							setKusciaJobStatus(now, &job.Status, kusciaapisv1alpha1.KusciaJobFailed, string(kusciaapisv1alpha1.CreateTaskFailed), getErr.Error())
							return true, nil
						}
					}

					if existTask.Labels == nil || existTask.Labels[common.LabelJobUID] != string(job.UID) {
						message := fmt.Sprintf("Failed to create task %v because of a task with the same name already exists", t.Name)
						nlog.Error(message)
						setKusciaJobStatus(now, &job.Status, kusciaapisv1alpha1.KusciaJobFailed, string(kusciaapisv1alpha1.CreateTaskFailed), message)
						return true, nil
					}
				} else {
					nlog.Errorf("Create kuscia task %s failed, %v", t.Name, err)
					return true, err
				}
			}
		}
	}

	if buildJobStatus(now, &job.Status, currentJobPhase) {
		needUpdateStatus = true
	}
	return needUpdateStatus, nil
}

func updateJobSubTaskStatus(kjStatus *kusciaapisv1alpha1.KusciaJobStatus, currentSubTasksStatus map[string]kusciaapisv1alpha1.KusciaTaskPhase) bool {
	if !reflect.DeepEqual(kjStatus.TaskStatus, currentSubTasksStatus) {
		kjStatus.TaskStatus = currentSubTasksStatus
		return true
	}
	return false
}
