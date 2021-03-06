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

// Coupon 
type Coupon struct {
	// Unique ID for this entity. Not to be confused with the Integration ID, which is set by your integration layer and used in most endpoints.
	Id int32 `json:"id"`
	// The exact moment this entity was created.
	Created time.Time `json:"created"`
	// The ID of the campaign that owns this entity.
	CampaignId int32 `json:"campaignId"`
	// The coupon code.
	Value string `json:"value"`
	// The number of times the coupon code can be redeemed. `0` means unlimited redemptions but any campaign usage limits will still apply. 
	UsageLimit int32 `json:"usageLimit"`
	// The amount of discounts that can be given with this coupon code. 
	DiscountLimit *float32 `json:"discountLimit,omitempty"`
	// Timestamp at which point the coupon becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiry date of the coupon. Coupon never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// Limits configuration for a coupon. These limits will override the limits set from the campaign.  **Note:** Only usable when creating a single coupon which is not tied to a specific recipient. Only per-profile limits are allowed to be configured. 
	Limits []LimitConfig `json:"limits,omitempty"`
	// The number of times this coupon has been successfully used.
	UsageCounter int32 `json:"usageCounter"`
	// The amount of discounts given on rules redeeming this coupon. Only usable if a coupon discount budget was set for this coupon.
	DiscountCounter *float32 `json:"discountCounter,omitempty"`
	// The remaining discount this coupon can give.
	DiscountRemainder *float32 `json:"discountRemainder,omitempty"`
	// Custom attributes associated with this coupon.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The integration ID of the referring customer (if any) for whom this coupon was created as an effect.
	ReferralId *int32 `json:"referralId,omitempty"`
	// The Integration ID of the customer that is allowed to redeem this coupon.
	RecipientIntegrationId *string `json:"recipientIntegrationId,omitempty"`
	// The ID of the Import which created this coupon.
	ImportId *int32 `json:"importId,omitempty"`
	// Defines the type of reservation: - `true`: The reservation is a soft reservation. Any customer can use the coupon. This is done via the [Create coupon reservation endpoint](/integration-api/#operation/createCouponReservation). - `false`: The reservation is a hard reservation. Only the associated customer (`recipientIntegrationId`) can use the coupon. This is done via the Campaign Manager when you create a coupon for a given `recipientIntegrationId`, the [Create coupons endpoint](/management-api/#operation/createCoupons) or [Create coupons for multiple recipients endpoint](/management-api/#operation/createCouponsForMultipleRecipients). 
	Reservation *bool `json:"reservation,omitempty"`
	// The id of the batch the coupon belongs to.
	BatchId *string `json:"batchId,omitempty"`
}

// NewCoupon instantiates a new Coupon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoupon(id int32, created time.Time, campaignId int32, value string, usageLimit int32, usageCounter int32) *Coupon {
	this := Coupon{}
	this.Id = id
	this.Created = created
	this.CampaignId = campaignId
	this.Value = value
	this.UsageLimit = usageLimit
	this.UsageCounter = usageCounter
	var reservation bool = true
	this.Reservation = &reservation
	return &this
}

// NewCouponWithDefaults instantiates a new Coupon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCouponWithDefaults() *Coupon {
	this := Coupon{}
	var reservation bool = true
	this.Reservation = &reservation
	return &this
}

// GetId returns the Id field value
func (o *Coupon) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Coupon) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Coupon) SetId(v int32) {
	o.Id = v
}

// GetCreated returns the Created field value
func (o *Coupon) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Coupon) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Coupon) SetCreated(v time.Time) {
	o.Created = v
}

// GetCampaignId returns the CampaignId field value
func (o *Coupon) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *Coupon) GetCampaignIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *Coupon) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetValue returns the Value field value
func (o *Coupon) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Coupon) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Coupon) SetValue(v string) {
	o.Value = v
}

// GetUsageLimit returns the UsageLimit field value
func (o *Coupon) GetUsageLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value
// and a boolean to check if the value has been set.
func (o *Coupon) GetUsageLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageLimit, true
}

// SetUsageLimit sets field value
func (o *Coupon) SetUsageLimit(v int32) {
	o.UsageLimit = v
}

// GetDiscountLimit returns the DiscountLimit field value if set, zero value otherwise.
func (o *Coupon) GetDiscountLimit() float32 {
	if o == nil || o.DiscountLimit == nil {
		var ret float32
		return ret
	}
	return *o.DiscountLimit
}

