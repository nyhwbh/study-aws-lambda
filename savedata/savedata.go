package savedata

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"studylambda/bringdata"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

const (
	// // test code
	// spreadsheetID_5500 = "1u929XLIGAoJ9wISI5S2rpqbIYxv0N7smV5EO-Q8CWOo"

	// main code
	spreadsheetID_80000  = "1miHqinWQGva5UTGXkV1nlFW_OVOWfjWJuNiTifQfNvc"
	spreadsheetID_100000 = "1W0-OYi2emrt6lLPe95yO7Nrfn-1-v7gQprXfD_v937Q"
	spreadsheetID_200000 = "11bkH00fDOh9f8PoRF0zUcikO2mC2wR7AswGyURw5UA4"
	spreadsheetID_790000 = "1FMu1Gr3Gw5qbmx33OP6vpIcPjxSogHiGZIVLdXn0Ok8"
)
const sheetID uint = 0

func SaveDataAll() {
	accessToken := bringdata.GetAccessToken()

	data := bringdata.Getdata(accessToken)
	// 가져오는 데이터에 변함이 생긴경우 함수 실행
	if data != nil {
		// 가져온 데이터 가격에따라 분리
		inputData := classifyData(data)

		// // test code
		// savedata(spreadsheetID_5500, sheetID, inputData.class5500)

		// main code
		// 분리된 데이터 각 시트에 삽입
		//class80000
		savedata(spreadsheetID_80000, sheetID, inputData.class80000)
		//class100000
		savedata(spreadsheetID_100000, sheetID, inputData.class100000)
		//class200000
		savedata(spreadsheetID_200000, sheetID, inputData.class200000)
		//class790000
		savedata(spreadsheetID_790000, sheetID, inputData.class790000)
	}
}

// 시트에 자료를 삽입하는 함수
func savedata(spreadSheetID string, sheetID uint, inputData [][]interface{}) {
	// client_secret.json 가져오기
	data, err := bringClientSecret()
	checkError(err)

	// google 토큰 가지고오기
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)

	//클라이언트 생성
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	var spreadsheet spreadsheet.Spreadsheet
	for i := 0; i < 3; i++ {
		// 시트 연결하기
		spreadsheet, err = service.FetchSpreadsheet(spreadSheetID)
		// 오류 발생시 제시도
		if err == nil {
			fmt.Println("시트 연결에 성공하였습니다.")
			break
		}
		// 새번의 재시도를 했지만 실패한 경우
		if i == 3 && err == nil {
			fmt.Println("시트 연결에 실패하였습니다.")
		}
		time.Sleep(15 * time.Second)
	}
	// 재시도후 오류가 없는 경우
	if err == nil {
		sheet, err := spreadsheet.SheetByIndex(0)
		checkError(err)

		lastRowIndex := len(sheet.Rows) - 1

		// 마지막 행(row) 가지고 오기 lastRow[4].Value
		lastRow := sheet.Rows[lastRowIndex]

		// 마지막 행 번호(no) 가지고 오기
		lastNo := lastRow[4].Row

		// 마지막 행과 일치하는 해 이후에 있는 데이터 삭제
		// 아임웹 데이터는 [최신 -> 과거] 시트 데이터는 [과거 -> 최신]
		saveData := duplicateDataCheck(lastRow[4].Value, inputData)

		// save data
		if len(saveData) > 0 {
			k := 1
			for i := len(saveData) - 1; i >= 0; i-- {
				// sheet NO 입력하기
				sheet.Update(lastRowIndex+k, 0, strconv.Itoa(int(lastNo)+1))
				for j := 1; j < len(saveData[i])-1; j++ {
					// Data 입력하기
					sheet.Update(lastRowIndex+k, j, fmt.Sprintf("%v", saveData[i][j]))
				}
				lastNo++
				k++
			}
			k = 1
		}

		// 시트에 변경 내용 적용  (+) 변경 내용 없을시(sheet가 비어있는 경우) 실항하지 않음
		if len(saveData) > 0 {
			err = sheet.Synchronize()
			if err != nil {
				panic(err.Error())
			} else {
				fmt.Println("시트에 자료 저장을 했습니다. 업데이트 시간 : ", time.Now().Format("2006-01-02 15:04:05"))
			}
		}
	}
}

// 중복되는 과거 데이터 삭제
func duplicateDataCheck(orderNo string, targetData [][]interface{}) [][]interface{} {
	// targetData가 없는경우 애러방지
	var newData [][]interface{}
	if len(targetData) > 0 {
		// 중복되는 것이 없는 경우 데이터를 삭제하지 않음
		deleteindex := len(targetData)
		for i, row := range targetData {
			for _, value := range row {
				if value == orderNo {
					deleteindex = i
					break
				}
			}
		}
		newData = append(newData, targetData[:deleteindex]...)

	} else {
		newData = targetData
	}
	return newData
}

func bringClientSecret() ([]byte, error) {
	secretName := "client_secret.json"
	region := "ap-northeast-2"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		log.Fatal(err.Error())
	}

	// Decrypts secret using the associated KMS key.
	var secretString string = *result.SecretString

	secretByte := []byte(secretString)

	return secretByte, err
}

// err 체크 함수
func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
