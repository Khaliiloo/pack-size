package packsize

import (
	"fmt"
	"testing"
)

type numberOfPacksTest struct {
	order    int
	expected map[int]int
}

var numberOfPacksTests = []numberOfPacksTest{
	{1, map[int]int{250: 1}},
	{250, map[int]int{250: 1}},
	{251, map[int]int{500: 1}},
	{501, map[int]int{500: 1, 250: 1}},
	{12001, map[int]int{5000: 2, 2000: 1, 250: 1}},
	{15000, map[int]int{5000: 3}},
	{21010, map[int]int{5000: 4, 1000: 1, 250: 1}},
}

func TestNumberOfPacks(t *testing.T) {

	for _, test := range numberOfPacksTests {
		got := NumberOfPacks(test.order)
		for i, v := range test.expected {
			if got[i] != v {
				t.Errorf("got %q, expected %q", got[i], v)
			}
		}
	}

}

func TestAddPackSize(t *testing.T) {
	expected := map[int]int{5000: 2, 2000: 1, 250: 1}
	got := NumberOfPacks(12001)
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("got %q, expected %q", got[i], v)
		}
	}

	AddPackSize(10000)
	expected = map[int]int{10000: 1, 2000: 1, 250: 1}
	got = NumberOfPacks(12001)
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("got %q, expected %q", got[i], v)
		}
	}
}

func TestRemovePackSize(t *testing.T) {
	expected := map[int]int{5000: 2, 2000: 1, 250: 1}
	got := NumberOfPacks(12001)
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("got %q, expected %q", got[i], v)
		}
	}

	RemovePackSize(5000)
	expected = map[int]int{2000: 6, 250: 1}
	got = NumberOfPacks(12001)
	fmt.Println(got)
	for i, v := range expected {
		if got[i] != v {
			t.Errorf("got %q, expected %q", got[i], v)
		}
	}
}
