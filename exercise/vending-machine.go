package main

import (
	"fmt"

	"github.com/ajjammy/vendingMachine"
)

var coins = map[string]int{"T": 10, "F": 5, "TW": 2, "O": 1}
var items = map[string]int{"SD": 18, "CC": 12, "DW": 7}

func main() {
	vm := vendingMachine.NewVendingMachine(coins, items)
	fmt.Println("Inserted Money:", vm.InsertedMoney())

	vm.InsertCoin("T")
	vm.InsertCoin("F")
	vm.InsertCoin("TW")
	vm.InsertCoin("O")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	can := vm.SelectSD() // SD
	fmt.Println(can)

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	can = vm.SelectCC()
	fmt.Println(can) // CC, F, TW, O

	vm.InsertCoin("T")
	vm.InsertCoin("T")
	vm.InsertCoin("F")
	fmt.Println("Inserted Money:", vm.InsertedMoney())
	can = vm.CoinReturn()
	fmt.Println(can) // T, T, F
}
