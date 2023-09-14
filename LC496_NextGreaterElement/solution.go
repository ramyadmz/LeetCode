package lc496

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	nextGreaterMap := make(map[int]int, len(nums2))
	s := &Stack{items: []int{}}

	for _, num := range nums2 {
		for !s.IsEmpty() && num > s.Top() {
			nextGreaterMap[s.Pop()] = num
		}
		s.Push(num)
	}

	for i := range nums1 {
		if value, exist := nextGreaterMap[nums1[i]]; exist {
			nums1[i] = value
		} else {
			nums1[i] = -1
		}
	}

	return nums1
}

type Stack struct {
	items []int
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return -1
	}
	n := len(s.items) - 1
	item := s.items[n]
	s.items = s.items[:n]
	return item
}

func (s *Stack) Top() int {
	if s.IsEmpty() {
		return -1
	}
	n := len(s.items) - 1
	return s.items[n]
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
