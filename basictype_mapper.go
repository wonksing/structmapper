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
	srcField := srcValue.FieldByIndex(srcIndex)
	destField := destValue.FieldByIndex(destIndex)

	if srcField.Type().AssignableTo(destField.Type()) {
		destField.Set(srcField)
	} else if srcField.Type().ConvertibleTo(destField.Type()) {
		destField.Set(srcField.Convert(destField.Type()))
	}
}
