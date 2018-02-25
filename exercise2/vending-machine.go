package main

import (
	"fmt"
	"strings"
)

type VendingMachine struct {
	coins []string
	items []string
}

func main() {
	coins := []string{}
	items := []string{}
	vm := NewVendingMachine(coins, items)

	//Buy SD(soft drink) with exact change
	fmt.Println()
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 18
	fmt.Println("Choose: Select SD")
	can := vm.SelectSD()
	fmt.Println("Return:", can) // SD

	vm = NewVendingMachine(coins, items)

	//Buy CC(canned coffee) without exact change
	fmt.Println()
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 20
	fmt.Println("Choose: Select CC")
	can = vm.SelectCC()
	fmt.Println("Return:", can) // CC, F, TW, O

	vm = NewVendingMachine(coins, items)

	// Start adding change but hit coin return
	fmt.Println()
	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.GetInsertedMoney()) // 25
	fmt.Println("Choose: Coin Return")
	change := vm.CoinReturn()
	fmt.Println("Return:", change) // T, T, F
}

func (vm *VendingMachine) CoinReturn() string {
	vm.items = []string{}
	return strings.Join(vm.coins, ", ")
}

func (vm *VendingMachine) SelectCC() string {
	vm.items = append(vm.items, "CC")
	insertedMoney := vm.GetInsertedMoney()
	priceCC := 12

	balance := insertedMoney - priceCC
	if balance > 0 {
		return vm.GetSelectedItems() + ", " + getDisplayMoney(balance)
	} else {
		return vm.GetSelectedItems()
	}
}

func (vm *VendingMachine) SelectSD() string {
	vm.items = append(vm.items, "SD")
	return vm.GetSelectedItems()
}

func (vm VendingMachine) GetSelectedItems() string {
	return strings.Join(vm.items, ", ")
}

func (vm VendingMachine) GetInsertedMoney() int {
	coin := map[string]int{
		"T":  10,
		"F":  5,
		"TW": 2,
		"O":  1,
	}

	var sumCoinValue int
	for _, value := range vm.coins {
		sumCoinValue += coin[value]
	}

	return sumCoinValue
}

func getDisplayMoney(money int) string {
	displayMoney := []string{}
	coinsValue := []int{10, 5, 2, 1}
	coinsDisplay := []string{"T", "F", "TW", "O"}

	balance := money

	for index := 0; index < len(coinsValue); index++ {
		if balance/coinsValue[index] > 0 {
			for i := 0; i < (balance / coinsValue[index]); i++ {
				displayMoney = append(displayMoney, coinsDisplay[index])
			}
			balance = balance % coinsValue[index]
		}
	}

	return strings.Join(displayMoney, ", ")
}

func (vm *VendingMachine) InsertCoin(coin string) {
	vm.coins = append(vm.coins, coin)
}

func NewVendingMachine(coins, items []string) *VendingMachine {
	return &VendingMachine{coins, items}
}
