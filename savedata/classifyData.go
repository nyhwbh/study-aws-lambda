package savedata

import (
	"studylambda/bringdata"
)

type clsfData struct {
	class5500 [][]interface{}
	// class80000  [][]interface{}
	// class100000 [][]interface{}
	// class200000 [][]interface{}
	// class790000 [][]interface{}
}

// 아임웹에서 가지고 온 데이터를 각 강의별로 분리
func classifyData(targetData []bringdata.ListData) clsfData {
	var classifiedData clsfData
	for _, data := range targetData {
		switch data.Price {
		case 5500:
			classifiedData.class5500 = append(classifiedData.class5500, changeData(data))
		// case 80000:
		// 	classifiedData.class80000 = append(classifiedData.class80000, changeData(data))
		// case 100000:
		// 	classifiedData.class100000 = append(classifiedData.class100000, changeData(data))
		// case 200000:
		// 	classifiedData.class200000 = append(classifiedData.class200000, changeData(data))
		// case 790000:
		// 	classifiedData.class790000 = append(classifiedData.class790000, changeData(data))
		default:
			// 위 분류에 해당하지 않는 경우 처리
		}
	}

	return classifiedData
}

func changeData(data bringdata.ListData) []interface{} {
	dataSlice := []interface{}{data.No, data.Name, data.Email, data.CallNum, data.OrderNo, data.OrderTime, data.FromEmail, data.Price}
	return dataSlice
}
