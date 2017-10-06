package form

type PaymentOneTimeForm struct {
	MerchantUID string `json:"merchant_uid" binding:"required"`
	Amount float32 `json:"amount" binding:"required"`
	Vat float32 `json:"vat,omitempty"`
	CardNumber string `json:"card_number"`
	Expiry string `json:"expiry" binding:"required"`
	Birth string `json:"birth" binding:"required"`
	Pwd2digit string `json:"pwd_2digit" binding:"required"`
	CustomerUID string `json:"customer_uid,omitempty"`
	Name string `json:"name,omitempty"`
	CardQuota int `json:"card_quota,omitempty"`
	BuyerName string `json:"buyer_name,omitempty"`
	BuyerEmail string `json:"buyer_email,omitempty"`
	BuyerTel string `json:"buyer_tel,omitempty"`
	BuyerAddr string `json:"buyer_addr,omitempty"`
	BuyerPostcode string `json:"buyer_postcode,omitempty"`
	CustomData string `json:"custom_data,omitempty"`
}

type PaymentAgainForm struct {
	CustomerUID string `json:"customer_uid" binding:"required"`
	MerchantUID string `json:"merchant_uid" binding:"required"`
	Amount float32 `json:"amount" binding:"required"`
	Vat float32 `json:"vat,omitempty"`
	Name string `json:"name,omitempty"`
	CardQuota int `json:"card_quota,omitempty"`
	BuyerName string `json:"buyer_name,omitempty"`
	BuyerEmail string `json:"buyer_email,omitempty"`
	BuyerTel string `json:"buyer_tel,omitempty"`
	BuyerAddr string `json:"buyer_addr,omitempty"`
	BuyerPostcode string `json:"buyer_postcode,omitempty"`
	CustomData string `json:"custom_data,omitempty"`
}

type PaymentCancel struct {
	ImpUID string `json:"imp_uid,omitempty"`
	MerchantUID string `json:"merchant_uid,omitempty"`
	Amount float32 `json:"amount,omitempty"`
	TaxFree float32 `json:"tax_free,omitempty"`
	Reason string `json:"reason,omitempty"`
	RefundHolder string `json:"refund_holder,omitempty"`
	RefundBank string `json:"refund_bank,omitempty"`
	RefundAccount string `json:"refund_account,omitempty"`
}