package helper

import "encoding/json"

// ConvertStruct2Map ...
func ConvertStruct2Map(data interface{}) (map[string]interface{}, error) {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	err = json.Unmarshal(dataByte, &result)
	return result, err
}

