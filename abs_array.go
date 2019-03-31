package collection

// 这个是一个虚函数，能实现的都实现，不能实现的panic
type AbsArray struct {
	IArray
	Parent IArray
}

func (arr *AbsArray) Append(obj interface{}) {
	panic("Append: not Implement")
}

// Search find string in arr, -1 present not found, >=0 present index
func (arr *AbsArray) Search(obj interface{}) int {
	panic("Search: not Implement")
}

func (arr *AbsArray) Column(key string) IArray {
	panic("Column: not Implement")
}

func (arr *AbsArray) Unique() IArray {
	panic("Unique: not Implement")
}

func (arr *AbsArray) Max() IMix {
	panic("Max: not Implement")
}

func (arr *AbsArray) Min() IMix {
	panic("Min: not Implement")
}

func (arr *AbsArray) ToString() []string {
	panic("ToString: not Implement")
}

func (arr *AbsArray) ToInt64() []int64 {
	panic("ToInt64: not Implement")
}

func (arr *AbsArray) ToInt() []int{
	panic("ToInt: not Implement")
}

func (arr *AbsArray) KeyBy(key string) IMap {
	panic("KeyBy: not Implement")
}

func (arr *AbsArray) Slice(start, end int) IArray {
	panic("Slice: not Implement")
}

func (arr *AbsArray) Index(i int) IMix {
	panic("Index: not Implement")
}

func (arr *AbsArray) Len() int {
	panic("Len: not Implement")
}

func (arr *AbsArray) Has(obj interface{}) bool {
	if arr.Parent.Search(obj) >= 0 {
		return true
	}
	return false
}

func (arr *AbsArray) Merge(bArr IArray) IArray {
	l := bArr.Len()
	for i := 0; i < l; i++{
		arr.Append(bArr.Index(i).ToInterface())
	}
	return arr
}

func (arr *AbsArray) NewEmptyIArray() IArray {
	panic("NewEmptyIArray: not Implement")
}

func (arr *AbsArray) Filter(f func(obj interface{}, index int) bool) IArray {
	ret := arr.Parent.NewEmptyIArray()
	l := arr.Parent.Len()
	for i := 0; i < l; i++ {
		obj := arr.Parent.Index(i).ToInterface()
		if f(obj, i) == true {
			ret.Append(obj)
		}
	}
	return ret
}

func (arr *AbsArray) First(f func(obj interface{}, index int) bool) *Mix {
	l := arr.Parent.Len()
	for i := 0; i < l; i++ {
		obj := arr.Parent.Index(i).ToInterface()
		if f(obj, i) == true {
			return NewMix(obj)
		}
	}
	return nil
}