// GetDiscountLimitOk returns a tuple with the DiscountLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetDiscountLimitOk() (*float32, bool) {
	if o == nil || o.DiscountLimit == nil {
		return nil, false
	}
	return o.DiscountLimit, true
}

// HasDiscountLimit returns a boolean if a field has been set.
func (o *Coupon) HasDiscountLimit() bool {
	if o != nil && o.DiscountLimit != nil {
		return true
	}

	return false
}

// SetDiscountLimit gets a reference to the given float32 and assigns it to the DiscountLimit field.
func (o *Coupon) SetDiscountLimit(v float32) {
	o.DiscountLimit = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Coupon) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *Coupon) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Coupon) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *Coupon) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *Coupon) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *Coupon) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *Coupon) GetLimits() []LimitConfig {
	if o == nil || o.Limits == nil {
		var ret []LimitConfig
		return ret
	}
	return o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetLimitsOk() ([]LimitConfig, bool) {
	if o == nil || o.Limits == nil {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *Coupon) HasLimits() bool {
	if o != nil && o.Limits != nil {
		return true
	}

	return false
}

// SetLimits gets a reference to the given []LimitConfig and assigns it to the Limits field.
func (o *Coupon) SetLimits(v []LimitConfig) {
	o.Limits = v
}

// GetUsageCounter returns the UsageCounter field value
func (o *Coupon) GetUsageCounter() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.UsageCounter
}

// GetUsageCounterOk returns a tuple with the UsageCounter field value
// and a boolean to check if the value has been set.
func (o *Coupon) GetUsageCounterOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsageCounter, true
}

// SetUsageCounter sets field value
func (o *Coupon) SetUsageCounter(v int32) {
	o.UsageCounter = v
}

// GetDiscountCounter returns the DiscountCounter field value if set, zero value otherwise.
func (o *Coupon) GetDiscountCounter() float32 {
	if o == nil || o.DiscountCounter == nil {
		var ret float32
		return ret
	}
	return *o.DiscountCounter
}

// GetDiscountCounterOk returns a tuple with the DiscountCounter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetDiscountCounterOk() (*float32, bool) {
	if o == nil || o.DiscountCounter == nil {
		return nil, false
	}
	return o.DiscountCounter, true
}

// HasDiscountCounter returns a boolean if a field has been set.
func (o *Coupon) HasDiscountCounter() bool {
	if o != nil && o.DiscountCounter != nil {
		return true
	}

	return false
}

// SetDiscountCounter gets a reference to the given float32 and assigns it to the DiscountCounter field.
func (o *Coupon) SetDiscountCounter(v float32) {
	o.DiscountCounter = &v
}

// GetDiscountRemainder returns the DiscountRemainder field value if set, zero value otherwise.
func (o *Coupon) GetDiscountRemainder() float32 {
	if o == nil || o.DiscountRemainder == nil {
		var ret float32
		return ret
	}
	return *o.DiscountRemainder
}

// GetDiscountRemainderOk returns a tuple with the DiscountRemainder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetDiscountRemainderOk() (*float32, bool) {
	if o == nil || o.DiscountRemainder == nil {
		return nil, false
	}
	return o.DiscountRemainder, true
}

// HasDiscountRemainder returns a boolean if a field has been set.
func (o *Coupon) HasDiscountRemainder() bool {
	if o != nil && o.DiscountRemainder != nil {
		return true
	}

	return false
}

// SetDiscountRemainder gets a reference to the given float32 and assigns it to the DiscountRemainder field.
func (o *Coupon) SetDiscountRemainder(v float32) {
	o.DiscountRemainder = &v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *Coupon) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *Coupon) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *Coupon) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetReferralId returns the ReferralId field value if set, zero value otherwise.
func (o *Coupon) GetReferralId() int32 {
	if o == nil || o.ReferralId == nil {
		var ret int32
		return ret
	}
	return *o.ReferralId
}

// GetReferralIdOk returns a tuple with the ReferralId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetReferralIdOk() (*int32, bool) {
	if o == nil || o.ReferralId == nil {
		return nil, false
	}
	return o.ReferralId, true
}

// HasReferralId returns a boolean if a field has been set.
func (o *Coupon) HasReferralId() bool {
	if o != nil && o.ReferralId != nil {
		return true
	}

	return false
}

