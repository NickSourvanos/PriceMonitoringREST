package utils

import (
	"PriceMonitoringService/models"
	"encoding/json"
	"github.com/mitchellh/mapstructure"
)

func MapToStruct(dict interface{}) (user *models.TokenUser){
	var tempMap = dict
	test := []byte(`{"key":"value"`)
	//test["UserInfo"] = val
	json.Unmarshal(test, &tempMap)
	//myMap := tempMap.(map[string]interface{})
	//userId := myMap["userId"]
	var str *models.TokenUser
	mapstructure.Decode(dict, &str)

	return str

}
