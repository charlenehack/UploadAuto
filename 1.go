/* ========== 封装 ========== */
// model/person.go
package model

import "fmt"

type person struct {
	Name string
	age int
	salary float64
}

func NewPerson(name string) *person  {
	return &person{
		Name: name,
	}
}

// 为了让外部能访问到首写小字母的本包，编写一对SetXxx方法和GetXxx方法用于赋值和取值
func (p *person) SetAge(age int)  {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确。")
	}
}

func (p * person) GetAge() int {
	return p.age
}

// main/main.go
package main

import (
"fmt"
"model"
)

func main()  {
	p := model.NewPerson("aaa")
	p.SetAge(151)
	fmt.Println(p.GetAge())
}
