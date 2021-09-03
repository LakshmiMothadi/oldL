package main

import "fmt"

type dataType struct {
	name string
	age int
}

func basic(){
	 person := dataType{"sonu", 18}
	 person2 := dataType{ "lakshmi", 20}
	 person3 := dataType{"boni", 29}
	 fmt.Println(person , "\n", person2 ,"\n", person3)
}
func basicLevel(){
data :=[]dataType{}
data.name


}

func main(){
	basic()
	basicLevel()
}