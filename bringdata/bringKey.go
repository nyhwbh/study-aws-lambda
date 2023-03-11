package bringdata

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type token struct {
	Msg          string
	Code         int
	Access_Token string
}

func GetAccessToken() string {
	// API key와 secret key 가져오기
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	// Request 객체 생성
	req, err := http.NewRequest("GET", "https://api.imweb.me/v2/auth&key="+apiKey+"&secret="+apiSecret, nil)
	if err != nil {
		panic(err)
	}

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 토근 반환
	var tokenObj = token{}
	bytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(bytes), &tokenObj)
	if err != nil {
		log.Fatal("Fail to unmarshalling data")
	}

	return tokenObj.Access_Token
}
