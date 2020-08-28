package utils

import (
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"io/ioutil"
	_ "io/ioutil"
)

type Resource struct {
	Name string `json:"name"` //名称
	Url  string `json:"url"`  //路径
}

func SwaggerJson(swagger string) (a map[string]interface{}) {
	xxx := make(map[string]interface{})
	bytes, err := ioutil.ReadFile(swagger)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.Unmarshal(bytes, &xxx)
	if err != nil {
		fmt.Println(err.Error())
	}
	birds := xxx["paths"].(map[string]interface{})
	return birds
}

func GetOpId(vv interface{}) (ret interface{}) {
	dd := make(map[string]interface{})
	_ = beego.HelperConvetInterface(vv, &dd)
	for kk, value := range dd {
		if kk == "operationId" {
			ret = value
			return value
		} else {
			ret = GetOpId(value)
			if ret != nil {
				return ret
			}
		}
	}
	return nil
}
