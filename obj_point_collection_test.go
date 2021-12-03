package collection

import (
	"testing"
)

type FooBar struct {
	Foo string
	Bar int
}

func FooBarCompare(a interface{}, b interface{}) int {
	aobj := a.(*FooBar)
	bobj := b.(*FooBar)
	return aobj.Bar - bobj.Bar
}

func InitFooObjPoints() []*FooBar {
	return []*FooBar{
		{
			Foo: "astring",
			Bar: 1,
		},
		{
			Foo: "bstring",
			Bar: 2,
		},
	}
}

func TestObjPointCollection_Normal(t *testing.T) {
	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	// 	[Append](#Append) 挂载一个元素到当前Collection
	{
		count := coll.Copy().Append(&FooBar{
			Foo: "cstring",
			Bar: 3,
		}).Count()
		if count != 3 {
			t.Fatal("append error")
		}
	}

	// [Avg](#Avg) 返回Collection的数值平均数
	{

	}

	// [Contain](#Contain) 判断一个元素是否在Collection中
	{
		obj := objs[0]
		if coll.Contains(obj) != true {
			t.Fatal("contains error")
		}
	}

	// [Copy](#Copy) 根据当前的数组，创造出一个同类型的数组
	{
		if coll.Copy().Count() != 2 {
			t.Fatal("copy error")
		}
	}

	// [DD](#DD) 按照友好的格式展示Collection
	{
		// coll.DD()
	}

	// [Diff](#Diff) 获取前一个Collection不在后一个Collection中的元素
	{
		coll2 := NewObjPointCollection(objs).SetCompare(FooBarCompare)
		newColl := coll.Diff(coll2)
		if newColl.Count() != 0 {
			t.Fatal("diff error")
		}
	}

	// [Each](#Each) 对Collection中的每个函数都进行一次函数调用
	{
		sum := 0
		coll.Each(func(item interface{}, key int) {
			obj := item.(*FooBar)
			sum = obj.Bar + sum
		})
		if sum != 3 {
			t.Fatal("each error")
		}
	}

	// [Every](#Every) 判断Collection中的每个元素是否都符合某个条件
	{
		check := coll.Every(func(item interface{}, key int) bool {
			obj := item.(*FooBar)
			if obj.Bar > 0 {
				return true
			}
			return false
		})
		if check != true {
			t.Fatal("every error")
		}

		check = coll.Every(func(item interface{}, key int) bool {
			obj := item.(*FooBar)
			if obj.Bar > 1 {
				return true
			}
			return false
		})
		if check != false {
			t.Fatal("every error")
		}
	}

	// [ForPage](#ForPage) 将Collection函数进行分页
	{
		newColl := coll.ForPage(1, 1)
		if newColl.Count() != 1 {
			t.Fatal("for page error")
		}
	}

	// [Filter](#Filter) 根据过滤函数获取Collection过滤后的元素
	{
		newColl := coll.Filter(func(obj interface{}, index int) bool {
			ob := obj.(*FooBar)
			if ob.Bar > 1 {
				return true
			}
			return false
		})
		if newColl.Count() != 1 {
			t.Fatal("filter error")
		}
	}

	// [First](#First) 获取符合过滤条件的第一个元素
	{
		first := coll.First(func(obj interface{}, index int) bool {
			ob := obj.(*FooBar)
			if ob.Bar > 1 {
				return true
			}
			return false
		})
		f := first.MustToInterface().(*FooBar)
		if f.Bar != 2 {
			t.Fatal("first error")
		}

	}

	// [Index](#Index) 获取元素中的第几个元素，下标从0开始
	{
		first := coll.Index(0)
		f := first.MustToInterface().(*FooBar)
		if f.Bar != 1 {
			t.Fatal("Index error")
		}

	}

	// [IsEmpty](#IsEmpty) 判断一个Collection是否为空
	{
		empty := coll.IsEmpty()
		if empty != false {
			t.Fatal("empty error")
		}
	}

	// [IsNotEmpty](#IsNotEmpty) 判断一个Collection是否为空
	{
		empty := coll.IsNotEmpty()
		if empty != true {
			t.Fatal("is not empty error")
		}
	}

	// [Join](#Join) 将Collection中的元素按照某种方式聚合成字符串
	{
		str := coll.Join(",", func(item interface{}) string {
			ob := item.(*FooBar)
			return ob.Foo
		})
		if str != "astring,bstring" {
			t.Fatal("join error")
		}
	}

	// [Last](#Last) 获取该Collection中满足过滤的最后一个元素
	{
		last := coll.Last().MustToInterface().(*FooBar)
		if last.Foo != "bstring" {
			t.Fatal("last error")
		}
	}

	// [Merge](#Merge) 将两个Collection的元素进行合并
	{
		collCopy := coll.Copy()
		foobar2 := []*FooBar{
			{
				Foo: "cstring",
				Bar: 3,
			},
			{
				Foo: "dstring",
				Bar: 4,
			},
		}
		coll2 := NewObjPointCollection(foobar2)
		coll3 := collCopy.Merge(coll2)
		if coll3.Count() != 4 {
			t.Fatal("merge error")
		}

	}

	// [Map](#Map) 对Collection中的每个函数都进行一次函数调用
	{
		collCopy := coll.Copy()
		newColl := collCopy.Map(func(item interface{}, key int) interface{} {
			ob := item.(*FooBar)
			return ob.Foo
		})
		strs := newColl.Join(",")
		if strs != "astring,bstring" {
			t.Fatal("map error")
		}
	}

	// [Mode](#Mode) 获取Collection中的众数
	{
		mod := coll.Mode().MustToInterface().(*FooBar)
		if mod.Bar != 1 {
			t.Fatal("mod error")
		}
	}

	// [Max](#Max) 获取Collection中的最大元素
	{
		max := coll.Max().MustToInterface().(*FooBar)
		if max.Bar != 2 {
			t.Fatal("max error")
		}
	}

	// [Min](#Min) 获取Collection中的最小元素
	{
		min := coll.Min().MustToInterface().(*FooBar)
		if min.Bar != 1 {
			t.Fatal("min error")
		}
	}

	// [Median](#Median) 获取Collection的中位数
	{
		err := coll.Median().Err()
		if err == nil {
			t.Fatal("median error")
		}
		coll.SetErr(nil)
	}

	// [Nth](#Nth) 获取从offset偏移量开始的每第n个
	{
		newColl := coll.Nth(2, 0)
		if newColl.Count() != 1 {
			t.Fatal("nth error")
		}
	}

	// [Pad](#Pad) 填充Collection数组
	{
		err := coll.Pad(5, &FooBar{}).Err()
		if err == nil {
			t.Fatal("pad error")
		}
		coll.SetErr(nil)
	}

	// [Pop](#Pop) 从Collection右侧弹出一个元素
	{
		obj := coll.Copy().Pop().MustToInterface().(*FooBar)
		if obj.Bar != 2 {
			t.Fatal("pop error")
		}
	}

	// [Push](#Push) 往Collection的右侧推入一个元素
	{
		newColl := coll.Copy().Push(&FooBar{
			Foo: "cstring",
			Bar: 3,
		})
		if newColl.Count() != 3 {
			t.Fatal("push error")
		}
	}

	// [Prepend](#Prepend) 往Collection左侧加入元素
	{
		newColl := coll.Copy().Prepend(&FooBar{
			Foo: "cstring",
			Bar: 3,
		})
		if newColl.Index(0).MustToInterface().(*FooBar).Bar != 3 {
			t.Fatal("prepend error")
		}
	}

	// [Pluck](#Pluck) 将对象数组中的某个元素提取出来组成一个新的Collection
	{
		newColl := coll.Pluck("Foo")
		if newColl.Index(0).MustToString() != "astring" {
			t.Fatal("pluck error")
		}
	}

	// [Reject](#Reject) 将满足过滤条件的元素删除
	{
		newColl := coll.Reject(func(item interface{}, key int) bool {
			obj := item.(*FooBar)
			if obj.Bar > 1 {
				return true
			}
			return false
		})
		if newColl.Index(0).MustToInterface().(*FooBar).Bar != 1 {
			t.Fatal("reject error")
		}
	}

	// [Reduce](#Reduce) 对Collection中的所有元素进行聚合计算
	{
		err := coll.Reduce(func(carry, item IMix) IMix {
			return carry
		}).Err()
		if err == nil {
			t.Fatal("reduce error")
		}
		coll.SetErr(nil)
	}

	// [Random](#Random) 随机获取Collection中的元素
	{
		obj := coll.Random()
		foobar := obj.MustToInterface().(*FooBar)
		if foobar == nil {
			t.Fatal("random error")
		}
	}

	// [Reverse](#Reverse) 将Collection数组进行转置
	{
		newColl := coll.Copy().Reverse()
		if newColl.Index(0).MustToInterface().(*FooBar).Bar != 2 {
			t.Fatal("error")
		}
	}

	// [Slice](#Slice) 获取Collection中的片段
	{
		newColl := coll.Copy().Slice(1)
		if newColl.Count() != 1 {
			t.Fatal("slice error")
		}
	}

	// [Search](#Search) 查找Collection中第一个匹配查询元素的下标
	{
		search := &FooBar{
			Foo: "cstring",
			Bar: 2,
		}
		index := coll.Search(search)
		if index != 1 {
			t.Fatal("search error")
		}
	}

	// [Sort](#Sort) 将Collection中的元素进行升序排列输出
	{
		newColl := coll.Copy().Sort()
		if newColl.Index(1).MustToInterface().(*FooBar).Bar != 2 {
			t.Fatal("sort error")
		}
	}

	// [SortDesc](#SortDesc) 将Collection中的元素按照降序排列输出
	{
		newColl := coll.Copy().SortDesc()
		if newColl.Index(1).MustToInterface().(*FooBar).Bar != 1 {
			t.Fatal("sortDesc error")
		}
	}

	// [Sum](#Sum) 返回Collection中的元素的和
	{
		if coll.Sum().Err() == nil {
			t.Fatal("sum error")
		}
		coll.SetErr(nil)
	}

	// [Shuffle](#Shuffle) 将Collection中的元素进行乱序排列
	{
		newColl := coll.Shuffle()
		if newColl.Count() != 2 {
			t.Fatal("shuffle error")
		}
	}

	// [SortBy](#SortBy) 根据对象数组中的某个元素进行Collection升序排列
	{
		newColl := coll.Copy().SortBy("Bar")
		if newColl.Index(1).MustToInterface().(*FooBar).Bar != 2 {
			t.Fatal("sortby error")
		}
	}

	// [SortByDesc](#SortByDesc) 根据对象数组中的某个元素进行Collection降序排列
	{
		newColl := coll.Copy().SortByDesc("Bar")
		if newColl.Index(1).MustToInterface().(*FooBar).Bar != 1 {
			t.Fatal("sortbydesc error")
		}
	}

	// [ToInts](#ToInts) 将Collection变化为int数组

	// [ToInt64s](#ToInt64s) 将Collection变化为int64数组

	// [ToFloat64s](#ToFloat64s) 将Collection变化为float64数组

	// [ToFloat32s](#ToFloat32s) 将Collection变化为float32数组

	// [ToMixs](#ToMixs) 将Collection变化为Mix数组

	// [ToInterfaces]
	{
		objs, _ := coll.ToInterfaces()
		if len(objs) != 2 {
			t.Fatal("tointerface error")
		}
	}

	{
		out := []*FooBar{}
		err := coll.ToObjs(&out)
		if err != nil {
			t.Fatal("to objs error", err.Error())
		}
		if len(out) != 2 {
			t.Fatal("to objs len error")
		}
	}

	// [Unique](#Unique) 将Collection中重复的元素进行合并
	{
		c := &FooBar{
			Foo: "cstring",
			Bar: 2,
		}
		newColl := coll.Append(c)
		if newColl.Unique().Count() != 2 {
			t.Fatal("unique error")
		}

	}

	{
		newColl := coll.Copy().SortBy("Bar")
		newArr := []*FooBar{}
		err := newColl.ToObjs(&newArr)
		if err != nil {
			t.Fatal(err)
		}
		if len(newArr) != newColl.Count() {
			t.Fatal("len error")
		}

	}

}

