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
type AbsArray struct {
	compare func(interface{}, interface{}) int // 比较函数
	err error // 错误信息

	IArray
	Parent IArray
}

func (arr *AbsArray) Err() error {
	return arr.err
}

func (arr *AbsArray) SetErr(err error) IArray {
	arr.err = err
	return arr
}

/*
下面的几个函数必须要实现
 */
func (arr *AbsArray) NewEmpty(err ...error) IArray {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.NewEmpty(err...)
}

func (arr *AbsArray) Insert(index int, obj interface{}) IArray {
	if arr.Parent == nil {
		panic("no parent")
	}

	return arr.Parent.Insert(index, obj)
}


func (arr *AbsArray) Remove(index int) IArray {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Remove(index)
}

func (arr *AbsArray) Index(i int) IMix {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Index(i)
}

func (arr *AbsArray) Count() int {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Count()
}

func (arr *AbsArray) DD()  {
	if arr.Parent == nil {
		panic("DD: not Implement")
	}
	arr.Parent.DD()
}

func (arr *AbsArray) ToJson() []byte {
	if arr.Parent == nil {
		panic("ToJson: not Implement")
	}
	return arr.Parent.ToJson()
}

/*
下面这些函数是所有函数体都一样
 */

func (arr *AbsArray) Append(item interface{}) IArray {
	if arr.Err() != nil {
		return arr
	}
	return arr.Parent.Insert(arr.Count(), item)
}

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

func (arr *AbsArray) Reject(f func(item interface{}, key int) bool) IArray {
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

func (arr *AbsArray) Last(fs ...func(item interface{}, key int) bool) IMix {
	if len(fs) > 1 {
		panic("Last 参数个数错误")
	}

	if len(fs) == 0 {
		return arr.Index(arr.Count() - 1)
	}

	newArr := arr.Filter(fs[0])
	return newArr.Last()
}

func (arr *AbsArray) Slice(ps ...int) IArray {
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

func (arr *AbsArray) Merge(arr2 IArray) IArray {
	if arr.Err() != nil {
		return arr
	}

	for i := 0; i < arr2.Count(); i++ {
		arr.Append(arr2.Index(i).ToInterface())
	}
	return arr
}


func (arr *AbsArray) Each(f func(item interface{}, key int)) {
	for i := 0; i < arr.Count(); i++ {
		f(arr.Index(i).ToInterface(), i)
	}
}

func newMixArray(mix IMix) IArray {
	switch mix.Type().Kind() {
	case reflect.String:
		return NewStrArray([]string{})
	case reflect.Int:
		return NewIntArray([]int{})
	case reflect.Int64:
		return NewInt64Array([]int64{})
	case reflect.Float32:
		return NewFloat32Array([]float32{})
	case reflect.Float64:
		return NewFloat64Array([]float64{})
	}
	return nil
}

func (arr *AbsArray) Map(f func(item interface{}, key int) IMix) IArray {
	if arr.Err() != nil {
		return arr
	}

	// call first f for map type
	if arr.Count() == 0 {
		return nil
	}

	first := f(arr.Index(0).ToInterface(), 0)
	ret := newMixArray(first)
	ret.Append(first.ToInterface())
	for i := 1; i < arr.Count(); i++ {
		ret.Append(f(arr.Index(i).ToInterface(), 0).ToInterface())
	}
	return ret
}

func (arr *AbsArray) Reduce(f func(carry IMix, item IMix) IMix) IMix {
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

func (arr *AbsArray) Every(f func(item interface{}, key int) bool) bool {
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

func (arr *AbsArray) ForPage(page int, perPage int) IArray {
	if arr.Err() != nil {
		return arr
	}

	start := page * perPage
	return arr.Slice(start, perPage)
}

func (arr *AbsArray) Nth(n int, offset int) IArray {
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

func (arr *AbsArray) Pad(start int, def interface{}) IArray {
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

func (arr *AbsArray) Pop() IMix {
	if arr.Err() != nil {
		return nil
	}

	ret := arr.Index(arr.Count() - 1)
	arr.Remove(arr.Count() - 1)

	return ret
}

func (arr *AbsArray) Push(item interface{}) IArray {
	if arr.Err() != nil {
		return arr
	}

	return arr.Append(item)
}

func (arr *AbsArray) Prepend(item interface{}) IArray {
	if arr.Err() != nil {
		return arr
	}

	return arr.Insert(0, item)
}

func (arr *AbsArray) Random() IMix {
	if arr.Err() != nil {
		return nil
	}

	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	index := r.Intn(arr.Count())
	return arr.Index(index)
}

func (arr *AbsArray) Reverse() IArray {
	if arr.Err() != nil {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < newArr.Count() - 1; i++ {
		newArr.Append(arr.Index(newArr.Count() - 1 - i).ToInterface())
	}
	return newArr
}

func (arr *AbsArray) Shuffle() IArray {
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

func (arr *AbsArray) Pluck(key string) IArray {
	if arr.Err() != nil {
		return arr
	}

	arr.SetErr(errors.New("format not support"))
	return arr
}

func (arr *AbsArray) SortBy(key string) IArray {
	if arr.Err() != nil {
		return arr
	}

	arr.SetErr(errors.New("format not support"))
	return arr
}

func (arr *AbsArray) SortByDesc(key string) IArray {
	if arr.Err() != nil {
		return arr
	}

	arr.SetErr(errors.New("format not support"))
	return arr
}

func (arr *AbsArray) SetCompare(compare func(a interface{}, b interface{}) int) IArray {
	arr.compare = compare
	return arr
}

func (arr *AbsArray) Max() IMix {
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

func (arr *AbsArray) Min() IMix {
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

func (arr *AbsArray) Contains(obj interface{}) bool {
	for i := 0; i < arr.Count(); i++ {
		if arr.compare(arr.Index(i).ToInterface(), obj) == 0 {
			return true
		}
	}
	return false
}

func (arr *AbsArray) Diff(arr2 IArray) IArray {
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

func (arr *AbsArray) Sort() IArray {
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

func (arr *AbsArray) SortDesc() IArray {
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

func (arr *AbsArray) Join(split string, format ...func(item interface{}) string) string {
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

func (arr *AbsArray) Avg() IMix {
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

func (arr *AbsArray) Median() IMix {
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

func (arr *AbsArray) Mode() IMix {
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

func (arr *AbsArray) Sum() IMix {
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}

	mix := NewMix(arr.Index(0).ToInterface())
	for i := 0; i < arr.Count(); i++ {
		mix.Add(arr.Index(i))
	}
	return mix
}

func (arr *AbsArray) Filter(f func(obj interface{}, index int) bool) IArray {
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

func (arr *AbsArray) First(f ...func(obj interface{}, index int) bool) IMix {
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

func (arr *AbsArray) ToStrings() ([]string, error) {
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

func (arr *AbsArray) ToInt64s() ([]int64, error) {
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

func (arr *AbsArray) ToInts() ([]int, error) {
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

func (arr *AbsArray) ToMixs() ([]IMix, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]IMix, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		ret[i] = arr.Index(i)
	}
	return ret, nil
}

func (arr *AbsArray) ToFloat64s() ([]float64, error) {
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

func (arr *AbsArray) ToFloat32s() ([]float32, error) {
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
