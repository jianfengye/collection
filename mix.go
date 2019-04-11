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

	Add(mix IMix) (IMix, error) // 加法操作
	Sub(mix IMix) (IMix, error) // 减法操作
	Div(n int) (IMix, error) // 除法操作
	Multi(n int) (IMix, error) // 乘法操作

	ToString() (string, error)
	ToInt64() (int64, error)
	ToInt() (int, error)
	ToFloat64() (float64, error)
	ToFloat32() (float32, error)
	ToInterface() interface{} // 所有函数可用

	Format() string // 打印成string
	DD()
}

type Mix struct {
	IMix
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
			item1, err := m.ToString()
			if err != nil {
				return false
			}
			item2, err := n.ToString()
			if err != nil {
				return false
			}
			return item1 == item2
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
			item1, err := m.ToInt64()
			if err != nil {
				return false
			}
			item2, err := n.ToInt64()
			if err != nil {
				return false
			}
			return item1 == item2
		case reflect.Float64:
			item1, err := m.ToFloat64()
			if err != nil {
				return false
			}
			item2, err := n.ToFloat64()
			if err != nil {
				return false
			}
			return item1 == item2
		case reflect.Float32:
			item1, err := m.ToFloat32()
			if err != nil {
				return false
			}
			item2, err := n.ToFloat32()
			if err != nil {
				return false
			}
			return item1 == item2
		default:
			panic("Mix.Equal: not support kind")
		}
	}
	return false
}

func (m *Mix) Add(n IMix) (IMix, error) {
	panic("not implement")
}

func (m *Mix) Sub(n IMix) (IMix, error) {
	panic("not implement")
}

func (m *Mix) Div(n int) (IMix, error) {
	panic("not implement")
}

func (m *Mix) Multi(n int) (IMix, error) {
	panic("not implement")
}

func (m *Mix) ToString() (string, error){
	if ret, ok := m.real.(string); ok{
		return ret, nil
	}
	return "", errors.New("Mix can not covert to string")
}

func (m *Mix) ToInt64() (int64, error) {
	if ret, ok := m.real.(int64); ok {
		return ret, nil
	}
	return 0, errors.New("Mix can not covert to int64")
}

func (m *Mix) ToInt() (int, error) {
	if ret, ok := m.real.(int); ok {
		return ret, nil
	}
	return 0, errors.New("Mix can not convert to int")
}

func (m *Mix) ToFloat64() (float64, error) {
	if ret, ok := m.real.(float64); ok {
		return ret, nil
	}
	return 0, errors.New("Mix can not covert to float64")
}

func (m *Mix) ToFloat32() (float32, error) {
	if ret, ok := m.real.(float32); ok {
		return ret, nil
	}
	return 0, errors.New("Mix can not covert to float32")
}


func (m *Mix) ToInterface() interface{} {
	return m.real
}

func (m *Mix) Format() string {
	return fmt.Sprintf("%v", m.ToInterface())
}

func (m *Mix) DD() {
	ret := fmt.Sprintf("IMix(%s): %v \n", m.typ.Kind(), m.real)
	fmt.Print(ret)
}