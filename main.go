package main

import (
	"fmt"

	"github.com/fpinatares/couponEvaluator"
)

// Package Alias

func main() {
	requiredKeys := []string{"total", "amount"}
	parameters := make(map[string]interface{}, 8)
	parameters["total"] = 1000
	parameters["amount"] = 500
	fmt.Println("df")
	fdf := couponEvaluator.Evaluate(requiredKeys, parameters, "total < amount")
	fmt.Println(fdf)
}
