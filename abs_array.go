package collection

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"
)

// 这个是一个虚函数，能实现的都实现，不能实现的panic
type AbsCollection struct {
	compare func(interface{}, interface{}) int // 比较函数
	err error // 错误信息

	ICollection
	Parent ICollection
}

func (arr *AbsCollection) Err() error {
	return arr.err
}

func (arr *AbsCollection) SetErr(err error) ICollection {
	arr.err = err
	return arr
}

/*
下面的几个函数必须要实现
 */
func (arr *AbsCollection) NewEmpty(err ...error) ICollection {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.NewEmpty(err...)
}

func (arr *AbsCollection) Insert(index int, obj interface{}) ICollection {
	if arr.Parent == nil {
		panic("no parent")
	}

	return arr.Parent.Insert(index, obj)
}


func (arr *AbsCollection) Remove(index int) ICollection {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Remove(index)
}

func (arr *AbsCollection) Index(i int) IMix {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Index(i)
}

func (arr *AbsCollection) Count() int {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Count()
}

func (arr *AbsCollection) DD()  {
	if arr.Parent == nil {
		panic("DD: not Implement")
	}
	arr.Parent.DD()
}

func (arr *AbsCollection) ToJson() []byte {
	if arr.Parent == nil {
		panic("ToJson: not Implement")
	}
	return arr.Parent.ToJson()
}

/*
下面这些函数是所有函数体都一样
 */

func (arr *AbsCollection) Append(item interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	return arr.Parent.Insert(arr.Count(), item)
}

func (arr *AbsCollection) IsEmpty() bool {
	return arr.Count() == 0
}

func (arr *AbsCollection) IsNotEmpty() bool {
	return arr.Count() != 0
}

func (arr *AbsCollection) Search(item interface{}) int {
	for i := 0; i < arr.Count(); i++ {
		if arr.compare(arr.Index(i).ToInterface(), item) == 0 {
			return i
		}
	}
	return -1
}

