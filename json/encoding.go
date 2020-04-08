package json

import (
	"encoding/json"
)

func Encode(v interface{}) (string, error) {
	// 序列化成json 字节数组
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func Decode(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}
