# Split 

`Split(size int) []ICollection`

Split 按照size个数进行分组

```

func TestIntCollection_Split(t *testing.T) {
	intColl := NewIntCollection([]int{1, 2, 3, 4, 5, 6, 7, 8})
	ret := intColl.Split(3)

	if len(ret) != 3 {
		t.Fatal("split len not right")
	}

	ret[0].DD()
	ret[1].DD()
	ret[2].DD()

	if ret[0].Count() != 3 || ret[2].Count() != 2 {
		t.Fatal("split not right")
	}

	int2Coll := NewIntCollection([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	ret2 := int2Coll.Split(3)
	if len(ret2) != 3 {
		t.Fatal("split not right")
	}

	if ret2[2].Count() != 3 {
		t.Fatal("split not right")
	}

	ret2[0].DD()
	ret2[1].DD()
	ret2[2].DD()
}

/*
=== RUN   TestIntCollection_Split
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
IntCollection(3):{
	0:	4
	1:	5
	2:	6
}
IntCollection(2):{
	0:	7
	1:	8
}
IntCollection(3):{
	0:	1
	1:	2
	2:	3
}
IntCollection(3):{
	0:	4
	1:	5
	2:	6
}
IntCollection(3):{
	0:	7
	1:	8
	2:	9
}
--- PASS: TestIntCollection_Split (0.00s)
*/
```