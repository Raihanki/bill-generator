package main

import (
	"fmt"
	"os"
)

type bill struct {
	name string
	menu map[string]float64
	tip float64
}

func newBill (name string) bill {
	b := bill{
		name:name,
		menu: map[string]float64{},
		tip: 0,
	}

	return b
}

func (b *bill) formatBill() string {
	var fs = fmt.Sprintf("Halloo %v here is your bill : \n", b.name)

	var total float64 = 0

	for k,v := range b.menu {
		fs += fmt.Sprintf("%-25v %v %.2f \n", k, ":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v %v %.2f \n", "tip", ":", b.tip)

	total += b.tip

	fs += fmt.Sprintf("%-25v %v %.2f \n", "total", ":", total)

	return fs
}

func (b *bill) addTip(tip float64) {
	b.tip = tip
}

func (b *bill) addMenu(name string, price float64) {
	b.menu[name] = price
}

func (b *bill) save() {
	data := []byte(b.formatBill())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Println("Bill was saved to file ...")
}
