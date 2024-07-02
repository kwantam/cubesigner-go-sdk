package v0

// MfaRequestInfo struct returned for GetMfaRequest and ApproveMfaRequest
type MfaRequestInfo struct {
	// DateTime measured in seconds since unix epoch. A wrapper type for serialization that encodes a `SystemTime` as a `u64` representing the number of seconds since `SystemTime::UNIX_EPOCH`.
	ExpiresAt *int64 `json:"timestamp,omitempty"`
	// Approval request ID.
	Id *string `json:"id,omitempty"`
	// Receipt that an MFA request was approved.
	Receipt *MfaReceipt `json:"receipt,omitempty"`
	// Request
	Request *MfaHttpRequest `json:"request,omitempty"`
	// Status
	Status *MfaStatus `json:"status,omitempty"`
}

// NewMfaRequestInfo instantiates a new MfaRequestInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaRequestInfo(id *string, receipt *MfaReceipt, request *MfaHttpRequest, status *MfaStatus) *MfaRequestInfo {
	this := MfaRequestInfo{}
	this.Id = id
	this.Receipt = receipt
	this.Request = request
	this.Status = status
	return &this
}

// NewMfaRequestInfoWithDefaults instantiates a new MfaRequestInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaRequestInfoWithDefaults() *MfaRequestInfo {
	this := MfaRequestInfo{}
	return &this
}

// GetExpiresAt returns the ExpiresAt field value
func (o *MfaRequestInfo) GetExpiresAt() *int64 {
	if o == nil {
		var ret *int64
		return ret
	}

	return o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value
// and a boolean to check if the value has been set.
func (o *MfaRequestInfo) GetExpiresAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.ExpiresAt, true
}

// SetExpiresAt sets the ExpiresAt field value
func (o *MfaRequestInfo) SetExpiresAt(v *int64) {
	o.ExpiresAt = v
}

// GetId returns the Id field value
func (o *MfaRequestInfo) GetId() *string {
	if o == nil {
		var ret *string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MfaRequestInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Id, true
}

// SetId sets the Id field value
func (o *MfaRequestInfo) SetId(v *string) {
	o.Id = v
}

// GetReceipt returns the Receipt field value if set, zero value otherwise
func (o *MfaRequestInfo) GetReceipt() *MfaReceipt {
	if o == nil || o.Receipt == nil {
		return nil
	}
	return o.Receipt
}

// GetReceiptOk returns a tuple with the Receipt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaRequestInfo) GetReceiptOk() (*MfaReceipt, bool) {
	if o == nil || o.Receipt == nil {
		return nil, false
	}
	return o.Receipt, true
}

// SetReceipt sets the Receipt field value
func (o *MfaRequestInfo) SetReceipt(v *MfaReceipt) {
	o.Receipt = v
}

// GetRequest returns the Request field value
func (o *MfaRequestInfo) GetRequest() *MfaHttpRequest {
	if o == nil {
		var ret *MfaHttpRequest
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *MfaRequestInfo) GetRequestOk() (*MfaHttpRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Request, true
}

// SetRequest sets the Request field value
func (o *MfaRequestInfo) SetRequest(v *MfaHttpRequest) {
	o.Request = v
}

// GetStatus returns the Status field value
func (o *MfaRequestInfo) GetStatus() *MfaStatus {
	if o == nil {
		var ret *MfaStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *MfaRequestInfo) GetStatusOk() (*MfaStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status, true
}

// SetStatus sets the Status field value
func (o *MfaRequestInfo) SetStatus(v *MfaStatus) {
	o.Status = v
}
