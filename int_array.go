package collection

import (
	"math"
)

type IntArray struct{
	VArray
	objs []int
}

func NewIntArray(objs []int) *IntArray {
	arr := &IntArray{
		objs:objs,
	}
	arr.VArray.Parent = arr
	return arr
}

func (arr *IntArray) mustBeInt(obj interface{}) int {
	if i, ok := obj.(int); ok {
		return i
	} else {
		panic("obj must be int")
	}
}

func (arr *IntArray) Append(obj interface{}) {
	param := arr.mustBeInt(obj)
	arr.objs = append(arr.objs, param)
}

// Search find string in arr, -1 present not found, >=0 present index
func (arr *IntArray) Search(obj interface{}) int {
	param := arr.mustBeInt(obj)
	for i, t := range arr.objs {
		if t == param {
			return i
		}
	}
	return -1
}

func (arr *IntArray) Max() *Mix {
	var max int = math.MinInt32
	for _, obj := range arr.objs {
		if obj > max {
			max = obj
		}
	}
	return NewMix(max)
}

func (arr *IntArray) Min() *Mix {
	var min int = math.MaxInt32
	for _, obj := range arr.objs {
		if obj < min {
			min = obj
		}
	}
	return NewMix(min)
}

func (arr *IntArray) ToInt() []int{
	return arr.objs
}

func (arr *IntArray) Index(i int) *Mix {
	return NewMix(arr.objs[i])
}

func (arr *IntArray) Slice(start, end int) IArray {
	return NewIntArray(arr.objs[start:end])
}

func (arr *IntArray) Len() int {
	return len(arr.objs)
}

func (arr *IntArray) NewEmptyIArray() IArray {
	return NewIntArray([]int{})
}

func (arr *IntArray) Unique() IArray {
	objs := arr.ToInt()
	ret := NewIntArray([]int{})

	for _, s := range objs {
		if ret.Search(s) < 0 {
			ret.Append(s)
		}
	}

	return ret
}