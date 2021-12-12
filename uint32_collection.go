package collection

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type UInt32Collection struct {
	AbsCollection
	objs []uint32
}

func compareUInt32(i interface{}, i2 interface{}) int {
	int1 := i.(uint32)
	int2 := i2.(uint32)
	if int1 > int2 {
		return 1
	}
	if int1 < int2 {
		return -1
	}
	return 0
}

// NewUInt32Collection create a new UInt32Collection
func NewUInt32Collection(objs []uint32) *UInt32Collection {
	arr := &UInt32Collection{
		objs: objs,
	}
	arr.AbsCollection.Parent = arr
	arr.AbsCollection.eleType = Type_INT32
	arr.SetCompare(compareUInt32)
	return arr
}

// Copy copy collection
func (arr *UInt32Collection) Copy() ICollection {
	return NewUInt32Collection(arr.objs)
}

func (arr *UInt32Collection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(uint32); ok {
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

func (arr *UInt32Collection) Remove(i int) ICollection {
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

func (arr *UInt32Collection) NewEmpty(err ...error) ICollection {
	return NewUInt32Collection([]uint32{})
}

func (arr *UInt32Collection) Index(i int) IMix {
	if i < 0 || i >= arr.Count() {
		return NewErrorMix(errors.New("index exceeded"))
	}
	return NewMix(arr.objs[i]).SetCompare(arr.compare)
}
func (arr *UInt32Collection) SetIndex(i int, val interface{}) ICollection {
	if i < 0 || i >= arr.Count() {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs[i] = val.(uint32)
	return arr
}

func (arr *UInt32Collection) Count() int {
	return len(arr.objs)
}

func (arr *UInt32Collection) DD() {
	ret := fmt.Sprintf("UInt32Collection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%d\n", k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

func (arr *UInt32Collection) ToJson() ([]byte, error) {
	return json.Marshal(arr.objs)
}

func (arr *UInt32Collection) FromJson(data []byte) error {
	return json.Unmarshal(data, &(arr.objs))
}
