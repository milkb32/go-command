package common

import (
	"encoding/json"
	"time"
)

func StructToMap(data interface{}) (map[string]interface{}, error) {
	jsonString, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	newMap := map[string]interface{}{}
	err = json.Unmarshal(jsonString, &newMap)
	if err != nil {
		return nil, err
	}

	return newMap, nil
}

func TimestampToDate(timestamp int64) string {
	tm := time.Unix(timestamp, 0)
	return tm.Format("2006-01-02 15:04:05")
}