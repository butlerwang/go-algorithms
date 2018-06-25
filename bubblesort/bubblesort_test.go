package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || value[4] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 12345")
	}
}

func TestBubbleSort2(t *testing.T) {
	values = []int{5, 5, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || value[4] != 5 {
		t.Error("BubbleSort() failed. Got", values, "Expected 1 2 3 5 5")
	}
}

func TestBubbleSort3(t *testing.T) {
	values = []int{5}
	BubbleSort(values)
	if values[0] != 1 {
		t.Error("BubbleSort() failed. Got", values, "Expected  5")
	}
}
