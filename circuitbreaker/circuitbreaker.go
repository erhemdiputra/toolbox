package circuitbreaker

import (
	"errors"
	"sync"

	circuit "github.com/rubyist/circuitbreaker"
)

var (
	mapBreaker = make(map[string]*circuit.Breaker)
	mutex      = &sync.Mutex{}
)

// call this function in init()
func SetupMapBreaker(maps map[string]*circuit.Breaker) {
	mapBreaker = maps
}

func GetCircuitBreaker(name string) *circuit.Breaker {
	return mapBreaker[name]
}

func CircuitBreakerWrap(name string, processFunc func() error) error {
	breaker := GetCircuitBreaker(name)

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
