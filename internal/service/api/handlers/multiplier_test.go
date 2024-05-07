package handlers

import (
	"github.com/magiconair/properties/assert"
	"math/big"
	"testing"
)

var (
	gasMultiplier    = 1.3
	gasPriceZero     = big.NewInt(0)
	gasPriceAverage  = big.NewInt(7709)
	gasPriceLow      = big.NewInt(150)
	gasPriceNegative = big.NewInt(-5543)
)

func TestZeroGasMultiplier(t *testing.T) {
	multiplied := multiplyGasPrice(gasPriceZero, gasMultiplier)
	assert.Equal(t, multiplied, big.NewInt(0))
}

func TestAverageGasMultiplier(t *testing.T) {
	multiplied := multiplyGasPrice(gasPriceAverage, gasMultiplier)
	assert.Equal(t, multiplied, big.NewInt(10021))
}

func TestLowGasMultiplier(t *testing.T) {
	multiplied := multiplyGasPrice(gasPriceLow, gasMultiplier)
	assert.Equal(t, multiplied, big.NewInt(195))
}

func TestNegativeGasMultiplier(t *testing.T) {
	multiplied := multiplyGasPrice(gasPriceNegative, gasMultiplier)
	assert.Equal(t, multiplied, big.NewInt(-7206))
}
