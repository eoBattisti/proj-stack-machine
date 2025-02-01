package src

import (
	"errors"
	"log"
)


type Stack struct {
  Maxsize int
  Values []float64
  length int
}

/*
Method responsible to insert a new value at the top of the stack.
It will also updates the stack's length.
*/
func (s *Stack) Push(v float64) {

  if s.length == s.Maxsize {
    log.Fatal("The stack is full")
  }

  s.Values = append(s.Values, v)
  s.length++
}

/*
Method responsible to pop off the stack's last element so it has the
LIFO behavior. It will also updates the stack's length. 
*/
func (s *Stack) Pop() (v float64, err error) {

  if s.length == 0 {
    return 0, errors.New("Cannot pop empty stack")
  }

  var lastElement float64 = s.Values[s.length - 1]
  s.Values = s.Values[:s.length - 1]
  s.length--

  return lastElement, nil
}


/*
Returns the topmost item, but the stack size does not change.
*/
func (s *Stack) Peek() (float64) {
  return s.Values[s.length - 1]
}


/*
The two topmost items exchange positions.
*/
func (s *Stack) Swap() {

  if s.length <= 1 {
    log.Fatal("Cannot swap stack values")
  }

  var last float64 = s.Peek()
  var penult float64 = s.Values[s.length - 2]

  s.Values[s.length - 1] = penult
  s.Values[s.length - 2] = last
}

/*
Returns the top value from the stack and duplicates the value to add to
the stack.
*/
func (s *Stack) Duplicate() (float64) {
  var v, err = s.Pop()

  if err != nil {
    log.Fatalf("Error while duplicating values: %s", err)
  }

  s.Push(v)
  s.Push(v)

  return v
}
