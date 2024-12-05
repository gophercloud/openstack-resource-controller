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

// PortImportApplyConfiguration represents a declarative configuration of the PortImport type for use
// with apply.
type PortImportApplyConfiguration struct {
	ID     *string                       `json:"id,omitempty"`
	Filter *PortFilterApplyConfiguration `json:"filter,omitempty"`
}

// PortImportApplyConfiguration constructs a declarative configuration of the PortImport type for use with
// apply.
func PortImport() *PortImportApplyConfiguration {
	return &PortImportApplyConfiguration{}
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *PortImportApplyConfiguration) WithID(value string) *PortImportApplyConfiguration {
	b.ID = &value
	return b
}

// WithFilter sets the Filter field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Filter field is set to the value of the last call.
func (b *PortImportApplyConfiguration) WithFilter(value *PortFilterApplyConfiguration) *PortImportApplyConfiguration {
	b.Filter = value
	return b
}
