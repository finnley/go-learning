package demo_reflect

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// demo1:
// 反射的各种panic
func TestReflectPanic1(t *testing.T) {
	typ := reflect.TypeOf(123)
	//typ.NumField() // 这么写会报错，NumField只能用于结构体，所以一般会使用下面写法
	if typ.Kind() == reflect.Struct {
		typ.NumField()
	}

}

func TestReflectPanic2(t *testing.T) {
	//typ := reflect.TypeOf(User{})
	typ := reflect.TypeOf(&User{}) // 指针也不行，因为它是指针，不是struct
	//typ.NumField() // 这么写会报错，NumField只能用于结构体，所以一般会使用下面写法
	if typ.Kind() == reflect.Struct {
		typ.NumField()
	}
}

func TestReflectPanic3(t *testing.T) {
	typ := reflect.TypeOf(&User{})
	if typ.Kind() == reflect.Struct {
		fmt.Println("结构体")
	} else if typ.Kind() == reflect.Ptr {
		fmt.Println("指针")
	}
}

type User struct {
	Name string
	// todo 还要考虑小写非公共字段，比如
	//age int
}

// demo2:
// 反射输出所有的字段名字（关键点是只有Struct才有字段）
func InterateFields(val any) {
	res, err := interateFields(val)
	if err != nil {
		fmt.Println(err)
		return
	}
	for k, v := range res {
		fmt.Println(k, v)
	}
}

func interateFields(val any) (map[string]any, error) {
	//fmt.Println("字段1")
	//fmt.Println("字段2")
	if val == nil {
		return nil, errors.New("不能为 nil")
	}
	typ := reflect.TypeOf(val)
	refVal := reflect.ValueOf(val)
	// 怎么拿到指针指向的对象
	//if typ.Kind() == reflect.Ptr {
	//	typ = typ.Elem()
	//	refVal = refVal.Elem()
	//}
	for typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		refVal = refVal.Elem()
	}

	numField := typ.NumField()
	res := make(map[string]any, numField)

	for i := 0; i < numField; i++ {
		fdType := typ.Field(i)
		//res[fdType.Name] = refVal.Field(i)
		res[fdType.Name] = refVal.Field(i).Interface()
	}
	return res, nil
}

func TestInterateFields(t *testing.T) {
	u1 := &User{
		Name: "夜猫子",
	}
	u2 := &u1
	tests := []struct {
		// 名字
		name string
		// 输入部分
		val any
		// 输出部分
		wantRes map[string]any
		wantErr error
	}{
		// 构造一个个测试用例
		{
			name:    "nil",
			val:     nil,
			wantErr: errors.New("不能为 nil"),
		},
		{
			name:    "user",
			val:     User{Name: "Tom"},
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "Tom",
			},
		},
		{
			name: "pointer",
			val:  &User{Name: "Jerry"},
			// 要支持指针就是nil
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "Jerry",
			},
		},
		// 如果是多重指针
		{
			name: "multiple pointer",
			val:  u2,
			// 要支持指针就是nil
			wantErr: nil,
			wantRes: map[string]any{
				"Name": "夜猫子",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := interateFields(tt.val)
			assert.Equal(t, tt.wantErr, err)
			if err != nil {
				return
			}
			assert.Equal(t, tt.wantRes, res)
		})
	}
}
