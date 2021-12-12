package collection

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type UInt64Collection struct {
	AbsCollection
	objs []uint64
}

func compareUInt64(i interface{}, i2 interface{}) int {
	int1 := i.(uint64)
	int2 := i2.(uint64)
	if int1 > int2 {
		return 1
	}
	if int1 < int2 {
		return -1
	}
	return 0
}

// NewUInt64Collection create a new UInt64Collection
func NewUInt64Collection(objs []uint64) *UInt64Collection {
	arr := &UInt64Collection{
		objs: objs,
	}
	arr.AbsCollection.Parent = arr
	arr.AbsCollection.eleType = Type_INT64
	arr.SetCompare(compareUInt64)
	return arr
}

// Copy copy collection
func (arr *UInt64Collection) Copy() ICollection {
	return NewUInt64Collection(arr.objs)
}

func (arr *UInt64Collection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(uint64); ok {
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

func (arr *UInt64Collection) Remove(i int) ICollection {
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

func (arr *UInt64Collection) NewEmpty(err ...error) ICollection {
	return NewUInt64Collection([]uint64{})
}

func (arr *UInt64Collection) Index(i int) IMix {
	if i < 0 || i >= arr.Count() {
		return NewErrorMix(errors.New("index exceeded"))
	}
	return NewMix(arr.objs[i]).SetCompare(arr.compare)
}

func (arr *UInt64Collection) SetIndex(i int, val interface{}) ICollection {
	if i < 0 || i >= arr.Count() {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs[i] = val.(uint64)
	return arr
}
func (arr *UInt64Collection) Count() int {
	return len(arr.objs)
}

func (arr *UInt64Collection) DD() {
	ret := fmt.Sprintf("UInt64Collection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%d\n", k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

func (arr *UInt64Collection) ToJson() ([]byte, error) {
	return json.Marshal(arr.objs)
}

func (arr *UInt64Collection) FromJson(data []byte) error {
	return json.Unmarshal(data, &(arr.objs))
}
