package main

import (
	"encoding/json"
	"fmt"
)

type AlertReceivers struct {
	WechatReceivers []WechatReceiverItem `json:"wechat_receivers"`
}

type WechatReceiverItem struct {
	WechatCorpId      string `json:"wechat_corp_id"`
	WechatCorpSecret  string `json:"wechat_corp_secret"`
	WechatAgentId     int64  `json:"wechat_agent_id"`
	WechatSafeEnabled bool   `json:"wechat_safe_enabled"`
	WechatProxyAddr   string `json:"wechat_proxy_addr"`
	WechatUserId      string `json:"wechat_user_id"`
	WechatPartyId     string `json:"wechat_party_id"`
	WechatTagId       string `json:"wechat_tag_id"`
}

func main() {
	wechatReceivers := []WechatReceiverItem{
		{
			WechatCorpId:      "1",
			WechatCorpSecret:  "aa",
			WechatAgentId:     0,
			WechatSafeEnabled: false,
			WechatProxyAddr:   "192.168.1.1",
			WechatUserId:      "10001",
			WechatPartyId:     "1",
			WechatTagId:       "101",
		},
		{
			WechatCorpId:      "2",
			WechatCorpSecret:  "aa",
			WechatAgentId:     0,
			WechatSafeEnabled: false,
			WechatProxyAddr:   "192.168.1.2",
			WechatUserId:      "10002",
			WechatPartyId:     "2",
			WechatTagId:       "102",
		},
	}
	alertReceivers := AlertReceivers{wechatReceivers}

	//for _, v := range alertReceivers.WechatReceivers {
	//	fmt.Printf("%#v\n", v)
	//	t := reflect.TypeOf(v)
	//	fmt.Println(t.Name())
	//	fmt.Println(t.Kind())
	//	//for i := 0; i < t.NumField(); i++ {
	//	//	field := t.Field(i)
	//	//	tag := field.Tag.Get("json")
	//	//	fmt.Printf("%d.%v (%v), tag: '%v'\n", i+1, field.Name, field.Type.Name(), tag)
	//	//	if tag == "wechat_proxy_addr" {
	//	//		fmt.Println(field.Name)
	//	//	}
	//	//}

	slice, err := json.Marshal(alertReceivers)
	if err != nil {
		fmt.Println("json转换失败")
	} else {
		fmt.Println(string(slice))
	}

	var result map[string]interface{}
	err = json.Unmarshal(slice, &result)
	if err != nil {
		fmt.Println("map转换失败")
	}

	nodes := result["wechat_receivers"].([]interface{})
	for _, node := range nodes {
		m := node.(map[string]interface{})
		//if _, exists := m["wechat_proxy_addr"]; exists {
		//	m["wechat_proxy_addr"] = "new-value1"
		//	fmt.Println("success")
		//	//if name == "node1" {
		//	//	m["location"] = "new-value1"
		//	//} else if name == "node2" {
		//	//	m["location"] = "new-value2"
		//	//}
		//}
		if _, exists := m["wechat_proxy_ip"]; exists {
			m["wechat_proxy_addr"] = "new-value1"
		}
	}
	fmt.Println(result)

	// Convert golang object back to byte
	//byteValue, err = json.Marshal(result)
	//if err != nil {
	//	fmt.Println("转换失败")
	//}
	//
	//err := json.Unmarshal(slice, &alertReceivers)
	//if err != nil {
	//	fmt.Println("json转换失败")
	//} else {
	//	fmt.Println(stu)
	//}
}
