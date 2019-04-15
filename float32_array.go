package collection

import (
	"fmt"
	"github.com/pkg/errors"
)

type Float32Array struct{
	AbsArray
	objs []float32
}

func NewFloat32Array(objs []float32) *Float32Array {
	arr := &Float32Array{
		objs:objs,
	}
	arr.AbsArray.Parent = arr
	arr.AbsArray.compare = func(i interface{}, i2 interface{}) int {
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
	return arr
}

func (arr *Float32Array) Insert(index int, obj interface{}) IArray {
	if arr.Err() != nil {
		return arr
	}
	if i, ok := obj.(float32); ok {
		length := len(arr.objs)
		tail := arr.objs[index:length]
		arr.objs = append(arr.objs[0:index], i)
		arr.objs = append(arr.objs, tail...)
	} else {
		return arr.SetErr(errors.New("Insert: type error"))
	}
	return arr
}

func (arr *Float32Array) Remove(i int) IArray {
	if arr.Err() != nil {
		return arr
	}

	len := arr.Count()
	if i >= len {
		return arr.SetErr(errors.New("index exceeded"))
	}
	arr.objs = append(arr.objs[0:i], arr.objs[i+1: len]...)
	return arr
}

func (arr *Float32Array) NewEmpty(err ...error) IArray {
	intArr := NewFloat32Array([]float32{})
	if len(err) != 0 {
		intArr.err = err[0]
	}
	return intArr
}


func (arr *Float32Array) Index(i int) IMix {
	return NewMix(arr.objs[i])
}

func (arr *Float32Array) Count() int {
	return len(arr.objs)
}

func (arr *Float32Array) DD() {
	ret := fmt.Sprintf("IntArray(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%v\n",k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}