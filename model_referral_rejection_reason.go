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

// ReferralRejectionReason Holds a reference to the campaign, the referral and the reason for which that referral was rejected. Should only be present when there is a 'rejectReferral' effect.
type ReferralRejectionReason struct {
	CampaignId int32 `json:"campaignId"`
	ReferralId int32 `json:"referralId"`
	Reason string `json:"reason"`
}

// NewReferralRejectionReason instantiates a new ReferralRejectionReason object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferralRejectionReason(campaignId int32, referralId int32, reason string) *ReferralRejectionReason {
	this := ReferralRejectionReason{}
	this.CampaignId = campaignId
	this.ReferralId = referralId
	this.Reason = reason
	return &this
}

// NewReferralRejectionReasonWithDefaults instantiates a new ReferralRejectionReason object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferralRejectionReasonWithDefaults() *ReferralRejectionReason {
	this := ReferralRejectionReason{}
	return &this
}

// GetCampaignId returns the CampaignId field value
func (o *ReferralRejectionReason) GetCampaignId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CampaignId
}

// GetCampaignIdOk returns a tuple with the CampaignId field value
// and a boolean to check if the value has been set.
func (o *ReferralRejectionReason) GetCampaignIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CampaignId, true
}

// SetCampaignId sets field value
func (o *ReferralRejectionReason) SetCampaignId(v int32) {
	o.CampaignId = v
}

// GetReferralId returns the ReferralId field value
func (o *ReferralRejectionReason) GetReferralId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReferralId
}

// GetReferralIdOk returns a tuple with the ReferralId field value
// and a boolean to check if the value has been set.
func (o *ReferralRejectionReason) GetReferralIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferralId, true
}

// SetReferralId sets field value
func (o *ReferralRejectionReason) SetReferralId(v int32) {
	o.ReferralId = v
}

// GetReason returns the Reason field value
func (o *ReferralRejectionReason) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ReferralRejectionReason) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ReferralRejectionReason) SetReason(v string) {
	o.Reason = v
}

func (o ReferralRejectionReason) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["campaignId"] = o.CampaignId
	}
	if true {
		toSerialize["referralId"] = o.ReferralId
	}
	if true {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableReferralRejectionReason struct {
	value *ReferralRejectionReason
	isSet bool
}

func (v NullableReferralRejectionReason) Get() *ReferralRejectionReason {
	return v.value
}

func (v *NullableReferralRejectionReason) Set(val *ReferralRejectionReason) {
	v.value = val
	v.isSet = true
}

func (v NullableReferralRejectionReason) IsSet() bool {
	return v.isSet
}

func (v *NullableReferralRejectionReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferralRejectionReason(val *ReferralRejectionReason) *NullableReferralRejectionReason {
	return &NullableReferralRejectionReason{value: val, isSet: true}
}

func (v NullableReferralRejectionReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferralRejectionReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


