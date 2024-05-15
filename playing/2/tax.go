package tax

import "errors"

type Repository interface {
	SaveTax(tax float64) error
}

func CalculateTax(amount float64) (float64, error) {
	if amount <= 0 {
		return 0.0, errors.New("amount must be greater than zero")
	}

	if amount >= 1000 {
		return 10.0, nil
	}

	return 5.0, nil
}

func CalculateTaxAndSave(amount float64, repository Repository) error {
	tax := CalculateTaxWithoutError(amount)
	err := repository.SaveTax(tax)
	if err != nil {
		return err
	}

	return nil
}

func CalculateTaxWithoutError(amount float64) float64 {
	if amount <= 0 {
		return 0.0
	}

	if amount >= 1000 {
		return 10.0
	}

	return 5.0
}