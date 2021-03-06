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

// CustomerProfileIntegrationRequestV2 
type CustomerProfileIntegrationRequestV2 struct {
	// Arbitrary properties associated with this item.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	AudiencesChanges *ProfileAudiencesChanges `json:"audiencesChanges,omitempty"`
	// Optional list of extra data that you want to get in the response. Use this property to get as much data as you need in one request instead of sending extra requests to other endpoints.  **Note:** `ruleFailureReasons` is always part of the response when the [Application type](https://docs.talon.one/docs/product/applications/overview#application-types) is `sandbox`. 
	ResponseContent []string `json:"responseContent,omitempty"`
}

// NewCustomerProfileIntegrationRequestV2 instantiates a new CustomerProfileIntegrationRequestV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerProfileIntegrationRequestV2() *CustomerProfileIntegrationRequestV2 {
	this := CustomerProfileIntegrationRequestV2{}
	return &this
}

// NewCustomerProfileIntegrationRequestV2WithDefaults instantiates a new CustomerProfileIntegrationRequestV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerProfileIntegrationRequestV2WithDefaults() *CustomerProfileIntegrationRequestV2 {
	this := CustomerProfileIntegrationRequestV2{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *CustomerProfileIntegrationRequestV2) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileIntegrationRequestV2) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *CustomerProfileIntegrationRequestV2) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *CustomerProfileIntegrationRequestV2) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetAudiencesChanges returns the AudiencesChanges field value if set, zero value otherwise.
func (o *CustomerProfileIntegrationRequestV2) GetAudiencesChanges() ProfileAudiencesChanges {
	if o == nil || o.AudiencesChanges == nil {
		var ret ProfileAudiencesChanges
		return ret
	}
	return *o.AudiencesChanges
}

// GetAudiencesChangesOk returns a tuple with the AudiencesChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileIntegrationRequestV2) GetAudiencesChangesOk() (*ProfileAudiencesChanges, bool) {
	if o == nil || o.AudiencesChanges == nil {
		return nil, false
	}
	return o.AudiencesChanges, true
}

// HasAudiencesChanges returns a boolean if a field has been set.
func (o *CustomerProfileIntegrationRequestV2) HasAudiencesChanges() bool {
	if o != nil && o.AudiencesChanges != nil {
		return true
	}

	return false
}

// SetAudiencesChanges gets a reference to the given ProfileAudiencesChanges and assigns it to the AudiencesChanges field.
func (o *CustomerProfileIntegrationRequestV2) SetAudiencesChanges(v ProfileAudiencesChanges) {
	o.AudiencesChanges = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *CustomerProfileIntegrationRequestV2) GetResponseContent() []string {
	if o == nil || o.ResponseContent == nil {
		var ret []string
		return ret
	}
	return o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerProfileIntegrationRequestV2) GetResponseContentOk() ([]string, bool) {
	if o == nil || o.ResponseContent == nil {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *CustomerProfileIntegrationRequestV2) HasResponseContent() bool {
	if o != nil && o.ResponseContent != nil {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given []string and assigns it to the ResponseContent field.
func (o *CustomerProfileIntegrationRequestV2) SetResponseContent(v []string) {
	o.ResponseContent = v
}

func (o CustomerProfileIntegrationRequestV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.AudiencesChanges != nil {
		toSerialize["audiencesChanges"] = o.AudiencesChanges
	}
	if o.ResponseContent != nil {
		toSerialize["responseContent"] = o.ResponseContent
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerProfileIntegrationRequestV2 struct {
	value *CustomerProfileIntegrationRequestV2
	isSet bool
}

func (v NullableCustomerProfileIntegrationRequestV2) Get() *CustomerProfileIntegrationRequestV2 {
	return v.value
}

func (v *NullableCustomerProfileIntegrationRequestV2) Set(val *CustomerProfileIntegrationRequestV2) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerProfileIntegrationRequestV2) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerProfileIntegrationRequestV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerProfileIntegrationRequestV2(val *CustomerProfileIntegrationRequestV2) *NullableCustomerProfileIntegrationRequestV2 {
	return &NullableCustomerProfileIntegrationRequestV2{value: val, isSet: true}
}

func (v NullableCustomerProfileIntegrationRequestV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerProfileIntegrationRequestV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


