package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

// 儲存已處理過的位址
var processedAddresses = make(map[unsafe.Pointer]bool)

func TrimAllStrings(a any) {
	elem := reflect.ValueOf(a)

	// 檢查是否已處理過
	addr := unsafe.Pointer(elem.Pointer())
	if processedAddresses[addr] {
		return
	}
	processedAddresses[addr] = true

	// 持續解開指標，直到不是指標為止
	for elem.Kind() == reflect.Ptr {
		if elem.IsNil() {
			return
		}
		elem = elem.Elem()
	}

	// 檢查是否為結構體
	if elem.Kind() != reflect.Struct {
		return
	}

	// 取得欄位數量
	numFields := elem.NumField()

	// 遍歷每個欄位
	for i := 0; i < numFields; i++ {
		field := elem.Field(i)

		// 如果是字串，進行 trim
		if field.Kind() == reflect.String {
			if field.CanSet() {
				field.SetString(strings.TrimSpace(field.String()))
			}
		}

		// 如果是指標類型，遞迴處理（處理 Next 欄位）
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			TrimAllStrings(field.Interface())
		}
	}
}
func main() {

	type Person struct {
		Name string
		Age  int
		Next *Person
	}

	a := &Person{
		Name: " name ",
		Age:  20,
		Next: &Person{
			Name: " name2 ",
			Age:  21,
			Next: &Person{
				Name: " name3 ",
				Age:  22,
			},
		},
	}

	TrimAllStrings(&a)

	m, _ := json.Marshal(a)

	fmt.Println(string(m))

	a.Next = a

	TrimAllStrings(&a)

	fmt.Println(a.Next.Next.Name == "name")
}
