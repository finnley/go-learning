package main

import (
	"fmt"
)

func main() {
	// Creating the maps for JSON
	//m := map[string]interface{}{}
	//
	//// Parsing/Unmarshalling JSON encoding/json
	//err := json.Unmarshal([]byte(input), &m)
	//
	//if err != nil {
	//	panic(err)
	//}
	//parseMap(m)

	fmt.Println(input)
}

func parseMap(aMap map[string]interface{}) {
	for key, val := range aMap {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println(key)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println(key)
			parseArray(val.([]interface{}))
		default:
			fmt.Println(key, ":", concreteVal)
		}
	}
}

func parseArray(anArray []interface{}) {
	for i, val := range anArray {
		switch concreteVal := val.(type) {
		case map[string]interface{}:
			fmt.Println("Index:", i)
			parseMap(val.(map[string]interface{}))
		case []interface{}:
			fmt.Println("Index:", i)
			parseArray(val.([]interface{}))
		default:
			fmt.Println("Index", i, ":", concreteVal)

		}
	}
}

const input = `
{
    "data": [
        {
            "domain": "ni.suock.com",
            "method": "GET",
            "data": "empty data",
            "path": "/node/clientnodelist",
            "reply": "reply data1",
            "base64reply": "reply data2"
        },
		{
            "domain": "hi.nihao.com",
            "method": "POST",
            "data": "",
            "path": "/base/captcha",
            "reply": "reply data1",
            "base64reply": ""
        }
    ]
}
`
