package mypkg

import (
	"fmt"
	"reflect"
)

type TaxCalculator interface {
	Calc() float64
}

type IncludeTaxMode struct {
	Price   float64
	TaxRate float64
}

type WithoutTaxMode struct {
	Price float64
}

func (t *IncludeTaxMode) Calc() float64 {
	return t.Price - t.Price/(1+t.TaxRate)
}

func (t *WithoutTaxMode) Calc() float64 {
	return t.Price * 0.06
}

func CalcTax(tc TaxCalculator) float64 {
	return tc.Calc()
}


func test1()  {
	siteMode := "IncludeTax"

	//if siteMode == "IncludeTax" {
	//	mode1 := &IncludeTaxMode{Price: 1000, TaxRate: 0.2}
	//	tax1 := CalcTax(mode1)
	//	fmt.Println(tax1)
	//} else {
	//	mode2 := &WithoutTaxMode{Price: 1000}
	//	tax2 := CalcTax(mode2)
	//	fmt.Println(tax2)
	//}

	var mode TaxCalculator
	if siteMode == "IncludeTax" {
		mode = &IncludeTaxMode{Price: 1000, TaxRate: 0.2}
	} else {
		mode = &WithoutTaxMode{Price: 1000}
	}
	tax := CalcTax(mode)
	fmt.Println(tax)
}

func test2()  {
	structName := "IncludeTaxMode"
	structType := reflect.TypeOf(structName)

	//structType := reflect.TypeOf(IncludeTaxMode{})
	structPtr := reflect.New(structType)
	structValue := structPtr.Elem()
	structValue.FieldByName("Price").SetFloat(2000)

	fmt.Println(structValue.Interface().(IncludeTaxMode))
}


var typeRegistry = make(map[string]reflect.Type)

func init() {
	registerType((*IncludeTaxMode)(nil))
	registerType((*WithoutTaxMode)(nil))
}

func registerType(elem interface{}) {
	t := reflect.TypeOf(elem).Elem()
	typeRegistry[t.Name()] = t
}

func newStruct(name string) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	//fmt.Println(elem)
	if !ok {
		return nil, false
	}
	return reflect.New(elem).Elem().Interface(), true
}

func test3()  {
	structName := "IncludeTaxMode"
	s, ok := newStruct(structName)
	if !ok {
		return
	}

	//t, ok := s.(IncludeTaxMode)
	//if !ok {
	//	return
	//}
	//t.Price = 1000
	//t.TaxRate = 0.2
	//
	//tax1 := CalcTax(&t)
	//fmt.Println(tax1)

	p, ok := s.(TaxCalculator)
	fmt.Println(p)
	p.Calc()

}


