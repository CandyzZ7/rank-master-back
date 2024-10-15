package xtransform

import (
	"fmt"
	"reflect"
)

func ProcessStructFields(input interface{}, tagKey string, tagValue string, transformFunc func(interface{}) (interface{}, error)) (interface{}, error) {
	// 获取输入结构体的反射值
	inputValue := reflect.ValueOf(input)
	// 检查输入值是否为指针类型，如果是，则解引用
	if inputValue.Kind() == reflect.Ptr {
		// 解引用指针，获取指向的值
		inputValue = inputValue.Elem()
		// 检查解引用后的值是否为结构体类型
		if inputValue.Kind() != reflect.Struct {
			return nil, fmt.Errorf("input is not a struct")
		}
	} else {
		// 如果不是指针类型，则直接检查是否为结构体类型
		if inputValue.Kind() != reflect.Struct {
			return nil, fmt.Errorf("input is not a struct")
		}
	}

	// 创建一个新的结构体类型的实例
	outputValue := reflect.New(inputValue.Type()).Elem()

	// 遍历输入结构体的字段
	for i := 0; i < inputValue.NumField(); i++ {
		// 获取字段值
		fieldValue := inputValue.Field(i)

		// 获取字段的类型
		fieldType := inputValue.Type().Field(i)

		// 获取字段的标签
		tag := fieldType.Tag.Get(tagKey)

		// 如果字段有指定的tag，则进行处理
		if tag == tagValue {
			// 使用传入的处理方法进行转换
			transformedValue, err := transformFunc(fieldValue.Interface())
			if err != nil {
				return nil, err
			}

			// 将转换后的值设置到新结构体中的相应字段上
			outputValue.Field(i).Set(reflect.ValueOf(transformedValue))
		} else {
			// 如果字段没有指定的tag，则直接复制字段值到新结构体中
			outputValue.Field(i).Set(fieldValue)
		}
	}

	// 返回新结构体的实例
	return outputValue.Addr().Interface(), nil
}
