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

// SamlConnectionMetadata struct for SamlConnectionMetadata
type SamlConnectionMetadata struct {
	// ID of the SAML service.
	Name string `json:"name"`
	// Determines if this SAML connection active.
	Enabled bool `json:"enabled"`
	AccountId float32 `json:"accountId"`
	// Identity Provider metadata XML document.
	MetadataDocument string `json:"metadataDocument"`
}

// NewSamlConnectionMetadata instantiates a new SamlConnectionMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlConnectionMetadata(name string, enabled bool, accountId float32, metadataDocument string) *SamlConnectionMetadata {
	this := SamlConnectionMetadata{}
	this.Name = name
	this.Enabled = enabled
	this.AccountId = accountId
	this.MetadataDocument = metadataDocument
	return &this
}

// NewSamlConnectionMetadataWithDefaults instantiates a new SamlConnectionMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlConnectionMetadataWithDefaults() *SamlConnectionMetadata {
	this := SamlConnectionMetadata{}
	return &this
}

// GetName returns the Name field value
func (o *SamlConnectionMetadata) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SamlConnectionMetadata) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SamlConnectionMetadata) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *SamlConnectionMetadata) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SamlConnectionMetadata) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SamlConnectionMetadata) SetEnabled(v bool) {
	o.Enabled = v
}

// GetAccountId returns the AccountId field value
func (o *SamlConnectionMetadata) GetAccountId() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *SamlConnectionMetadata) GetAccountIdOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *SamlConnectionMetadata) SetAccountId(v float32) {
	o.AccountId = v
}

// GetMetadataDocument returns the MetadataDocument field value
func (o *SamlConnectionMetadata) GetMetadataDocument() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataDocument
}

// GetMetadataDocumentOk returns a tuple with the MetadataDocument field value
// and a boolean to check if the value has been set.
func (o *SamlConnectionMetadata) GetMetadataDocumentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataDocument, true
}

// SetMetadataDocument sets field value
func (o *SamlConnectionMetadata) SetMetadataDocument(v string) {
	o.MetadataDocument = v
}

func (o SamlConnectionMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["metadataDocument"] = o.MetadataDocument
	}
	return json.Marshal(toSerialize)
}

type NullableSamlConnectionMetadata struct {
	value *SamlConnectionMetadata
	isSet bool
}

func (v NullableSamlConnectionMetadata) Get() *SamlConnectionMetadata {
	return v.value
}

func (v *NullableSamlConnectionMetadata) Set(val *SamlConnectionMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlConnectionMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlConnectionMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlConnectionMetadata(val *SamlConnectionMetadata) *NullableSamlConnectionMetadata {
	return &NullableSamlConnectionMetadata{value: val, isSet: true}
}

func (v NullableSamlConnectionMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlConnectionMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


