package collection

import "reflect"

type Mix struct {
	real interface{}
	typ reflect.Type
}

func NewMix(real interface{}) *Mix {
	return &Mix{
		real: real,
		typ: reflect.TypeOf(real),
	}
}

// Equal 判断两个Mix是否相等
func (m *Mix) Equal(n *Mix) bool {
	if m.typ == n.typ {
		switch m.typ.Kind() {
		case reflect.String:
			return m.ToString() == n.ToString()
		case reflect.Int:
			return m.ToInt() == n.ToInt()
		case reflect.Int64:
			return m.ToInt64() == n.ToInt64()
		case reflect.Float64:
			return m.ToFloat64() == n.ToFloat64()
		case reflect.Float32:
			return m.ToFloat32() == n.ToFloat32()
		default:
			panic("Mix.Equal: not support kind")
		}
	}
	return false
}

func (m *Mix) ToString() string {
	if ret, ok := m.real.(string); ok{
		return ret
	}
	panic("Mix can not covert to string")
}

func (m *Mix) ToInt64() int64 {
	if ret, ok := m.real.(int64); ok {
		return ret
	}
	panic("Mix can not convert to int64")
}

func (m *Mix) ToInt() int {
	if ret, ok := m.real.(int); ok {
		return ret
	}
	panic("Mix can not convert to int")
}

func (m *Mix) ToFloat64() float64 {
	if ret, ok := m.real.(float64); ok {
		return ret
	}
	panic("Mix can not convert to float64")
}

func (m *Mix) ToFloat32() float32 {
	if ret, ok := m.real.(float32); ok {
		return ret
	}
	panic("Mix can not convert to float64")
}


func (m *Mix) ToInterface() interface{} {
	return m.real
}

