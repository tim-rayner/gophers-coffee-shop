package main

import (
	"strings"
)

type Drink struct {
	name  string
	price float64
}

type Syrup struct {
	name  string
	price float64
}

type Size struct {
	name  string
	price float64
}

type Temperature struct {
	name  string
	price float64
}

func DrinkAvailable(drinks map[string]float64, drink string) bool {
	_, ok := findPrice(drinks, drink)
	return ok
}

func SyrupAvailable(syrups map[string]float64, syrup string) bool {
	_, ok := findPrice(syrups, syrup)
	return ok
}

func SizeAvailable(sizes map[string]float64, size string) bool {
	_, ok := findPrice(sizes, size)
	return ok
}

func TempAvailable(temps map[string]float64, temp string) bool {
	_, ok := findPrice(temps, temp)
	return ok
}

func findPrice(items map[string]float64, name string) (float64, bool) {
	for itemName, price := range items {
		if strings.EqualFold(itemName, name) {
			return price, true
		}
	}
	return 0, false
}
