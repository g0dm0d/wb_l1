package main

import (
	"fmt"
	"reflect"
)

/*
Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/

// метод из стандартной библиотеки
func task14() {
	var teststr string
	fmt.Println(reflect.TypeOf(teststr))

	var testint int
	fmt.Println(reflect.TypeOf(testint))

	var testbool bool
	fmt.Println(reflect.TypeOf(testbool))

	testchan := make(chan interface{})
	fmt.Println(reflect.TypeOf(testchan))
}

// Format Specifier
func task14_2() {
	var teststr string
	fmt.Printf("%T\n", teststr)

	var testint int
	fmt.Printf("%T\n", testint)

	var testbool bool
	fmt.Printf("%T\n", testbool)

	testchan := make(chan interface{})
	fmt.Printf("%T\n", testchan)
}

// через type в inteface
func task14_3(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case chan interface{}:
		fmt.Println("chan interface{}")
	}
}
