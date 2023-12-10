package structmapper

import "reflect"

type StructTypeMapper struct {
	SrcIndex  []int
	DestIndex []int

	structMapper *StructMapper
}

func (c *StructTypeMapper) Map(srcValue, destValue reflect.Value) {
	if srcValue.Type().Kind() == reflect.Pointer {
		if srcValue.IsNil() {
			return
		}

		if destValue.Type().Kind() == reflect.Pointer && destValue.IsNil() {
			destValue.Set(reflect.New(destValue.Type().Elem()))
		}
		c.structMapper.Map(srcValue.Interface(), destValue.Interface())

	} else {
		c.structMapper.Map(srcValue.Addr().Interface(), destValue.Addr().Interface())
	}
}
