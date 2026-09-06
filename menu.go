package main

import (
	"fmt"
	"sort"
	"strings"
	"unicode/utf8"
)

type Menu struct {
	drinks       map[string]float64
	syrups       map[string]float64
	sizes        map[string]float64
	temperatures map[string]float64
}

type MenuItem struct {
	name  string
	price float64
}

var CafeMenu = Menu{
	drinks: map[string]float64{
		"Decaf":       2.00,
		"Espresso":    3.00,
		"Latte":       4.00,
		"Cappuccino":  5.00,
		"Americano":   6.00,
		"Mocha":       7.00,
		"Macchiato":   8.00,
		"Iced Coffee": 9.00,
	},
	syrups: map[string]float64{
		"Hazelnut":              1.00,
		"Caramel":               1.00,
		"Vanilla":               1.00,
		"Chocolate":             1.00,
		"Maple":                 1.00,
		"Cinnamon":              1.05,
		"Pumpkin Spice":         1.00,
		"Irish Cream":           1.00,
		"Peppermint":            1.00,
		"Gophers Special Syrup": 200.00,
	},
	sizes: map[string]float64{
		"Small":  0.50,
		"Medium": 1.00,
		"Grande": 1.50,
		"Venti":  2.00,
	},
	temperatures: map[string]float64{
		"Hot":  0.00,
		"Warm": 0.00,
		"Iced": 0.00,
	},
}

const menuInnerWidth = 44
const menuFullInnerWidth = menuInnerWidth*2 + 1 // two columns plus the middle ║

/**
* ok you got me... I used cursor to format my menu
 */

func (m Menu) PrintMenu() {
	printFullRule("╔", "╗")
	printFullCentered("GOPHERS COFFEE SHOP")
	printFullCentered("~ freshly brewed ~")
	printSplitRule("╠", "╦", "╣")
	printColumns(sectionLines("DRINKS", m.drinks), sectionLines("SYRUPS", m.syrups))
	printSplitRule("╠", "╬", "╣")
	printColumns(sectionLines("SIZES", m.sizes), sectionLines("TEMPERATURE", m.temperatures))
	printSplitRule("╚", "╩", "╝")
}

func sectionLines(title string, items map[string]float64) []string {
	lines := []string{
		centerText("· "+title+" ·", menuInnerWidth),
		blankInner(),
	}
	for _, item := range sortedMenuItems(items) {
		lines = append(lines, padMenuRow(formatMenuItem(item, menuInnerWidth-2)))
	}
	return append(lines, blankInner())
}

func sortedMenuItems(items map[string]float64) []MenuItem {
	out := make([]MenuItem, 0, len(items))
	for name, price := range items {
		out = append(out, MenuItem{name: name, price: price})
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].price != out[j].price {
			return out[i].price < out[j].price
		}
		return out[i].name < out[j].name
	})
	return out
}

func printColumns(left, right []string) {
	n := len(left)
	if len(right) > n {
		n = len(right)
	}
	for i := 0; i < n; i++ {
		l, r := blankInner(), blankInner()
		if i < len(left) {
			l = left[i]
		}
		if i < len(right) {
			r = right[i]
		}
		fmt.Println("║" + l + "║" + r + "║")
	}
}

func formatMenuItem(item MenuItem, width int) string {
	name := item.name
	if item.price <= 0 {
		return name
	}

	price := fmt.Sprintf("£%.2f", item.price)
	gap := width - utf8.RuneCountInString(name) - utf8.RuneCountInString(price)
	if gap < 2 {
		gap = 2
	}
	return name + " " + strings.Repeat("·", gap-2) + " " + price
}

func padMenuRow(text string) string {
	max := menuInnerWidth - 2
	runes := []rune(text)
	if len(runes) > max {
		runes = runes[:max]
	}
	text = string(runes)
	pad := max - utf8.RuneCountInString(text)
	return " " + text + strings.Repeat(" ", pad) + " "
}

func centerText(text string, width int) string {
	padding := width - utf8.RuneCountInString(text)
	if padding < 0 {
		padding = 0
	}
	left := padding / 2
	right := padding - left
	return strings.Repeat(" ", left) + text + strings.Repeat(" ", right)
}

func printFullCentered(text string) {
	fmt.Println("║" + centerText(text, menuFullInnerWidth) + "║")
}

func printFullRule(left, right string) {
	fmt.Println(left + strings.Repeat("═", menuFullInnerWidth) + right)
}

func printSplitRule(left, mid, right string) {
	col := strings.Repeat("═", menuInnerWidth)
	fmt.Println(left + col + mid + col + right)
}

func blankInner() string {
	return strings.Repeat(" ", menuInnerWidth)
}


func ValdidateItem(orderItem *OrderItem, usersChoice string, itemType string) bool {
	switch itemType {
	case "drink":
		if price, ok := findPrice(CafeMenu.drinks, usersChoice); ok {
			orderItem.drink = Drink{
				name:  usersChoice,
				price: price,
			}
			return true
		}
	case "syrup":
		if price, ok := findPrice(CafeMenu.syrups, usersChoice); ok {
			orderItem.syrup = Syrup{
				name:  usersChoice,
				price: price,
			}
			return true
		}
	case "size":
		if price, ok := findPrice(CafeMenu.sizes, usersChoice); ok {
			orderItem.size = Size{
				name:  usersChoice,
				price: price,
			}
			return true
		}
	case "temperature":
		if price, ok := findPrice(CafeMenu.temperatures, usersChoice); ok {
			orderItem.temperature = Temperature{
				name:  usersChoice,
				price: price,
			}
			return true
		}
	}
	return false
}