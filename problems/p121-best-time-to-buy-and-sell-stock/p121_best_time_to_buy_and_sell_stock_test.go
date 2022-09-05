package best_time_to_buy_and_sell_stock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	type Input struct {
		Para1 []int
	}

	type Output struct {
		Para1 int
	}

	input := []Input{
		{
			Para1: []int{7, 1, 5, 3, 6, 4},
		},
		{
			Para1: []int{7, 6, 4, 3, 1},
		},
	}

	output := []Output{
		{Para1: 5},
		{Para1: 0},
	}

	for i := range input {
		assert.Equal(t, output[i].Para1, maxProfit(input[i].Para1))
	}
}
