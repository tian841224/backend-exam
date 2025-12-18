package main

import (
	"fmt"
	"reflect"
)

func swap[T any](a, b T) {
	// 使用 reflect 來交換值
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	// 檢查是否為指標
	if va.Kind() != reflect.Ptr || vb.Kind() != reflect.Ptr {
		panic("需為指標")
	}

	// 檢查是否為相同型別
	if va.Elem().Kind() != vb.Elem().Kind() {
		panic("兩者型別需一致")
	}

	// 交換值
	tmp := reflect.New(va.Elem().Type()).Elem()
	tmp.Set(va.Elem())
	va.Elem().Set(vb.Elem())
	vb.Elem().Set(tmp)
}

func main() {
	a := 10
	b := 20

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)

	swap(&a, &b)

	fmt.Printf("a = %d, &a = %p\n", a, &a)
	fmt.Printf("b = %d, &b = %p\n", b, &b)
}
