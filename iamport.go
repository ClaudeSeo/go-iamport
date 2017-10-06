package iamport

import (
	"encoding/json"
	"fmt"
	"github.com/ClaudeSeo/go-iamport/form"
	"github.com/parnurzeal/gorequest"
	"reflect"
)

type Iamport struct {
	ImpKey    string
	ImpSecret string
	ImpURL    string
}

type iamportResponse struct {
	Code     int
	Message  string
	Response map[string]interface{}
}

type responseError struct {
	Code    int
	Message string
}

func (err *responseError) Error() string {
	return fmt.Sprintf("%d - %s", err.Code, err.Message)
}

func getResponse(resp gorequest.Response, body string, errs []error) (map[string]interface{}, error) {
	if len(errs) > 0 {
		return nil, errs[0]
	}
	var result iamportResponse
	err := json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 || result.Code != 0 {
		return nil, &responseError{result.Code, result.Message}
	}
	return result.Response, nil
}

func (im Iamport) get(url string, payload interface{}) (map[string]interface{}, error) {
	token, err := im.getToken()
	if err != nil {
		return nil, err
	}
	request := gorequest.New()
	if v := reflect.ValueOf(payload); v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			val := v.MapIndex(k)
			request = request.Param(k.Interface().(string), val.Interface().(string))
		}
	}
	resp, body, errs := request.Get(url).AppendHeader("X-ImpTokenHeader", token).End()
	result, err := getResponse(resp, body, errs)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (im Iamport) post(url string, payload interface{}) (map[string]interface{}, error) {
	token, err := im.getToken()
	if err != nil {
		return nil, err
	}
	serialize, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request := gorequest.New()
	resp, body, errs := request.Post(url).AppendHeader("X-ImpTokenHeader", token).Send(string(serialize)).End()
	result, err := getResponse(resp, body, errs)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (im Iamport) getToken() (string, error) {
	url := fmt.Sprintf("%s/users/getToken", im.ImpURL)
	request := gorequest.New()
	payload := map[string]interface{}{
		"imp_key":    im.ImpKey,
		"imp_secret": im.ImpSecret,
	}
	serialize, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	resp, body, errs := request.Post(url).Send(string(serialize)).End()
	result, err := getResponse(resp, body, errs)
	if err != nil {
		return "", err
	}
	return result["access_token"].(string), nil
}

func (im Iamport) FindByMerchantUID(merchantUID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/payments/find/%s", im.ImpURL, merchantUID)
	return im.get(url, nil)
}

func (im Iamport) FindByImpUID(impUID string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/payments/%s", im.ImpURL, impUID)
	return im.get(url, nil)
}

func (im Iamport) PayOneTime(f form.PaymentOneTimeForm) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/subscribe/payments/onetime", im.ImpURL)
	valid, err := form.Validate(f)
	if !valid || err != nil {
		return nil, err
	}
	return im.post(url, f)
}

func (im Iamport) PayAgain(f form.PaymentAgainForm) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/subscribe/payments/again", im.ImpURL)
	valid, err := form.Validate(f)
	if !valid || err != nil {
		return nil, err
	}
	return im.post(url, f)
}

func (im Iamport) PaySchedule(f form.PaymentSchedule) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/subscribe/payments/schedule", im.ImpURL)
	valid, err := form.Validate(f)
	if !valid || err != nil {
		return nil, err
	}
	return im.post(url, f)
}

func (im Iamport) PayUnschedule(f form.PaymentUnschedule) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/subscribe/payments/unschedule", im.ImpURL)
	valid, err := form.Validate(f)
	if !valid || err != nil {
		return nil, err
	}
	return im.post(url, f)
}

func (im Iamport) Cancel(f form.PaymentCancel) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s/payments/cancel", im.ImpURL)
	valid, err := form.Validate(f)
	if !valid || err != nil {
		return nil, err
	}
	return im.post(url, f)
}

func New(impKey string, impSecret string, impURL ...string) *Iamport {
	url := "https://api.iamport.kr"
	if len(impURL) == 1 {
		url = impURL[0]
	}
	return &Iamport{impKey, impSecret, url}
}
