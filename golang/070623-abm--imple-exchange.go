package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Action string

const (
	LIMIT_BUY   Action = "LIMIT_BUY"
	LIMIT_SELL  Action = "LIMIT_SELL"
	MARKET_BUY  Action = "MARKET_BUY"
	MARKET_SELL Action = "MARKET_SELL"
)

type Order struct {
	Action Action
	Price  float64
	Amount float64
}

type State struct {
	Price      float64
	BuyOrders  []Order
	SellOrders []Order
}

func main() {
	// Initialize constants
	totalSteps := 1_000
	totalAgents := 50

	// Initialize world
	state := initializeState()

	// Create empty history
	priceHistory := []float64{}

	// Iterate over steps
	for step := 0; step < totalSteps; step++ {
		stepActions := generateStepActions(state, totalAgents)
		state = applyActionsEffects(stepActions, state)
		priceHistory = append(priceHistory, state.Price)
	}

	plotPriceHistory(priceHistory)
}

func initializeState() State {
	state := State{
		Price:      1.0,
		BuyOrders:  []Order{},
		SellOrders: []Order{},
	}
	return state
}

func generateStepActions(state State, totalAgents int) []Order {
	stepActions := []Order{}
	totalActions := generateRandomInt(0, totalAgents)

	for n := 0; n < totalActions; n++ {
		action := generateRandomAction(state)
		stepActions = append(stepActions, action)
	}

	return stepActions
}

func generateRandomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func generateRandomAction(state State) Order {
	price := state.Price
	if rand.Float64() < 0.5 {
		price = generateSamplePrice(state.Price)
	}

	amount := rand.Float64()

	actionType := generateRandomInt(0, 3)
	switch actionType {
	case 0:
		return Order{Action: LIMIT_BUY, Price: price, Amount: amount}
	case 1:
		return Order{Action: LIMIT_SELL, Price: price, Amount: amount}
	case 2:
		return Order{Action: MARKET_BUY, Amount: amount}
	case 3:
		return Order{Action: MARKET_SELL, Amount: amount}
	default:
		return Order{}
	}
}

func generateSamplePrice(currentPrice float64) float64 {
	// Use the `sampleuv` library to generate a sample price
	// Replace the following code with the appropriate function calls using the library
	samplePrice := currentPrice * 1.01 // Placeholder code, replace with actual implementation

	return samplePrice
}

func applyActionsEffects(actions []Order, state State) State {
	for _, action := range actions {
		switch action.Action {
		case LIMIT_BUY:
			state.BuyOrders = append(state.BuyOrders, action)
			sort.Slice(state.BuyOrders, func(i, j int) bool {
				return state.BuyOrders[i].Price > state.BuyOrders[j].Price
			})
		case LIMIT_SELL:
			state.SellOrders = append(state.SellOrders, action)
			sort.Slice(state.SellOrders, func(i, j int) bool {
				return state.SellOrders[i].Price < state.SellOrders[j].Price
			})
		case MARKET_BUY:
			n := 0
			totalAmount := 0.0
			for n < len(state.BuyOrders) && totalAmount < action.Amount {
				totalAmount += state.BuyOrders[n].Amount
				n++
			}

			if len(state.BuyOrders) == 0 {
				state.BuyOrders = []Order{}
				state.Price = 1_000_000
			} else if n > len(state.BuyOrders) {
				state.BuyOrders = []Order{}
				state.Price = 1_000_000
			} else {
				state.BuyOrders = state.BuyOrders[n-1:]
				state.Price = state.BuyOrders[0].Price
			}
		case MARKET_SELL:
			n := 0
			totalAmount := 0.0
			for n < len(state.SellOrders) && totalAmount < action.Amount {
				totalAmount += state.SellOrders[n].Amount
				n++
			}

			if len(state.SellOrders) == 0 {
				state.SellOrders = []Order{}
				state.Price = 0
			} else if n > len(state.SellOrders) {
				state.SellOrders = []Order{}
				state.Price = 0
			} else {
				state.SellOrders = state.SellOrders[n-1:]
				state.Price = state.SellOrders[0].Price
			}
		}
	}

	return state
}

func plotPriceHistory(priceHistory []float64) {
	// Implement plotting logic here
	fmt.Println(priceHistory)
}
