package main

import (
//	"fmt"
	"github.com/widaT/skiplist/skiplist"
	"fmt"
)
func main()  {

	s :=skiplist.New()
	s.Set(10,"dddd")
	s.Set(1,"dddd")
	s.Set(9,"dddd")
	s.Set(9,"ddddd")
	s.Set(9,"ddddddddd")
	s.Set(9,"ddddddd")
	s.Set(9,"ddddddddd")
	s.Set(9,"ddddddddddddddd")
	s.Set(1,"ddddddd")
	s.Set(11,"dddd")
	s.Set(1,"dddd")
	s.Delete(9,"dddd")


	s.DeleteAll(9)



	for i := s.Iterator();i.Next(); {
		fmt.Println(i.Key(),i.Value())
	}

	//s.DeleteAll(9)
	//s.Delete(11,"dddd")


/*	iter :=s.Seek(1)
	for iter.Next() {
		fmt.Println(iter.Value(),iter.Key())
	}
	iter.Close()*/


/*	iter :=  s.SeekToFirst()

	fmt.Println(iter.Key(),iter.Value())

	iter =  s.SeekToLast()

	fmt.Println(iter.Key(),iter.Value())*/
/*	iter =s.Range(1,11)
	for iter.Next() {
		fmt.Println(iter.Value(),iter.Key())
	}
	iter.Close()*/
}