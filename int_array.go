package collection

import (
	"fmt"
	"github.com/pkg/errors"
)

type IntArray struct{
	AbsArray
	objs []int
}

func NewIntArray(objs []int) *IntArray {
	arr := &IntArray{
		objs:objs,
	}
	arr.AbsArray.Parent = arr
	arr.AbsArray.compare = func(i interface{}, i2 interface{}) int {
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
	return arr
}

func (arr *IntArray) Insert(index int, obj interface{}) (IArray, error) {
	if i, ok := obj.(int); ok {
		length := len(arr.objs)
		tail := arr.objs[index:length]
		arr.objs = append(arr.objs[0:index], i)
		arr.objs = append(arr.objs, tail...)
	} else {
		return arr, errors.New("Insert: type error")
	}
	return arr, nil
}

func (arr *IntArray) Remove(i int) (IArray, error) {
	len := arr.Count()
	if i >= len {
		return arr, errors.New("index exceeded")
	}
	arr.objs = append(arr.objs[0:i], arr.objs[i+1: len]...)
	return arr, nil
}

func (arr *IntArray) NewEmpty() IArray {
	return NewIntArray([]int{})
}


func (arr *IntArray) Index(i int) IMix {
	return NewMix(arr.objs[i])
}

func (arr *IntArray) Count() int {
	return len(arr.objs)
}

func (arr *IntArray) DD() {
	ret := fmt.Sprintf("IntArray(%d):{\n", arr.Count())
	for k, v := range arr.objs {
		ret = ret + fmt.Sprintf("\t%d:\t%d\n",k, v)
	}
	ret = ret + "}\n"
	fmt.Print(ret)
}