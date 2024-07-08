package lifo

type Stack struct {
    size int
}

func New() *Stack {
    return &Stack{}
}

func (s *Stack) Add(element string) error {
    return nil
}