package collection


type IArray interface {
	Append(obj interface{})

	Search(obj interface{}) int
	Column(key string) IArray
	Unique() IArray

	KeyBy(key string) *Map

	Max() *Mix
	Min() *Mix

	Slice(start, end int) IArray
	Index(i int) *Mix
	Len() int
	Has(obj interface{}) bool
	Merge(arr IArray) IArray

	ToString() []string
	ToInt64() []int64
	ToInt() []int
}


type VArray struct {
	IArray
	Parent IArray
}

func (arr *VArray) Append(obj interface{}) {
	panic("Append: not Implement")
}

// Search find string in arr, -1 present not found, >=0 present index
func (arr *VArray) Search(obj interface{}) int {
	panic("Search: not Implement")
}

func (arr *VArray) Column(key string) IArray {
	panic("Column: not Implement")
}

func (arr *VArray) Unique() IArray {
	panic("Unique: not Implement")
}

func (arr *VArray) Max() *Mix {
	panic("Max: not Implement")
}

func (arr *VArray) Min() *Mix {
	panic("Min: not Implement")
}

func (arr *VArray) ToString() []string {
	panic("ToString: not Implement")
}

func (arr *VArray) ToInt64() []int64 {
	panic("ToInt64: not Implement")
}

func (arr *VArray) ToInt() []int{
	panic("ToInt: not Implement")
}

func (arr *VArray) KeyBy(key string) *Map {
	panic("KeyBy: not Implement")
}

func (arr *VArray) Slice(start, end int) IArray {
	panic("Slice: not Implement")
}

func (arr *VArray) Index(i int) *Mix {
	panic("Index: not Implement")
}

func (arr *VArray) Len() int {
	panic("Len: not Implement")
}

func (arr *VArray) Has(obj interface{}) bool {
	if arr.Parent.Search(obj) >= 0 {
		return true
	}
	return false
}

func (arr *VArray) Merge(bArr IArray) {
	l := bArr.Len()
	for i := 0; i < l; i++{
		arr.Append(bArr.Index(i).ToInterface())
	}
}