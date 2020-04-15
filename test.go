package main

import "fmt"

type A struct {
  Ptr1 *B
  Ptr2 *B
  Val B
}

type B struct {
  Str string
}

func main() {
  a := A{
    Ptr1: &B{"ptr-str-1"},
    Ptr2: &B{"ptr-str-2"},
    Val: B{"val-str"},
  }
  fmt.Println(a.Ptr1)
  fmt.Println(a.Ptr2)
  fmt.Println(a.Val)
  demo(a)
  fmt.Println(a.Ptr1)
  fmt.Println(a.Ptr2)
  fmt.Println(a.Val)
}

func demo(a A) {
  // Update a value of a pointer and changes will persist
  a.Ptr1.Str = "new-ptr-str1"
  // Use an entirely new B object and changes won't persist
  a.Ptr2 = &B{"new-ptr-str-2"}
  a.Val.Str = "new-val-str"
}
