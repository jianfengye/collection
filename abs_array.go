package collection

// 这个是一个虚函数，能实现的都实现，不能实现的panic
type AbsArray struct {
	compare func(interface{}, interface{}) int // 比较函数

	IArray
	Parent IArray
}

/*
下面的几个函数必须要实现
 */
func (arr *AbsArray) NewEmpty() IArray {
	panic("NewEmpty: not Implement")
}

func (arr *AbsArray) Append(item interface{}) error {
	panic("Append: not Implement")
}

/*
下面这些函数是所有函数体都一样
 */
func (arr *AbsArray) IsEmpty() bool {
	return arr.Count() == 0
}

func (arr *AbsArray) IsNotEmpty() bool {
	return arr.Count() != 0
}

func (arr *AbsArray) Search(item interface{}) int {
	for i := 0; i < arr.Count(); i++ {
		if arr.compare(arr.Index(i).ToInterface(), item) == 0 {
			return i
		}
	}
	return -1
}

func (arr *AbsArray) Unique() IArray {
	newArr := arr.NewEmpty()
	for i := 0; i < arr.Count(); i++ {
		if newArr.Contains(arr.Index(i).ToInterface()) == false {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}
	return newArr
}

func (arr *AbsArray) Reject(func(item interface{}, key int) bool) IArray {
	panic("Reject: not Implement")
}

func (arr *AbsArray) Last(...func(item interface{}, key int) bool) IMix {
	panic("Last: not Implement")
}

func (arr *AbsArray) Slice(start, end int) IArray {
	panic("Slice: not Implement")
}

func (arr *AbsArray) Index(i int) IMix {
	panic("Index: not Implement")
}

func (arr *AbsArray) Count() int {
	panic("Count: not Implement")
}

func (arr *AbsArray) Merge(arr2 IArray) IArray {
	panic("Merge: not Implement")
}

func (arr *AbsArray) Chunk(count int) {
	panic("Chunk: not Implement")
}

func (arr *AbsArray) Combine(arr2 IArray) IMap {
	panic("Combine: not Implement")
}

func (arr *AbsArray) CrossJoin(arr2 IArray) IMap {
	panic("CrossJoin: not Implement")
}

func (arr *AbsArray) Each(f func(item interface{}, key int)) {
	panic("Each: not Implement")
}

func (arr *AbsArray) Map(func(item interface{}, key int)) IArray {
	panic("Map: not Implement")
}

func (arr *AbsArray) Reduce(func(carry IMix, item IMix) IMix) IMix {
	panic("Reduce: not Implement")
}

func (arr *AbsArray) Every(func(item interface{}, key int) bool) {
	panic("Every: not Implement")
}

func (arr *AbsArray) ForPage(page int, perPage int) IArray {
	panic("ForPage: not Implement")
}

func (arr *AbsArray) Nth(n int) IArray {
	panic("Nth: not Implement")
}

func (arr *AbsArray) Pad(start int, def interface{}) IArray {
	panic("Pad: not Implement")
}

func (arr *AbsArray) Pop() IMix {
	panic("Pop: not Implement")
}

func (arr *AbsArray) Push(item interface{}) {
	panic("Push: not Implement")
}

func (arr *AbsArray) Prepend(item interface{}) IArray {
	panic("Prepend: not Implement")
}

func (arr *AbsArray) Random() IMix {
	panic("Random: not Implement")
}

func (arr *AbsArray) Reverse() IArray {
	panic("Reverse: not Implement")
}

func (arr *AbsArray) Shuffle() IArray {
	panic("Shuffle: not Implement")
}

func (arr *AbsArray) DD() string {
	panic("DD: not Implement")
}

func (arr *AbsArray) ToJson() string {
	panic("ToJson: not Implement")
}

func (arr *AbsArray) Column(string) (IArray, error) {
	panic("Column: not Implement")
}

func (arr *AbsArray) KeyBy(key string) (IMap, error) {
	panic("KeyBy: not Implement")
}

func (arr *AbsArray) Pluck(val string, key string) (IMap, error) {
	panic("Pluck: not Implement")
}

func (arr *AbsArray) SortBy(key string) (IArray, error) {
	panic("SortBy: not Implement")
}

func (arr *AbsArray) SortByDesc(key string) (IArray, error) {
	panic("Reverse: not Implement")
}

func (arr *AbsArray) SetCompare(func(a interface{}, b interface{}) int) {
	panic("Reverse: not Implement")
}

func (arr *AbsArray) Max() IMix {
	panic("Max: not Implement")
}

func (arr *AbsArray) Min() IMix {
	panic("Min: not Implement")
}

func (arr *AbsArray) Contains(obj interface{}) bool {
	panic("Contains: not Implement")
}

func (arr *AbsArray) CountBy() IMap {
	panic("CountBy: not Implement")
}

func (arr *AbsArray) Diff(arr2 IArray) (IArray, error) {
	panic("Diff: not Implement")
}

func (arr *AbsArray) Sort() IArray {
	panic("Sort: not Implement")
}

func (arr *AbsArray) SortDesc() IArray {
	panic("ToString: not Implement")
}

func (arr *AbsArray) Join(split string) string {
	panic("Join: not Implement")
}

func (arr *AbsArray) Avg() IMix{
	panic("Avg: not Implement")
}

func (arr *AbsArray) Median() IMix {
	panic("Median: not Implement")
}

func (arr *AbsArray) Mode() IMix {
	panic("Mode: not Implement")
}

func (arr *AbsArray) Sum() IMix {
	panic("Sum: not Implement")
}

func (arr *AbsArray) Filter(f func(obj interface{}, index int) bool) IArray {
	ret := arr.Parent.NewEmpty()
	l := arr.Parent.Count()
	for i := 0; i < l; i++ {
		obj := arr.Parent.Index(i).ToInterface()
		if f(obj, i) == true {
			ret.Append(obj)
		}
	}
	return ret
}

func (arr *AbsArray) First(f ...func(obj interface{}, index int) bool) IMix {
	l := arr.Parent.Count()
	for i := 0; i < l; i++ {
		obj := arr.Parent.Index(i).ToInterface()
		if f[0](obj, i) == true {
			return NewMix(obj)
		}
	}
	return nil
}

func (arr *AbsArray) ToString() ([]string, error) {
	panic("Sum: not Implement")
}

func (arr *AbsArray) ToInt64() ([]int64, error) {
	panic("Sum: not Implement")
}

func (arr *AbsArray) ToInt() ([]int, error) {
	panic("Sum: not Implement")
}

func (arr *AbsArray) ToMix() []IMix {
	panic("Sum: not Implement")
}

func (arr *AbsArray) ToFloat64() ([]float64, error) {
	panic("Sum: not Implement")
}

func (arr *AbsArray) ToFloat32() ([]float32, error) {
	panic("Sum: not Implement")
}
