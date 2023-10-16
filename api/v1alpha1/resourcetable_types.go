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

// ResourceTableColumn defines the desired state of ResourceTable
type ResourceTableColumn struct {
	Name string `json:"name,omitempty"`
}

// ResourceTableDefaultOrder defines the desired state of ResourceTable
type ResourceTableDefaultOrder struct {
	Name string `json:"name,omitempty"`
}

// ResourceTableSpec defines the desired state of ResourceTable
type ResourceTableSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Selector     ResourceKindSelector      `json:"selector,omitempty"`
	Columns      []ResourceTableColumn     `json:"columns,omitempty"`
	DefaultOrder ResourceTableDefaultOrder `json:"defaultOrder,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ResourceTable is the Schema for the resourcetables API
type ResourceTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ResourceTableSpec `json:"spec,omitempty"`
}

//+kubebuilder:object:root=true

// ResourceTableList contains a list of ResourceTable
type ResourceTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceTable `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceTable{}, &ResourceTableList{})
}
