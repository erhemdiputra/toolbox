package circuitbreaker

import (
	"errors"

	circuit "github.com/rubyist/circuitbreaker"
)

func CircuitBreakerWrap(breaker *circuit.Breaker, processFunc func() error) error {
	if breaker.Ready() {
		err := processFunc()
		if err != nil {
			breaker.Fail()
			return err
		}

		breaker.Success()
		return nil
	}

	err := errors.New("circuit breaker tripped")
	return err
}
