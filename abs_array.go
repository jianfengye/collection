package collection

import (
	"fmt"
	"math/rand"
	"reflect"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type EleType int

const (
	TYPE_UNKNWON EleType = iota
	Type_INT
	Type_INT64
	Type_INT32
	TYPE_STRING
	TYPE_FLOAT32
	TYPE_FLOAT64
	TYPE_OBJ
	TYPE_OBJ_POINT
)

// 这个是一个虚函数，能实现的都实现
type AbsCollection struct {
	compare func(interface{}, interface{}) int // 比较函数
	err     error                              // 错误信息

	eleType EleType // 元素类型

	ICollection
	Parent ICollection //用于调用子类
}

func (arr *AbsCollection) Err() error {
	return arr.err
}

func (arr *AbsCollection) SetErr(err error) ICollection {
	arr.err = err
	return arr
}

/*
下面几个函数是内部函数
*/
func (arr *AbsCollection) mustSetCompare() *AbsCollection {
	if arr.compare == nil {
		err := errors.New("compare function must be set")
		arr.SetErr(err)
	}
	return arr
}

func (arr *AbsCollection) mustBeNumType() *AbsCollection {
	switch arr.eleType {
	case TYPE_OBJ, TYPE_OBJ_POINT, TYPE_STRING, TYPE_UNKNWON:
		err := errors.New("collection type must be num type")
		arr.SetErr(err)
	}
	return arr
}

func (arr *AbsCollection) mustBeBaseType() *AbsCollection {
	switch arr.eleType {
	case TYPE_OBJ, TYPE_OBJ_POINT, TYPE_UNKNWON:
		err := errors.New("collection type must be base type")
		arr.SetErr(err)
	}
	return arr
}

func (arr *AbsCollection) mustNotBeBaseType() *AbsCollection {
	switch arr.eleType {
	case TYPE_OBJ, TYPE_OBJ_POINT, TYPE_UNKNWON:
		return arr
	}
	err := errors.New("collection type must not be base type")
	arr.SetErr(err)
	return arr
}

func (arr *AbsCollection) mustNotBeEmpty() *AbsCollection {
	if arr.Parent.Count() == 0 {
		err := errors.New("collection should not be empty")
		arr.SetErr(err)
	}
	return arr
}

/*
下面的几个函数必须要实现
*/
func (arr *AbsCollection) NewEmpty(err ...error) ICollection {
	if arr.Parent == nil {
		panic("no parent")
	}
	empty := arr.Parent.NewEmpty(err...)
	empty.SetCompare(arr.compare)
	return empty
}

func (arr *AbsCollection) Insert(index int, obj interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if arr.Parent == nil {
		panic("no parent")
	}

	return arr.Parent.Insert(index, obj)
}

func (arr *AbsCollection) Remove(index int) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Remove(index)
}

func (arr *AbsCollection) Index(i int) IMix {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Index(i).SetCompare(arr.compare)
}

func (arr *AbsCollection) SetIndex(i int, val interface{}) ICollection {
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.SetIndex(i, val)
}

func (arr *AbsCollection) Copy() ICollection {
	if arr.Parent == nil {
		panic("no parent")
	}
	newArr := arr.Parent.Copy()
	if newArr.GetCompare() == nil {
		newArr.SetCompare(arr.compare)
	}
	if arr.Err() != nil {
		newArr.SetErr(arr.Err())
	}
	return newArr
}

func (arr *AbsCollection) Count() int {
	if arr.Err() != nil {
		return 0
	}
	if arr.Parent == nil {
		panic("no parent")
	}
	return arr.Parent.Count()
}

func (arr *AbsCollection) DD() {
	if arr.Parent == nil {
		panic("DD: not Implement")
	}
	arr.Parent.DD()
}

/*
下面这些函数是所有子类都一样
*/

func (arr *AbsCollection) Append(item interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}
	return arr.Insert(arr.Count(), item)
}

func (arr *AbsCollection) IsEmpty() bool {
	return arr.Count() == 0
}

func (arr *AbsCollection) IsNotEmpty() bool {
	return arr.Count() != 0
}

func (arr *AbsCollection) Search(item interface{}) int {
	arr.mustSetCompare()
	if arr.Err() != nil {
		return -1
	}

	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(i).ToInterface()
		if arr.compare(o, item) == 0 {
			return i
		}
	}
	return -1
}

