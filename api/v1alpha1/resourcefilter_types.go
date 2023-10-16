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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ResourceFilter defines the desired state of ResourceTable
type ResourceFilterSpecFilter struct {
	Name    string                           `json:"name,omitempty"`
	Options []ResourceFilterSpecFilterOption `json:"options,omitempty"`
}

// ResourceFilterSpecFilterOption defines the desired state of ResourceTable
type ResourceFilterSpecFilterOption struct {
	Name string `json:"name,omitempty"`
}

// ResourceFilterSpec defines the desired state of ResourceFilter
type ResourceFilterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Selector ResourceKindSelector     `json:"selector,omitempty"`
	Filter   ResourceFilterSpecFilter `json:"filter,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ResourceFilter is the Schema for the resourcefilters API
type ResourceFilter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ResourceFilterSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// ResourceFilterList contains a list of ResourceFilter
type ResourceFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceFilter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceFilter{}, &ResourceFilterList{})
}
