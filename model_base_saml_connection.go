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

// BaseSamlConnection struct for BaseSamlConnection
type BaseSamlConnection struct {
	// The ID of the account that owns this entity.
	AccountId int32 `json:"accountId"`
	// ID of the SAML service.
	Name string `json:"name"`
	// Determines if this SAML connection active.
	Enabled bool `json:"enabled"`
	// Identity Provider Entity ID.
	Issuer string `json:"issuer"`
	// Single Sign-On URL.
	SignOnURL string `json:"signOnURL"`
	// Single Sign-Out URL.
	SignOutURL *string `json:"signOutURL,omitempty"`
	// Metadata URL.
	MetadataURL *string `json:"metadataURL,omitempty"`
	// The application-defined unique identifier that is the intended audience of the SAML assertion. This is most often the SP Entity ID of your application. When not specified, the ACS URL will be used. 
	AudienceURI *string `json:"audienceURI,omitempty"`
}

// NewBaseSamlConnection instantiates a new BaseSamlConnection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaseSamlConnection(accountId int32, name string, enabled bool, issuer string, signOnURL string) *BaseSamlConnection {
	this := BaseSamlConnection{}
	this.AccountId = accountId
	this.Name = name
	this.Enabled = enabled
	this.Issuer = issuer
	this.SignOnURL = signOnURL
	return &this
}

// NewBaseSamlConnectionWithDefaults instantiates a new BaseSamlConnection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaseSamlConnectionWithDefaults() *BaseSamlConnection {
	this := BaseSamlConnection{}
	return &this
}

// GetAccountId returns the AccountId field value
func (o *BaseSamlConnection) GetAccountId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *BaseSamlConnection) GetAccountIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *BaseSamlConnection) SetAccountId(v int32) {
	o.AccountId = v
}

// GetName returns the Name field value
func (o *BaseSamlConnection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BaseSamlConnection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BaseSamlConnection) SetName(v string) {
	o.Name = v
}

// GetEnabled returns the Enabled field value
func (o *BaseSamlConnection) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *BaseSamlConnection) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *BaseSamlConnection) SetEnabled(v bool) {
	o.Enabled = v
}

// GetIssuer returns the Issuer field value
func (o *BaseSamlConnection) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *BaseSamlConnection) GetIssuerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *BaseSamlConnection) SetIssuer(v string) {
	o.Issuer = v
}

// GetSignOnURL returns the SignOnURL field value
func (o *BaseSamlConnection) GetSignOnURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SignOnURL
}

// GetSignOnURLOk returns a tuple with the SignOnURL field value
// and a boolean to check if the value has been set.
func (o *BaseSamlConnection) GetSignOnURLOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignOnURL, true
}

// SetSignOnURL sets field value
func (o *BaseSamlConnection) SetSignOnURL(v string) {
	o.SignOnURL = v
}

// GetSignOutURL returns the SignOutURL field value if set, zero value otherwise.
func (o *BaseSamlConnection) GetSignOutURL() string {
	if o == nil || o.SignOutURL == nil {
		var ret string
		return ret
	}
	return *o.SignOutURL
}

// GetSignOutURLOk returns a tuple with the SignOutURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSamlConnection) GetSignOutURLOk() (*string, bool) {
	if o == nil || o.SignOutURL == nil {
		return nil, false
	}
	return o.SignOutURL, true
}

// HasSignOutURL returns a boolean if a field has been set.
func (o *BaseSamlConnection) HasSignOutURL() bool {
	if o != nil && o.SignOutURL != nil {
		return true
	}

	return false
}

// SetSignOutURL gets a reference to the given string and assigns it to the SignOutURL field.
func (o *BaseSamlConnection) SetSignOutURL(v string) {
	o.SignOutURL = &v
}

// GetMetadataURL returns the MetadataURL field value if set, zero value otherwise.
func (o *BaseSamlConnection) GetMetadataURL() string {
	if o == nil || o.MetadataURL == nil {
		var ret string
		return ret
	}
	return *o.MetadataURL
}

// GetMetadataURLOk returns a tuple with the MetadataURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSamlConnection) GetMetadataURLOk() (*string, bool) {
	if o == nil || o.MetadataURL == nil {
		return nil, false
	}
	return o.MetadataURL, true
}

// HasMetadataURL returns a boolean if a field has been set.
func (o *BaseSamlConnection) HasMetadataURL() bool {
	if o != nil && o.MetadataURL != nil {
		return true
	}

	return false
}

// SetMetadataURL gets a reference to the given string and assigns it to the MetadataURL field.
func (o *BaseSamlConnection) SetMetadataURL(v string) {
	o.MetadataURL = &v
}

// GetAudienceURI returns the AudienceURI field value if set, zero value otherwise.
func (o *BaseSamlConnection) GetAudienceURI() string {
	if o == nil || o.AudienceURI == nil {
		var ret string
		return ret
	}
	return *o.AudienceURI
}

// GetAudienceURIOk returns a tuple with the AudienceURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaseSamlConnection) GetAudienceURIOk() (*string, bool) {
	if o == nil || o.AudienceURI == nil {
		return nil, false
	}
	return o.AudienceURI, true
}

// HasAudienceURI returns a boolean if a field has been set.
func (o *BaseSamlConnection) HasAudienceURI() bool {
	if o != nil && o.AudienceURI != nil {
		return true
	}

	return false
}

// SetAudienceURI gets a reference to the given string and assigns it to the AudienceURI field.
func (o *BaseSamlConnection) SetAudienceURI(v string) {
	o.AudienceURI = &v
}

func (o BaseSamlConnection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["accountId"] = o.AccountId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["issuer"] = o.Issuer
	}
	if true {
		toSerialize["signOnURL"] = o.SignOnURL
	}
	if o.SignOutURL != nil {
		toSerialize["signOutURL"] = o.SignOutURL
	}
	if o.MetadataURL != nil {
		toSerialize["metadataURL"] = o.MetadataURL
	}
	if o.AudienceURI != nil {
		toSerialize["audienceURI"] = o.AudienceURI
	}
	return json.Marshal(toSerialize)
}

type NullableBaseSamlConnection struct {
	value *BaseSamlConnection
	isSet bool
}

func (v NullableBaseSamlConnection) Get() *BaseSamlConnection {
	return v.value
}

func (v *NullableBaseSamlConnection) Set(val *BaseSamlConnection) {
	v.value = val
	v.isSet = true
}

func (v NullableBaseSamlConnection) IsSet() bool {
	return v.isSet
}

func (v *NullableBaseSamlConnection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaseSamlConnection(val *BaseSamlConnection) *NullableBaseSamlConnection {
	return &NullableBaseSamlConnection{value: val, isSet: true}
}

func (v NullableBaseSamlConnection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaseSamlConnection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


