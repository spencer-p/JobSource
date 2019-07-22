/*
Copyright 2019 The Knative Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/apis"
)

var condSet = apis.NewLivingConditionSet(
	JobSourceConditionSinkProvided,
	JobSourceConditionStarted,
)

// GetGroupVersionKind implements kmeta.OwnerRefable
func (js *JobSource) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("JobSource")
}

func (jss *JobSourceStatus) InitializeConditions() {
	condSet.Manage(jss).InitializeConditions()
}

func (jss *JobSourceStatus) MarkServiceUnavailable(name string) {
	condSet.Manage(jss).MarkFalse(
		JobSourceConditionReady,
		"ServiceUnavailable",
		"Service %q wasn't found.", name)
}

func (jss *JobSourceStatus) MarkServiceAvailable() {
	condSet.Manage(jss).MarkTrue(JobSourceConditionReady)
}

// GetCondition returns the condition currently associated with the given type, or nil.
func (jss *JobSourceStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return condSet.Manage(jss).GetCondition(t)
}

// IsReady returns true if the resource is ready overall.
func (jss *JobSourceStatus) IsReady() bool {
	return condSet.Manage(jss).IsHappy()
}

// MarkSink sets the condition that the source has a sink configured.
func (jss *JobSourceStatus) MarkSink(uri string) {
	jss.SinkURI = uri
	if len(uri) > 0 {
		condSet.Manage(jss).MarkTrue(JobSourceConditionSinkProvided)
	} else {
		condSet.Manage(jss).MarkUnknown(JobSourceConditionSinkProvided, "SinkEmpty", "Sink has resolved to empty.%s", "")
	}
}

// MarkNoSink sets the condition that the source does not have a sink configured.
func (jss *JobSourceStatus) MarkNoSink(reason, messageFormat string, messageA ...interface{}) {
	condSet.Manage(jss).MarkFalse(JobSourceConditionSinkProvided, reason, messageFormat, messageA...)
}

// IsStarted returns true if the Started condition has status true, otherwise
// false.
func (jss *JobSourceStatus) IsStarted() bool {
	c := condSet.Manage(jss).GetCondition(JobSourceConditionStarted)
	if c != nil {
		return c.IsTrue()
	}
	return false
}

// MarkStarted sets the condition that the source has been deployed.
func (jss *JobSourceStatus) MarkStarted() {
	condSet.Manage(jss).MarkTrue(JobSourceConditionStarted)
}

// MarkStarting sets the condition that the source is deploying.
func (jss *JobSourceStatus) MarkStarting(reason, messageFormat string, messageA ...interface{}) {
	condSet.Manage(jss).MarkUnknown(JobSourceConditionStarted, reason, messageFormat, messageA...)
}

// MarkNotStarted sets the condition that the source has not been deployed.
func (jss *JobSourceStatus) MarkNotStarted(reason, messageFormat string, messageA ...interface{}) {
	condSet.Manage(jss).MarkFalse(JobSourceConditionStarted, reason, messageFormat, messageA...)
}
