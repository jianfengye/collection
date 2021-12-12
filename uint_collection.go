package collection

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type UIntCollection struct {
	AbsCollection
	objs []uint
}

func compareUInt(i interface{}, i2 interface{}) int {
	int1 := i.(int)
	int2 := i2.(int)
	if int1 > int2 {
		return 1
	}
	if int1 < int2 {
		return -1
	}
	return 0
}

// NewUIntCollection create a new IntCollection
func NewUIntCollection(objs []uint) *UIntCollection {
	arr := &UIntCollection{
		objs: objs,
	}
	arr.AbsCollection.Parent = arr
	arr.AbsCollection.eleType = Type_INT
	arr.SetCompare(compareInt)
	return arr
}

// Copy copy collection
func (arr *UIntCollection) Copy() ICollection {
	return NewUIntCollection(arr.objs)
}

func (arr *UIntCollection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(uint); ok {
		length := len(arr.objs)

		// 如果是append操作，直接调用系统的append，不新创建collection
		if index >= length {
			arr.objs = append(arr.objs, i)
			return arr
		}

		arr.objs = append(arr.objs, 0)
		copy(arr.objs[index+1:], arr.objs[index:])
		arr.objs[index] = i
	} else {
		return arr.SetErr(errors.New("Insert: type error"))
	}
	return arr
}

func (arr *UIntCollection) Remove(i int) ICollection {
	if arr.Err() != nil {
		return arr
	}

	len := arr.Count()
	if i < 0 || i >= len {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs = append(arr.objs[0:i], arr.objs[i+1:len]...)
	return arr
}

func (arr *UIntCollection) NewEmpty(err ...error) ICollection {
	return NewIntCollection([]int{})
}

func (arr *UIntCollection) Index(i int) IMix {
	if i < 0 || i >= arr.Count() {
		return NewErrorMix(errors.New("index exceeded"))
	}
	return NewMix(arr.objs[i]).SetCompare(arr.compare)
}

func (arr *UIntCollection) SetIndex(i int, val interface{}) ICollection {
	if i < 0 || i >= arr.Count() {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs[i] = val.(uint)
	return arr
}

func (arr *UIntCollection) Count() int {
	return len(arr.objs)
}

func (arr *UIntCollection) DD() {
	ret := fmt.Sprintf("UIntCollection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%d\n", k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

func (arr *UIntCollection) ToJson() ([]byte, error) {
	return json.Marshal(arr.objs)
}

func (arr *UIntCollection) FromJson(data []byte) error {
	return json.Unmarshal(data, &(arr.objs))
}
