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

// UpdateCoupon 
type UpdateCoupon struct {
	// The number of times the coupon code can be redeemed. `0` means unlimited redemptions but any campaign usage limits will still apply. 
	UsageLimit *int32 `json:"usageLimit,omitempty"`
	// The amount of discounts that can be given with this coupon code. 
	DiscountLimit *float32 `json:"discountLimit,omitempty"`
	// Timestamp at which point the coupon becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured. 
	Limits []LimitConfig `json:"limits,omitempty"`
	// The integration ID for this coupon's beneficiary's profile.
	RecipientIntegrationId *string `json:"recipientIntegrationId,omitempty"`
	// Arbitrary properties associated with this item.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// NewUpdateCoupon instantiates a new UpdateCoupon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCoupon() *UpdateCoupon {
	this := UpdateCoupon{}
	return &this
}

// NewUpdateCouponWithDefaults instantiates a new UpdateCoupon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCouponWithDefaults() *UpdateCoupon {
	this := UpdateCoupon{}
	return &this
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *UpdateCoupon) GetUsageLimit() int32 {
	if o == nil || o.UsageLimit == nil {
		var ret int32
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCoupon) GetUsageLimitOk() (*int32, bool) {
	if o == nil || o.UsageLimit == nil {
		return nil, false
	}
	return o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *UpdateCoupon) HasUsageLimit() bool {
	if o != nil && o.UsageLimit != nil {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.
func (o *UpdateCoupon) SetUsageLimit(v int32) {
	o.UsageLimit = &v
}

// GetDiscountLimit returns the DiscountLimit field value if set, zero value otherwise.
func (o *UpdateCoupon) GetDiscountLimit() float32 {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret
	}
	return *o.DiscountLimit
}

// GetDiscountLimitOk returns a tuple with the DiscountLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCoupon) GetDiscountLimitOk() (*float32, bool) {
	if o == nil || o.DiscountLimit == nil {
		return nil, false
	}
	return o.DiscountLimit, true
}

// HasDiscountLimit returns a boolean if a field has been set.
func (o *UpdateCoupon) HasDiscountLimit() bool {
	if o != nil && o.DiscountLimit != nil {
		return true
	}

	return false
}

// SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.
func (o *UpdateCoupon) SetDiscountLimit(v float32) {
	o.DiscountLimit = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UpdateCoupon) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCoupon) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UpdateCoupon) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *UpdateCoupon) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *UpdateCoupon) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCoupon) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *UpdateCoupon) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *UpdateCoupon) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *UpdateCoupon) GetLimits() []LimitConfig {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret
	}
	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCoupon) GetLimitsOk() ([]LimitConfig, bool) {
	if o == nil || o.Limits == nil {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *UpdateCoupon) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.
func (o *UpdateCoupon) SetLimits(v []LimitConfig) {
	o.Limits = v
}

// GetRecipientIntegrationId returns the RecipientIntegrationId field value if set, zero value otherwise.
func (o *UpdateCoupon) GetRecipientIntegrationId() string {
	if o == nil || o.RecipientIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.RecipientIntegrationId
}

// GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCoupon) GetRecipientIntegrationIdOk() (*string, bool) {
	if o == nil || o.RecipientIntegrationId == nil {
		return nil, false
	}
	return o.RecipientIntegrationId, true
}

// HasRecipientIntegrationId returns a boolean if a field has been set.
func (o *UpdateCoupon) HasRecipientIntegrationId() bool {
	if o != nil && o.RecipientIntegrationId != nil {
		return true
	}

	return false
}

// SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.
func (o *UpdateCoupon) SetRecipientIntegrationId(v string) {
	o.RecipientIntegrationId = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateCoupon) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCoupon) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateCoupon) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *UpdateCoupon) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

func (o UpdateCoupon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UsageLimit != nil {
		toSerialize["usageLimit"] = o.UsageLimit
	}
	if o.DiscountLimit != nil {
		toSerialize["discountLimit"] = o.DiscountLimit
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.ExpiryDate != nil {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if o.Limits != nil {
		toSerialize["limits"] = o.Limits
	}
	if o.RecipientIntegrationId != nil {
		toSerialize["recipientIntegrationId"] = o.RecipientIntegrationId
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateCoupon struct {
	value *UpdateCoupon
	isSet bool
}

func (v NullableUpdateCoupon) Get() *UpdateCoupon {
	return v.value
}

func (v *NullableUpdateCoupon) Set(val *UpdateCoupon) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCoupon) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCoupon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCoupon(val *UpdateCoupon) *NullableUpdateCoupon {
	return &NullableUpdateCoupon{value: val, isSet: true}
}

func (v NullableUpdateCoupon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCoupon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