// SetReferralId gets a reference to the given int32 and assigns it to the ReferralId field.
func (o *Coupon) SetReferralId(v int32) {
	o.ReferralId = &v
}

// GetRecipientIntegrationId returns the RecipientIntegrationId field value if set, zero value otherwise.
func (o *Coupon) GetRecipientIntegrationId() string {
	if o == nil || o.RecipientIntegrationId == nil {
		var ret string
		return ret
	}
	return *o.RecipientIntegrationId
}

// GetRecipientIntegrationIdOk returns a tuple with the RecipientIntegrationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetRecipientIntegrationIdOk() (*string, bool) {
	if o == nil || o.RecipientIntegrationId == nil {
		return nil, false
	}
	return o.RecipientIntegrationId, true
}

// HasRecipientIntegrationId returns a boolean if a field has been set.
func (o *Coupon) HasRecipientIntegrationId() bool {
	if o != nil && o.RecipientIntegrationId != nil {
		return true
	}

	return false
}

// SetRecipientIntegrationId gets a reference to the given string and assigns it to the RecipientIntegrationId field.
func (o *Coupon) SetRecipientIntegrationId(v string) {
	o.RecipientIntegrationId = &v
}

// GetImportId returns the ImportId field value if set, zero value otherwise.
func (o *Coupon) GetImportId() int32 {
	if o == nil || o.ImportId == nil {
		var ret int32
		return ret
	}
	return *o.ImportId
}

// GetImportIdOk returns a tuple with the ImportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetImportIdOk() (*int32, bool) {
	if o == nil || o.ImportId == nil {
		return nil, false
	}
	return o.ImportId, true
}

// HasImportId returns a boolean if a field has been set.
func (o *Coupon) HasImportId() bool {
	if o != nil && o.ImportId != nil {
		return true
	}

	return false
}

// SetImportId gets a reference to the given int32 and assigns it to the ImportId field.
func (o *Coupon) SetImportId(v int32) {
	o.ImportId = &v
}

// GetReservation returns the Reservation field value if set, zero value otherwise.
func (o *Coupon) GetReservation() bool {
	if o == nil || o.Reservation == nil {
		var ret bool
		return ret
	}
	return *o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetReservationOk() (*bool, bool) {
	if o == nil || o.Reservation == nil {
		return nil, false
	}
	return o.Reservation, true
}

// HasReservation returns a boolean if a field has been set.
func (o *Coupon) HasReservation() bool {
	if o != nil && o.Reservation != nil {
		return true
	}

	return false
}

// SetReservation gets a reference to the given bool and assigns it to the Reservation field.
func (o *Coupon) SetReservation(v bool) {
	o.Reservation = &v
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *Coupon) GetBatchId() string {
	if o == nil || o.BatchId == nil {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Coupon) GetBatchIdOk() (*string, bool) {
	if o == nil || o.BatchId == nil {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *Coupon) HasBatchId() bool {
	if o != nil && o.BatchId != nil {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *Coupon) SetBatchId(v string) {
	o.BatchId = &v
}

func (o Coupon) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["campaignId"] = o.CampaignId
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
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
	if true {
		toSerialize["usageCounter"] = o.UsageCounter
	}
	if o.DiscountCounter != nil {
		toSerialize["discountCounter"] = o.DiscountCounter
	}
	if o.DiscountRemainder != nil {
		toSerialize["discountRemainder"] = o.DiscountRemainder
	}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.ReferralId != nil {
		toSerialize["referralId"] = o.ReferralId
	}
	if o.RecipientIntegrationId != nil {
		toSerialize["recipientIntegrationId"] = o.RecipientIntegrationId
	}
	if o.ImportId != nil {
		toSerialize["importId"] = o.ImportId
	}
	if o.Reservation != nil {
		toSerialize["reservation"] = o.Reservation
	}
	if o.BatchId != nil {
		toSerialize["batchId"] = o.BatchId
	}
	return json.Marshal(toSerialize)
}

type NullableCoupon struct {
	value *Coupon
	isSet bool
}

func (v NullableCoupon) Get() *Coupon {
	return v.value
}

func (v *NullableCoupon) Set(val *Coupon) {
	v.value = val
	v.isSet = true
}

func (v NullableCoupon) IsSet() bool {
	return v.isSet
}

func (v *NullableCoupon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoupon(val *Coupon) *NullableCoupon {
	return &NullableCoupon{value: val, isSet: true}
}

func (v NullableCoupon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoupon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


