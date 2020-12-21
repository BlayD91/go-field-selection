package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	dynamicstruct "github.com/ompluscator/dynamic-struct"
)

var fields = []string{"Id", "FirstName", "Age", "Roles.Id", "Roles.Images.Url", "Roles.Images.ImageGroups.Id", "Image.Url", "Birthday"}

var user = []User{
	{
		Id:        1,
		FirstName: "John ",
		LastName:  "Doe",
		Age:       24,
		Password:  "123456",
		Roles: []Role{
			{
				Id:   1,
				Name: "User",
				Images: []Image{
					{
						Id:   4,
						Url:  "img4",
						Size: 44,
						ImageGroups: []ImageGroup{
							{
								Id:   1,
								Name: "Type 1",
							},
							{
								Id:   2,
								Name: "Type 2",
							},
						},
					},
					{
						Id:   5,
						Url:  "img5",
						Size: 55,
						ImageGroups: []ImageGroup{
							{
								Id:   3,
								Name: "Type 3",
							},
							{
								Id:   4,
								Name: "Type 4",
							},
						},
					},
					{
						Id:   6,
						Url:  "img6",
						Size: 66,
						ImageGroups: []ImageGroup{
							{
								Id:   2,
								Name: "Type 1",
							},
							{
								Id:   3,
								Name: "Type 2",
							},
						},
					},
				},
			},
			{
				Id:   2,
				Name: "Admin",
			},
		},
		Image: Image{
			Id:   14,
			Url:  "img14",
			Size: 123,
		},
		Birthday: time.Now(),
	},
}

func main() {
	b, _ := json.Marshal(&user)

	resMap := fieldPreprocessing(fields)

	resStruct := buildStruct(resMap, user)

	err := json.Unmarshal(b, &resStruct)
	if err != nil {
		fmt.Println(err)
	}

	b, err = json.Marshal(&resStruct)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))

}

// Формирование карты полей
func fieldPreprocessing(fields []string) map[string]interface{} {
	res := map[string]interface{}{}

	nextFields := map[string][]string{}

	for _, field := range fields {
		if !strings.Contains(field, ".") {
			res[field] = true
		} else {
			arr := strings.Split(field, ".")
			nextFields[arr[0]] = append(nextFields[arr[0]], strings.Join(arr[1:], "."))
		}
	}
	for k, v := range nextFields {
		res[k] = fieldPreprocessing(v)
	}
	return res
}

// Построение стуктуры на основе карты полей
func buildStruct(structMap map[string]interface{}, data interface{}) interface{} {
	var i *interface{}
	isArr := isArray(data)
	builder := dynamicstruct.NewStruct()
	for k, v := range structMap {
		if nextMap, isOk := v.(map[string]interface{}); isOk {
			builder.AddField(k, buildStruct(nextMap, getDataFromStructByFieldName(data, k)), "")
		} else {
			builder.AddField(k, i, "")
		}
	}
	if isArr {
		return builder.Build().NewSliceOfStructs()
	}
	return builder.Build().New()
}

//Получаем вложенные данные из структуры
func getDataFromStructByFieldName(data interface{}, field string) interface{} {
	r := reflect.ValueOf(data)
	if r.Kind() == reflect.Slice {
		r := reflect.TypeOf(data)
		el := reflect.New(r.Elem())
		return reflect.Indirect(el).FieldByName(field).Interface()
	} else {
		return reflect.Indirect(r).FieldByName(field).Interface()
	}
}

func isArray(data interface{}) (res bool) {
	rt := reflect.TypeOf(data)
	switch rt.Kind() {
	case reflect.Slice:
		res = true
	case reflect.Array:
		res = true
	default:
		res = false
	}
	return res
}
