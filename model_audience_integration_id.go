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

// AudienceIntegrationID struct for AudienceIntegrationID
type AudienceIntegrationID struct {
	// The ID of this audience in the third-party integration.
	IntegrationId *string `json:"integrationId,omitempty"`
}

// NewAudienceIntegrationID instantiates a new AudienceIntegrationID object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAudienceIntegrationID() *AudienceIntegrationID {
	this := AudienceIntegrationID{}
	return &this
}

// NewAudienceIntegrationIDWithDefaults instantiates a new AudienceIntegrationID object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAudienceIntegrationIDWithDefaults() *AudienceIntegrationID {
	this := AudienceIntegrationID{}
	return &this
}

// GetIntegrationId returns the IntegrationId field value if set, zero value otherwise.
func (o *AudienceIntegrationID) GetIntegrationId() string {
	if o == nil || o.IntegrationId == nil {
		var ret string
		return ret
	}
	return *o.IntegrationId
}

// GetIntegrationIdOk returns a tuple with the IntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AudienceIntegrationID) GetIntegrationIdOk() (*string, bool) {
	if o == nil || o.IntegrationId == nil {
		return nil, false
	}
	return o.IntegrationId, true
}

// HasIntegrationId returns a boolean if a field has been set.
func (o *AudienceIntegrationID) HasIntegrationId() bool {
	if o != nil && o.IntegrationId != nil {
		return true
	}

	return false
}

// SetIntegrationId gets a reference to the given string and assigns it to the IntegrationId field.
func (o *AudienceIntegrationID) SetIntegrationId(v string) {
	o.IntegrationId = &v
}

func (o AudienceIntegrationID) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IntegrationId != nil {
		toSerialize["integrationId"] = o.IntegrationId
	}
	return json.Marshal(toSerialize)
}

type NullableAudienceIntegrationID struct {
	value *AudienceIntegrationID
	isSet bool
}

func (v NullableAudienceIntegrationID) Get() *AudienceIntegrationID {
	return v.value
}

func (v *NullableAudienceIntegrationID) Set(val *AudienceIntegrationID) {
	v.value = val
	v.isSet = true
}

func (v NullableAudienceIntegrationID) IsSet() bool {
	return v.isSet
}

func (v *NullableAudienceIntegrationID) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAudienceIntegrationID(val *AudienceIntegrationID) *NullableAudienceIntegrationID {
	return &NullableAudienceIntegrationID{value: val, isSet: true}
}

func (v NullableAudienceIntegrationID) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAudienceIntegrationID) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


