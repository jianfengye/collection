package collection

import (
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

type Float64Collection struct {
	AbsCollection
	objs []float64
}

func compareFloat64(i interface{}, i2 interface{}) int {
	int1 := i.(float64)
	int2 := i2.(float64)
	if int1 > int2 {
		return 1
	}
	if int1 < int2 {
		return -1
	}
	return 0
}

// NewFloat64Collection create a new Float64Collection
func NewFloat64Collection(objs []float64) *Float64Collection {
	arr := &Float64Collection{
		objs: objs,
	}
	arr.AbsCollection.Parent = arr
	arr.AbsCollection.eleType = TYPE_FLOAT64
	arr.SetCompare(compareFloat64)
	return arr
}

// Copy copy collection
func (arr *Float64Collection) Copy() ICollection {
	return NewFloat64Collection(arr.objs)
}

func (arr *Float64Collection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(float64); ok {
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

func (arr *Float64Collection) Remove(i int) ICollection {
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

func (arr *Float64Collection) NewEmpty(err ...error) ICollection {
	return NewFloat64Collection([]float64{})
}

func (arr *Float64Collection) Index(i int) IMix {
	if i < 0 || i >= arr.Count() {
		return NewErrorMix(errors.New("index exceeded"))
	}
	return NewMix(arr.objs[i]).SetCompare(arr.compare)
}

func (arr *Float64Collection) SetIndex(i int, val interface{}) ICollection {
	if i < 0 || i >= arr.Count() {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs[i] = val.(float64)
	return arr
}

func (arr *Float64Collection) Count() int {
	return len(arr.objs)
}

func (arr *Float64Collection) DD() {
	ret := fmt.Sprintf("Float64Collection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%f\n", k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}

func (arr *Float64Collection) ToJson() ([]byte, error) {
	return json.Marshal(arr.objs)
}
func (arr *Float64Collection) FromJson(data []byte) error {
	return json.Unmarshal(data, &(arr.objs))
}
