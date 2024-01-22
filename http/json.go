package http

import (
	jsonIter "github.com/json-iterator/go"
	"strings"
)

var cjson = jsonIter.ConfigCompatibleWithStandardLibrary

// 滴滴开源json库
func JsonEncode(v interface{}) ([]byte, error) {

	return cjson.Marshal(v)
}

func JsonDecode(v string) (map[string]interface{}, error) {

	reader := strings.NewReader(v)
	decoder := cjson.NewDecoder(reader)
	params := make(map[string]interface{})
	err := decoder.Decode(&params)

	return params, err

}

func JsonMarshal(v interface{}) ([]byte, error) {

	return cjson.Marshal(v)
}

func JsonUnmarshal(data []byte, v interface{}) error {
	jsonIter.Unmarshal(data, v)
	return cjson.Unmarshal(data, v)
}

func JsonUnmarshalFromString(data string, v interface{}) error {
	return cjson.UnmarshalFromString(data, v)
}

func MarshalToString(v interface{}) (string, error) {
	return cjson.MarshalToString(v)
}

func GetValueByKeyFromJson(data []byte, key string) string {

	return cjson.Get(data, key).ToString()
}
