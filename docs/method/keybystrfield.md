# KeyByStrField

`KeyByStrField(key string) (map[string]interface{}, error)`

KeyByStrField 根据某个字段为key，返回一个map,要求key对应的field是string


``` golang
func TestObjCollection_KeyByStrField(t *testing.T) {
	a1 := Foo{A: "a1", B: 1}
	a2 := Foo{A: "a2", B: 2}
	a3 := Foo{A: "a3", B: 3}
	a4 := Foo{A: "a3", B: 2}
	objColl := NewObjCollection([]Foo{a1, a2, a3, a4})
	maps, err := objColl.KeyByStrField("A")
	if err != nil {
		t.Errorf("err = %v", err)
	}
	if len(maps) != 3 {
		t.Errorf("expected 3 values, got %v", len(maps))
	}

	// find "a2" value must be a2
	if _, ok := maps["a2"]; !ok {
		t.Errorf("expected contains a2 key but not found")
	}

	if _, ok := maps["a2"].(Foo); !ok {
		t.Errorf("expected a2 kay if Foo object but not found")
	}

	outA2 := maps["a2"].(Foo)
	if !reflect.DeepEqual(a2, outA2) {
		t.Errorf("expected a2 foo but not found")
	}
}
```