func (arr *AbsCollection) Unique() ICollection {
	arr.mustSetCompare()
	if arr.Err() != nil {
		return arr
	}
	if arr.Count() == 0 {
		return arr
	}
	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(i).ToInterface()
		if newArr.Contains(o) == false {
			newArr.Append(o)
		}
	}
	return newArr
}

func (arr *AbsCollection) Reject(f func(item interface{}, key int) bool) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if arr.Count() == 0 {
		return arr
	}
	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(i).ToInterface()
		if f(o, i) == false {
			if arr.Err() != nil {
				break
			}
			newArr.Append(o)
		}
	}
	return newArr
}

func (arr *AbsCollection) Last(fs ...func(item interface{}, key int) bool) IMix {
	arr.mustNotBeEmpty()
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}
	if len(fs) > 1 {
		return NewErrorMix(errors.New("param error"))
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
	if len(ps) > 2 || len(ps) == 0 {
		panic("Slice params count error")
	}
	if arr.Count() == 0 {
		return arr
	}

	start := ps[0]
	count := arr.Count()
	if len(ps) == 2 && ps[1] != -1 {
		count = ps[1]
	}

	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < arr.Count(); i++ {
		if i >= start {
			o, _ := arr.Index(i).ToInterface()
			newArr.Append(o)
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
	if arr2.Err() != nil {
		return arr.SetErr(errors.New("merge error collection"))
	}

	for i := 0; i < arr2.Count(); i++ {
		o, _ := arr2.Index(i).ToInterface()
		arr.Append(o)
	}
	return arr
}

func (arr *AbsCollection) Each(f func(item interface{}, key int)) {
	if arr.Err() != nil {
		return
	}
	if arr.Count() == 0 {
		return
	}
	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(i).ToInterface()
		f(o, i)
		// if f want to stop Each, it just set err for AbsCollection
		if arr.Err() != nil {
			return
		}
	}
}

func (arr *AbsCollection) GroupBy(f func(item interface{}, key int) interface{}) map[interface{}]ICollection {
	if arr.Err() != nil {
		return nil
	}

	size := arr.Count()
	if size == 0 {
		return nil
	}
	objMap := make(map[interface{}]ICollection, size)
	for i := 0; i < size; i++ {
		o, _ := arr.Index(i).ToInterface()
		key := f(o, i)
		// 如果返回空，则默认为continue
		if key == nil {
			continue
		}
		if value, isOk := objMap[key]; isOk {
			value.Append(o)
		} else {
			ret := arr.NewEmpty()
			ret.Append(o)
			objMap[key] = ret
		}
	}

	return objMap
}

func (arr *AbsCollection) Map(f func(item interface{}, key int) interface{}) ICollection {
	if arr.Err() != nil {
		return arr
	}

	// call first f for map type
	if arr.Count() == 0 {
		return arr
	}

	ret := make([]interface{}, 0, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(i).ToInterface()
		o2 := f(o, i)

		if arr.Err() != nil {
			break
		}

		// 如果返回空，则默认为continue
		if o2 == nil {
			continue
		}

		ret = append(ret, o2)
	}

	if len(ret) == 0 {
		return nil
	}
	newColl := NewMixCollection(reflect.TypeOf(ret[0]))
	for _, v := range ret {
		newColl.Append(v)
	}
	newColl.SetErr(arr.err)
	return newColl
}

func (arr *AbsCollection) Reduce(f func(carry IMix, item IMix) IMix) IMix {
	arr.mustNotBeEmpty().mustBeBaseType()
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}
	if arr.Count() == 1 {
		o, _ := arr.Index(0).ToInterface()
		return NewMix(o)
	}

	o0, _ := arr.Index(0).ToInterface()
	o1, _ := arr.Index(1).ToInterface()
	carry := f(NewMix(o0), NewMix(o1))
	if arr.Err() != nil {
		return carry
	}

	for i := 2; i < arr.Count(); i++ {
		oi, _ := arr.Index(i).ToInterface()
		carry = f(carry, NewMix(oi))
		if arr.Err() != nil {
			return carry
		}
	}
	return carry
}

func (arr *AbsCollection) Every(f func(item interface{}, key int) bool) bool {
	if arr.Err() != nil {
		return false
	}
	if arr.Count() == 0 {
		return true
	}

	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(i).ToInterface()
		if f(o, i) == false {
			return false
		}
		if arr.Err() != nil {
			return false
		}
	}
	return true
}

