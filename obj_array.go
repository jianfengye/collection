package collection

import (
	"reflect"
)

type ObjArray struct{
	AbsArray
	objs reflect.Value // 数组对象，是一个slice
	typ reflect.Type // 数组对象每个元素类型
	ptr reflect.Value // 指向数组对象的指针
}

// 根据对象数组创建
func NewObjArray(objs interface{}) *ObjArray {
	vals := reflect.ValueOf(objs)
	typ := reflect.TypeOf(objs).Elem()
	arr := &ObjArray{
		objs: vals,
		typ: typ,
	}
	arr.AbsArray.Parent = arr
	return arr
}

// 根据Values和单个元素的type创建
func NewObjArrayWithValType(vals reflect.Value, typ reflect.Type) *ObjArray {
	arr := &ObjArray{
		objs: vals,
		typ: typ,
	}
	arr.AbsArray.Parent = arr
	return arr
}

func (arr *ObjArray) Append(obj interface{}) {
	arr.objs = reflect.Append(arr.objs, reflect.ValueOf(obj))
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
	return NewObjArrayWithValType(arr.objs.Slice(start, end), arr.typ)
}

func (arr *ObjArray) Len() int {
	return arr.objs.Len()
}

func (arr *ObjArray) NewEmptyIArray() IArray {
	objs := reflect.MakeSlice(arr.objs.Type(), 0, 0)
	ret := &ObjArray{
		objs: objs,
		typ: arr.typ,
	}
	ret.AbsArray.Parent = ret
	return ret
}