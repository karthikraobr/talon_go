/*
Integration API reference docs

Use the Integration API to push data to and retrieve data from Talon.One in real time. For more background information about this API, see [Integration API overview](/docs/dev/integration-api/overview)  For example, use this API to share shopping cart information as a session with Talon.One and evaluate promotion rules. You can also create custom events to track specific actions that do not fit into the session data model.  Ensure you [authenticate](#section/Authentication) to make requests to the API.  <div class=\"redoc-section\">   <p class=\"title\">Are you looking for a different API?</p>    If you need the API to:    - Interact with the Campaign Manager for backoffice operations, see [the Management API reference docs](https://docs.talon.one/management-api).   - Integrate with Talon.One from a CEP or CDP platform, see [the Third-party API reference docs](https://docs.talon.one/third-party-api).  </div>  # Authentication  <SecurityDefinitions /> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package talon

import (
	"encoding/json"
)

// CreateReferralsForMultipleAdvocates201Response struct for CreateReferralsForMultipleAdvocates201Response
type CreateReferralsForMultipleAdvocates201Response struct {
	TotalResultSize int32 `json:"totalResultSize"`
	Data []Referral `json:"data"`
}

// NewCreateReferralsForMultipleAdvocates201Response instantiates a new CreateReferralsForMultipleAdvocates201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateReferralsForMultipleAdvocates201Response(totalResultSize int32, data []Referral) *CreateReferralsForMultipleAdvocates201Response {
	this := CreateReferralsForMultipleAdvocates201Response{}
	this.TotalResultSize = totalResultSize
	this.Data = data
	return &this
}

// NewCreateReferralsForMultipleAdvocates201ResponseWithDefaults instantiates a new CreateReferralsForMultipleAdvocates201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateReferralsForMultipleAdvocates201ResponseWithDefaults() *CreateReferralsForMultipleAdvocates201Response {
	this := CreateReferralsForMultipleAdvocates201Response{}
	return &this
}

// GetTotalResultSize returns the TotalResultSize field value
func (o *CreateReferralsForMultipleAdvocates201Response) GetTotalResultSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TotalResultSize
}

// GetTotalResultSizeOk returns a tuple with the TotalResultSize field value
// and a boolean to check if the value has been set.
func (o *CreateReferralsForMultipleAdvocates201Response) GetTotalResultSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalResultSize, true
}

// SetTotalResultSize sets field value
func (o *CreateReferralsForMultipleAdvocates201Response) SetTotalResultSize(v int32) {
	o.TotalResultSize = v
}

// GetData returns the Data field value
func (o *CreateReferralsForMultipleAdvocates201Response) GetData() []Referral {
	if o == nil {
		var ret []Referral
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreateReferralsForMultipleAdvocates201Response) GetDataOk() ([]Referral, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CreateReferralsForMultipleAdvocates201Response) SetData(v []Referral) {
	o.Data = v
}

func (o CreateReferralsForMultipleAdvocates201Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["totalResultSize"] = o.TotalResultSize
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableCreateReferralsForMultipleAdvocates201Response struct {
	value *CreateReferralsForMultipleAdvocates201Response
	isSet bool
}

func (v NullableCreateReferralsForMultipleAdvocates201Response) Get() *CreateReferralsForMultipleAdvocates201Response {
	return v.value
}

func (v *NullableCreateReferralsForMultipleAdvocates201Response) Set(val *CreateReferralsForMultipleAdvocates201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateReferralsForMultipleAdvocates201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateReferralsForMultipleAdvocates201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateReferralsForMultipleAdvocates201Response(val *CreateReferralsForMultipleAdvocates201Response) *NullableCreateReferralsForMultipleAdvocates201Response {
	return &NullableCreateReferralsForMultipleAdvocates201Response{value: val, isSet: true}
}

func (v NullableCreateReferralsForMultipleAdvocates201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateReferralsForMultipleAdvocates201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


