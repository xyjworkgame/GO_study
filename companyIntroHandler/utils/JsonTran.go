package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

func TranJson(param io.ReadCloser) (map[string]interface{}, error) {
	params, _ := ioutil.ReadAll(param)
	var data map[string]interface{}
	if err := json.Unmarshal([]byte(params), &data); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil
}
