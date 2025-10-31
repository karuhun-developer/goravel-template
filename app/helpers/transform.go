package helpers

import (
	"encoding/json"
	"strconv"

	"github.com/goravel/framework/contracts/http"
)

func ModelToMap(model any) http.Json {
	responseByte, _ := json.Marshal(model)
	var responeMap http.Json
	_ = json.Unmarshal(responseByte, &responeMap)

	return responeMap
}

func StringToBool(str string) bool {
	boolean, err := strconv.ParseBool(str)

	if err != nil {
		return false
	}

	return boolean
}