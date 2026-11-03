package errorhandling

import (
	"errors"
	"fmt"
)

var ErrOutOfStock = errors.New("out of stock") // exported, compared by identity

func buy(stock int) error {
	if stock == 0 {
		return fmt.Errorf("buy: %w", ErrOutOfStock)
	}
	return nil
}

func Example_sentinel() {
	if err := buy(0); errors.Is(err, ErrOutOfStock) {
		fmt.Println("show sold-out badge")
	}
	// Output: show sold-out badge
}
