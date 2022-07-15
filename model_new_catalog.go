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

// NewCatalog 
type NewCatalog struct {
	// The cart item catalog name.
	Name string `json:"name"`
	// A description of this cart item catalog.
	Description string `json:"description"`
	// A list of the IDs of the applications that are subscribed to this catalog.
	SubscribedApplicationsIds []int32 `json:"subscribedApplicationsIds,omitempty"`
}

// NewNewCatalog instantiates a new NewCatalog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewCatalog(name string, description string) *NewCatalog {
	this := NewCatalog{}
	this.Name = name
	this.Description = description
	return &this
}

// NewNewCatalogWithDefaults instantiates a new NewCatalog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewCatalogWithDefaults() *NewCatalog {
	this := NewCatalog{}
	return &this
}

// GetName returns the Name field value
func (o *NewCatalog) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *NewCatalog) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *NewCatalog) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *NewCatalog) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *NewCatalog) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *NewCatalog) SetDescription(v string) {
	o.Description = v
}

// GetSubscribedApplicationsIds returns the SubscribedApplicationsIds field value if set, zero value otherwise.
func (o *NewCatalog) GetSubscribedApplicationsIds() []int32 {
	if o == nil || o.SubscribedApplicationsIds == nil {
		var ret []int32
		return ret
	}
	return o.SubscribedApplicationsIds
}

// GetSubscribedApplicationsIdsOk returns a tuple with the SubscribedApplicationsIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewCatalog) GetSubscribedApplicationsIdsOk() ([]int32, bool) {
	if o == nil || o.SubscribedApplicationsIds == nil {
		return nil, false
	}
	return o.SubscribedApplicationsIds, true
}

// HasSubscribedApplicationsIds returns a boolean if a field has been set.
func (o *NewCatalog) HasSubscribedApplicationsIds() bool {
	if o != nil && o.SubscribedApplicationsIds != nil {
		return true
	}

	return false
}

// SetSubscribedApplicationsIds gets a reference to the given []int32 and assigns it to the SubscribedApplicationsIds field.
func (o *NewCatalog) SetSubscribedApplicationsIds(v []int32) {
	o.SubscribedApplicationsIds = v
}

func (o NewCatalog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.SubscribedApplicationsIds != nil {
		toSerialize["subscribedApplicationsIds"] = o.SubscribedApplicationsIds
	}
	return json.Marshal(toSerialize)
}

type NullableNewCatalog struct {
	value *NewCatalog
	isSet bool
}

func (v NullableNewCatalog) Get() *NewCatalog {
	return v.value
}

func (v *NullableNewCatalog) Set(val *NewCatalog) {
	v.value = val
	v.isSet = true
}

func (v NullableNewCatalog) IsSet() bool {
	return v.isSet
}

func (v *NullableNewCatalog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewCatalog(val *NewCatalog) *NullableNewCatalog {
	return &NullableNewCatalog{value: val, isSet: true}
}

func (v NullableNewCatalog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewCatalog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


