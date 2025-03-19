/*
Copyright 2019 The Crossplane Authors.

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

// Package protection implements the Crossplane deletion protection controller.
package protection

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/crossplane-runtime/pkg/controller"

	"github.com/crossplane/crossplane/internal/controller/protection/usage"
	"github.com/crossplane/crossplane/internal/features"
)

// Setup API extensions controllers.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	if !o.Features.Enabled(features.EnableBetaUsages) {
		return nil
	}

	for _, fn := range []func(mgr ctrl.Manager, o controller.Options) error{
		usage.SetupUsage,
		usage.SetupClusterUsage,
		usage.SetupLegacyUsage,
	} {
		if err := fn(mgr, o); err != nil {
			return err
		}
	}

	return nil
}
