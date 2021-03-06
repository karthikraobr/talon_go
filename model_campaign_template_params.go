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

// CampaignTemplateParams struct for CampaignTemplateParams
type CampaignTemplateParams struct {
	// Name of the campaign template parameter.
	Name string `json:"name"`
	// Defines the type of parameter value.
	Type string `json:"type"`
	// Explains the meaning of this template parameter and the placeholder value that will define it. It is used on campaign creation from this template.
	Description string `json:"description"`
	// ID of the corresponding attribute.
	AttributeId *int32 `json:"attributeId,omitempty"`
}

// NewCampaignTemplateParams instantiates a new CampaignTemplateParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignTemplateParams(name string, type_ string, description string) *CampaignTemplateParams {
	this := CampaignTemplateParams{}
	this.Name = name
	this.Type = type_
	this.Description = description
	return &this
}

// NewCampaignTemplateParamsWithDefaults instantiates a new CampaignTemplateParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignTemplateParamsWithDefaults() *CampaignTemplateParams {
	this := CampaignTemplateParams{}
	return &this
}

// GetName returns the Name field value
func (o *CampaignTemplateParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CampaignTemplateParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CampaignTemplateParams) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CampaignTemplateParams) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CampaignTemplateParams) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CampaignTemplateParams) SetType(v string) {
	o.Type = v
}

// GetDescription returns the Description field value
func (o *CampaignTemplateParams) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *CampaignTemplateParams) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *CampaignTemplateParams) SetDescription(v string) {
	o.Description = v
}

// GetAttributeId returns the AttributeId field value if set, zero value otherwise.
func (o *CampaignTemplateParams) GetAttributeId() int32 {
	if o == nil || o.AttributeId == nil {
		var ret int32
		return ret
	}
	return *o.AttributeId
}

// GetAttributeIdOk returns a tuple with the AttributeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CampaignTemplateParams) GetAttributeIdOk() (*int32, bool) {
	if o == nil || o.AttributeId == nil {
		return nil, false
	}
	return o.AttributeId, true
}

// HasAttributeId returns a boolean if a field has been set.
func (o *CampaignTemplateParams) HasAttributeId() bool {
	if o != nil && o.AttributeId != nil {
		return true
	}

	return false
}

// SetAttributeId gets a reference to the given int32 and assigns it to the AttributeId field.
func (o *CampaignTemplateParams) SetAttributeId(v int32) {
	o.AttributeId = &v
}

func (o CampaignTemplateParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.AttributeId != nil {
		toSerialize["attributeId"] = o.AttributeId
	}
	return json.Marshal(toSerialize)
}

type NullableCampaignTemplateParams struct {
	value *CampaignTemplateParams
	isSet bool
}

func (v NullableCampaignTemplateParams) Get() *CampaignTemplateParams {
	return v.value
}

func (v *NullableCampaignTemplateParams) Set(val *CampaignTemplateParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignTemplateParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignTemplateParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignTemplateParams(val *CampaignTemplateParams) *NullableCampaignTemplateParams {
	return &NullableCampaignTemplateParams{value: val, isSet: true}
}

func (v NullableCampaignTemplateParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignTemplateParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


