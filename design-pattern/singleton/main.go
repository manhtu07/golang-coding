package main

import (
	"fmt"
	"sync"
)

type db struct {
	name string
}

var once sync.Once
var instance *db

func InitInstance() *db {
	if instance == nil {
		once.Do(func() {
			t := db{name: "connect string"}
			instance = &t
			fmt.Println("initialzie instance")
		})
	} else {
		fmt.Println("instance existed")
	}
	return instance
}

type dependencyInjection struct {
	ChildI ChildInterface
}

type ChildInterface interface {
}

type DIInterface interface {
	D1() bool
}

func NewDependencyInjection(ChildI ChildInterface) DIInterface {
	return &dependencyInjection{
		ChildI: ChildI,
	}
}

func (d *dependencyInjection) D1() bool {
	return true
}

func main() {
	InitInstance()
	InitInstance()
	InitInstance()
}
