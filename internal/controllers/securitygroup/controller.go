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

package securitygroup

import (
	"context"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	orcv1alpha1 "github.com/k-orc/openstack-resource-controller/api/v1alpha1"

	ctrlexport "github.com/k-orc/openstack-resource-controller/internal/controllers/export"
	"github.com/k-orc/openstack-resource-controller/internal/controllers/generic"
	"github.com/k-orc/openstack-resource-controller/internal/scope"
)

// +kubebuilder:rbac:groups=openstack.k-orc.cloud,resources=securitygroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=openstack.k-orc.cloud,resources=securitygroups/status,verbs=get;update;patch

type securitygroupReconcilerConstructor struct {
	scopeFactory scope.Factory
}

func New(scopeFactory scope.Factory) ctrlexport.Controller {
	return securitygroupReconcilerConstructor{
		scopeFactory: scopeFactory,
	}
}

func (securitygroupReconcilerConstructor) GetName() string {
	return "securitygroup"
}

// SetupWithManager sets up the controller with the Manager.
func (c securitygroupReconcilerConstructor) SetupWithManager(_ context.Context, mgr ctrl.Manager, options controller.Options) error {
	reconciler := generic.NewController(c.GetName(), mgr.GetClient(), c.scopeFactory, securityGroupHelperFactory{}, securityGroupStatusWriter{})

	return ctrl.NewControllerManagedBy(mgr).
		For(&orcv1alpha1.SecurityGroup{}).
		WithOptions(options).
		Complete(&reconciler)

}
