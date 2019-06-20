package collection

import "testing"

func TestStrCollection_FromJson(t *testing.T) {
	data := `["aa", "bb"]`
	objColl := NewStrCollection([]string{})
	err := objColl.FromJson([]byte(data))
	if err != nil {
		t.Error(err)
	}
	objColl.DD()
}