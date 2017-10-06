package iamport

import (
	"github.com/claudeseo/go-iamport/form"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var im *Iamport

func TestPaymentOneTime(t *testing.T) {
	payload := form.PaymentOneTimeForm{
		MerchantUID: "1234",
		Amount:      1,
		CardNumber:  "4092-0230-1234-1234",
		Expiry:      "2019-03",
		Birth:       "500203",
		Pwd2digit:   "19",
	}
	result, err := im.PayOneTime(payload)
	assert.Nil(t, result)
	assert.Containsf(t, err.Error(), "이미 주문이 이루어진 건입니다", "Error Message contains `이미 주문이 이루어진 건입니다`")
}

func TestPaymentAgain(t *testing.T) {
	payload := form.PaymentAgainForm{
		CustomerUID: "stkn0121",
		MerchantUID: "0000000101010001",
		Amount:      5000,
	}
	result, err := im.PayAgain(payload)
	assert.Nil(t, result)
	assert.Containsf(t, err.Error(), "등록되지 않은 구매자입니다", "Error Message contains `등록되지 않은 구매자입니다`")
}

func TestPaySchedule(t *testing.T) {
	payload := form.PaymentSchedule{
		CustomerUID: "stkn0121",
		Schedules: []form.ScheduleForm{
			form.ScheduleForm{
				MerchantUID: "0000000101010001",
				Amount:      5000,
				ScheduleAt:  int32(time.Now().Unix()) + int32(1000),
			},
		},
	}
	result, err := im.PaySchedule(payload)
	assert.Nil(t, result)
	assert.Containsf(t, err.Error(), "등록된 고객정보가 없습니다", "Error Message contains `등록된 고객정보가 없습니다`")
}

func TestPayUnschedule(t *testing.T) {
	payload := form.PaymentUnschedule{
		CustomerUID: "stkn0121",
		MerchantUID: "0000000101010001",
	}
	result, err := im.PayUnschedule(payload)
	assert.Nil(t, result)
	assert.Containsf(t, err.Error(), "취소할 예약결제 기록이 존재하지 않습니다", "Error Message contains `취소할 예약결제 기록이 존재하지 않습니다`")
}

func init() {
	im = New("imp_apikey", "ekKoeW8RyKuT0zgaZsUtXXTLQ4AhPFW3ZGseDA6bkA5lamv9OqDMnxyeB9wqOsuO9W3Mx9YSJ4dTqJ3f")
}
