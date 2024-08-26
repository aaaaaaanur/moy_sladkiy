package task_07

type Stack struct {
	Ages []int
}

func (s *Stack) Push(v int) {
	s.Ages = append(s.Ages, v)
}

func (s *Stack) Pop() int {
	a := s.Ages[len(s.Ages)-1]

	s.Ages = s.Ages[:len(s.Ages)-1]

	return a
}
