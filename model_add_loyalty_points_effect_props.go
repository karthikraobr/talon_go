/*
Integration API reference docs

Use the Integration API to push data to and retrieve data from Talon.One in real time. For more background information about this API, see [Integration API overview](/docs/dev/integration-api/overview)  For example, use this API to share shopping cart information as a session with Talon.One and evaluate promotion rules. You can also create custom events to track specific actions that do not fit into the session data model.  Ensure you [authenticate](#section/Authentication) to make requests to the API.  <div class=\"redoc-section\">   <p class=\"title\">Are you looking for a different API?</p>    If you need the API to:    - Interact with the Campaign Manager for backoffice operations, see [the Management API reference docs](https://docs.talon.one/management-api).   - Integrate with Talon.One from a CEP or CDP platform, see [the Third-party API reference docs](https://docs.talon.one/third-party-api).  </div>  # Authentication  <SecurityDefinitions /> 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package talon

import (
	"encoding/json"
	"time"
)

// AddLoyaltyPointsEffectProps The properties specific to the \"addLoyaltyPoints\" effect. This gets triggered whenever a validated rule contained an \"add loyalty\" effect. These points are automatically stored and managed inside Talon.One. 
type AddLoyaltyPointsEffectProps struct {
	// The name/description of this loyalty point addition.
	Name string `json:"name"`
	// The ID of the loyalty program where these points were added.
	ProgramId int32 `json:"programId"`
	// The ID of the subledger within the loyalty program where these points were added.
	SubLedgerId string `json:"subLedgerId"`
	// The amount of points that were added.
	Value float32 `json:"value"`
	// The original amount of loyalty points to be awarded.
	DesiredValue *float32 `json:"desiredValue,omitempty"`
	// The user for whom these points were added.
	RecipientIntegrationId string `json:"recipientIntegrationId"`
	// Date after which points will be valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Date after which points will expire.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The identifier of this addition in the loyalty ledger.
	TransactionUUID string `json:"transactionUUID"`
	// The index of the item in the cart items list on which the loyal points addition should be applied.
	CartItemPosition *float32 `json:"cartItemPosition,omitempty"`
	// The sub position is triggered when application flattening is enabled. It indicates to which item the loyalty points addition applies, for cart items with `quantity` > 1. 
	CartItemSubPosition *float32 `json:"cartItemSubPosition,omitempty"`
	// The card on which these points were added.
	CardIdentifier *string `json:"cardIdentifier,omitempty"`
}

// NewAddLoyaltyPointsEffectProps instantiates a new AddLoyaltyPointsEffectProps object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddLoyaltyPointsEffectProps(name string, programId int32, subLedgerId string, value float32, recipientIntegrationId string, transactionUUID string) *AddLoyaltyPointsEffectProps {
	this := AddLoyaltyPointsEffectProps{}
	this.Name = name
	this.ProgramId = programId
	this.SubLedgerId = subLedgerId
	this.Value = value
	this.RecipientIntegrationId = recipientIntegrationId
	this.TransactionUUID = transactionUUID
	return &this
}

// NewAddLoyaltyPointsEffectPropsWithDefaults instantiates a new AddLoyaltyPointsEffectProps object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddLoyaltyPointsEffectPropsWithDefaults() *AddLoyaltyPointsEffectProps {
	this := AddLoyaltyPointsEffectProps{}
	return &this
}

// GetName returns the Name field value
func (o *AddLoyaltyPointsEffectProps) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddLoyaltyPointsEffectProps) SetName(v string) {
	o.Name = v
}

// GetProgramId returns the ProgramId field value
func (o *AddLoyaltyPointsEffectProps) GetProgramId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ProgramId
}

// GetProgramIdOk returns a tuple with the ProgramId field value
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetProgramIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProgramId, true
}

// SetProgramId sets field value
func (o *AddLoyaltyPointsEffectProps) SetProgramId(v int32) {
	o.ProgramId = v
}

// GetSubLedgerId returns the SubLedgerId field value
func (o *AddLoyaltyPointsEffectProps) GetSubLedgerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubLedgerId
}

// GetSubLedgerIdOk returns a tuple with the SubLedgerId field value
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetSubLedgerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubLedgerId, true
}

// SetSubLedgerId sets field value
func (o *AddLoyaltyPointsEffectProps) SetSubLedgerId(v string) {
	o.SubLedgerId = v
}

// GetValue returns the Value field value
func (o *AddLoyaltyPointsEffectProps) GetValue() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetValueOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *AddLoyaltyPointsEffectProps) SetValue(v float32) {
	o.Value = v
}

// GetDesiredValue returns the DesiredValue field value if set, zero value otherwise.
func (o *AddLoyaltyPointsEffectProps) GetDesiredValue() float32 {
	if o == nil || o.DesiredValue == nil {
		var ret float32
		return ret
	}
	return *o.DesiredValue
}

// GetDesiredValueOk returns a tuple with the DesiredValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetDesiredValueOk() (*float32, bool) {
	if o == nil || o.DesiredValue == nil {
		return nil, false
	}
	return o.DesiredValue, true
}

// HasDesiredValue returns a boolean if a field has been set.
func (o *AddLoyaltyPointsEffectProps) HasDesiredValue() bool {
	if o != nil && o.DesiredValue != nil {
		return true
	}

	return false
}

// SetDesiredValue gets a reference to the given float32 and assigns it to the DesiredValue field.
func (o *AddLoyaltyPointsEffectProps) SetDesiredValue(v float32) {
	o.DesiredValue = &v
}

// GetRecipientIntegrationId returns the RecipientIntegrationId field value
func (o *AddLoyaltyPointsEffectProps) GetRecipientIntegrationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientIntegrationId
}

// GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field value
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetRecipientIntegrationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecipientIntegrationId, true
}

// SetRecipientIntegrationId sets field value
func (o *AddLoyaltyPointsEffectProps) SetRecipientIntegrationId(v string) {
	o.RecipientIntegrationId = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *AddLoyaltyPointsEffectProps) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *AddLoyaltyPointsEffectProps) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *AddLoyaltyPointsEffectProps) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *AddLoyaltyPointsEffectProps) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *AddLoyaltyPointsEffectProps) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *AddLoyaltyPointsEffectProps) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetTransactionUUID returns the TransactionUUID field value
func (o *AddLoyaltyPointsEffectProps) GetTransactionUUID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransactionUUID
}

// GetTransactionUUIDOk returns a tuple with the TransactionUUID field value
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetTransactionUUIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransactionUUID, true
}

// SetTransactionUUID sets field value
func (o *AddLoyaltyPointsEffectProps) SetTransactionUUID(v string) {
	o.TransactionUUID = v
}

// GetCartItemPosition returns the CartItemPosition field value if set, zero value otherwise.
func (o *AddLoyaltyPointsEffectProps) GetCartItemPosition() float32 {
	if o == nil || o.CartItemPosition == nil {
		var ret float32
		return ret
	}
	return *o.CartItemPosition
}

// GetCartItemPositionOk returns a tuple with the CartItemPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetCartItemPositionOk() (*float32, bool) {
	if o == nil || o.CartItemPosition == nil {
		return nil, false
	}
	return o.CartItemPosition, true
}

// HasCartItemPosition returns a boolean if a field has been set.
func (o *AddLoyaltyPointsEffectProps) HasCartItemPosition() bool {
	if o != nil && o.CartItemPosition != nil {
		return true
	}

	return false
}

// SetCartItemPosition gets a reference to the given float32 and assigns it to the CartItemPosition field.
func (o *AddLoyaltyPointsEffectProps) SetCartItemPosition(v float32) {
	o.CartItemPosition = &v
}

// GetCartItemSubPosition returns the CartItemSubPosition field value if set, zero value otherwise.
func (o *AddLoyaltyPointsEffectProps) GetCartItemSubPosition() float32 {
	if o == nil || o.CartItemSubPosition == nil {
		var ret float32
		return ret
	}
	return *o.CartItemSubPosition
}

// GetCartItemSubPositionOk returns a tuple with the CartItemSubPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetCartItemSubPositionOk() (*float32, bool) {
	if o == nil || o.CartItemSubPosition == nil {
		return nil, false
	}
	return o.CartItemSubPosition, true
}

// HasCartItemSubPosition returns a boolean if a field has been set.
func (o *AddLoyaltyPointsEffectProps) HasCartItemSubPosition() bool {
	if o != nil && o.CartItemSubPosition != nil {
		return true
	}

	return false
}

// SetCartItemSubPosition gets a reference to the given float32 and assigns it to the CartItemSubPosition field.
func (o *AddLoyaltyPointsEffectProps) SetCartItemSubPosition(v float32) {
	o.CartItemSubPosition = &v
}

// GetCardIdentifier returns the CardIdentifier field value if set, zero value otherwise.
func (o *AddLoyaltyPointsEffectProps) GetCardIdentifier() string {
	if o == nil || o.CardIdentifier == nil {
		var ret string
		return ret
	}
	return *o.CardIdentifier
}

// GetCardIdentifierOk returns a tuple with the CardIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddLoyaltyPointsEffectProps) GetCardIdentifierOk() (*string, bool) {
	if o == nil || o.CardIdentifier == nil {
		return nil, false
	}
	return o.CardIdentifier, true
}

// HasCardIdentifier returns a boolean if a field has been set.
func (o *AddLoyaltyPointsEffectProps) HasCardIdentifier() bool {
	if o != nil && o.CardIdentifier != nil {
		return true
	}

	return false
}

// SetCardIdentifier gets a reference to the given string and assigns it to the CardIdentifier field.
func (o *AddLoyaltyPointsEffectProps) SetCardIdentifier(v string) {
	o.CardIdentifier = &v
}

func (o AddLoyaltyPointsEffectProps) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["programId"] = o.ProgramId
	}
	if true {
		toSerialize["subLedgerId"] = o.SubLedgerId
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if o.DesiredValue != nil {
		toSerialize["desiredValue"] = o.DesiredValue
	}
	if true {
		toSerialize["recipientIntegrationId"] = o.RecipientIntegrationId
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.ExpiryDate != nil {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if true {
		toSerialize["transactionUUID"] = o.TransactionUUID
	}
	if o.CartItemPosition != nil {
		toSerialize["cartItemPosition"] = o.CartItemPosition
	}
	if o.CartItemSubPosition != nil {
		toSerialize["cartItemSubPosition"] = o.CartItemSubPosition
	}
	if o.CardIdentifier != nil {
		toSerialize["cardIdentifier"] = o.CardIdentifier
	}
	return json.Marshal(toSerialize)
}

type NullableAddLoyaltyPointsEffectProps struct {
	value *AddLoyaltyPointsEffectProps
	isSet bool
}

func (v NullableAddLoyaltyPointsEffectProps) Get() *AddLoyaltyPointsEffectProps {
	return v.value
}

func (v *NullableAddLoyaltyPointsEffectProps) Set(val *AddLoyaltyPointsEffectProps) {
	v.value = val
	v.isSet = true
}

func (v NullableAddLoyaltyPointsEffectProps) IsSet() bool {
	return v.isSet
}

func (v *NullableAddLoyaltyPointsEffectProps) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddLoyaltyPointsEffectProps(val *AddLoyaltyPointsEffectProps) *NullableAddLoyaltyPointsEffectProps {
	return &NullableAddLoyaltyPointsEffectProps{value: val, isSet: true}
}

func (v NullableAddLoyaltyPointsEffectProps) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddLoyaltyPointsEffectProps) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

