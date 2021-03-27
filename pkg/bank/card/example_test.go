package card

import (
	"bank/pkg/bank/types"	
	"fmt"
)


func ExampleTotal() {
	card := []types.Card{
		{
			Balance: 1_000_00,
			Active: true,
		},
	}
	fmt.Println(Total(card))
	// Output: 100000
}