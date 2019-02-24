package main

import "fmt"

func main() {
	t1 := &[]string{"aaaaa"}
	t2 := "ddd"
	t3 := make([]string, len(append(*t1, t2)))
	res := copy(t3, append(*t1, t2))
	fmt.Println(res)
	fmt.Println(append(*t1, t2))
}
