package bringdata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type ResponseData struct {
	Msg          string `json:"msg"`
	Code         int    `json:"code"`
	RequestCount int    `json:"request_count"`
	Version      string `json:"version"`
	Data         struct {
		List []Order `json:"list"`
	} `json:"data"`
}

type Order struct {
	OrderNo   string `json:"order_no"`
	OrderType string `json:"order_type"`
	Device    struct {
		Type string `json:"type"`
	} `json:"device"`
	OrderTime    int64 `json:"order_time"`
	CompleteTime int64 `json:"complete_time"`
	Orderer      struct {
		MemberCode string `json:"member_code"`
		Name       string `json:"name"`
		Email      string `json:"email"`
		Call       string `json:"call"`
		Call2      string `json:"call2"`
	} `json:"orderer"`
	Delivery struct {
		Country     string `json:"country"`
		CountryText string `json:"country_text"`
		Address     struct {
			Name            string `json:"name"`
			Phone           string `json:"phone"`
			Phone2          string `json:"phone2"`
			Postcode        string `json:"postcode"`
			Address         string `json:"address"`
			AddressDetail   string `json:"address_detail"`
			AddressStreet   string `json:"address_street"`
			AddressBuilding string `json:"address_building"`
			AddressCity     string `json:"address_city"`
			AddressState    string `json:"address_state"`
			LogisticsType   string `json:"logistics_type"`
		} `json:"address"`
		Memo string `json:"memo"`
	} `json:"delivery"`
	Payment struct {
		PayType       string `json:"pay_type"`
		PgType        string `json:"pg_type"`
		DelivType     string `json:"deliv_type"`
		DelivPayType  string `json:"deliv_pay_type"`
		PriceCurrency string `json:"price_currency"`
		TotalPrice    int64  `json:"total_price"`
		DelivPrice    int64  `json:"deliv_price"`
		Point         int64  `json:"point"`
	} `json:"payment"`
	Form []struct {
		Type            string      `json:"type"`
		Title           string      `json:"title"`
		Desc            string      `json:"desc"`
		Value           string      `json:"value"`
		FormConfigValue interface{} `json:"form_config_value"`
	} `json:"form"`
}

type ListData struct {
	No        string
	Name      string
	Email     string
	CallNum   string
	OrderNo   string
	OrderTime string
	FromEmail string
	Price     int64
}

func Getdata(accessToken string) []ListData {
	// Request 객체 생성
	req, err := http.NewRequest("GET", "https://api.imweb.me/v2/shop/orders?type=normal&status=COMPLETE", nil)
	if err != nil {
		panic(err)
	}

	// 필요한 해더 추가
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("access-token", accessToken)

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	var respData ResponseData
	err = json.NewDecoder(resp.Body).Decode(&respData)
	if err != nil {
		fmt.Printf("Error while decoding response: %v", err)
		return nil
	}

	if len(respData.Data.List) == 0 {
		fmt.Println("List is empty")
		return nil
	}

	listDatas := make([]ListData, 0)

	for _, order := range respData.Data.List {
		temp := ListData{
			OrderNo:   order.OrderNo,
			Name:      order.Orderer.Name,
			Email:     order.Orderer.Email,
			CallNum:   order.Orderer.Call,
			OrderTime: changeTime(order.OrderTime),
			FromEmail: order.Form[0].Value,
			Price:     order.Payment.TotalPrice,
		}

		listDatas = append(listDatas, temp)
	}

	return listDatas
}

func changeTime(timeData int64) string {
	timeObj := time.Unix(timeData, 0)
	formattedTime := timeObj.Format("2006-01-02 15:04:05")

	return formattedTime
}
