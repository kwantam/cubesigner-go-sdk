package v0

import (
	"encoding/json"
	"github.com/lombard-finance/cubesigner-sdk/api"
)

// BabylonStakingWithdrawalAllOf struct for BabylonStakingWithdrawalAllOf
type BabylonStakingWithdrawalAllOf struct {
	// The transaction fee value. The `fee_type` field determines whether this is a fixed fee in sats or a rate in sats per (estimated) virtual byte of transaction weight (i.e., sats per vb).
	Fee int64 `json:"fee"`
	FeeType api.FeeType `json:"fee_type"`
	// The withdrawal recipient, specified as a Bitcoin spend script
	Recipient string `json:"recipient"`
	// Transaction-id of the staking transaction to withdraw from
	Txid string `json:"txid"`
	// An optional lock height (in blocks) for this transaction. The resulting transaction cannot be mined before the specified block height.
	TxnLockHeight api.NullableInt32 `json:"txn_lock_height,omitempty"`
	// The value in sats that is staked in the transaction being withdrawn
	Value int64 `json:"value"`
	// Transaction output index of the staking transaction from which to withdraw. For staking transactions generated by CubeSigner, this will always be zero.
	Vout int32 `json:"vout"`
}

// NewBabylonStakingWithdrawalAllOf instantiates a new BabylonStakingWithdrawalAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBabylonStakingWithdrawalAllOf(fee int64, feeType api.FeeType, recipient string, txid string, value int64, vout int32) *BabylonStakingWithdrawalAllOf {
	this := BabylonStakingWithdrawalAllOf{}
	this.Fee = fee
	this.FeeType = feeType
	this.Recipient = recipient
	this.Txid = txid
	this.Value = value
	this.Vout = vout
	return &this
}

// NewBabylonStakingWithdrawalAllOfWithDefaults instantiates a new BabylonStakingWithdrawalAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBabylonStakingWithdrawalAllOfWithDefaults() *BabylonStakingWithdrawalAllOf {
	this := BabylonStakingWithdrawalAllOf{}
	return &this
}

// GetFee returns the Fee field value
func (o *BabylonStakingWithdrawalAllOf) GetFee() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Fee
}

// GetFeeOk returns a tuple with the Fee field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingWithdrawalAllOf) GetFeeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fee, true
}

// SetFee sets field value
func (o *BabylonStakingWithdrawalAllOf) SetFee(v int64) {
	o.Fee = v
}

// GetFeeType returns the FeeType field value
func (o *BabylonStakingWithdrawalAllOf) GetFeeType() api.FeeType {
	if o == nil {
		var ret api.FeeType
		return ret
	}

	return o.FeeType
}

// GetFeeTypeOk returns a tuple with the FeeType field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingWithdrawalAllOf) GetFeeTypeOk() (*api.FeeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FeeType, true
}

// SetFeeType sets field value
func (o *BabylonStakingWithdrawalAllOf) SetFeeType(v api.FeeType) {
	o.FeeType = v
}

// GetRecipient returns the Recipient field value
func (o *BabylonStakingWithdrawalAllOf) GetRecipient() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Recipient
}

// GetRecipientOk returns a tuple with the Recipient field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingWithdrawalAllOf) GetRecipientOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Recipient, true
}

// SetRecipient sets field value
func (o *BabylonStakingWithdrawalAllOf) SetRecipient(v string) {
	o.Recipient = v
}

// GetTxid returns the Txid field value
func (o *BabylonStakingWithdrawalAllOf) GetTxid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Txid
}

// GetTxidOk returns a tuple with the Txid field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingWithdrawalAllOf) GetTxidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Txid, true
}

// SetTxid sets field value
func (o *BabylonStakingWithdrawalAllOf) SetTxid(v string) {
	o.Txid = v
}

// GetTxnLockHeight returns the TxnLockHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BabylonStakingWithdrawalAllOf) GetTxnLockHeight() int32 {
	if o == nil || o.TxnLockHeight.Get() == nil {
		var ret int32
		return ret
	}
	return *o.TxnLockHeight.Get()
}

// GetTxnLockHeightOk returns a tuple with the TxnLockHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BabylonStakingWithdrawalAllOf) GetTxnLockHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TxnLockHeight.Get(), o.TxnLockHeight.IsSet()
}

// HasTxnLockHeight returns a boolean if a field has been set.
func (o *BabylonStakingWithdrawalAllOf) HasTxnLockHeight() bool {
	if o != nil && o.TxnLockHeight.IsSet() {
		return true
	}

	return false
}

// SetTxnLockHeight gets a reference to the given NullableInt32 and assigns it to the TxnLockHeight field.
func (o *BabylonStakingWithdrawalAllOf) SetTxnLockHeight(v int32) {
	o.TxnLockHeight.Set(&v)
}
// SetTxnLockHeightNil sets the value for TxnLockHeight to be an explicit nil
func (o *BabylonStakingWithdrawalAllOf) SetTxnLockHeightNil() {
	o.TxnLockHeight.Set(nil)
}

// UnsetTxnLockHeight ensures that no value is present for TxnLockHeight, not even an explicit nil
func (o *BabylonStakingWithdrawalAllOf) UnsetTxnLockHeight() {
	o.TxnLockHeight.Unset()
}

// GetValue returns the Value field value
func (o *BabylonStakingWithdrawalAllOf) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingWithdrawalAllOf) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *BabylonStakingWithdrawalAllOf) SetValue(v int64) {
	o.Value = v
}

// GetVout returns the Vout field value
func (o *BabylonStakingWithdrawalAllOf) GetVout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vout
}

// GetVoutOk returns a tuple with the Vout field value
// and a boolean to check if the value has been set.
func (o *BabylonStakingWithdrawalAllOf) GetVoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vout, true
}

// SetVout sets field value
func (o *BabylonStakingWithdrawalAllOf) SetVout(v int32) {
	o.Vout = v
}

func (o BabylonStakingWithdrawalAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fee"] = o.Fee
	}
	if true {
		toSerialize["fee_type"] = o.FeeType
	}
	if true {
		toSerialize["recipient"] = o.Recipient
	}
	if true {
		toSerialize["txid"] = o.Txid
	}
	if o.TxnLockHeight.IsSet() {
		toSerialize["txn_lock_height"] = o.TxnLockHeight.Get()
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["vout"] = o.Vout
	}
	return json.Marshal(toSerialize)
}

type NullableBabylonStakingWithdrawalAllOf struct {
	value *BabylonStakingWithdrawalAllOf
	isSet bool
}

func (v NullableBabylonStakingWithdrawalAllOf) Get() *BabylonStakingWithdrawalAllOf {
	return v.value
}

func (v *NullableBabylonStakingWithdrawalAllOf) Set(val *BabylonStakingWithdrawalAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBabylonStakingWithdrawalAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBabylonStakingWithdrawalAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBabylonStakingWithdrawalAllOf(val *BabylonStakingWithdrawalAllOf) *NullableBabylonStakingWithdrawalAllOf {
	return &NullableBabylonStakingWithdrawalAllOf{value: val, isSet: true}
}

func (v NullableBabylonStakingWithdrawalAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBabylonStakingWithdrawalAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

