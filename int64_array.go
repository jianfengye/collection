package collection

import (
	"math"
)

type Int64Array struct{
	VArray
	objs []int64
}

func NewInt64Array(objs []int64) *Int64Array {
	arr := &Int64Array{
		objs:objs,
	}
	arr.VArray.Parent = arr
	return arr
}

func (arr *Int64Array) mustBeInt64(obj interface{}) int64 {
	if i, ok := obj.(int64); ok {
		return i
	} else {
		panic("obj must be int64")
	}
}

func (arr *Int64Array) Append(obj interface{}) {
	param := arr.mustBeInt64(obj)
	arr.objs = append(arr.objs, param)
}

// Search find string in arr, -1 present not found, >=0 present index
func (arr *Int64Array) Search(obj interface{}) int {
	param := arr.mustBeInt64(obj)
	for i, t := range arr.objs {
		if t == param {
			return i
		}
	}
	return -1
}

func (arr *Int64Array) Max() *Mix {
	var max int64 = math.MinInt64
	for _, obj := range arr.objs {
		if obj > max {
			max = obj
		}
	}
	return NewMix(max)
}

func (arr *Int64Array) Min() *Mix {
	var min int64 = math.MaxInt64
	for _, obj := range arr.objs {
		if obj < min {
			min = obj
		}
	}
	return NewMix(min)
}

func (arr *Int64Array) ToInt64() []int64 {
	return arr.objs
}

func (arr *Int64Array) Len() int {
	return len(arr.objs)
}

func (arr *Int64Array) NewEmptyIArray() IArray {
	return NewInt64Array([]int64{})
}


func (arr *Int64Array) Unique() IArray {
	objs := arr.ToInt64()
	ret := NewInt64Array([]int64{})

	for _, s := range objs {
		if ret.Search(s) < 0 {
			ret.Append(s)
		}
	}

	return ret
}