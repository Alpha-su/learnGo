package main

import (
	"errors"
	"fmt"
)

func test() {
	//切片学习
	//slice := make([]string, 4, 4)
	//slice2 := []string{"a", "b","c","d","e","f"}
	//slice1 := append(slice, slice2...)
	//println(len(slice1), cap(slice1))
	//for i,v:=range slice1{
	//	println(i,v)
	//}
	//函数学习
	//result, err := sumPositive(3, 5,2,-10)
	//if err == nil{
	//	fmt.Println(result)
	//} else {
	//	fmt.Printf("%T\n", err)
	//	fmt.Println(err)
	//}
	s := &testString{s: "123"}
	fmt.Printf("%T\n", s)
	fmt.Println(s)
	//方法学习
	//age:=Age(25)
	//age.String()
	//age.Modify()
	//age.String()
}

type testString struct {
	s string
}

func sumPositive(params ...int) (int, error) {
	sum := 0
	for _, value := range params {
		if value < 0 {
			//println(reflect.TypeOf(errors.New("cant set any params as a negative number")).String())
			fmt.Printf("%T\n", errors.New("cant set any params as a negative number"))
			return 0, errors.New("cant set any params as a negative number")
		} else {
			sum += value
		}
	}
	return sum, nil
}

type Age uint

func (age Age) String() {
	fmt.Println("the age is", age)
}

func (age *Age) Modify() {
	*age = Age(30)
}
