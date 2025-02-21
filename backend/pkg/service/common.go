package service

import (
	"fmt"
	"math"

	"github.com/Knetic/govaluate"
)

type commonEventService struct{}

func NewCommonEventService() commonEventService {
	return commonEventService{}
}

func (s *commonEventService) eval(expString string, count float64) (float64, error) {
	functions := map[string]govaluate.ExpressionFunction{
		"log": func(args ...interface{}) (interface{}, error) {
			if len(args) != 2 {
				return nil, fmt.Errorf("log function requires exactly two arguments")
			}

			base, ok := args[0].(float64)
			if !ok {
				return nil, fmt.Errorf("first argument of log function must be a number")
			}

			value, ok := args[1].(float64)
			if !ok {
				return nil, fmt.Errorf("second argument of log function must be a number")
			}

			if base <= 0 {
				return nil, fmt.Errorf("base of log must be greater than 0")
			}

			if value <= 0 {
				return nil, fmt.Errorf("value of log must be greater than 0")
			}

			result := math.Log(value) / math.Log(base)
			return result, nil
		},
	}

	expression, err := govaluate.NewEvaluableExpressionWithFunctions(expString, functions)
	if err != nil {
		return 0, err
	}

	result, err := expression.Evaluate(map[string]interface{}{"x": count})
	if err != nil {
		return 0, err
	}

	val, ok := result.(float64)
	if !ok {
		return 0, fmt.Errorf("can't get result of expression")
	}

	return val, nil
}

func (s *commonEventService) getMark(maxMark float64, incorrect, all int) int {
	var mark int
	if all == 0 {
		mark = 0
	} else {
		mark = int(math.Round(float64(maxMark) * (1 - 2*float64(incorrect)/float64(all))))
	}
	if mark < 0 {
		return 0
	}

	return mark
}
