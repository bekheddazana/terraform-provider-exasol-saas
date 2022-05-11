/*
Exasol SaaS REST-API

## Authentication  The REST API can be used with your Personal Access Token (PAT). You don't know what a PAT is, check our documentation  [here](https://docs.exasol.com/saas/administration/access_mngt/access_token.htm).  After you created a PAT click on Authorize and add your PAT under BearerAuth.

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// UpdateCluster struct for UpdateCluster
type UpdateCluster struct {
	Name     string    `json:"name"`
	AutoStop *AutoStop `json:"autoStop,omitempty"`
}

// NewUpdateCluster instantiates a new UpdateCluster object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCluster(name string) *UpdateCluster {
	this := UpdateCluster{}
	this.Name = name
	return &this
}

// NewUpdateClusterWithDefaults instantiates a new UpdateCluster object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateClusterWithDefaults() *UpdateCluster {
	this := UpdateCluster{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateCluster) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateCluster) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateCluster) SetName(v string) {
	o.Name = v
}

// GetAutoStop returns the AutoStop field value if set, zero value otherwise.
func (o *UpdateCluster) GetAutoStop() AutoStop {
	if o == nil || o.AutoStop == nil {
		var ret AutoStop
		return ret
	}
	return *o.AutoStop
}

// GetAutoStopOk returns a tuple with the AutoStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCluster) GetAutoStopOk() (*AutoStop, bool) {
	if o == nil || o.AutoStop == nil {
		return nil, false
	}
	return o.AutoStop, true
}

// HasAutoStop returns a boolean if a field has been set.
func (o *UpdateCluster) HasAutoStop() bool {
	if o != nil && o.AutoStop != nil {
		return true
	}

	return false
}

// SetAutoStop gets a reference to the given AutoStop and assigns it to the AutoStop field.
func (o *UpdateCluster) SetAutoStop(v AutoStop) {
	o.AutoStop = &v
}

func (o UpdateCluster) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.AutoStop != nil {
		toSerialize["autoStop"] = o.AutoStop
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCluster struct {
	value *UpdateCluster
	isSet bool
}

func (v NullableUpdateCluster) Get() *UpdateCluster {
	return v.value
}

func (v *NullableUpdateCluster) Set(val *UpdateCluster) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCluster) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCluster) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCluster(val *UpdateCluster) *NullableUpdateCluster {
	return &NullableUpdateCluster{value: val, isSet: true}
}

func (v NullableUpdateCluster) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCluster) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
