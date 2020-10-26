package collection

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/pkg/errors"
)

// ObjCollection 代表数组集合
type ObjPointCollection struct {
	AbsCollection
	objs reflect.Value // 数组对象，是一个slice
	typ  reflect.Type  // 数组对象每个元素类型
}

// NewObjCollection 根据对象数组创建
func NewObjPointCollection(objs interface{}) *ObjPointCollection {

	vals := reflect.ValueOf(objs)
	typ := reflect.TypeOf(objs).Elem()

	arr := &ObjPointCollection{
		objs: vals,
		typ:  typ,
	}
	arr.AbsCollection.Parent = arr
	arr.AbsCollection.eleType = TYPE_OBJ_POINT
	return arr
}

// Copy 复制到新的数组
func (arr *ObjPointCollection) Copy() ICollection {

	objs2 := reflect.MakeSlice(arr.objs.Type(), arr.objs.Len(), arr.objs.Len())
	reflect.Copy(objs2, arr.objs)
	arr2 := &ObjPointCollection{
		objs: objs2,
		typ:  arr.objs.Type(),
	}
	if arr.Err() != nil {
		arr2.SetErr(arr.Err())
	}
	if arr.compare != nil {
		arr2.compare = arr.compare
	}
	arr2.Parent = arr2
	arr2.eleType = arr.eleType

	return arr2
}

func (arr *ObjPointCollection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}

	ret := arr.objs.Slice(0, index)

	length := arr.objs.Len()
	tail := arr.objs.Slice(index, length)

	tailNew := reflect.MakeSlice(arr.objs.Type(), length-index, length-index)
	reflect.Copy(tailNew, tail)
	ret = reflect.Append(ret, reflect.ValueOf(obj))
	for i := 0; i < tail.Len(); i++ {
		ret = reflect.Append(ret, tailNew.Index(i))
	}
	arr.objs = ret
	arr.AbsCollection.Parent = arr
	return arr
}

func (arr *ObjPointCollection) Index(i int) IMix {
	return NewMix(arr.objs.Index(i).Interface()).SetCompare(arr.compare)
}

func (arr *ObjPointCollection) SetIndex(i int, val interface{}) ICollection {
	arr.objs.Index(i).Set(reflect.ValueOf(val))
	return arr
}

func (arr *ObjPointCollection) NewEmpty(err ...error) ICollection {
	objs2 := reflect.MakeSlice(arr.objs.Type(), 0, 0)
	arr2 := &ObjPointCollection{
		objs: objs2,
		typ:  arr.objs.Type(),
	}
	arr2.compare = arr.compare
	arr2.eleType = arr.eleType
	arr2.Parent = arr2

	return arr2
}

func (arr *ObjPointCollection) Remove(i int) ICollection {
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

func (arr *ObjPointCollection) Count() int {
	return arr.objs.Len()
}

func (arr *ObjPointCollection) DD() {
	if arr.Err() != nil {
		fmt.Println(arr.Err().Error())
	}
	ret := fmt.Sprintf("ObjCollection(%d)(%s):{\n", arr.Count(), arr.typ.String())
	for i := 0; i < arr.objs.Len(); i++ {
		ret = ret + fmt.Sprintf("\t%d:\t%+v\n", i, arr.objs.Index(i))
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

// 将对象的某个key作为Slice的value，作为slice返回
func (arr *ObjPointCollection) Pluck(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	var objs ICollection

	field, found := arr.typ.Elem().FieldByName(key)
	if !found {
		err := errors.New("ObjCollection.Pluck:key not found")
		arr.SetErr(err)
		return arr
	}

	objs = NewMixCollection(field.Type)
	for i := 0; i < arr.objs.Len(); i++ {
		v := arr.objs.Index(i).Elem().FieldByName(key).Interface()
		objs.Append(v)
	}

	return objs
}

// 按照某个字段进行排序
func (arr *ObjPointCollection) SortBy(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	compare := func(a interface{}, b interface{}) int {
		mixA := NewMix(reflect.ValueOf(a).Elem().FieldByName(key).Interface())
		mixB := NewMix(reflect.ValueOf(b).Elem().FieldByName(key).Interface())
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
func (arr *ObjPointCollection) SortByDesc(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	compare := func(a interface{}, b interface{}) int {
		mixA := NewMix(reflect.ValueOf(a).Elem().FieldByName(key).Interface())
		mixB := NewMix(reflect.ValueOf(b).Elem().FieldByName(key).Interface())
		ret, _ := mixB.Compare(mixA)
		return ret
	}

	oldCompare := arr.compare
	arr.compare = compare
	newArr := arr.Sort()
	newArr.SetCompare(oldCompare)
	return newArr
}

func (arr *ObjPointCollection) ToJson() ([]byte, error) {
	return json.Marshal(arr.objs.Interface())
}

func (arr *ObjPointCollection) FromJson(data []byte) error {
	vals := reflect.MakeSlice(reflect.SliceOf(arr.typ), 0, 0)
	objs := vals.Interface()
	err := json.Unmarshal(data, &objs)
	if err != nil {
		return err
	}
	arr.objs = reflect.ValueOf(objs)
	return nil
}

func (arr *ObjPointCollection) ToObjs(objs interface{}) error {
	arr.mustNotBeBaseType()

	objVal := reflect.ValueOf(objs)
	if objVal.Elem().CanSet() {
		objVal.Elem().Set(arr.objs)
		return nil
	}
	return errors.New("element should be can set")
}
