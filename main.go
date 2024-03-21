package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInfo(r *bufio.Reader, promt string) (string, error) {
	fmt.Print(promt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promtOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInfo(reader, "Choose Option (a - add item, s - save bill, t - add tip) : ")

	switch opt {
		case "a":
			itemName, _ := getInfo(reader, "Enter menu name : ")
			itemPrice, _ := getInfo(reader, "Enter menu price : ")

			p, err := strconv.ParseFloat(itemPrice, 64)
			if err != nil {
				fmt.Println("The price must be a number ...")
				promtOptions(b)
			}
			b.addMenu(itemName, p)

			fmt.Println("Item added - ", itemName, p)
			promtOptions(b)
		case "t":
			tip, _ := getInfo(reader, "Enter tip amount ($) : ")
			t, err := strconv.ParseFloat(tip, 64)
			if err != nil {
				fmt.Println("The tip must be a number ...")
				promtOptions(b)
			}
			b.addTip(t)

			fmt.Println("Tip added - ", t)
			promtOptions(b)
		case "s":
			fmt.Println("Saving the bill ... "+ b.name)
			b.save()
		default:
			fmt.Println("You choose invalid option ...")
			promtOptions(b)
	}
}

func generateBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInfo(reader, "Enter your name for the bill: ")

	b := newBill(name)
	fmt.Println("Creating bill for", name)

	promtOptions(b)

	return b
}

func main() {
	generateBill()
}
