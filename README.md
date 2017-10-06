[![Build Status](https://travis-ci.org/ClaudeSeo/go-iamport.svg?branch=master)](https://travis-ci.org/ClaudeSeo/go-iamport)

# go-iamport
GoLang 아임포트(http://iamport.kr/) REST API Client

## Installation
```shell
$ go get github.com/claudeseo/go-iamport
```

## Usage
```go
import (
  "github.com/claudeseo/go-iamport"
)

func main() {
  im := iamport.New("imp_apikey" "ekKoeW8RyKuT0zgaZsUtXXTLQ4AhPFW3ZGseDA6bkA5lamv9OqDMnxyeB9wqOsuO9W3Mx9YSJ4dTqJ3f")

  // 아임포트 고유번호로 결제 정보 조회
  result, err := im.FindByImpUID("imp_uid")

  // 상품 아이디로 결제 정보 조회
  result, err := im.FindByMerchantUID("merchant_uid")
}
```

## Features
- [x] get token
- [x] find by merchant_uid
- [x] find by imp_uid
- [x] payment one time
- [x] payment again
- [x] schedule
- [x] unschedule
- [x] payment cancel
- [ ] prepare
- [ ] receipts
- [ ] subscribe.customer
- [ ] vbanks
- [ ] sms certification


