package collection

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Int32Collection struct {
	AbsCollection
	objs []int32
}

func compareInt32(i interface{}, i2 interface{}) int {
	int1 := i.(int32)
	int2 := i2.(int32)
	if int1 > int2 {
		return 1
	}
	if int1 < int2 {
		return -1
	}
	return 0
}

// NewInt32Collection create a new Int32Collection
func NewInt32Collection(objs []int32) *Int32Collection {
	arr := &Int32Collection{
		objs: objs,
	}
	arr.AbsCollection.Parent = arr
	arr.AbsCollection.eleType = Type_INT32
	arr.SetCompare(compareInt32)
	return arr
}

// Copy copy collection
func (arr *Int32Collection) Copy() ICollection {
	return NewInt32Collection(arr.objs)
}

func (arr *Int32Collection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(int32); ok {
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

func (arr *Int32Collection) Remove(i int) ICollection {
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

func (arr *Int32Collection) NewEmpty(err ...error) ICollection {
	return NewInt32Collection([]int32{})
}

func (arr *Int32Collection) Index(i int) IMix {
	if i < 0 || i >= arr.Count() {
		return NewErrorMix(errors.New("index exceeded"))
	}
	return NewMix(arr.objs[i]).SetCompare(arr.compare)
}
func (arr *Int32Collection) SetIndex(i int, val interface{}) ICollection {
	if i < 0 || i >= arr.Count() {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs[i] = val.(int32)
	return arr
}

func (arr *Int32Collection) Count() int {
	return len(arr.objs)
}

func (arr *Int32Collection) DD() {
	ret := fmt.Sprintf("Int32Collection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%d\n", k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

func (arr *Int32Collection) ToJson() ([]byte, error) {
	return json.Marshal(arr.objs)
}

func (arr *Int32Collection) FromJson(data []byte) error {
	return json.Unmarshal(data, &(arr.objs))
}
