package src

import (
	"log"
	"strconv"
	"strings"
)

type Calculator struct {
	operands Stack
}

type Operator int

const  (
  MultOperator Operator = 42
  SumOperator Operator = 43
  SubOperator Operator = 45
  DivOperator Operator = 47
)

func (c *Calculator) sum(a float64, b float64) (v float64) {
	v = a + b
	return v
}

func (c *Calculator) sub(a float64, b float64) (v float64) {
	v = a - b
	return v
}

func (c *Calculator) mult(a float64, b float64) (v float64) {
	v = a * b
	return v
}

func (c *Calculator) div(a float64, b float64) (v float64) {
	v = a / b
	return v
}

func (c *Calculator) Parse(t string) (float64) {
  c.operands = Stack{ Maxsize: 100 }

	var size int
	for range t {
		size++
	}

	var previous_index int = 0
	for i := 0; i < size; i++ {

		if t[i] != ' ' {

      if (t[i] == byte(MultOperator) ||
          t[i] == byte(SumOperator) ||
          t[i] == byte(SubOperator) || 
          t[i] == byte(DivOperator)) {

        left, _ := c.operands.Pop()
        right, _ := c.operands.Pop()

        switch t[i] {
        case byte(MultOperator):
          value := c.mult(right, left)
          c.operands.Push(value)
        case byte(SumOperator):
          value := c.sum(right, left)
          c.operands.Push(value)
        case byte(SubOperator):
          value := c.sub(right, left)
          c.operands.Push(value)
        case byte(DivOperator):
          value := c.div(right, left)
          c.operands.Push(value)
        }

        previous_index = i + 1
        i++
      }

			continue
		}

		supposed_value := strings.TrimSpace(t[previous_index:i])
		value, err := strconv.ParseFloat(supposed_value, 64)

    if err != nil {
      log.Fatalf("Could not parse `%s`, this should be a valid number", supposed_value)
    }

		c.operands.Push(value)
		previous_index = i
	}

	return c.operands.Peek()

}
