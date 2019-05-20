package collection

import (
	"errors"
	"fmt"
)

type Float32Collection struct {
	AbsCollection
	objs []float32
}

func compareFloat32(i interface{}, i2 interface{}) int {
	int1 := i.(float32)
	int2 := i2.(float32)
	if int1 > int2 {
		return 1
	}
	if int1 < int2 {
		return -1
	}
	return 0
}

// NewFloat32Collection create a new Float32Collection
func NewFloat32Collection(objs []float32) *Float32Collection {
	arr := &Float32Collection{
		objs: objs,
	}
	arr.AbsCollection.Parent = arr
	arr.SetCompare(compareFloat32)
	return arr
}

// Copy copy collection
func (arr *Float32Collection) Copy() ICollection {
	objs2 := make([]float32, len(arr.objs))
	copy(objs2, arr.objs)
	arr.objs = objs2
	return arr
}

func (arr *Float32Collection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(float32); ok {
		length := len(arr.objs)

		// 如果是append操作，直接调用系统的append，不新创建collection
		if index == length {
			arr.objs = append(arr.objs, i)
			return arr
		}

		new := arr.objs[0:index]
		new = append(new, i)
		new = append(new, arr.objs[index:length]...)
		arr.objs = new
	} else {
		return arr.SetErr(errors.New("Insert: type error"))
	}
	return arr
}

func (arr *Float32Collection) Remove(i int) ICollection {
	if arr.Err() != nil {
		return arr
	}

	len := arr.Count()
	if i >= len {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs = append(arr.objs[0:i], arr.objs[i+1:len]...)
	return arr
}

func (arr *Float32Collection) NewEmpty(err ...error) ICollection {
	intArr := NewFloat32Collection([]float32{})
	if len(err) != 0 {
		intArr.err = err[0]
	}
	return intArr
}

func (arr *Float32Collection) Index(i int) IMix {
	return NewMix(arr.objs[i])
}

func (arr *Float32Collection) Count() int {
	return len(arr.objs)
}

func (arr *Float32Collection) DD() {
	ret := fmt.Sprintf("Float32Collection(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%f\n", k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}
