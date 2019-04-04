package collection

import (
	"reflect"
)

type ObjArray struct{
	AbsArray
	objs reflect.Value // 数组对象，是一个slice
	typ reflect.Type // 数组对象每个元素类型
	ptr reflect.Value // 指向数组对象的指针
}
