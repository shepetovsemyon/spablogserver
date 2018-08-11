package main

import (
	"./modules/mycontext"
)

func main() {
	ctx := mycontext.CreateDefContext()

	Run("../client/dist", ctx)
}

//func printer(in <-chan string) {
//	//fmt.Println(<-in)
//	for v := range in {
//		fmt.Println(v)
//	}
//
