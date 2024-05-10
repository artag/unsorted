package classes

import "math"

type Dealer struct {
	FeePercent float64
}

func NewDealer(feePercent int) *Dealer {
	return &Dealer{FeePercent: (float64(feePercent) / 100)}
}

func (d *Dealer) RollTheDice(dices *Dices) *Dices {
	dices.RollTheDice()
	return dices
}

func (d *Dealer) GetFee(bet int) int {
	dealerFee := float64(bet) * d.FeePercent
	fee := int(math.Round(dealerFee))
	return fee
}