func (arr *AbsCollection) ForPage(page int, perPage int) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if arr.Count() == 0 {
		return arr
	}

	start := page * perPage
	return arr.Slice(start, perPage)
}

func (arr *AbsCollection) Nth(n int, offset int) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if arr.Count() == 0 {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())
	for i := offset; i < arr.Count(); i++ {
		if (i-offset)%n == 0 {
			o, _ := arr.Index(i).ToInterface()
			newArr.Append(o)
		}
	}
	return newArr
}

func (arr *AbsCollection) Pad(count int, def interface{}) ICollection {
	arr.mustBeNumType()
	if arr.Err() != nil {
		return arr
	}
	if arr.Count() == 0 {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())

	if count > 0 {
		if count <= arr.Count() {
			return arr.Slice(0, count)
		}

		for i := 0; i < arr.Count(); i++ {
			o, _ := arr.Index(i).ToInterface()
			newArr.Append(o)
		}
		for i := arr.Count(); i < count; i++ {
			newArr = newArr.Append(def)
		}

		return newArr
	}

	if count < 0 {
		if count >= -arr.Count() {
			startIndex := arr.Count() + count
			return arr.Slice(startIndex, arr.Count()-startIndex)
		}

		for i := count; i < -arr.Count(); i++ {
			newArr.Append(def)
		}

		for i := 0; i < arr.Count(); i++ {
			o, _ := arr.Index(i).ToInterface()
			newArr.Append(o)
		}
	}

	return newArr
}

func (arr *AbsCollection) Pop() IMix {
	if arr.Err() != nil {
		return nil
	}

	if arr.Count() == 0 {
		return NewErrorMix(errors.New("Collection can not be empty"))
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

	if arr.Count() == 0 {
		return NewErrorMix(errors.New("Collection can not be empty"))
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
	if arr.Count() == 0 {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(arr.Count() - 1 - i).ToInterface()
		newArr.Append(o)
	}
	return newArr
}

func (arr *AbsCollection) Shuffle() ICollection {
	if arr.Err() != nil {
		return arr
	}
	if arr.Count() == 0 {
		return arr
	}

	indexs := make([]int, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		indexs[i] = i
	}
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	r.Shuffle(arr.Count(), func(i, j int) {
		indexs[i], indexs[j] = indexs[j], indexs[i]
	})

	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < len(indexs); i++ {
		o, _ := arr.Index(indexs[i]).ToInterface()
		newArr.Append(o)
	}
	return newArr
}

func (arr *AbsCollection) Pluck(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	return arr.Parent.Pluck(key)
}

func (arr *AbsCollection) SortBy(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	if arr.Count() == 0 {
		return arr
	}

	arr.SetErr(errors.New("format not support"))
	return arr
}

func (arr *AbsCollection) SortByDesc(key string) ICollection {
	if arr.Err() != nil {
		return arr
	}

	if arr.Count() == 0 {
		return arr
	}

	arr.SetErr(errors.New("format not support"))
	return arr
}

func (arr *AbsCollection) SetCompare(compare func(a interface{}, b interface{}) int) ICollection {
	arr.compare = compare
	return arr
}

func (arr *AbsCollection) GetCompare() func(a interface{}, b interface{}) int {
	return arr.compare
}

func (arr *AbsCollection) Max() IMix {
	arr.mustSetCompare().mustNotBeEmpty()
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}

	if arr.Count() == 0 {
		return NewErrorMix(errors.New("max: arr count can not be zero"))
	}

	max := 0
	for i := 0; i < arr.Count(); i++ {
		oi, _ := arr.Index(i).ToInterface()
		omax, _ := arr.Index(max).ToInterface()
		if arr.compare(oi, omax) > 0 {
			max = i
		}
	}
	return arr.Index(max)
}

func (arr *AbsCollection) Min() IMix {
	arr.mustSetCompare().mustNotBeEmpty()
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}

	if arr.Count() == 0 {
		return NewErrorMix(errors.New("min: arr count can not be zero"))
	}

	min := 0
	for i := 0; i < arr.Count(); i++ {
		oi, _ := arr.Index(i).ToInterface()
		omin, _ := arr.Index(min).ToInterface()
		if arr.compare(oi, omin) < 0 {
			min = i
		}
	}
	return arr.Index(min)
}

