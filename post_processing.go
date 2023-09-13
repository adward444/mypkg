package mypkg

import "fmt"

type PostProcessingBase struct {

}

func (* PostProcessingBase) AfterCheckout() {
	fmt.Println("This is base operate")
}

type PostProcessingAU struct {
	PostProcessingBase
}

func (* PostProcessingAU) AfterCheckout() {
	fmt.Println("This is AU different operate")
}

type PostProcessingJP struct {
	PostProcessingBase
}

func (* PostProcessingJP) AfterCheckout() {
	fmt.Println("This is JP different operate")
}

func PostProcessingTest() {
	p1 := PostProcessingAU{}
	p1.PostProcessingBase.AfterCheckout()
	p1.AfterCheckout()

	p2 := PostProcessingJP{}
	p2.PostProcessingBase.AfterCheckout()
	p2.AfterCheckout()
}
