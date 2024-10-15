package xtransform

import (
	"fmt"
	"testing"
)

type Person struct {
	Name    string `tag:"age"`
	Age     int    `tag:"age"`
	Country string `tag:"country"`
}

func TestTransformStructFieldsWithTag(t *testing.T) {
	// 定义一个转换函数，将整数转换为字符串
	transformFunc := func(value interface{}) (interface{}, error) {
		if _, ok := value.(int); ok {
			return 100, nil
		}
		return "100", nil
	}

	// 输入结构体
	input := Person{
		Name:    "John",
		Age:     30,
		Country: "USA",
	}

	// 执行转换
	output, err := ProcessStructFields(input, "tag", "age", transformFunc)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 输出转换后的结果
	fmt.Printf("%+v\n", output)
}
