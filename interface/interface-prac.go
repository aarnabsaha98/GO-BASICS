package main

import (
	"fmt"
)

type permanent struct {
	empID    int
	basicPay int
	bonus    int
}

type contract struct {
	empID    int
	basicPay int
}

// design a interface

type calculatepayment interface {
	calculate() int
}

func (c contract) calculate() int {
	return c.basicPay
}

func (p permanent) calculate() int {
	return p.basicPay + p.bonus
}

func calculateExpense(emp []calculatepayment) int {
	expense := 0

	for _, v := range emp {
		expense += v.calculate()
	}

	return expense
}

func main() {

	pemp1 := permanent{
		empID:    1,
		basicPay: 5000,
		bonus:    20,
	}
	pemp2 := permanent{
		empID:    2,
		basicPay: 6000,
		bonus:    30,
	}
	// cemp1 := contract{
	// 	empID:    3,
	// 	basicPay: 3000,
	// }

	employees := []calculatepayment{pemp1, pemp2}
	res := calculateExpense(employees)
	fmt.Println(res)

}
