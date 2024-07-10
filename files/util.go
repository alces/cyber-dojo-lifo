package lifo

func (s *Stack) index(element string) int {
    for i := 0; i < s.Size(); i++ {
        if s.items[i] == element {
            return i
        }
    }
    
    return -1
}
