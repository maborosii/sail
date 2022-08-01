package util

import (
	"fmt"
	"reflect"
)

func validateStructForReflect(in interface{}) (interface{}, error) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("StructToMap only accepts struct or struct pointer; got %T", v)
	}
	return v.Interface(), nil
}

// 结构体转为Map[string]interface{}
// 将结构体属性拉平，即map的层级只有一层
// 需要原结构的所有属性tagName是不同的
// bug: 若 value 为结构体切片则无法展开，数组本身属性都不适合拉平展开
// 建议若该属性为结构体切片，则不要进行展开，自行定义方法获取想要的值。
func SpreadToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	newIn, err := validateStructForReflect(in)
	// 校验参数是否是结构体
	if err != nil {
		return nil, err
	}

	// 存放待解析结构体
	queue := make([]interface{}, 0)
	offset := 0
	// fmt.Println(offset)
	queue = append(queue, newIn)

	// QUEUE_LOOP:
	for len(queue) != offset {
		// 移出队首部元素
		v := reflect.ValueOf(queue[offset])
		offset++
		t := v.Type()
		// 遍历结构体字段
		// 指定tagName值为map中key;字段值为map中value

		// FIELD_LOOP:
		for i := 0; i < v.NumField(); i++ {
			element := v.Field(i)
			// fmt.Printf("%+v\n", element)
			switch element.Kind() {
			case reflect.Ptr:
				// 结构体指针属性
				el := element.Elem()
				if el.Kind() == reflect.Struct {
					//! notice
					// 反射第二定律:将反射类型对象转换为接口类型变量
					queue = append(queue, el.Interface())
				} else {
					ti := t.Field(i)
					if tagValue := ti.Tag.Get(tagName); tagValue != "" {
						out[tagValue] = el.Interface()
					}
				}
			case reflect.Struct:
				// 结构体属性
				queue = append(queue, element.Interface())
			default:
				// 一般属性
				ti := t.Field(i)
				if tagValue := ti.Tag.Get(tagName); tagValue != "" {
					out[tagValue] = element.Interface()
				}
			}
		}
	}
	return out, nil
}
