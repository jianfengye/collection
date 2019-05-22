package collection

import (
	"errors"
	"fmt"
	"reflect"
)

// IMix是一个混合结构
type IMix interface {
	Err() error
	SetErr(err error) IMix

	Compare(n IMix) (int, error) // 两个IMix结构是否相同
	Type() reflect.Type          // 获取类型
	SetCompare(func(a interface{}, b interface{}) int) IMix

	Add(mix IMix) (IMix, error) // 加法操作
	Sub(mix IMix) (IMix, error) // 减法操作
	Div(n int) (IMix, error)    // 除法操作
	Multi(n int) (IMix, error)  // 乘法操作

	ToString() (string, error)
	ToInt64() (int64, error)
	ToInt32() (int32, error)
	ToInt() (int, error)
	ToFloat64() (float64, error)
	ToFloat32() (float32, error)
	ToInterface() (interface{}, error) // 所有函数可用

	Format() string // 打印成string
	DD()
}

type Mix struct {
	IMix

	err  error
	real interface{}
	typ  reflect.Type

	compare func(interface{}, interface{}) int // 比较函数
}

func NewErrorMix(err error) *Mix {
	mix := &Mix{}
	mix.SetErr(err)
	return mix
}

func NewMix(real interface{}) *Mix {
	m := &Mix{
		real: real,
		typ:  reflect.TypeOf(real),
	}
	switch m.typ.Kind() {
	case reflect.String:
		m.compare = compareString
	case reflect.Int:
		m.compare = compareInt
	case reflect.Int32:
		m.compare = compareInt32
	case reflect.Int64:
		m.compare = compareInt64
	case reflect.Float32:
		m.compare = compareFloat32
	case reflect.Float64:
		m.compare = compareFloat64
	}
	return m
}

// 根据typ 创建一个新的Mix数组
func NewMixCollection(typ reflect.Type) ICollection {
	switch typ.Kind() {
	case reflect.String:
		return NewStrCollection([]string{})
	case reflect.Int:
		return NewIntCollection([]int{})
	case reflect.Int64:
		return NewInt64Collection([]int64{})
	case reflect.Int32:
		return NewInt32Collection([]int32{})
	case reflect.Float32:
		return NewFloat32Collection([]float32{})
	case reflect.Float64:
		return NewFloat64Collection([]float64{})
	case reflect.Ptr:
		return NewObjCollectionByType(typ)
	}
	return nil
}

func (m *Mix) Err() error {
	return m.err
}

func (m *Mix) SetErr(err error) IMix {
	m.err = err
	return m
}

func (m *Mix) Type() reflect.Type {
	return m.typ
}

// Equal 判断两个Mix是否相等
func (m *Mix) Compare(n IMix) (ret int, err error) {
	if m.typ != n.Type() {
		return 0, errors.New("type not match")
	}
	if m.compare == nil {
		return 0, errors.New("compare does not exist in Mix")
	}

	mR, err := m.ToInterface()
	if err != nil {
		return 0, err
	}
	nR, err := n.ToInterface()
	if err != nil {
		return 0, err
	}

	return m.compare(mR, nR), nil
}

func (m *Mix) SetCompare(compare func(a interface{}, b interface{}) int) IMix {
	m.compare = compare
	return m
}

func (m *Mix) Add(n IMix) (IMix, error) {
	if m.Err() != nil {
		return m, m.Err()
	}
	switch m.typ.Kind() {
	case reflect.String:
		item1, err := m.ToString()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToString()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 + item2), nil
	case reflect.Int:
		item1, err := m.ToInt()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToInt()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 + item2), nil
	case reflect.Int64:
		item1, err := m.ToInt64()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToInt64()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 + item2), nil
	case reflect.Int32:
		item1, err := m.ToInt32()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToInt32()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 + item2), nil
	case reflect.Float64:
		item1, err := m.ToFloat64()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToFloat64()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 + item2), nil
	case reflect.Float32:
		item1, err := m.ToFloat32()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToFloat32()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 + item2), nil
	default:
		return nil, errors.New("format not support")
	}
}

