package slice_test

import "testing"

func TestSliceInit(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2), s2[2])
}

func TestSliceGrowing(t *testing.T) {
	var s []int
	for i := 1; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s))
	}
}

func TestSliceShareMemory(t *testing.T) {
	month := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := month[5:8]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := month[6:9]
	t.Log(summer, len(summer), cap(summer))
	summer[0] = "unknown"
	t.Log(Q2)
}
