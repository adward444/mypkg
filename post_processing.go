package mypkg

import "fmt"

type PostProcessingBase struct {

}

func (* PostProcessingBase) afterCheckout() {
	fmt.Println("This is base operate")
}

type PostProcessingAU struct {
	PostProcessingBase
}

func (* PostProcessingAU) afterCheckout() {
	fmt.Println("This is AU different operate")
}

type PostProcessingJP struct {
	PostProcessingBase
}

func (* PostProcessingJP) afterCheckout() {
	fmt.Println("This is JP different operate")
}

func test() {
	p1 := PostProcessingAU{}
	p1.PostProcessingBase.afterCheckout()
	p1.afterCheckout()

	p2 := PostProcessingJP{}
	p2.PostProcessingBase.afterCheckout()
	p2.afterCheckout()
}
