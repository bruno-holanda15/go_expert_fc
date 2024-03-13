package tax

import "errors"

// import "time"

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than zero")
	}

	if amount >= 1000 {
		return 10.0, nil
	}

	return 5.0, nil
}