func (arr *AbsCollection) Contains(obj interface{}) bool {
	arr.mustSetCompare()
	if arr.Err() != nil {
		return false
	}
	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(i).ToInterface()
		if arr.compare(o, obj) == 0 {
			return true
		}
	}
	return false
}

func (arr *AbsCollection) ContainsCount(obj interface{}) int {
	arr.mustSetCompare()
	if arr.Err() != nil {
		return 0
	}
	sum := 0
	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(i).ToInterface()
		if arr.compare(o, obj) == 0 {
			sum++
		}
	}
	return sum
}

func (arr *AbsCollection) Diff(arr2 ICollection) ICollection {
	arr.mustSetCompare()
	if arr.Err() != nil {
		return arr
	}

	newArr := arr.NewEmpty(arr.Err())
	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(i).ToInterface()
		if arr2.Contains(o) == false {
			newArr.Append(o)
		}
	}
	return newArr
}

func (arr *AbsCollection) qsort(left, right int, isAscOrder bool) {
	if arr.Err() != nil {
		return
	}
	tmp := arr.Index(left)
	p := left
	i, j := left, right
	for i <= j {
		for j >= p {
			c, err := arr.Index(j).Compare(tmp)
			if err != nil {
				arr.SetErr(err)
				return
			}
			if isAscOrder && c >= 0 {
				j--
				continue
			}
			if !isAscOrder && c <= 0 {
				j--
				continue
			}

			break
		}

		if j >= p {
			t, _ := arr.Index(j).ToInterface()
			arr.SetIndex(p, t)
			p = j
		}

		for i <= p {
			c, err := arr.Index(i).Compare(tmp)
			if err != nil {
				arr.SetErr(err)
				return
			}
			if isAscOrder && c <= 0 {
				i++
				continue
			}
			if !isAscOrder && c >= 0 {
				i++
				continue
			}
			break
		}

		if i <= p {
			t, _ := arr.Index(i).ToInterface()
			arr.SetIndex(p, t)
			p = i
		}
	}

	t, _ := tmp.ToInterface()
	arr.SetIndex(p, t)

	if p-left > 1 {
		arr.qsort(left, p-1, isAscOrder)
	}

	if right-p > 1 {
		arr.qsort(p+1, right, isAscOrder)
	}
}

func (arr *AbsCollection) Sort() ICollection {
	arr.mustSetCompare()
	if arr.Err() != nil {
		return arr
	}

	if arr.Count() == 0 {
		return arr
	}

	arr.qsort(0, arr.Count()-1, true)
	return arr
}

func (arr *AbsCollection) SortDesc() ICollection {
	arr.mustSetCompare()
	if arr.Err() != nil {
		return arr
	}

	if arr.Count() == 0 {
		return arr
	}

	arr.qsort(0, arr.Count()-1, false)
	return arr
}

func (arr *AbsCollection) Join(split string, format ...func(item interface{}) string) string {
	if arr.Err() != nil {
		return ""
	}

	var ret strings.Builder
	for i := 0; i < arr.Count(); i++ {
		if len(format) == 0 {
			o, _ := arr.Index(i).ToInterface()
			ret.WriteString(fmt.Sprintf("%v", o))
		} else {
			f := format[0]
			o, _ := arr.Index(i).ToInterface()
			ret.WriteString(f(o))
		}

		if i != arr.Count()-1 {
			ret.WriteString(split)
		}
	}
	return ret.String()
}

func (arr *AbsCollection) Avg() IMix {
	arr.mustSetCompare().mustNotBeEmpty().mustBeNumType()
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}
	var sum IMix
	var err error
	o0, _ := arr.Index(0).ToInterface()
	sum = NewMix(o0)
	for i := 1; i < arr.Count(); i++ {
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
	arr.mustSetCompare().mustNotBeEmpty().mustBeNumType()
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}

	newArr := arr.Sort()
	if newArr.Count()%2 == 0 {
		imax, err := newArr.Index(newArr.Count()/2 - 1).Add(newArr.Index(newArr.Count() / 2))
		if err != nil {
			return NewErrorMix(err)
		}
		div, err := imax.Div(2)
		if err != nil {
			return NewErrorMix(err)
		}
		return div
	}

	return newArr.Index(newArr.Count()/2 + 1)
}

