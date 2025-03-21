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

// PlacementApplyConfiguration represents a declarative configuration of the Placement type for use
// with apply.
type PlacementApplyConfiguration struct {
	Size  *int32 `json:"size,omitempty"`
	Start *int32 `json:"start,omitempty"`
}

// PlacementApplyConfiguration constructs a declarative configuration of the Placement type for use with
// apply.
func Placement() *PlacementApplyConfiguration {
	return &PlacementApplyConfiguration{}
}

// WithSize sets the Size field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Size field is set to the value of the last call.
func (b *PlacementApplyConfiguration) WithSize(value int32) *PlacementApplyConfiguration {
	b.Size = &value
	return b
}

// WithStart sets the Start field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Start field is set to the value of the last call.
func (b *PlacementApplyConfiguration) WithStart(value int32) *PlacementApplyConfiguration {
	b.Start = &value
	return b
}
