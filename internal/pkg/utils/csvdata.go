package utils

import (
	"fmt"
	makves "github.com/cucumberjaye/makves_testovoe"
	"reflect"
	"strconv"
)

func DataToItems(data [][]string) ([]makves.Item, error) {
	var items = []makves.Item{}

	for i := 0; i < len(data)-1; i++ {
		item := makves.Item{}
		obj := reflect.ValueOf(&item).Elem()
		for j := 0; j < obj.NumField(); j++ {
			if _, ok := obj.Field(j).Interface().(int); ok {
				tmp, err := strconv.Atoi(data[i][j])
				if err != nil {
					fmt.Println("lol", data[i][j], item)
					fmt.Printf("%#v\n", data[i])
					return nil, err
				}
				obj.Field(j).SetInt(int64(tmp))
			} else if _, ok = obj.Field(j).Interface().(string); ok {
				if data[i][j] == "\"\"" {
					obj.Field(j).SetString("")
				} else {
					obj.Field(j).SetString(data[i][j])
				}
			} else if _, ok = obj.Field(j).Interface().(bool); ok {
				tmp, err := strconv.ParseBool(data[i][j])
				if err != nil {
					obj.Field(j).SetBool(false)
					continue
				}
				obj.Field(j).SetBool(tmp)
			}
		}
		items = append(items, item)
	}
	return items, nil
}
