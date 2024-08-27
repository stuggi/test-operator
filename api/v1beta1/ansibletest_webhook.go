/*
Copyright 2023.

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

//
// Generated by:
//
// operator-sdk create webhook --group test --version v1beta1 --kind AnsibleTest --programmatic-validation --defaulting
//

package v1beta1

import (
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// log is for logging in this package.
var ansibletestlog = logf.Log.WithName("ansibletest-resource")

func (r *AnsibleTest) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

//+kubebuilder:webhook:path=/mutate-test-openstack-org-v1beta1-ansibletest,mutating=true,failurePolicy=fail,sideEffects=None,groups=test.openstack.org,resources=ansibletests,verbs=create;update,versions=v1beta1,name=mansibletest.kb.io,admissionReviewVersions=v1

var _ webhook.Defaulter = &AnsibleTest{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *AnsibleTest) Default() {
	ansibletestlog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-test-openstack-org-v1beta1-ansibletest,mutating=false,failurePolicy=fail,sideEffects=None,groups=test.openstack.org,resources=ansibletests,verbs=create;update,versions=v1beta1,name=vansibletest.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &AnsibleTest{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *AnsibleTest) ValidateCreate() (admission.Warnings, error) {
	ansibletestlog.Info("validate create", "name", r.Name)

	var allWarnings admission.Warnings

	if r.Spec.Privileged {
		allWarnings = append(allWarnings, fmt.Sprintf(WarnPrivilegedModeOn, "AnsibleTest"))
	}

	return allWarnings, nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *AnsibleTest) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	ansibletestlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil, nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *AnsibleTest) ValidateDelete() (admission.Warnings, error) {
	ansibletestlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil, nil
}