func TestObjPointCollection_ToObjs(t *testing.T) {
	a1 := &Foo{A: "a1", B: 1}
	a2 := &Foo{A: "a2", B: 2}
	a3 := &Foo{A: "a3", B: 3}

	bArr := []*Foo{}
	objColl := NewObjPointCollection([]*Foo{a1, a2, a3})
	err := objColl.ToObjs(&bArr)
	if err != nil {
		t.Fatal(err)
	}
	if len(bArr) != 3 {
		t.Fatal("toObjs error len")
	}
	if bArr[1].A != "a2" {
		t.Fatal("toObjs error copy")
	}
}

func TestObjCollection_ToObjs2(t *testing.T) {
	a1 := Foo{A: "a1", B: 1}
	a2 := Foo{A: "a2", B: 2}
	a3 := Foo{A: "a3", B: 3}

	bArr := []Foo{}

	objColl := NewObjCollection([]Foo{a1, a2, a3})
	err := objColl.ToObjs(&bArr)
	if err != nil {
		t.Fatal(err)
	}
	if len(bArr) != 3 {
		t.Fatal("toObjs error len")
	}
	if bArr[1].A != "a2" {
		t.Fatal("toObjs error copy")
	}
}

func TestObjCollection_GroupBy2(t *testing.T) {
	a1 := &Foo{A: "a1", B: 1}
	a2 := &Foo{A: "a2", B: 2}
	a3 := &Foo{A: "a3", B: 3}
	a4 := &Foo{A: "a3", B: 2}
	objColl := NewObjPointCollection([]*Foo{a1, a2, a3, a4})
	groupBy := objColl.GroupBy(func(item interface{}, i2 int) interface{} {
		foo := item.(*Foo)
		return foo.A
	})
	for k, collection := range groupBy {
		t.Log(k)
		collection.DD()
	}
}
