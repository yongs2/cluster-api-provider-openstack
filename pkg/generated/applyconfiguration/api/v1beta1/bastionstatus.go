/*
Copyright 2024 The Kubernetes Authors.

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

package v1beta1

import (
	v1beta1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1beta1"
)

// BastionStatusApplyConfiguration represents an declarative configuration of the BastionStatus type for use
// with apply.
type BastionStatusApplyConfiguration struct {
	ID         *string                                `json:"id,omitempty"`
	Name       *string                                `json:"name,omitempty"`
	SSHKeyName *string                                `json:"sshKeyName,omitempty"`
	State      *v1beta1.InstanceState                 `json:"state,omitempty"`
	IP         *string                                `json:"ip,omitempty"`
	FloatingIP *string                                `json:"floatingIP,omitempty"`
	Resolved   *ResolvedMachineSpecApplyConfiguration `json:"resolved,omitempty"`
	Resources  *MachineResourcesApplyConfiguration    `json:"resources,omitempty"`
}

// BastionStatusApplyConfiguration constructs an declarative configuration of the BastionStatus type for use with
// apply.
func BastionStatus() *BastionStatusApplyConfiguration {
	return &BastionStatusApplyConfiguration{}
}

// WithID sets the ID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ID field is set to the value of the last call.
func (b *BastionStatusApplyConfiguration) WithID(value string) *BastionStatusApplyConfiguration {
	b.ID = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *BastionStatusApplyConfiguration) WithName(value string) *BastionStatusApplyConfiguration {
	b.Name = &value
	return b
}

// WithSSHKeyName sets the SSHKeyName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the SSHKeyName field is set to the value of the last call.
func (b *BastionStatusApplyConfiguration) WithSSHKeyName(value string) *BastionStatusApplyConfiguration {
	b.SSHKeyName = &value
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *BastionStatusApplyConfiguration) WithState(value v1beta1.InstanceState) *BastionStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithIP sets the IP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IP field is set to the value of the last call.
func (b *BastionStatusApplyConfiguration) WithIP(value string) *BastionStatusApplyConfiguration {
	b.IP = &value
	return b
}

// WithFloatingIP sets the FloatingIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FloatingIP field is set to the value of the last call.
func (b *BastionStatusApplyConfiguration) WithFloatingIP(value string) *BastionStatusApplyConfiguration {
	b.FloatingIP = &value
	return b
}

// WithResolved sets the Resolved field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resolved field is set to the value of the last call.
func (b *BastionStatusApplyConfiguration) WithResolved(value *ResolvedMachineSpecApplyConfiguration) *BastionStatusApplyConfiguration {
	b.Resolved = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *BastionStatusApplyConfiguration) WithResources(value *MachineResourcesApplyConfiguration) *BastionStatusApplyConfiguration {
	b.Resources = value
	return b
}