package collection

import (
	"testing"
)

// [Append](#Append) 挂载一个元素到当前Collection
func Benchmark_Append(b *testing.B) {
	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Append(&FooBar{
			Foo: "cstring",
			Bar: 3,
		})
	}
}

// [Avg](#Avg) 返回Collection的数值平均数，只能数值类型coll调用

// [Contain](#Contain) 判断一个元素是否在Collection中。非数值类型必须设置对象compare方法。
func Benchmark_Contain(b *testing.B) {
	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Contains(&FooBar{
			Foo: "cstring",
			Bar: 2,
		})
	}
}

// [Copy](#Copy) 根据当前的数组，创造出一个同类型的数组
func Benchmark_Copy(b *testing.B) {
	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Copy()
	}
}

// [DD](#DD) 按照友好的格式展示Collection

// [Diff](#Diff) 获取前一个Collection不在后一个Collection中的元素, 只能数值类型Diff调用

func Benchmark_Diff(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)
	coll2 := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Diff(coll2)
	}
}

// [Each](#Each) 对Collection中的每个函数都进行一次函数调用

func Benchmark_Each(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Each(func(item interface{}, key int) {
		})
	}
}

// [Every](#Every) 判断Collection中的每个元素是否都符合某个条件
func Benchmark_Every(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Every(func(item interface{}, key int) bool {
			return true
		})
	}
}

// [ForPage](#ForPage) 将Collection函数进行分页
func Benchmark_ForPage(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.ForPage(0, 1)
	}
}

// [Filter](#Filter) 根据过滤函数获取Collection过滤后的元素
func Benchmark_Filter(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Filter(func(item interface{}, key int) bool {
			return true
		})
	}
}

// [First](#First) 获取符合过滤条件的第一个元素
func Benchmark_First(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.First()
	}
}

// [Index](#Index) 获取元素中的第几个元素，下标从0开始
func Benchmark_Index(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Index(0)
	}
}

// [IsEmpty](#IsEmpty) 判断一个Collection是否为空
func Benchmark_IsEmpty(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.IsEmpty()
	}
}

// [IsNotEmpty](#IsNotEmpty) 判断一个Collection是否为空
func Benchmark_IsNotEmpty(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.IsNotEmpty()
	}
}

// [Join](#Join) 将Collection中的元素按照某种方式聚合成字符串
func Benchmark_Join(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Join(",", func(item interface{}) string {
			ob := item.(*FooBar)
			return ob.Foo
		})
	}
}

// [Last](#Last) 获取该Collection中满足过滤的最后一个元素
func Benchmark_Last(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Last()
	}
}

// [Merge](#Merge) 将两个Collection的元素进行合并
func Benchmark_Merge(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)
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
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Copy().Merge(coll2)
	}
}

// [Map](#Map) 对Collection中的每个函数都进行一次函数调用
func Benchmark_Map(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Map(func(item interface{}, key int) interface{} {
			ob := item.(*FooBar)
			return ob.Foo
		})
	}
}

// [Mode](#Mode) 获取Collection中的众数

// [Max](#Max) 获取Collection中的最大元素
func Benchmark_Max(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Max()
	}
}

// [Min](#Min) 获取Collection中的最小元素
func Benchmark_Min(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Min()
	}
}

// [Median](#Median) 获取Collection的中位数
func Benchmark_Median(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Median()
	}
}

// [Nth](#Nth) 获取从offset偏移量开始的每第n个
func Benchmark_Nth(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Nth(2, 1)
	}
}

// [Pad](#Pad) 填充Collection数组

// [Pop](#Pop) 从Collection右侧弹出一个元素
func Benchmark_Pop(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Pop()
	}
}

// [Push](#Push) 往Collection的右侧推入一个元素
func Benchmark_Push(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Push(&FooBar{
			Foo: "",
			Bar: 0,
		})
	}
}

// [Prepend](#Prepend) 往Collection左侧加入元素
func Benchmark_Prepend(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Prepend(&FooBar{
			Foo: "",
			Bar: 0,
		})
	}
}

// [Pluck](#Pluck) 将对象数组中的某个元素提取出来组成一个新的Collection
func Benchmark_Pluck(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Pluck("Foo")
	}
}

// [Reject](#Reject) 将满足过滤条件的元素删除
func Benchmark_Reject(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Reject(func(item interface{}, key int) bool {
			return true
		})
	}
}

// [Reduce](#Reduce) 对Collection中的所有元素进行聚合计算

// [Random](#Random) 随机获取Collection中的元素
func Benchmark_Random(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Random()
	}
}

// [Reverse](#Reverse) 将Collection数组进行转置
func Benchmark_Reverse(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Reverse()
	}
}

// [Slice](#Slice) 获取Collection中的片段
func Benchmark_Slice(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Slice(0, 1)
	}
}

// [Search](#Search) 查找Collection中第一个匹配查询元素的下标
func Benchmark_Search(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Search(&FooBar{
			Foo: "cstring",
			Bar: 3,
		})
	}
}

// [Sort](#Sort) 将Collection中的元素进行升序排列输出
func Benchmark_Sort(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Sort()
	}
}

// [SortDesc](#SortDesc) 将Collection中的元素按照降序排列输出
func Benchmark_SortDesc(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.SortDesc()
	}
}

// [Sum](#Sum) 返回Collection中的元素的和

// [Shuffle](#Shuffle) 将Collection中的元素进行乱序排列
func Benchmark_Shuffle(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Shuffle()
	}
}

// [SortBy](#SortBy) 根据对象数组中的某个元素进行Collection升序排列
func Benchmark_SortBy(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.SortBy("Bar")
	}
}

// [SortByDesc](#SortByDesc) 根据对象数组中的某个元素进行Collection降序排列
func Benchmark_SortByDesc(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.SortByDesc("Bar")
	}
}

// [ToInts](#ToInts) 将Collection变化为int数组

// [ToInt64s](#ToInt64s) 将Collection变化为int64数组

// [ToFloat64s](#ToFloat64s) 将Collection变化为float64数组

// [ToFloat32s](#ToFloat32s) 将Collection变化为float32数组

// [ToMixs](#ToMixs) 将Collection变化为Mix数组

// [ToInterfaces](#ToInterfaces) 将collection变化为interface{}数组

// [Unique](#Unique) 将Collection中重复的元素进行合并
func Benchmark_Unique(b *testing.B) {

	objs := InitFooObjPoints()
	coll := NewObjPointCollection(objs).SetCompare(FooBarCompare)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		coll.Unique()
	}
}
