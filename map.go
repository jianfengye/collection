package collection

import (
	"fmt"
	"reflect"
)

type IMap interface {

	// 设置一个Map的key和value，如果key存在，则覆盖
	Set(key interface{}, value interface{})
	// 删除一个Map的key
	Remove(key interface{})
	// 根据key获取一个Map的value
	Get(key interface{}) *Mix
	// 获取一个Map的长度
	Len() int

	// 查询一个Value，返回第一个key
	Search(val interface{}) *Mix

	// key和value进行对调
	Flip() (IMap, error)
	// 判断是否有这个key，或者多个key
	Has(keys ...interface{}) bool

	// 获取Map的所有key组成的集合
	Keys() IArray
	// 获取Map的所有value组成的集合
	Values() IArray

	Only(keys ...interface{}) IMap

	// 根据Key进行升序排列
	SortKeys() IMap
	// 根据key进行降序排列
	SortKeysDesc() IMap

	DD()
}

type Map struct {
	IMap

	objs map[*Mix]*Mix

	keyType reflect.Type
	valType reflect.Type
}

func NewEmptyMap(key reflect.Type, val reflect.Type, compare ...func(key1, key2 interface{}) bool) *Map {
	m := make(map[*Mix]*Mix)
	return &Map{
		objs: m,
		keyType: key,
		valType: val,
	}
}



func (m *Map) mustBeKeyType(key interface{}) {
	if reflect.TypeOf(key) != m.keyType {
		panic("key type wrong")
	}
}

func (m *Map) mustBeValueType(val interface{}) {
	if reflect.TypeOf(val) != m.valType {
		panic("val type wrong")
	}
}

func (m *Map) Set(key interface{}, value interface{}) {
	m.mustBeKeyType(key)
	m.mustBeValueType(value)
	k := NewMix(key)
	v := NewMix(value)

	// 不管有没有，都设置v
	m.objs[k] = v
}

func (m *Map) Get(key interface{}) *Mix {
	m.mustBeKeyType(key)
	kParam := NewMix(key)
	for k, v := range m.objs {
		if k.Equal(kParam) {
			return v
		}
	}
	return nil
}


func (m *Map) Search(value interface{}) IMix {
	panic("not implement")
}

func (m *Map) Remove(key interface{}) {
	m.mustBeKeyType(key)
	kParam := NewMix(key)

	for k := range m.objs {
		if k.Equal(kParam) {
			delete(m.objs, k)
		}
	}
}

func (m *Map) Len() int {
	return len(m.objs)
}

func (m *Map) Keys() IArray {
	var objs IArray
	switch m.keyType.Kind() {
	case reflect.String:
		objs = NewStrArray([]string{})
	case reflect.Int64:
		objs = NewInt64Array([]int64{})
	case reflect.Int:
		objs = NewIntArray([]int{})
	default:
		panic("ObjArray.Column: not support kind")
	}

	for k := range m.objs {
		objs.Append(k)
	}
	return objs
}

func (m *Map) Values() IArray {
	var objs IArray
	switch m.valType.Kind() {
	case reflect.String:
		objs = NewStrArray([]string{})
	case reflect.Int64:
		objs = NewInt64Array([]int64{})
	case reflect.Int:
		objs = NewIntArray([]int{})
	default:
		panic("ObjArray.Column: not support kind")
	}

	for _, v := range m.objs {
		objs.Append(v)
	}
	return objs
}

func (m *Map) DD() {
	ret := fmt.Sprintf("Map(%d):{\n", len(m.objs))
	for k, v := range m.objs {
		ret = ret + fmt.Sprintf("\t%s:\t%s\n",k.Format(), v.Format())
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

