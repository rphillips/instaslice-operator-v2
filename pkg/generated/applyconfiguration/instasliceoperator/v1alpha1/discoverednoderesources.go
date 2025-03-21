/*
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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
)

// DiscoveredNodeResourcesApplyConfiguration represents a declarative configuration of the DiscoveredNodeResources type for use
// with apply.
type DiscoveredNodeResourcesApplyConfiguration struct {
	NodeGPUs      []DiscoveredGPUApplyConfiguration `json:"nodeGpus,omitempty"`
	MigPlacement  map[string]MigApplyConfiguration  `json:"migPlacement,omitempty"`
	NodeResources *v1.ResourceList                  `json:"nodeResources,omitempty"`
}

// DiscoveredNodeResourcesApplyConfiguration constructs a declarative configuration of the DiscoveredNodeResources type for use with
// apply.
func DiscoveredNodeResources() *DiscoveredNodeResourcesApplyConfiguration {
	return &DiscoveredNodeResourcesApplyConfiguration{}
}

// WithNodeGPUs adds the given value to the NodeGPUs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the NodeGPUs field.
func (b *DiscoveredNodeResourcesApplyConfiguration) WithNodeGPUs(values ...*DiscoveredGPUApplyConfiguration) *DiscoveredNodeResourcesApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithNodeGPUs")
		}
		b.NodeGPUs = append(b.NodeGPUs, *values[i])
	}
	return b
}

// WithMigPlacement puts the entries into the MigPlacement field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the MigPlacement field,
// overwriting an existing map entries in MigPlacement field with the same key.
func (b *DiscoveredNodeResourcesApplyConfiguration) WithMigPlacement(entries map[string]MigApplyConfiguration) *DiscoveredNodeResourcesApplyConfiguration {
	if b.MigPlacement == nil && len(entries) > 0 {
		b.MigPlacement = make(map[string]MigApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.MigPlacement[k] = v
	}
	return b
}

// WithNodeResources sets the NodeResources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeResources field is set to the value of the last call.
func (b *DiscoveredNodeResourcesApplyConfiguration) WithNodeResources(value v1.ResourceList) *DiscoveredNodeResourcesApplyConfiguration {
	b.NodeResources = &value
	return b
}
