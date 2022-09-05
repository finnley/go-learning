package main

import (
	"fmt"
	"unsafe"
)

// 空结构体不占内存空间大小，空间大小为0，但是也是有地址的
// 用法在channel中
type NULL struct {
}

func main5() {
	// 如果定义多个空结构体，都指向同一个内存地址
	//null := NULL{}
	null := struct {
	}{}
	fmt.Println(unsafe.Sizeof(null)) // 0
	fmt.Printf("%p\n", &null)        // 0x11aac78

	null1 := struct {
	}{}
	fmt.Printf("%p\n", &null1) // 0x11aac78
}

type S struct {
	A struct{}
	B struct{}
}

func main4() {
	var s S
	fmt.Printf("%p\n", &s.A)      // 0x11aac78
	fmt.Printf("%p\n", &s.B)      // 0x11aac78
	fmt.Println(unsafe.Sizeof(s)) // 结果： 0
}


type UmonThreholds []UmonThrehold
type UmonThrehold struct {
	Id              string            `json:"id"`
	Expr            string            `json:"expr"`
	Description     string            `json:"description"`
	DurationSeconds int               `json:"duration_seconds"`
	Labels          map[string]string `json:"labels"`
	Summary         string            `json:"summary"`
	Detail          string            `json:"detail"`
	Enable          bool              `json:"enable"`
}

type UmonConfig struct {
	Thresholds UmonThreholds `json:"thresholds"`
}
func main()  {
	//noExistEmergencySet := map[string]struct{}{
	//	"PD_cluster_offline_tikv_nums":     {},
	//	"TiKV_memory_used_too_fast":        {},
	//	"TiKV_GC_can_not_work":             {},
	//	"TiDB_schema_error":                {},
	//	"TiDB_tikvclient_region_err_total": {},
	//	"TiDB_binlog_error_total":          {},
	//	"TiDB_domain_load_schema_total":    {},
	//	"TiDB_monitor_keep_alive":          {},
	//}

}