func (m *Mix) Sub(n IMix) (IMix, error) {
	if m.Err() != nil {
		return m, m.Err()
	}
	switch m.typ.Kind() {
	case reflect.String:
		return nil, errors.New("format not support")
	case reflect.Int:
		item1, err := m.ToInt()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToInt()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 - item2), nil
	case reflect.Int64:
		item1, err := m.ToInt64()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToInt64()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 - item2), nil
	case reflect.Int32:
		item1, err := m.ToInt32()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToInt32()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 - item2), nil
	case reflect.Float64:
		item1, err := m.ToFloat64()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToFloat64()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 - item2), nil
	case reflect.Float32:
		item1, err := m.ToFloat32()
		if err != nil {
			return nil, errors.New("format error")
		}
		item2, err := n.ToFloat32()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 - item2), nil
	default:
		return nil, errors.New("format not support")
	}
}

func (m *Mix) Div(n int) (IMix, error) {
	if m.Err() != nil {
		return m, m.Err()
	}
	switch m.typ.Kind() {
	case reflect.String:
		return nil, errors.New("format not support")
	case reflect.Int:
		item1, err := m.ToInt()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(float64(item1) / float64(n)), nil
	case reflect.Int64:
		item1, err := m.ToInt64()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(float64(item1) / float64(n)), nil
	case reflect.Int32:
		item1, err := m.ToInt32()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(float64(item1) / float64(n)), nil
	case reflect.Float64:
		item1, err := m.ToFloat64()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 - float64(n)), nil
	case reflect.Float32:
		item1, err := m.ToFloat32()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 / float32(n)), nil
	default:
		return nil, errors.New("format not support")
	}
}

func (m *Mix) Multi(n int) (IMix, error) {
	if m.Err() != nil {
		return m, m.Err()
	}
	switch m.typ.Kind() {
	case reflect.String:
		return nil, errors.New("format not support")
	case reflect.Int:
		item1, err := m.ToInt()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 * n), nil
	case reflect.Int64:
		item1, err := m.ToInt64()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 * int64(n)), nil
	case reflect.Int32:
		item1, err := m.ToInt32()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 * int32(n)), nil
	case reflect.Float64:
		item1, err := m.ToFloat64()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 * float64(n)), nil
	case reflect.Float32:
		item1, err := m.ToFloat32()
		if err != nil {
			return nil, errors.New("format error")
		}
		return NewMix(item1 * float32(n)), nil
	default:
		return nil, errors.New("format not support")
	}
}

func (m *Mix) ToString() (string, error) {
	if m.err != nil {
		return "", m.err
	}
	if ret, ok := m.real.(string); ok {
		return ret, nil
	}
	return "", errors.New("Mix can not covert to string")
}

func (m *Mix) ToInt64() (int64, error) {
	if m.err != nil {
		return 0, m.err
	}
	if ret, ok := m.real.(int64); ok {
		return ret, nil
	}
	return 0, errors.New("Mix can not covert to int64")
}

func (m *Mix) ToInt32() (int32, error) {
	if m.err != nil {
		return 0, m.err
	}
	if ret, ok := m.real.(int32); ok {
		return ret, nil
	}
	return 0, errors.New("Mix can not covert to int32")
}

func (m *Mix) ToInt() (int, error) {
	if m.err != nil {
		return 0, m.err
	}
	if ret, ok := m.real.(int); ok {
		return ret, nil
	}
	return 0, errors.New("Mix can not convert to int")
}

func (m *Mix) ToFloat64() (float64, error) {
	if m.err != nil {
		return 0.0, m.err
	}
	if ret, ok := m.real.(float64); ok {
		return ret, nil
	}
	return 0, errors.New("Mix can not covert to float64")
}

func (m *Mix) ToFloat32() (float32, error) {
	if m.err != nil {
		return 0.0, m.err
	}
	if ret, ok := m.real.(float32); ok {
		return ret, nil
	}
	return 0, errors.New("Mix can not covert to float32")
}

func (m *Mix) ToInterface() (interface{}, error) {
	return m.real, m.Err()
}

func (m *Mix) Format() string {
	o, _ := m.ToInterface()
	return fmt.Sprintf("%v", o)
}

func (m *Mix) DD() {
	ret := fmt.Sprintf("IMix(%s): %+v \n", m.typ.Kind(), m.real)
	fmt.Print(ret)
}
