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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	gentype "k8s.io/client-go/gentype"
	v1beta1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1beta1"
	apiv1beta1 "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/applyconfiguration/api/v1beta1"
	typedapiv1beta1 "sigs.k8s.io/cluster-api-provider-openstack/pkg/generated/clientset/clientset/typed/api/v1beta1"
)

// fakeOpenStackMachineTemplates implements OpenStackMachineTemplateInterface
type fakeOpenStackMachineTemplates struct {
	*gentype.FakeClientWithListAndApply[*v1beta1.OpenStackMachineTemplate, *v1beta1.OpenStackMachineTemplateList, *apiv1beta1.OpenStackMachineTemplateApplyConfiguration]
	Fake *FakeInfrastructureV1beta1
}

func newFakeOpenStackMachineTemplates(fake *FakeInfrastructureV1beta1, namespace string) typedapiv1beta1.OpenStackMachineTemplateInterface {
	return &fakeOpenStackMachineTemplates{
		gentype.NewFakeClientWithListAndApply[*v1beta1.OpenStackMachineTemplate, *v1beta1.OpenStackMachineTemplateList, *apiv1beta1.OpenStackMachineTemplateApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta1.SchemeGroupVersion.WithResource("openstackmachinetemplates"),
			v1beta1.SchemeGroupVersion.WithKind("OpenStackMachineTemplate"),
			func() *v1beta1.OpenStackMachineTemplate { return &v1beta1.OpenStackMachineTemplate{} },
			func() *v1beta1.OpenStackMachineTemplateList { return &v1beta1.OpenStackMachineTemplateList{} },
			func(dst, src *v1beta1.OpenStackMachineTemplateList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta1.OpenStackMachineTemplateList) []*v1beta1.OpenStackMachineTemplate {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1beta1.OpenStackMachineTemplateList, items []*v1beta1.OpenStackMachineTemplate) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