func (arr *AbsCollection) Unique() ICollection {
	if arr.Err() != nil {
		return arr
	}
	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < arr.Count(); i++ {
		if newArr.Contains(arr.Index(i).ToInterface()) == false {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}
	return newArr
}

func (arr *AbsCollection) Reject(f func(item interface{}, key int) bool) ICollection {
	if arr.Err() != nil {
		return arr
	}
	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < arr.Count(); i++ {
		if f(arr.Index(i).ToInterface(), i) == false {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}
	return newArr
}

func (arr *AbsCollection) Last(fs ...func(item interface{}, key int) bool) IMix {
	if len(fs) > 1 {
		panic("Last 参数个数错误")
	}

	if len(fs) == 0 {
		return arr.Index(arr.Count() - 1)
	}

	newArr := arr.Filter(fs[0])
	return newArr.Last()
}

func (arr *AbsCollection) Slice(ps ...int) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if len(ps) > 2 || len(ps) == 0{
		panic("Slice params count error")
	}

	start := ps[0]
	count := arr.Count()
	if len(ps) == 2 && ps[1] != -1 {
		count = ps[1]
	}

	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < arr.Count(); i++ {
		if i >= start {
			newArr.Append(arr.Index(i).ToInterface())
			if newArr.Count() >= count {
				break
			}
		}
	}
	return newArr
}

func (arr *AbsCollection) Merge(arr2 ICollection) ICollection {
	if arr.Err() != nil {
		return arr
	}

	for i := 0; i < arr2.Count(); i++ {
		arr.Append(arr2.Index(i).ToInterface())
	}
	return arr
}


func (arr *AbsCollection) Each(f func(item interface{}, key int)) {
	for i := 0; i < arr.Count(); i++ {
		f(arr.Index(i).ToInterface(), i)
	}
}

func newMixCollection(mix IMix) ICollection {
	switch mix.Type().Kind() {
	case reflect.String:
		return NewStrCollection([]string{})
	case reflect.Int:
		return NewIntCollection([]int{})
	case reflect.Int64:
		return NewInt64Collection([]int64{})
	case reflect.Float32:
		return NewFloat32Collection([]float32{})
	case reflect.Float64:
		return NewFloat64Collection([]float64{})
	}
	return nil
}

func (arr *AbsCollection) Map(f func(item interface{}, key int) IMix) ICollection {
	if arr.Err() != nil {
		return arr
	}

	// call first f for map type
	if arr.Count() == 0 {
		return nil
	}

	first := f(arr.Index(0).ToInterface(), 0)
	ret := newMixCollection(first)
	ret.Append(first.ToInterface())
	for i := 1; i < arr.Count(); i++ {
		ret.Append(f(arr.Index(i).ToInterface(), 0).ToInterface())
	}
	return ret
}

func (arr *AbsCollection) Reduce(f func(carry IMix, item IMix) IMix) IMix {
	if arr.Err() != nil {
		return nil
	}

	if arr.Count() == 0 {
		return nil
	}

	if arr.Count() == 1 {
		return NewMix(arr.Index(0).ToInterface())
	}

	carry := f(NewMix(arr.Index(0).ToInterface()), NewMix(arr.Index(1).ToInterface()))

	for i := 2; i < arr.Count(); i++ {
		carry = f(carry, NewMix(arr.Index(i).ToInterface()))
	}
	return carry
}

func (arr *AbsCollection) Every(f func(item interface{}, key int) bool) bool {
	if arr.Count() == 0 {
		return true
	}

	for i := 0; i < arr.Count(); i++ {
		if f(arr.Index(i).ToInterface(), i) == false {
			return false
		}
	}
	return true
}

func (arr *AbsCollection) ForPage(page int, perPage int) ICollection {
	if arr.Err() != nil {
		return arr
	}

	start := page * perPage
	return arr.Slice(start, perPage)
}

func (arr *AbsCollection) Nth(n int, offset int) ICollection {
	if arr.Err() != nil {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < arr.Count(); i++ {
		if (i - offset) % n == 0 {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}
	return newArr
}

func (arr *AbsCollection) Pad(start int, def interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())

	if start > 0 {
		if start <= arr.Count() {
			return arr.Slice(0, start)
		}

		for i:=0; i < arr.Count(); i++ {
			newArr.Append(arr.Index(i).ToInterface())
		}
		for i := arr.Count(); i < start; i++ {
			newArr = newArr.Append(def)
		}

		return newArr
	}

	if start < 0 {
		if start >= -arr.Count() {
			return arr.Slice(arr.Count() + start, arr.Count())
		}

		for i := start; i < -arr.Count(); i++ {
			newArr.Append(def)
		}

		for i := 0; i < arr.Count(); i++ {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}

	return newArr
}

func (arr *AbsCollection) Pop() IMix {
	if arr.Err() != nil {
		return nil
	}

	ret := arr.Index(arr.Count() - 1)
	arr.Remove(arr.Count() - 1)

	return ret
}

func (arr *AbsCollection) Push(item interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}

	return arr.Append(item)
}

func (arr *AbsCollection) Prepend(item interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}

	return arr.Insert(0, item)
}

func (arr *AbsCollection) Random() IMix {
	if arr.Err() != nil {
		return nil
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	index := r.Intn(arr.Count())
	return arr.Index(index)
}

func (arr *AbsCollection) Reverse() ICollection {
	if arr.Err() != nil {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < newArr.Count() - 1; i++ {
		newArr.Append(arr.Index(newArr.Count() - 1 - i).ToInterface())
	}
	return newArr
}

func (arr *AbsCollection) Shuffle() ICollection {
	if arr.Err() != nil {
		return arr
	}

	indexs := make([]int, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		indexs[i] = i
	}
	rand.Shuffle(arr.Count(), func(i, j int) {
		indexs[i], indexs[j] = indexs[j], indexs[i]
	})

	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < len(indexs); i ++ {
		newArr.Append(arr.Index(indexs[i]).ToInterface())
	}
	return newArr
}

func (arr *AbsCollection) Pluck(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	arr.SetErr(errors.New("format not support"))
	return arr
}

func (arr *AbsCollection) SortBy(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	arr.SetErr(errors.New("format not support"))
	return arr
}

func (arr *AbsCollection) SortByDesc(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	arr.SetErr(errors.New("format not support"))
	return arr
}

func (arr *AbsCollection) SetCompare(compare func(a interface{}, b interface{}) int) ICollection {
	arr.compare = compare
	return arr
}

func (arr *AbsCollection) Max() IMix {
	if arr.Err() != nil {
		return nil
	}

	max := 0
	for i := 0; i < arr.Count(); i++ {
		if arr.compare(arr.Index(i).ToInterface(), arr.Index(max).ToInterface()) > 0 {
			max = i
		}
	}
	return arr.Index(max)
}

func (arr *AbsCollection) Min() IMix {
	if arr.Err() != nil {
		return nil
	}

	min := 0
	for i := 0; i < arr.Count(); i++ {
		if arr.compare(arr.Index(i).ToInterface(), arr.Index(min).ToInterface()) < 0 {
			min = i
		}
	}
	return arr.Index(min)
}

func (arr *AbsCollection) Contains(obj interface{}) bool {
	for i := 0; i < arr.Count(); i++ {
		if arr.compare(arr.Index(i).ToInterface(), obj) == 0 {
			return true
		}
	}
	return false
}

func (arr *AbsCollection) Diff(arr2 ICollection) ICollection {
	if arr.Err() != nil {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < arr.Count(); i++ {
		if arr2.Contains(arr.Index(i).ToInterface()) == false {
			newArr.Append(arr.Index(i).ToInterface())
		}
	}
	return newArr
}

func (arr *AbsCollection) Sort() ICollection {
	if arr.Err() != nil {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())

	isContained := func(arr []int, item int) bool {
		for i := 0; i < len(arr); i++ {
			if item == arr[i] {
				return true
			}
		}
		return false
	}


	sorted := make([]int, 0, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		min := -1
		for j := 0; j < arr.Count(); j++ {
			if isContained(sorted, j) {
				continue
			}

			if min == -1 {
				min = j
				continue
			}

			if arr.compare(arr.Index(j).ToInterface(), arr.Index(min).ToInterface()) <= 0 {
				min = j
				continue
			}
		}

		sorted = append(sorted, min)
		newArr.Append(arr.Index(min).ToInterface())
	}
	return newArr
}

func (arr *AbsCollection) SortDesc() ICollection {
	if arr.Err() != nil {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())

	isContained := func(arr []int, item int) bool {
		for i := 0; i < len(arr); i++ {
			if item == arr[i] {
				return true
			}
		}
		return false
	}


	sorted := make([]int, 0, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		max := -1
		for j := 0; j < arr.Count(); j++ {
			if isContained(sorted, j) {
				continue
			}

			if max == -1 {
				max = j
				continue
			}

			if arr.compare(arr.Index(j).ToInterface(), arr.Index(max).ToInterface()) >= 0 {
				max = j
				continue
			}
		}

		sorted = append(sorted, max)
		newArr.Append(arr.Index(max).ToInterface())
	}
	return newArr
}

func (arr *AbsCollection) Join(split string, format ...func(item interface{}) string) string {
	if arr.Err() != nil {
		return ""
	}

	var ret strings.Builder
	for i := 0; i < arr.Count(); i++ {
		if len(format) == 0 {
			ret.WriteString(fmt.Sprintf("%v", arr.Index(i).ToInterface()))
		} else {
			f := format[0]
			ret.WriteString(f(arr.Index(i).ToInterface()))
		}

		if i != arr.Count() - 1 {
			ret.WriteString(split)
		}
	}
	return ret.String()
}

func (arr *AbsCollection) Avg() IMix {
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}
	if arr.Count() == 0 {
		return NewErrorMix(errors.New("arr count can not be empty"))
	}

	var sum IMix
	var err error
	sum = NewMix(arr.Index(0).ToInterface())
	for i:= 1; i < arr.Count(); i++ {
		sum, err = sum.Add(arr.Index(i))
		if err != nil {
			return NewErrorMix(err)
		}
	}
	div, err := sum.Div(arr.Count())
	if err != nil {
		return NewErrorMix(err)
	}
	return div
}

func (arr *AbsCollection) Median() IMix {
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}
	newArr := arr.Sort()
	if newArr.Count() % 2 == 0 {
		imax, err := newArr.Index(newArr.Count() / 2 - 1).Add(newArr.Index(newArr.Count() / 2 + 1))
		if err != nil {
			return NewErrorMix(err)
		}
		div, err := imax.Div(2)
		if err != nil {
			return NewErrorMix(err)
		}
		return div
	}

	return newArr.Index(newArr.Count() / 2 + 1)
}

func (arr *AbsCollection) Mode() IMix {
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}

	uniqColl := arr.Unique()
	max := 0
	retIndex := 0
	uniqColl.Each(func(item interface{}, key int) {
		sum := 0
		arr.Each(func(obj interface{}, index int) {
			if arr.compare(item, obj) == 0 {
				sum++
			}
		})
		if sum > max {
			max = sum
			retIndex = key
		}
	})

	return uniqColl.Index(retIndex)
}

func (arr *AbsCollection) Sum() IMix {
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}

	mix := NewMix(arr.Index(0).ToInterface())
	for i := 0; i < arr.Count(); i++ {
		mix.Add(arr.Index(i))
	}
	return mix
}

