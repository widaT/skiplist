package main

import (
	"fmt"
	"github.com/widaT/skiplist/skiplist"
)
func main()  {

	s :=skiplist.New()
	s.Set(10,"dddd")
	s.Set(9,"dddd")
	s.Set(9,"ddddd")
	s.Set(9,"ddddddddd")
	s.Set(9,"ddddddd")
	s.Set(9,"ddddddddd")
	s.Set(9,"ddddddddddddddd")

	s.Set(11,"dddd")
	//s.Delete(9,"dddd")



	/*s.DeleteAll(9)
	s.Delete(11,"dddd")*/
	/*for iter := s.Iterator();iter.Next(); {
		fmt.Println(iter.Value(),iter.Key())
	}*/


	iter :=s.Seek(9)
	for ;iter.Next(); {
		fmt.Println(iter.Value(),iter.Key())
	}
	iter.Close()



}