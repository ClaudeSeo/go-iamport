package form

type ScheduleForm struct {
	MerchantUID string `json:"merchant_uid" binding:"required"`
	ScheduleAt int32 `json:"schedule_at" binding:"required"`
	Amount float32 `json:"amount" binding:"required"`
	Name string `json:"name,omitempty"`
	BuyerName string `json:"buyer_name,omitempty"`
	BuyerEmail string `json:"buyer_email,omitempty"`
	BuyerTel string `json:"buyer_tel,omitempty"`
	BuyerAddr string `json:"buyer_addr,omitempty"`
	BuyerPostcode string `json:"buyer_postcode,omitempty"`
}


type PaymentSchedule struct {
	CustomerUID string `json:"customer_uid" binding:"required"`
	CheckingAmount float32 `json:"checking_amount,omitempty"`
	CardNumber string `json:"card_number,omitempty"`
	Expiry string `json:"expiry,omitempty"`
	Birth string `json:"birth,omitempty"`
	Pwd2digit string `json:"pwd_2digit,omitempty"`
	Schedules []ScheduleForm `json:"schedules" binidng:"required"`
}

type PaymentUnschedule struct {
	CustomerUID string `json:"customer_uid" binding:"required"`
	MerchantUID string `json:"merchant_uid,omitempty"`
}