func (arr *AbsCollection) Filter(f func(obj interface{}, index int) bool) ICollection {
	if arr.Err() != nil {
		return arr
	}

	ret := arr.Parent.NewEmpty(arr.Err())
	l := arr.Parent.Count()
	for i := 0; i < l; i++ {
		obj := arr.Parent.Index(i).ToInterface()
		if f(obj, i) == true {
			ret.Append(obj)
		}
	}
	return ret
}

func (arr *AbsCollection) First(f ...func(obj interface{}, index int) bool) IMix {
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}

	if len(f) == 0 {
		return arr.Parent.Index(0)
	}
	fun := f[0]

	l := arr.Parent.Count()
	for i := 0; i < l; i++ {
		obj := arr.Parent.Index(i).ToInterface()
		if fun(obj, i) == true {
			return arr.Parent.Index(i)
		}
	}
	return nil
}

func (arr *AbsCollection) ToStrings() ([]string, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}

	ret := make([]string, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t , err := arr.Index(i).ToString()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsCollection) ToInt64s() ([]int64, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]int64, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t , err := arr.Index(i).ToInt64()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsCollection) ToInts() ([]int, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]int, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t , err := arr.Index(i).ToInt()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsCollection) ToMixs() ([]IMix, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]IMix, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		ret[i] = arr.Index(i)
	}
	return ret, nil
}

func (arr *AbsCollection) ToFloat64s() ([]float64, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]float64, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t , err := arr.Index(i).ToFloat64()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsCollection) ToFloat32s() ([]float32, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]float32, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t , err := arr.Index(i).ToFloat32()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}
