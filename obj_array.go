package collection

import (
	"reflect"
)

type ObjArray struct{
	IntArray
	objs reflect.Value
	typ reflect.Type
}

func NewObjArray(objs reflect.Value) *ObjArray {
	var typ reflect.Type
	if objs.Len() <= 0 {
		panic("ObjArray.NewObjArray: objs can not be empty")
	} else {
		typ = objs.Index(0).Type()
	}

	arr := &ObjArray{
		objs: objs,
		typ: typ,
	}
	arr.VArray.Parent = arr
	return arr
}

func NewObjArrayWithType(objs reflect.Value, typ reflect.Type) *ObjArray {

	arr := &ObjArray{
		objs: objs,
		typ: typ,
	}
	arr.VArray.Parent = arr
	return arr
}

func (arr *ObjArray) Append(obj interface{}) {
	reflect.Append(arr.objs, reflect.ValueOf(obj))
}


// Column return some key by column
func (arr *ObjArray) Column(key string) IArray {
	var objs IArray

	field, found := arr.typ.FieldByName(key)
	if !found  {
		panic("ObjArray.Column:key not found")
	}

	switch field.Type.Kind() {
	case reflect.String:
		objs = NewStrArray([]string{})
	case reflect.Int64:
		objs = NewInt64Array([]int64{})
	case reflect.Int:
		objs = NewIntArray([]int{})
	default:
		panic("ObjArray.Column: not support kind")
	}

	for i := 0; i < arr.objs.Len(); i++ {
		v := arr.objs.Index(i).FieldByName(key).Interface()
		objs.Append(v)
	}

	return objs
}

func (arr *ObjArray) KeyBy(key string) *Map {

	field, found := arr.typ.FieldByName(key)
	if !found  {
		panic("ObjArray.KeyBy: key not found")
	}
	m := NewEmptyMap(field.Type, arr.typ)
	for i := 0; i < arr.objs.Len(); i++ {
		v := arr.objs.Index(i).FieldByName(key).Interface()
		m.Set(v, arr.objs.Index(i).Interface())
	}
	return m
}

func (arr *ObjArray) Index(i int) *Mix {
	return NewMix(arr.objs.Index(i).Interface())
}

func (arr *ObjArray) Slice(start, end int) IArray {
	return NewObjArray(arr.objs.Slice(start, end))
}

func (arr *ObjArray) Len() int {
	return arr.objs.Len()
}