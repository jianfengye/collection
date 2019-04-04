package collection

import (
	"fmt"
	"github.com/pkg/errors"
	"reflect"
)

// IMix是一个混合结构
type IMix interface {
	Equal(n IMix) bool // 两个IMix结构是否相同
	Type() reflect.Type // 获取类型

	ToString() string
	ToInt64() int64
	ToInt() (int, error)
	ToFloat64() float64
	ToFloat32() float32
	ToInterface() interface{} // 所有函数可用

	Format() string // 打印成string
}

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

func (m *Mix)Type() reflect.Type {
	return m.typ
}

// Equal 判断两个Mix是否相等
func (m *Mix) Equal(n IMix) bool {
	if m.typ == reflect.TypeOf(n) {
		switch m.typ.Kind() {
		case reflect.String:
			return m.ToString() == n.ToString()
		case reflect.Int:
			int1, err := m.ToInt()
			if err != nil {
				return false
			}
			int2, err := n.ToInt()
			if err != nil {
				return false
			}
			return int1 == int2
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

func (m *Mix) ToInt() (int, error) {
	if ret, ok := m.real.(int); ok {
		return ret, nil
	}
	return 0, errors.New("Mix can not convert to int")
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

func (m *Mix) Format() string {
	return fmt.Sprintf("%v", m.ToInterface())
}