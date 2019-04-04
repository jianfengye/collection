package collection

type Int64Array struct{
	AbsArray
	objs []int64
}

func NewInt64Array(objs []int64) *Int64Array {
	arr := &Int64Array{
		objs:objs,
	}
	arr.AbsArray.Parent = arr
	return arr
}