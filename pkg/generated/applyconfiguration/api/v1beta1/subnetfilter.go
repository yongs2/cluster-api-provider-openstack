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
	apiv1beta1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1beta1"
)

// SubnetFilterApplyConfiguration represents a declarative configuration of the SubnetFilter type for use
// with apply.
type SubnetFilterApplyConfiguration struct {
	Name                                  *string `json:"name,omitempty"`
	Description                           *string `json:"description,omitempty"`
	ProjectID                             *string `json:"projectID,omitempty"`
	IPVersion                             *int    `json:"ipVersion,omitempty"`
	GatewayIP                             *string `json:"gatewayIP,omitempty"`
	CIDR                                  *string `json:"cidr,omitempty"`
	IPv6AddressMode                       *string `json:"ipv6AddressMode,omitempty"`
	IPv6RAMode                            *string `json:"ipv6RAMode,omitempty"`
	FilterByNeutronTagsApplyConfiguration `json:",inline"`
}

// SubnetFilterApplyConfiguration constructs a declarative configuration of the SubnetFilter type for use with
// apply.
func SubnetFilter() *SubnetFilterApplyConfiguration {
	return &SubnetFilterApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *SubnetFilterApplyConfiguration) WithName(value string) *SubnetFilterApplyConfiguration {
	b.Name = &value
	return b
}

// WithDescription sets the Description field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Description field is set to the value of the last call.
func (b *SubnetFilterApplyConfiguration) WithDescription(value string) *SubnetFilterApplyConfiguration {
	b.Description = &value
	return b
}

// WithProjectID sets the ProjectID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProjectID field is set to the value of the last call.
func (b *SubnetFilterApplyConfiguration) WithProjectID(value string) *SubnetFilterApplyConfiguration {
	b.ProjectID = &value
	return b
}

// WithIPVersion sets the IPVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPVersion field is set to the value of the last call.
func (b *SubnetFilterApplyConfiguration) WithIPVersion(value int) *SubnetFilterApplyConfiguration {
	b.IPVersion = &value
	return b
}

// WithGatewayIP sets the GatewayIP field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GatewayIP field is set to the value of the last call.
func (b *SubnetFilterApplyConfiguration) WithGatewayIP(value string) *SubnetFilterApplyConfiguration {
	b.GatewayIP = &value
	return b
}

// WithCIDR sets the CIDR field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CIDR field is set to the value of the last call.
func (b *SubnetFilterApplyConfiguration) WithCIDR(value string) *SubnetFilterApplyConfiguration {
	b.CIDR = &value
	return b
}

// WithIPv6AddressMode sets the IPv6AddressMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPv6AddressMode field is set to the value of the last call.
func (b *SubnetFilterApplyConfiguration) WithIPv6AddressMode(value string) *SubnetFilterApplyConfiguration {
	b.IPv6AddressMode = &value
	return b
}

// WithIPv6RAMode sets the IPv6RAMode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the IPv6RAMode field is set to the value of the last call.
func (b *SubnetFilterApplyConfiguration) WithIPv6RAMode(value string) *SubnetFilterApplyConfiguration {
	b.IPv6RAMode = &value
	return b
}

// WithTags adds the given value to the Tags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Tags field.
func (b *SubnetFilterApplyConfiguration) WithTags(values ...apiv1beta1.NeutronTag) *SubnetFilterApplyConfiguration {
	for i := range values {
		b.FilterByNeutronTagsApplyConfiguration.Tags = append(b.FilterByNeutronTagsApplyConfiguration.Tags, values[i])
	}
	return b
}

// WithTagsAny adds the given value to the TagsAny field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TagsAny field.
func (b *SubnetFilterApplyConfiguration) WithTagsAny(values ...apiv1beta1.NeutronTag) *SubnetFilterApplyConfiguration {
	for i := range values {
		b.FilterByNeutronTagsApplyConfiguration.TagsAny = append(b.FilterByNeutronTagsApplyConfiguration.TagsAny, values[i])
	}
	return b
}

// WithNotTags adds the given value to the NotTags field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NotTags field.
func (b *SubnetFilterApplyConfiguration) WithNotTags(values ...apiv1beta1.NeutronTag) *SubnetFilterApplyConfiguration {
	for i := range values {
		b.FilterByNeutronTagsApplyConfiguration.NotTags = append(b.FilterByNeutronTagsApplyConfiguration.NotTags, values[i])
	}
	return b
}

// WithNotTagsAny adds the given value to the NotTagsAny field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NotTagsAny field.
func (b *SubnetFilterApplyConfiguration) WithNotTagsAny(values ...apiv1beta1.NeutronTag) *SubnetFilterApplyConfiguration {
	for i := range values {
		b.FilterByNeutronTagsApplyConfiguration.NotTagsAny = append(b.FilterByNeutronTagsApplyConfiguration.NotTagsAny, values[i])
	}
	return b
}
