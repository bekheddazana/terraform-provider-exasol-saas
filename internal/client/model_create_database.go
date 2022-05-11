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

// CreateDatabase struct for CreateDatabase
type CreateDatabase struct {
	Name           string        `json:"name"`
	InitialCluster CreateCluster `json:"initialCluster"`
	Provider       string        `json:"provider"`
	Region         string        `json:"region"`
}

// NewCreateDatabase instantiates a new CreateDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatabase(name string, initialCluster CreateCluster, provider string, region string) *CreateDatabase {
	this := CreateDatabase{}
	this.Name = name
	this.InitialCluster = initialCluster
	this.Provider = provider
	this.Region = region
	return &this
}

// NewCreateDatabaseWithDefaults instantiates a new CreateDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatabaseWithDefaults() *CreateDatabase {
	this := CreateDatabase{}
	return &this
}

// GetName returns the Name field value
func (o *CreateDatabase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDatabase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDatabase) SetName(v string) {
	o.Name = v
}

// GetInitialCluster returns the InitialCluster field value
func (o *CreateDatabase) GetInitialCluster() CreateCluster {
	if o == nil {
		var ret CreateCluster
		return ret
	}

	return o.InitialCluster
}

// GetInitialClusterOk returns a tuple with the InitialCluster field value
// and a boolean to check if the value has been set.
func (o *CreateDatabase) GetInitialClusterOk() (*CreateCluster, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialCluster, true
}

// SetInitialCluster sets field value
func (o *CreateDatabase) SetInitialCluster(v CreateCluster) {
	o.InitialCluster = v
}

// GetProvider returns the Provider field value
func (o *CreateDatabase) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *CreateDatabase) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *CreateDatabase) SetProvider(v string) {
	o.Provider = v
}

// GetRegion returns the Region field value
func (o *CreateDatabase) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *CreateDatabase) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *CreateDatabase) SetRegion(v string) {
	o.Region = v
}

func (o CreateDatabase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["initialCluster"] = o.InitialCluster
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["region"] = o.Region
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDatabase struct {
	value *CreateDatabase
	isSet bool
}

func (v NullableCreateDatabase) Get() *CreateDatabase {
	return v.value
}

func (v *NullableCreateDatabase) Set(val *CreateDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatabase(val *CreateDatabase) *NullableCreateDatabase {
	return &NullableCreateDatabase{value: val, isSet: true}
}

func (v NullableCreateDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
