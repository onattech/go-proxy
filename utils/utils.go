package utils

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
)

// Converts comma separated float like string to float64
// @example 13,5682 ==> 13.5682
func EasyFloat(floatLikeString string) float64 {
	dotString := strings.Replace(floatLikeString, ",", ".", 1)
	converted, err := strconv.ParseFloat(dotString, 8)
	if err != nil {
		panic("Can't convert string to float")
	}

	return converted
}

// Prettifies json
func PrettyStruct(data interface{}) string {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Println(err)
	}
	return string(val)
}
