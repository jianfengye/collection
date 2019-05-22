package collection

import (
	"errors"
	"fmt"
	"reflect"
)

// ObjCollection 代表数组集合
type ObjCollection struct {
	AbsCollection
	objs reflect.Value // 数组对象，是一个slice
	typ  reflect.Type  // 数组对象每个元素类型
}

// NewObjCollection 根据对象数组创建
func NewObjCollection(objs interface{}) *ObjCollection {

	vals := reflect.ValueOf(objs)
	typ := reflect.TypeOf(objs).Elem()

	arr := &ObjCollection{
		objs: vals,
		typ:  typ,
	}
	arr.AbsCollection.Parent = arr
	return arr
}

// NewObjCollectionByType 根据类型创建一个空的数组
func NewObjCollectionByType(typ reflect.Type) *ObjCollection {
	vals := reflect.MakeSlice(typ, 0, 0)
	arr := &ObjCollection{
		objs: vals,
		typ:  typ,
	}
	arr.AbsCollection.Parent = arr
	return arr
}

// Copy 复制到新的数组
func (arr *ObjCollection) Copy() ICollection {

	objs2 := reflect.MakeSlice(arr.objs.Type(), arr.objs.Len(), arr.objs.Len())
	reflect.Copy(objs2, arr.objs)
	arr.objs = objs2

	return arr
}

func (arr *ObjCollection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}

	ret := arr.objs.Slice(0, index)
	length := arr.objs.Len()
	tail := arr.objs.Slice(index, length)
	ret = reflect.Append(ret, reflect.ValueOf(obj))
	for i := 0; i < tail.Len(); i++ {
		ret = reflect.Append(ret, tail.Index(i))
	}
	arr.objs = ret
	arr.AbsCollection.Parent = arr
	return arr
}

func (arr *ObjCollection) Index(i int) IMix {
	return NewMix(arr.objs.Index(i).Interface()).SetCompare(arr.compare)
}

func (arr *ObjCollection) SetIndex(i int, val interface{}) ICollection {
	arr.objs.Index(i).Set(reflect.ValueOf(val))
	return arr
}

func (arr *ObjCollection) NewEmpty(err ...error) ICollection {
	objs := reflect.MakeSlice(arr.objs.Type(), 0, 0)
	ret := &ObjCollection{
		objs: objs,
		typ:  arr.typ,
	}
	ret.AbsCollection.Parent = ret
	if len(err) != 0 {
		ret.SetErr(err[0])
	}
	return ret
}

func (arr *ObjCollection) Remove(i int) ICollection {
	if arr.Err() != nil {
		return arr
	}

	len := arr.Count()
	if i >= len {
		return arr.SetErr(errors.New("index exceeded"))
	}

	ret := arr.objs.Slice(0, i)
	length := arr.objs.Len()
	tail := arr.objs.Slice(i+1, length)
	for i := 0; i < tail.Len(); i++ {
		ret = reflect.Append(ret, tail.Index(i))
	}
	arr.objs = ret
	arr.AbsCollection.Parent = arr
	return arr
}

func (arr *ObjCollection) Count() int {
	return arr.objs.Len()
}

func (arr *ObjCollection) DD() {
	ret := fmt.Sprintf("ObjCollection(%d)(%s):{\n", arr.Count(), arr.typ.String())
	for i := 0; i < arr.objs.Len(); i++ {
		ret = ret + fmt.Sprintf("\t%d:\t%+v\n", i, arr.objs.Index(i))
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

// 将对象的某个key作为Slice的value，作为slice返回
func (arr *ObjCollection) Pluck(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	var objs ICollection

	field, found := arr.typ.FieldByName(key)
	if !found {
		err := errors.New("ObjCollection.Pluck:key not found")
		arr.SetErr(err)
		return arr
	}

	objs = NewMixCollection(field.Type)
	for i := 0; i < arr.objs.Len(); i++ {
		v := arr.objs.Index(i).FieldByName(key).Interface()
		objs.Append(v)
	}

	return objs
}

// 按照某个字段进行排序
func (arr *ObjCollection) SortBy(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	compare := func(a interface{}, b interface{}) int {
		mixA := NewMix(reflect.ValueOf(a).FieldByName(key).Interface())
		mixB := NewMix(reflect.ValueOf(b).FieldByName(key).Interface())
		ret, _ := mixA.Compare(mixB)
		return ret
	}

	oldCompare := arr.compare
	arr.compare = compare
	newArr := arr.Sort()
	newArr.SetCompare(oldCompare)
	return newArr
}

// 按照某个字段进行排序,倒序
func (arr *ObjCollection) SortByDesc(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	compare := func(a interface{}, b interface{}) int {
		mixA := NewMix(reflect.ValueOf(a).FieldByName(key).Interface())
		mixB := NewMix(reflect.ValueOf(b).FieldByName(key).Interface())
		ret, _ := mixB.Compare(mixA)
		return ret
	}

	oldCompare := arr.compare
	arr.compare = compare
	newArr := arr.Sort()
	newArr.SetCompare(oldCompare)
	return newArr
}
