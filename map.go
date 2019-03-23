package collection

import "reflect"

type IMap interface {

	// 设置一个Map的key和value，如果key存在，则覆盖
	Set(key interface{}, value interface{})
	// 删除一个Map的key
	Remove(key interface{})
	// 根据key获取一个Map的value
	Get(key interface{}) *Mix
	// 获取一个Map的长度
	Len() int

	// 获取Map的所有key组成的集合
	Keys() IArray
	// 获取Map的所有value组成的集合
	Values() IArray
}

type Map struct {
	IMap

	objs map[*Mix]*Mix

	keyType reflect.Type
	valType reflect.Type
}

func NewEmptyMap(key, val reflect.Type) *Map {
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

