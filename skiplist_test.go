package skiplist

import "testing"

func TestSkipList_Len(t *testing.T) {

	sk := NewSkipList()
	if sk.Len() != 0 {
		t.Fatalf("skiplist length expected '%d', got '%d'",  0, string(sk.Len()))
	}

	sk.Set(1,"test1")
	sk.Set(2,"test2")
	if sk.Len() != 2 {
		t.Fatalf("skiplist length expected '%d', got '%d'",  2, string(sk.Len()))
	}

	sk.Set(1,"test1")
	sk.Set(2,"test2")


	if sk.Len() != 2 {
		t.Fatalf("skiplist length expected '%d', got '%d'",  2, string(sk.Len()))
	}


	sk.Delete(1,"test1")
	if sk.Len() != 1 {
		t.Fatalf("skiplist length expected '%d', got '%d'",  1, string(sk.Len()))
	}

	sk.Set(1,"test11")
	sk.Set(1,"test111")

	if sk.Len() != 3 {
		t.Fatalf("skiplist length expected '%d', got '%d'",  3, string(sk.Len()))
	}


	sk.DeleteAll(1)
	if sk.Len() != 1 {
		t.Fatalf("skiplist length expected '%d', got '%d'",  1, string(sk.Len()))
	}
}

func TestSkipList_Set(t *testing.T) {
	sk := NewSkipList()

	sk.Set(1,"test1")
	sk.Set(1,"test2")


	iter :=sk.Iterator()

	iter.Seek(1)

	if iter.Value() != "test1" && iter.Key() != float64(1) {
		t.Fatalf("skiplist length expected '%f:%s', got '%f:%s'",  float64(1), "test1",iter.Key(),iter.Value())
	}


	iter.Next()

	if iter.Value() != "test2" && iter.Key() != float64(1) {
		t.Fatalf("skiplist length expected '%f:%s', got '%f:%s'",  float64(1), "test2",iter.Key(),iter.Value())
	}

}


func TestIter_Seek(t *testing.T) {
	sk := NewSkipList()
	sk.Set(1,"test1")
	sk.Set(2,"test2")
	sk.Set(3,"test2")
	sk.Set(3,"test3")
	sk.Set(4,"test4")
	iter :=sk.SeekToFirst()

	if iter.Value() != "test1" && iter.Key() != float64(1) {
		t.Fatalf("skiplist length expected '%f:%s', got '%f:%s'",  float64(1), "test1",iter.Key(),iter.Value())
	}

	iter =sk.SeekToLast()
	if iter.Value() != "test4" && iter.Key() != float64(4) {
		t.Fatalf("skiplist length expected '%f:%s', got '%f:%s'",  float64(4), "test1",iter.Key(),iter.Value())
	}


	sk.Set(5,"test2")
	iter =sk.SeekToLast()
	if iter.Value() != "test4" && iter.Key() != float64(5) {
		t.Fatalf("skiplist length expected '%f:%s', got '%f:%s'",  float64(5), "test1",iter.Key(),iter.Value())
	}
}

func TestSkipList_Range(t *testing.T) {
	sk := NewSkipList()
	mp := map[float64]string {1:"test1",2:"test2",3:"test3",4:"test4"}
	for k ,v := range mp {
		sk.Set(k,v)
	}

	iter := sk.Range(1,4)

	i := 0
	for iter.Next() {
		i++
		if iter.Value() != mp[float64(i)] && iter.Key() != float64(i) {
			t.Fatalf("skiplist range expected '%f:%s', got '%f:%s'", float64(i), mp[float64(i)], iter.Key(), iter.Value())
		}
	}

	if i != 4 {
		t.Fatalf("skiplist range expected '%f', got '%f'", float64(4), mp[float64(i)])
	}
}

func TestSkipList_IndexRange(t *testing.T) {
	sk := NewSkipList()
	mp := map[float64]string {1:"test1",2:"test2",3:"test3",4:"test4"}
	for k ,v := range mp {
		sk.Set(k,v)
	}

	iter := sk.IndexRange(0,4)

	i := 0
	for iter.Next() {
		i++
		if iter.Value() != mp[float64(i)] && iter.Key() != float64(i) {
			t.Fatalf("skiplist range expected '%f:%s', got '%f:%s'", float64(i), mp[float64(i)], iter.Key(), iter.Value())
		}
	}

	if i != 4 {
		t.Fatalf("skiplist range expected '%f', got '%f'", float64(4), mp[float64(i)])
	}
}