package collection

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type IntCollection struct {
	AbsCollection
	objs []int
}

func compareInt(i interface{}, i2 interface{}) int {
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

// NewIntCollection create a new IntCollection
func NewIntCollection(objs []int) *IntCollection {
	arr := &IntCollection{
		objs: objs,
	}
	arr.AbsCollection.Parent = arr
	arr.AbsCollection.eleType = Type_INT
	arr.SetCompare(compareInt)
	return arr
}

// Copy copy collection
func (arr *IntCollection) Copy() ICollection {
	return NewIntCollection(arr.objs)
}

func (arr *IntCollection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(int); ok {
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

func (arr *IntCollection) Remove(i int) ICollection {
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

func (arr *IntCollection) NewEmpty(err ...error) ICollection {
	return NewIntCollection([]int{})
}

func (arr *IntCollection) Index(i int) IMix {
	if i < 0 || i >= arr.Count() {
		return NewErrorMix(errors.New("index exceeded"))
	}
	return NewMix(arr.objs[i]).SetCompare(arr.compare)
}

func (arr *IntCollection) SetIndex(i int, val interface{}) ICollection {
	if i < 0 || i >= arr.Count() {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs[i] = val.(int)
	return arr
}

func (arr *IntCollection) Count() int {
	return len(arr.objs)
}

func (arr *IntCollection) DD() {
	ret := fmt.Sprintf("IntCollection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%d\n", k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

func (arr *IntCollection) ToJson() ([]byte, error) {
	return json.Marshal(arr.objs)
}

func (arr *IntCollection) FromJson(data []byte) error {
	return json.Unmarshal(data, &(arr.objs))
}
