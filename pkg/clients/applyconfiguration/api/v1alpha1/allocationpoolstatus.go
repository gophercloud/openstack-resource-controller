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

// AllocationPoolStatusApplyConfiguration represents a declarative configuration of the AllocationPoolStatus type for use
// with apply.
type AllocationPoolStatusApplyConfiguration struct {
	Start *string `json:"start,omitempty"`
	End   *string `json:"end,omitempty"`
}

// AllocationPoolStatusApplyConfiguration constructs a declarative configuration of the AllocationPoolStatus type for use with
// apply.
func AllocationPoolStatus() *AllocationPoolStatusApplyConfiguration {
	return &AllocationPoolStatusApplyConfiguration{}
}

// WithStart sets the Start field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Start field is set to the value of the last call.
func (b *AllocationPoolStatusApplyConfiguration) WithStart(value string) *AllocationPoolStatusApplyConfiguration {
	b.Start = &value
	return b
}

// WithEnd sets the End field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the End field is set to the value of the last call.
func (b *AllocationPoolStatusApplyConfiguration) WithEnd(value string) *AllocationPoolStatusApplyConfiguration {
	b.End = &value
	return b
}
