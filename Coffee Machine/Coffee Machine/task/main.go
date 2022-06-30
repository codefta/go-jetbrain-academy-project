package main

import "fmt"

// Task 1
func howToCreateCopy() {
	fmt.Println("Starting to make a coffee")
	fmt.Println("Grinding coffee beans")
	fmt.Println("Boiling water")
	fmt.Println("Mixing boiled water with crushed coffee beans")
	fmt.Println("Pouring coffee into the cup")
	fmt.Println("Pouring some milk into the cup")
	fmt.Println("Coffee is ready!")
}

// Task 2
func ingredientInfo() {
	fmt.Println("Write how many cups of coffee you will need:")
	var cup int
	fmt.Scanln(&cup)
	var water = cup * 200
	var milk = cup * 50
	var beans = cup * 15
	fmt.Printf(`For 25 cups of coffee you will need:
%d ml of water
%d ml of milk
%d g of coffee beans`, water, milk, beans)
}

// Task 3
func calculateStockOfCoffee(water, milk, coffeeBean, cupOfCoffee int) string {
	var waterNeed = 200
	var milkNeed = 50
	var beansNeed = 15
	var totalWater = water / waterNeed
	var totalMilk = milk / milkNeed
	var totalBeans = coffeeBean / beansNeed
	var totalOfCoffee = (totalWater + totalMilk + totalBeans) / 3
	totalOfCoffee = findMinValue(totalWater, totalMilk, totalBeans)

	if totalOfCoffee > cupOfCoffee {
		return fmt.Sprintf("Yes, I can make that amount of coffee (and even %d more than that)", totalOfCoffee-cupOfCoffee)
	} else if totalOfCoffee == cupOfCoffee {
		return "Yes, I can make that amount of coffee"
	} else {
		return fmt.Sprintf("No, I can make only %d cups of coffee", totalOfCoffee)
	}
}

func findMinValue(value ...int) int {
	result := value[0]
	for _, val := range value {
		if val < result {
			result = val
		}
	}

	return result
}

func coffeeMachineSimulation() {
	var water int
	var milk int
	var coffeeBean int
	var cupOfCoffee int
	fmt.Println("Write how many ml of water the coffee machine has:")
	fmt.Scanln(&water)
	fmt.Println("Write how many ml of milk the coffee machine has:")
	fmt.Scanln(&milk)
	fmt.Println("Write how many grams of coffee beans the coffee machine has:")
	fmt.Scanln(&coffeeBean)
	fmt.Println("Write how many cups of coffee you will need:")
	fmt.Scanln(&cupOfCoffee)
	fmt.Println(calculateStockOfCoffee(water, milk, coffeeBean, cupOfCoffee))
}

// Task 4
func coffeeInfo(water, milk, coffeeBean, availableCup, totalMoney *int) {
	fmt.Printf("The coffee machine has:\n")
	fmt.Printf("%d ml of water\n", *water)
	fmt.Printf("%d ml of milk\n", *milk)
	fmt.Printf("%d ml of coffee beans\n", *coffeeBean)
	fmt.Printf("%d disposable cups\n", *availableCup)
	fmt.Printf("$%d of money\n", *totalMoney)
}

// Task 5
func coffeeMachine() {
	var water = 400
	var milk = 540
	var coffeeBean = 120
	var availableCup = 9
	var totalMoney = 550
	chooseAction(&water, &milk, &coffeeBean, &availableCup, &totalMoney)
}

func chooseAction(water, milk, coffeeBean, availableCup, totalMoney *int) {
	var action string

	for action != "exit" {
		fmt.Println("Write action (buy, fill, take, remaining, exit):")
		fmt.Scanln(&action)
		switch {
		case action == "buy":
			var buyAction int
			fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino:")
			fmt.Scanln(&buyAction)
			switch buyAction {
			case 1:
				switch {
				case *water < 250:
					fmt.Println("Sorry, not enough water!")
				case *coffeeBean < 16:
					fmt.Println("Sorry, not enough coffee beans!")
				case *availableCup < 1:
					fmt.Println("Sorry, not enough disposable cup!")
				default:
					fmt.Println("I have enough resources, making you a coffee!")
					if *water >= 250 && *coffeeBean >= 16 && *availableCup >= 1 {
						*water -= 250
						*coffeeBean -= 16
						*totalMoney += 4
						*availableCup--
					}
				}
			case 2:
				switch {
				case *water < 350:
					fmt.Println("Sorry, not enough water!")
				case *coffeeBean < 16:
					fmt.Println("Sorry, not enough coffee beans!")
				case *milk < 75:
					fmt.Println("Sorry, not enough milk!")
				case *availableCup < 1:
					fmt.Println("Sorry, not enough disposable cup!")
				default:
					fmt.Println("I have enough resources, making you a coffee!")
					*water -= 350
					*coffeeBean -= 20
					*milk -= 75
					*totalMoney += 7
					*availableCup--
				}
			case 3:
				switch {
				case *water < 200:
					fmt.Println("Sorry, not enough water!")
				case *coffeeBean < 12:
					fmt.Println("Sorry, not enough coffee beans!")
				case *milk < 100:
					fmt.Println("Sorry, not enough milk!")
				case *availableCup < 1:
					fmt.Println("Sorry, not enough disposable cup!")
				default:
					fmt.Println("I have enough resources, making you a coffee!")
					*water -= 200
					*milk -= 100
					*coffeeBean -= 12
					*totalMoney += 6
					*availableCup--
				}
			}
		case action == "fill":
			var newWater int
			var newMilk int
			var newCoffeeBean int
			var newAvailableCup int
			fmt.Println("WWrite how many ml of water you want to add:")
			fmt.Scanln(&newWater)
			*water += newWater
			fmt.Println("Write how many ml of milk you want to add:")
			fmt.Scanln(&newMilk)
			*milk += newMilk
			fmt.Println("Write how many grams of coffee beans you want to add:")
			fmt.Scanln(&newCoffeeBean)
			*coffeeBean += newCoffeeBean
			fmt.Println("Write how many disposable cups of coffee you want to add:")
			fmt.Scanln(&newAvailableCup)
			*availableCup += newAvailableCup
		case action == "take":
			fmt.Printf("I gave you %d\n", *totalMoney)
			*totalMoney = 0
		case action == "remaining":
			coffeeInfo(water, milk, coffeeBean, availableCup, totalMoney)
		}
	}
}

func main() {
	// write your code here
	coffeeMachine()
}
