// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/config/v1"
)

// IBMCloudServiceEndpointApplyConfiguration represents a declarative configuration of the IBMCloudServiceEndpoint type for use
// with apply.
type IBMCloudServiceEndpointApplyConfiguration struct {
	Name *v1.IBMCloudServiceName `json:"name,omitempty"`
	URL  *string                 `json:"url,omitempty"`
}

// IBMCloudServiceEndpointApplyConfiguration constructs a declarative configuration of the IBMCloudServiceEndpoint type for use with
// apply.
func IBMCloudServiceEndpoint() *IBMCloudServiceEndpointApplyConfiguration {
	return &IBMCloudServiceEndpointApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *IBMCloudServiceEndpointApplyConfiguration) WithName(value v1.IBMCloudServiceName) *IBMCloudServiceEndpointApplyConfiguration {
	b.Name = &value
	return b
}

// WithURL sets the URL field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the URL field is set to the value of the last call.
func (b *IBMCloudServiceEndpointApplyConfiguration) WithURL(value string) *IBMCloudServiceEndpointApplyConfiguration {
	b.URL = &value
	return b
}