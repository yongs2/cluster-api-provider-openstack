/*
Copyright 2023 The Kubernetes Authors.

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

package webhooks

import (
	"context"
	"testing"

	. "github.com/onsi/gomega" //nolint:revive
	"k8s.io/utils/ptr"

	infrav1 "sigs.k8s.io/cluster-api-provider-openstack/api/v1beta1"
)

func TestOpenStackCluster_ValidateUpdate(t *testing.T) {
	tests := []struct {
		name        string
		oldTemplate *infrav1.OpenStackCluster
		newTemplate *infrav1.OpenStackCluster
		wantErr     bool
	}{
		{
			name: "Changing OpenStackCluster.Spec.IdentityRef.Name is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobarbaz",
						CloudName: "foobar",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Changing OpenStackCluster.Spec.IdentityRef.CloudName is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobarbaz",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Changing OpenStackCluster.Spec.Bastion is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					Bastion: &infrav1.Bastion{
						Spec: &infrav1.OpenStackMachineSpec{
							Image: infrav1.ImageParam{
								Filter: &infrav1.ImageFilter{
									Name: ptr.To("foobar"),
								},
							},
							Flavor: ptr.To("minimal"),
						},
						Enabled: ptr.To(true),
					},
				},
				Status: infrav1.OpenStackClusterStatus{
					Bastion: &infrav1.BastionStatus{
						Name: "foobar",
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					Bastion: &infrav1.Bastion{
						Spec: &infrav1.OpenStackMachineSpec{
							Image: infrav1.ImageParam{
								Filter: &infrav1.ImageFilter{
									Name: ptr.To("foobarbaz"),
								},
							},
							Flavor: ptr.To("medium"),
						},
						Enabled: ptr.To(true),
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Changing security group rules on the OpenStackCluster.Spec.ManagedSecurityGroups.AllNodesSecurityGroupRules is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ManagedSecurityGroups: &infrav1.ManagedSecurityGroups{
						AllNodesSecurityGroupRules: []infrav1.SecurityGroupRuleSpec{
							{
								Name:                "foobar",
								Description:         ptr.To("foobar"),
								PortRangeMin:        ptr.To(80),
								PortRangeMax:        ptr.To(80),
								Protocol:            ptr.To("tcp"),
								RemoteManagedGroups: []infrav1.ManagedSecurityGroupName{"controlplane"},
							},
						},
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ManagedSecurityGroups: &infrav1.ManagedSecurityGroups{
						AllNodesSecurityGroupRules: []infrav1.SecurityGroupRuleSpec{
							{
								Name:                "foobar",
								Description:         ptr.To("foobar"),
								PortRangeMin:        ptr.To(80),
								PortRangeMax:        ptr.To(80),
								Protocol:            ptr.To("tcp"),
								RemoteManagedGroups: []infrav1.ManagedSecurityGroupName{"controlplane", "worker"},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Changing security group rules on the OpenStackCluster.Spec.ManagedSecurityGroups.ControlPlaneNodesSecurityGroupRules is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ManagedSecurityGroups: &infrav1.ManagedSecurityGroups{
						ControlPlaneNodesSecurityGroupRules: []infrav1.SecurityGroupRuleSpec{
							{
								Name:                "foobar",
								Description:         ptr.To("foobar"),
								PortRangeMin:        ptr.To(80),
								PortRangeMax:        ptr.To(80),
								Protocol:            ptr.To("tcp"),
								RemoteManagedGroups: []infrav1.ManagedSecurityGroupName{"controlplane"},
							},
						},
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ManagedSecurityGroups: &infrav1.ManagedSecurityGroups{
						ControlPlaneNodesSecurityGroupRules: []infrav1.SecurityGroupRuleSpec{
							{
								Name:                "foobar",
								Description:         ptr.To("foobar"),
								PortRangeMin:        ptr.To(80),
								PortRangeMax:        ptr.To(80),
								Protocol:            ptr.To("tcp"),
								RemoteManagedGroups: []infrav1.ManagedSecurityGroupName{"controlplane", "worker"},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Changing security group rules on the OpenStackCluster.Spec.ManagedSecurityGroups.WorkerNodesSecurityGroupRules is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ManagedSecurityGroups: &infrav1.ManagedSecurityGroups{
						WorkerNodesSecurityGroupRules: []infrav1.SecurityGroupRuleSpec{
							{
								Name:                "foobar",
								Description:         ptr.To("foobar"),
								PortRangeMin:        ptr.To(80),
								PortRangeMax:        ptr.To(80),
								Protocol:            ptr.To("tcp"),
								RemoteManagedGroups: []infrav1.ManagedSecurityGroupName{"worker"},
							},
						},
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ManagedSecurityGroups: &infrav1.ManagedSecurityGroups{
						WorkerNodesSecurityGroupRules: []infrav1.SecurityGroupRuleSpec{
							{
								Name:                "foobar",
								Description:         ptr.To("foobar"),
								PortRangeMin:        ptr.To(80),
								PortRangeMax:        ptr.To(80),
								Protocol:            ptr.To("tcp"),
								RemoteManagedGroups: []infrav1.ManagedSecurityGroupName{"worker", "controlplane"},
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Changing CIDRs on the OpenStackCluster.Spec.APIServerLoadBalancer.AllowedCIDRs is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					APIServerLoadBalancer: &infrav1.APIServerLoadBalancer{
						Enabled: ptr.To(true),
						AllowedCIDRs: []string{
							"0.0.0.0/0",
							"192.168.10.0/24",
						},
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					APIServerLoadBalancer: &infrav1.APIServerLoadBalancer{
						Enabled: ptr.To(true),
						AllowedCIDRs: []string{
							"0.0.0.0/0",
							"192.168.10.0/24",
							"10.6.0.0/16",
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Adding OpenStackCluster.Spec.ControlPlaneAvailabilityZones is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ControlPlaneAvailabilityZones: []string{
						"alice",
						"bob",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Modifying OpenStackCluster.Spec.ControlPlaneAvailabilityZones is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ControlPlaneAvailabilityZones: []string{
						"alice",
						"bob",
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ControlPlaneAvailabilityZones: []string{
						"alice",
						"bob",
						"eve",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Removing OpenStackCluster.Spec.ControlPlaneAvailabilityZones is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ControlPlaneAvailabilityZones: []string{
						"alice",
						"bob",
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Modifying OpenstackCluster.Spec.ControlPlaneOmitAvailabilityZone is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ControlPlaneOmitAvailabilityZone: ptr.To(true),
				},
			},
			wantErr: false,
		},
		{
			name: "Changing OpenStackCluster.Spec.APIServerFixedIP is allowed when API Server Floating IP is disabled",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					DisableAPIServerFloatingIP: ptr.To(true),
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					DisableAPIServerFloatingIP: ptr.To(true),
					APIServerFixedIP:           ptr.To("20.1.56.1"),
				},
			},
			wantErr: false,
		},
		{
			name: "Changing OpenStackCluster.Spec.APIServerFixedIP is not allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					DisableAPIServerFloatingIP: ptr.To(false),
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					DisableAPIServerFloatingIP: ptr.To(false),
					APIServerFixedIP:           ptr.To("20.1.56.1"),
				},
			},
			wantErr: true,
		},

		{
			name: "Changing OpenStackCluster.Spec.APIServerPort is allowed when API Server Floating IP is disabled",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					DisableAPIServerFloatingIP: ptr.To(true),
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					DisableAPIServerFloatingIP: ptr.To(true),
					APIServerPort:              ptr.To(uint16(8443)),
				},
			},
			wantErr: false,
		},
		{
			name: "Changing OpenStackCluster.Spec.APIServerPort is not allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					DisableAPIServerFloatingIP: ptr.To(false),
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					DisableAPIServerFloatingIP: ptr.To(false),
					APIServerPort:              ptr.To(uint16(8443)),
				},
			},
			wantErr: true,
		},
		{
			name: "Changing OpenStackCluster.Spec.APIServerFloatingIP is allowed when it matches the current api server loadbalancer IP",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
				},
				Status: infrav1.OpenStackClusterStatus{
					APIServerLoadBalancer: &infrav1.LoadBalancer{
						IP: "1.2.3.4",
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					APIServerFloatingIP: ptr.To("1.2.3.4"),
				},
				Status: infrav1.OpenStackClusterStatus{
					APIServerLoadBalancer: &infrav1.LoadBalancer{
						IP: "1.2.3.4",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Changing OpenStackCluster.Spec.APIServerFloatingIP is not allowed when it doesn't matches the current api server loadbalancer IP",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
				},
				Status: infrav1.OpenStackClusterStatus{
					APIServerLoadBalancer: &infrav1.LoadBalancer{
						IP: "1.2.3.4",
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					APIServerFloatingIP: ptr.To("5.6.7.8"),
				},
				Status: infrav1.OpenStackClusterStatus{
					APIServerLoadBalancer: &infrav1.LoadBalancer{
						IP: "1.2.3.4",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Removing OpenStackCluster.Spec.Bastion when it is enabled is not allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					Bastion: &infrav1.Bastion{
						Enabled: ptr.To(true),
						Spec: &infrav1.OpenStackMachineSpec{
							Flavor: ptr.To("m1.small"),
							Image: infrav1.ImageParam{
								Filter: &infrav1.ImageFilter{
									Name: ptr.To("ubuntu"),
								},
							},
						},
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
				},
			},
			wantErr: true,
		},
		{
			name: "Removing OpenStackCluster.Spec.Bastion when it is disabled is allowed",
			oldTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					Bastion: &infrav1.Bastion{
						Enabled: ptr.To(false),
						Spec: &infrav1.OpenStackMachineSpec{
							Flavor: ptr.To("m1.small"),
							Image: infrav1.ImageParam{
								Filter: &infrav1.ImageFilter{
									Name: ptr.To("ubuntu"),
								},
							},
						},
					},
				},
			},
			newTemplate: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			ctx := context.TODO()

			webhook := &openStackClusterWebhook{}
			warn, err := webhook.ValidateUpdate(ctx, tt.oldTemplate, tt.newTemplate)
			if tt.wantErr {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).NotTo(HaveOccurred())
			}
			// Nothing emits warnings yet
			g.Expect(warn).To(BeEmpty())
		})
	}
}

func TestOpenStackCluster_ValidateCreate(t *testing.T) {
	tests := []struct {
		name     string
		template *infrav1.OpenStackCluster
		wantErr  bool
	}{
		{
			name: "OpenStackCluster.Spec.IdentityRef with correct spec on create",
			template: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "OpenStackCluster.Spec.ManagedSecurityGroups.AllNodesSecurityGroupRules with correct spec on create",
			template: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ManagedSecurityGroups: &infrav1.ManagedSecurityGroups{
						AllNodesSecurityGroupRules: []infrav1.SecurityGroupRuleSpec{
							{
								Name:         "foobar",
								Description:  ptr.To("foobar"),
								PortRangeMin: ptr.To(80),
								PortRangeMax: ptr.To(80),
								Protocol:     ptr.To("tcp"),
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "OpenStackCluster.Spec.ManagedSecurityGroups.AllNodesSecurityGroupRules with mutually exclusive fields on create",
			template: &infrav1.OpenStackCluster{
				Spec: infrav1.OpenStackClusterSpec{
					IdentityRef: infrav1.OpenStackIdentityReference{
						Name:      "foobar",
						CloudName: "foobar",
					},
					ManagedSecurityGroups: &infrav1.ManagedSecurityGroups{
						AllNodesSecurityGroupRules: []infrav1.SecurityGroupRuleSpec{
							{
								Name:                "foobar",
								Description:         ptr.To("foobar"),
								PortRangeMin:        ptr.To(80),
								PortRangeMax:        ptr.To(80),
								Protocol:            ptr.To("tcp"),
								RemoteManagedGroups: []infrav1.ManagedSecurityGroupName{"controlplane"},
								RemoteGroupID:       ptr.To("foobar"),
							},
						},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			ctx := context.TODO()

			webhook := &openStackClusterWebhook{}
			warn, err := webhook.ValidateCreate(ctx, tt.template)
			if tt.wantErr {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).NotTo(HaveOccurred())
			}
			// Nothing emits warnings yet
			g.Expect(warn).To(BeEmpty())
		})
	}
}
