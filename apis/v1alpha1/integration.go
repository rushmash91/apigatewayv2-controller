// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IntegrationSpec defines the desired state of Integration.
//
// Represents an integration.
type IntegrationSpec struct {
	APIID  *string                                  `json:"apiID,omitempty"`
	APIRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"apiRef,omitempty"`

	ConnectionID  *string                                  `json:"connectionID,omitempty"`
	ConnectionRef *ackv1alpha1.AWSResourceReferenceWrapper `json:"connectionRef,omitempty"`

	ConnectionType *string `json:"connectionType,omitempty"`

	ContentHandlingStrategy *string `json:"contentHandlingStrategy,omitempty"`

	CredentialsARN *string `json:"credentialsARN,omitempty"`

	Description *string `json:"description,omitempty"`

	IntegrationMethod *string `json:"integrationMethod,omitempty"`

	IntegrationSubtype *string `json:"integrationSubtype,omitempty"`

	// +kubebuilder:validation:Required
	IntegrationType *string `json:"integrationType"`

	IntegrationURI *string `json:"integrationURI,omitempty"`

	PassthroughBehavior *string `json:"passthroughBehavior,omitempty"`

	PayloadFormatVersion *string `json:"payloadFormatVersion,omitempty"`

	RequestParameters map[string]*string `json:"requestParameters,omitempty"`

	RequestTemplates map[string]*string `json:"requestTemplates,omitempty"`

	ResponseParameters map[string]map[string]*string `json:"responseParameters,omitempty"`

	TemplateSelectionExpression *string `json:"templateSelectionExpression,omitempty"`

	TimeoutInMillis *int64 `json:"timeoutInMillis,omitempty"`

	TLSConfig *TLSConfigInput `json:"tlsConfig,omitempty"`
}

// IntegrationStatus defines the observed state of Integration
type IntegrationStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`

	// +kubebuilder:validation:Optional
	APIGatewayManaged *bool `json:"apiGatewayManaged,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationID *string `json:"integrationID,omitempty"`

	// +kubebuilder:validation:Optional
	IntegrationResponseSelectionExpression *string `json:"integrationResponseSelectionExpression,omitempty"`
}

// Integration is the Schema for the Integrations API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Integration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IntegrationSpec   `json:"spec,omitempty"`
	Status            IntegrationStatus `json:"status,omitempty"`
}

// IntegrationList contains a list of Integration
// +kubebuilder:object:root=true
type IntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Integration `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Integration{}, &IntegrationList{})
}