func (arr *AbsCollection) Mode() IMix {
	arr.mustNotBeEmpty().mustSetCompare()
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
	arr.mustBeNumType().mustNotBeEmpty().mustSetCompare()
	if arr.Err() != nil {
		return NewErrorMix(arr.Err())
	}

	o0, _ := arr.Index(0).ToInterface()
	var mix IMix
	mix = NewMix(o0)
	for i := 1; i < arr.Count(); i++ {
		mix, _ = mix.Add(arr.Index(i))
	}
	return mix
}

func (arr *AbsCollection) Filter(f func(obj interface{}, index int) bool) ICollection {
	if arr.Err() != nil {
		return arr
	}
	if arr.Count() == 0 {
		return arr
	}

	ret := arr.Parent.NewEmpty(arr.Err())
	l := arr.Parent.Count()
	for i := 0; i < l; i++ {
		obj, _ := arr.Index(i).ToInterface()
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
	if arr.Count() == 0 {
		return NewEmptyMix()
	}
	l := arr.Count()
	for i := 0; i < l; i++ {
		obj, _ := arr.Index(i).ToInterface()
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
		t, err := arr.Index(i).ToString()
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
		t, err := arr.Index(i).ToInt64()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsCollection) ToInt32s() ([]int32, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]int32, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t, err := arr.Index(i).ToInt32()
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
		t, err := arr.Index(i).ToInt()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsCollection) ToUInt64s() ([]uint64, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]uint64, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t, err := arr.Index(i).ToUInt64()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsCollection) ToUInt32s() ([]uint32, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]uint32, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t, err := arr.Index(i).ToUInt32()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsCollection) ToUInts() ([]uint, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]uint, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t, err := arr.Index(i).ToUInt()
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
		t, err := arr.Index(i).ToFloat64()
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
		t, err := arr.Index(i).ToFloat32()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsCollection) ToJson() ([]byte, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	if arr.Parent == nil {
		panic("no parent")
	}

	return arr.Parent.ToJson()
}

func (arr *AbsCollection) ToInterfaces() ([]interface{}, error) {
	if arr.Err() != nil {
		return nil, arr.Err()
	}
	ret := make([]interface{}, arr.Count())
	for i := 0; i < arr.Count(); i++ {
		t, err := arr.Index(i).ToInterface()
		if err != nil {
			return nil, err
		}
		ret[i] = t
	}
	return ret, nil
}

func (arr *AbsCollection) ToObjs(objs interface{}) error {
	arr.mustNotBeBaseType()

	return arr.Parent.ToObjs(objs)
}

func (arr *AbsCollection) FromJson(data []byte) error {
	if arr.Err() != nil {
		return arr.Err()
	}
	if arr.Parent == nil {
		panic("no parent")
	}

	return arr.Parent.FromJson(data)
}

func (arr *AbsCollection) MarshalJSON() ([]byte, error) {
	return arr.ToJson()
}

func (arr *AbsCollection) UnmarshalJSON(data []byte) error {
	return arr.FromJson(data)
}

func (arr *AbsCollection) Split(size int) []ICollection {
	if arr.Err() != nil {
		return []ICollection{arr}
	}

	if size <= 0 {
		arr.SetErr(errors.New("size can not be zero"))
		return []ICollection{arr}
	}

	total := arr.Count()
	mod := total % size
	sliceCount := total / size
	if mod > 0 {
		sliceCount = sliceCount + 1
	}
	ret := make([]ICollection, sliceCount)
	for i := 0; i < arr.Count(); i++ {
		idx := i / size
		idm := i % size
		if idm == 0 {
			ret[idx] = arr.NewEmpty()
		}

		ret[idx].Append(arr.Index(i).MustToInterface())
	}
	return ret
}

// Union 两个集合的并集
func (arr *AbsCollection) Union(arr2 ICollection) ICollection {
	arr.mustSetCompare()
	if arr.Err() != nil {
		return arr
	}

	newArr := arr.Copy()
	for i := 0; i < arr2.Count(); i++ {
		o, _ := arr2.Index(i).ToInterface()
		if arr.Contains(o) == false {
			newArr.Append(o)
		}
	}

	return newArr
}

// Intersect 两个集合的交集
func (arr *AbsCollection) Intersect (arr2 ICollection) ICollection {
	arr.mustSetCompare()
	if arr.Err() != nil {
		return arr
	}

	newArr := arr.NewEmpty()
	for i := 0; i < arr.Count(); i++ {
		o, _ := arr.Index(i).ToInterface()
		if arr2.Contains(o) == true {
			newArr.Append(o)
		}
	}

	return newArr
}
