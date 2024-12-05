/*
Copyright 2024 The ORC Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ServerStatusApplyConfiguration represents a declarative configuration of the ServerStatus type for use
// with apply.
type ServerStatusApplyConfiguration struct {
	Conditions []v1.ConditionApplyConfiguration        `json:"conditions,omitempty"`
	ID         *string                                 `json:"id,omitempty"`
	Resource   *ServerResourceStatusApplyConfiguration `json:"resource,omitempty"`
}

// ServerStatusApplyConfiguration constructs a declarative configuration of the ServerStatus type for use with
// apply.
func ServerStatus() *ServerStatusApplyConfiguration {
	return &ServerStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ServerStatusApplyConfiguration) WithConditions(values ...*v1.ConditionApplyConfiguration) *ServerStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *ServerStatusApplyConfiguration) WithID(value string) *ServerStatusApplyConfiguration {
	b.ID = &value
	return b
}

// WithResource sets the Resource field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resource field is set to the value of the last call.
func (b *ServerStatusApplyConfiguration) WithResource(value *ServerResourceStatusApplyConfiguration) *ServerStatusApplyConfiguration {
	b.Resource = value
	return b
}
