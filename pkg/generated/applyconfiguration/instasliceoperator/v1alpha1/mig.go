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

// MigApplyConfiguration represents a declarative configuration of the Mig type for use
// with apply.
type MigApplyConfiguration struct {
	Placements     []PlacementApplyConfiguration `json:"placements,omitempty"`
	GIProfileID    *int32                        `json:"giProfileId,omitempty"`
	CIProfileID    *int32                        `json:"ciProfileId,omitempty"`
	CIEngProfileID *int32                        `json:"ciEngProfileId,omitempty"`
}

// MigApplyConfiguration constructs a declarative configuration of the Mig type for use with
// apply.
func Mig() *MigApplyConfiguration {
	return &MigApplyConfiguration{}
}

// WithPlacements adds the given value to the Placements field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Placements field.
func (b *MigApplyConfiguration) WithPlacements(values ...*PlacementApplyConfiguration) *MigApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithPlacements")
		}
		b.Placements = append(b.Placements, *values[i])
	}
	return b
}

// WithGIProfileID sets the GIProfileID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GIProfileID field is set to the value of the last call.
func (b *MigApplyConfiguration) WithGIProfileID(value int32) *MigApplyConfiguration {
	b.GIProfileID = &value
	return b
}

// WithCIProfileID sets the CIProfileID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CIProfileID field is set to the value of the last call.
func (b *MigApplyConfiguration) WithCIProfileID(value int32) *MigApplyConfiguration {
	b.CIProfileID = &value
	return b
}

// WithCIEngProfileID sets the CIEngProfileID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CIEngProfileID field is set to the value of the last call.
func (b *MigApplyConfiguration) WithCIEngProfileID(value int32) *MigApplyConfiguration {
	b.CIEngProfileID = &value
	return b
}
