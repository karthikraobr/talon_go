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

// UpdateReferralBatch struct for UpdateReferralBatch
type UpdateReferralBatch struct {
	// Arbitrary properties associated with this item.
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	// The id of the batch the referral belongs to.
	BatchID string `json:"batchID"`
	// Timestamp at which point the referral code becomes valid.
	StartDate *time.Time `json:"startDate,omitempty"`
	// Expiry date of the referral code. Referral never expires if this is omitted, zero, or negative.
	ExpiryDate *time.Time `json:"expiryDate,omitempty"`
	// The number of times a referral code can be used. This can be set to 0 for no limit, but any campaign usage limits will still apply. 
	UsageLimit *int32 `json:"usageLimit,omitempty"`
}

// NewUpdateReferralBatch instantiates a new UpdateReferralBatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateReferralBatch(batchID string) *UpdateReferralBatch {
	this := UpdateReferralBatch{}
	this.BatchID = batchID
	return &this
}

// NewUpdateReferralBatchWithDefaults instantiates a new UpdateReferralBatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateReferralBatchWithDefaults() *UpdateReferralBatch {
	this := UpdateReferralBatch{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UpdateReferralBatch) GetAttributes() map[string]interface{} {
	if o == nil || o.Attributes == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReferralBatch) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UpdateReferralBatch) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given map[string]interface{} and assigns it to the Attributes field.
func (o *UpdateReferralBatch) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetBatchID returns the BatchID field value
func (o *UpdateReferralBatch) GetBatchID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BatchID
}

// GetBatchIDOk returns a tuple with the BatchID field value
// and a boolean to check if the value has been set.
func (o *UpdateReferralBatch) GetBatchIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BatchID, true
}

// SetBatchID sets field value
func (o *UpdateReferralBatch) SetBatchID(v string) {
	o.BatchID = v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *UpdateReferralBatch) GetStartDate() time.Time {
	if o == nil || o.StartDate == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReferralBatch) GetStartDateOk() (*time.Time, bool) {
	if o == nil || o.StartDate == nil {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *UpdateReferralBatch) HasStartDate() bool {
	if o != nil && o.StartDate != nil {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *UpdateReferralBatch) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetExpiryDate returns the ExpiryDate field value if set, zero value otherwise.
func (o *UpdateReferralBatch) GetExpiryDate() time.Time {
	if o == nil || o.ExpiryDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiryDate
}

// GetExpiryDateOk returns a tuple with the ExpiryDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReferralBatch) GetExpiryDateOk() (*time.Time, bool) {
	if o == nil || o.ExpiryDate == nil {
		return nil, false
	}
	return o.ExpiryDate, true
}

// HasExpiryDate returns a boolean if a field has been set.
func (o *UpdateReferralBatch) HasExpiryDate() bool {
	if o != nil && o.ExpiryDate != nil {
		return true
	}

	return false
}

// SetExpiryDate gets a reference to the given time.Time and assigns it to the ExpiryDate field.
func (o *UpdateReferralBatch) SetExpiryDate(v time.Time) {
	o.ExpiryDate = &v
}

// GetUsageLimit returns the UsageLimit field value if set, zero value otherwise.
func (o *UpdateReferralBatch) GetUsageLimit() int32 {
	if o == nil || o.UsageLimit == nil {
		var ret int32
		return ret
	}
	return *o.UsageLimit
}

// GetUsageLimitOk returns a tuple with the UsageLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateReferralBatch) GetUsageLimitOk() (*int32, bool) {
	if o == nil || o.UsageLimit == nil {
		return nil, false
	}
	return o.UsageLimit, true
}

// HasUsageLimit returns a boolean if a field has been set.
func (o *UpdateReferralBatch) HasUsageLimit() bool {
	if o != nil && o.UsageLimit != nil {
		return true
	}

	return false
}

// SetUsageLimit gets a reference to the given int32 and assigns it to the UsageLimit field.
func (o *UpdateReferralBatch) SetUsageLimit(v int32) {
	o.UsageLimit = &v
}

func (o UpdateReferralBatch) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["batchID"] = o.BatchID
	}
	if o.StartDate != nil {
		toSerialize["startDate"] = o.StartDate
	}
	if o.ExpiryDate != nil {
		toSerialize["expiryDate"] = o.ExpiryDate
	}
	if o.UsageLimit != nil {
		toSerialize["usageLimit"] = o.UsageLimit
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateReferralBatch struct {
	value *UpdateReferralBatch
	isSet bool
}

func (v NullableUpdateReferralBatch) Get() *UpdateReferralBatch {
	return v.value
}

func (v *NullableUpdateReferralBatch) Set(val *UpdateReferralBatch) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateReferralBatch) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateReferralBatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateReferralBatch(val *UpdateReferralBatch) *NullableUpdateReferralBatch {
	return &NullableUpdateReferralBatch{value: val, isSet: true}
}

func (v NullableUpdateReferralBatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateReferralBatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

