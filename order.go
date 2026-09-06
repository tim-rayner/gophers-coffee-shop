package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type OrderItem struct {
	drink       Drink
	syrup       Syrup
	size        Size
	temperature Temperature
}

type Order struct {
	Items []OrderItem
}

type Receipt struct {
	Order Order
	Total float64
}

const receiptInnerWidth = 44

func CalculateTotal(order Order) float32 {
	var totalPrice float32

	for _, item := range order.Items {
		totalPrice = totalPrice + float32(item.drink.price) + float32(item.syrup.price) + float32(item.size.price) + float32(item.temperature.price)
	}

	return totalPrice
}

func (item OrderItem) lineTotal() float64 {
	return item.drink.price + item.syrup.price + item.size.price + item.temperature.price
}

func (order *Order) PrintReceipt() {
	fmt.Println()
	printReceiptRule("╔", "╗")
	printReceiptCentered("GOPHERS COFFEE SHOP")
	printReceiptCentered("~ freshly brewed ~")
	printReceiptCentered("RECEIPT")
	printReceiptRule("╠", "╣")

	if len(order.Items) == 0 {
		printReceiptCentered("(empty order)")
	}

	for i, item := range order.Items {
		if i > 0 {
			printReceiptInner(strings.Repeat("·", receiptInnerWidth-2))
		}

		title := fmt.Sprintf("#%d  %s %s %s", i+1, item.size.name, item.temperature.name, item.drink.name)
		printReceiptText(" " + title)
		printReceiptRow(item.drink.name, item.drink.price)
		printReceiptRow(item.syrup.name+" syrup", item.syrup.price)
		printReceiptRow(item.size.name, item.size.price)
		if item.temperature.price > 0 {
			printReceiptRow(item.temperature.name, item.temperature.price)
		}
		printReceiptRow("item", item.lineTotal())
	}

	printReceiptRule("╠", "╣")
	printReceiptRow("TOTAL", float64(CalculateTotal(*order)))
	printReceiptRule("╠", "╣")
	printReceiptCentered("Thank you for your custom")
	printReceiptCentered("See you next brew!")
	printReceiptRule("╚", "╝")
	fmt.Println()
}

func printReceiptRule(left, right string) {
	fmt.Println(left + strings.Repeat("═", receiptInnerWidth) + right)
}

func printReceiptCentered(text string) {
	fmt.Println("║" + centerText(text, receiptInnerWidth) + "║")
}

func printReceiptText(text string) {
	runes := []rune(text)
	if len(runes) > receiptInnerWidth {
		runes = runes[:receiptInnerWidth]
		text = string(runes)
	}
	pad := receiptInnerWidth - utf8.RuneCountInString(text)
	fmt.Println("║" + text + strings.Repeat(" ", pad) + "║")
}

func printReceiptInner(text string) {
	printReceiptText(" " + text + " ")
}

func printReceiptRow(name string, price float64) {
	priceText := fmt.Sprintf("£%.2f", price)
	gap := receiptInnerWidth - 4 - utf8.RuneCountInString(name) - utf8.RuneCountInString(priceText)
	if gap < 2 {
		gap = 2
	}
	line := "  " + name + " " + strings.Repeat("·", gap-2) + " " + priceText + "  "
	printReceiptText(line)
}
