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
	v1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"
)

// SecurityGroupRuleApplyConfiguration represents a declarative configuration of the SecurityGroupRule type for use
// with apply.
type SecurityGroupRuleApplyConfiguration struct {
	Description    *v1alpha1.NeutronDescription     `json:"description,omitempty"`
	Direction      *v1alpha1.RuleDirection          `json:"direction,omitempty"`
	RemoteIPPrefix *v1alpha1.CIDR                   `json:"remoteIPPrefix,omitempty"`
	Protocol       *v1alpha1.Protocol               `json:"protocol,omitempty"`
	Ethertype      *v1alpha1.Ethertype              `json:"ethertype,omitempty"`
	PortRange      *PortRangeSpecApplyConfiguration `json:"portRange,omitempty"`
}

// SecurityGroupRuleApplyConfiguration constructs a declarative configuration of the SecurityGroupRule type for use with
// apply.
func SecurityGroupRule() *SecurityGroupRuleApplyConfiguration {
	return &SecurityGroupRuleApplyConfiguration{}
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *SecurityGroupRuleApplyConfiguration) WithDescription(value v1alpha1.NeutronDescription) *SecurityGroupRuleApplyConfiguration {
	b.Description = &value
	return b
}

// WithDirection sets the Direction field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Direction field is set to the value of the last call.
func (b *SecurityGroupRuleApplyConfiguration) WithDirection(value v1alpha1.RuleDirection) *SecurityGroupRuleApplyConfiguration {
	b.Direction = &value
	return b
}

// WithRemoteIPPrefix sets the RemoteIPPrefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RemoteIPPrefix field is set to the value of the last call.
func (b *SecurityGroupRuleApplyConfiguration) WithRemoteIPPrefix(value v1alpha1.CIDR) *SecurityGroupRuleApplyConfiguration {
	b.RemoteIPPrefix = &value
	return b
}

// WithProtocol sets the Protocol field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Protocol field is set to the value of the last call.
func (b *SecurityGroupRuleApplyConfiguration) WithProtocol(value v1alpha1.Protocol) *SecurityGroupRuleApplyConfiguration {
	b.Protocol = &value
	return b
}

// WithEthertype sets the Ethertype field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ethertype field is set to the value of the last call.
func (b *SecurityGroupRuleApplyConfiguration) WithEthertype(value v1alpha1.Ethertype) *SecurityGroupRuleApplyConfiguration {
	b.Ethertype = &value
	return b
}

// WithPortRange sets the PortRange field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PortRange field is set to the value of the last call.
func (b *SecurityGroupRuleApplyConfiguration) WithPortRange(value *PortRangeSpecApplyConfiguration) *SecurityGroupRuleApplyConfiguration {
	b.PortRange = value
	return b
}
