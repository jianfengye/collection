package collection

import (
	"errors"
	"fmt"
	"reflect"
	"sort"
)

type ObjCollection struct{
	AbsCollection
	objs reflect.Value // 数组对象，是一个slice
	typ reflect.Type // 数组对象每个元素类型
}

// 根据对象数组创建
func NewObjCollection(objs interface{}) *ObjCollection {

	vals := reflect.ValueOf(objs)
	typ := reflect.TypeOf(objs).Elem()

	objs2 := reflect.MakeSlice(reflect.TypeOf(objs), vals.Len(), vals.Len())
	reflect.Copy(objs2, reflect.ValueOf(objs))

	arr := &ObjCollection{
		objs: objs2,
		typ: typ,
	}
	arr.AbsCollection.Parent = arr
	return arr
}

// 根据类型创建一个空的数组
func NewObjCollectionByType(typ reflect.Type) *ObjCollection {
	vals := reflect.MakeSlice(typ, 0, 0)
	arr := &ObjCollection{
		objs: vals,
		typ: typ,
	}
	arr.AbsCollection.Parent = arr
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
	return NewMix(arr.objs.Index(i).Interface())
}

func (arr *ObjCollection) NewEmpty(err ...error) ICollection {
	objs := reflect.MakeSlice(arr.objs.Type(), 0, 0)
	ret := &ObjCollection{
		objs: objs,
		typ: arr.typ,
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
	tail := arr.objs.Slice(i + 1, length)
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
	for i:= 0; i< arr.objs.Len(); i++ {
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
	if !found  {
		err := errors.New("ObjCollection.Pluck:key not found")
		arr.SetErr(err)
		return arr
	}

	switch field.Type.Kind() {
	case reflect.String:
		objs = NewStrCollection([]string{})
	case reflect.Int64:
		objs = NewInt64Collection([]int64{})
	case reflect.Int:
		objs = NewIntCollection([]int{})
	case reflect.Float32:
		objs = NewFloat32Collection([]float32{})
	case reflect.Float64:
		objs = NewFloat64Collection([]float64{})
	default:
		err := errors.New("ObjCollection.Pluck: not support kind")
		arr.SetErr(err)
		return arr
	}

	for i := 0; i < arr.objs.Len(); i++ {
		v := arr.objs.Index(i).FieldByName(key).Interface()
		objs.Append(v)
	}

	return objs
}


type ByFieldSort ObjCollection
func (a ByFieldSort) Len() int { return a.objs.Len() }
func (a ByFieldSort) Swap(i, j int) {
	t := a.objs.Index(i).Interface()
	a.objs.Index(i).Set(a.objs.Index(j))
	a.objs.Index(j).Set(reflect.ValueOf(t))
}
func (a ByFieldSort) Less(i, j int) bool {
	iInterface := a.objs.Index(i).Interface()
	jInterface := a.objs.Index(j).Interface()

	if a.compare == nil {
		panic("Less compare does not exist")
	}

	if  a.compare(iInterface, jInterface) < 0 {
		return true
	}
	return false
}

type ByFieldSortDesc ObjCollection
func (a ByFieldSortDesc) Len() int { return a.objs.Len() }
func (a ByFieldSortDesc) Swap(i, j int) {
	t := a.objs.Index(i).Interface()
	a.objs.Index(i).Set(a.objs.Index(j))
	a.objs.Index(j).Set(reflect.ValueOf(t))
}
func (a ByFieldSortDesc) Less(i, j int) bool {
	iInterface := a.objs.Index(i).Interface()
	jInterface := a.objs.Index(j).Interface()

	if a.compare == nil {
		panic("Less compare does not exist")
	}

	if  a.compare(iInterface, jInterface) > 0 {
		return true
	}
	return false
}

// 按照某个字段进行排序
func (arr *ObjCollection) SortBy(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	sort.Sort(ByFieldSort(*arr))
	return arr
}

// 按照某个字段进行排序,倒序
func (arr *ObjCollection) SortByDesc(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	sort.Sort(ByFieldSortDesc(*arr))
	return arr
}