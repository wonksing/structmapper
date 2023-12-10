package structmapper

import (
	"reflect"
)

type BasicTypeMapper struct {
	SrcIndex  []int
	DestIndex []int
}

func (c *BasicTypeMapper) Map(srcValue, destValue reflect.Value) {
	mapBasicType(c.SrcIndex, c.DestIndex, srcValue, destValue)
}

func mapBasicType(srcIndex, destIndex []int, srcValue, destValue reflect.Value) {

	srcField, err := srcValue.FieldByIndexErr(srcIndex)
	if err != nil {
		return
	}
	var destField reflect.Value
	for i := 0; i < len(destIndex); i++ {
		destField = destValue.FieldByIndex(destIndex[:i+1])
		if destField.Type().Kind() == reflect.Pointer && destField.IsNil() {
			destField.Set(reflect.New(destField.Type().Elem()))
		}
	}

	if srcField.Type().AssignableTo(destField.Type()) {
		destField.Set(srcField)
	} else if srcField.Type().ConvertibleTo(destField.Type()) {
		destField.Set(srcField.Convert(destField.Type()))
	}
}
