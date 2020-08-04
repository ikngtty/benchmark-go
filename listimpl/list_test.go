package listimpl

import (
	"fmt"
	"reflect"
	"testing"
)

type intList interface {
	Add(int)
	ToA() []int
}

var newFromArrayFuncs = []struct {
	name string
	body func([]int) intList
}{
	{"IntList1", func(a []int) intList { return NewIntList1FromArray(a) }},
	{"IntList2", func(a []int) intList { return NewIntList2FromArray(a) }},
}

func TestNewIntListFromArray(t *testing.T) {
	cases := []struct {
		a []int
	}{
		{[]int{}},
		{[]int{20}},
		{[]int{20, 30}},
		{[]int{20, 30, 10}},
	}
	for _, f := range newFromArrayFuncs {
		for _, c := range cases {
			testName := fmt.Sprintf("%s:from_%d_length_array", f.name, len(c.a))
			t.Run(testName, func(t *testing.T) {
				list := f.body(c.a)

				got := list.ToA()
				want := c.a
				if !reflect.DeepEqual(got, want) {
					t.Errorf("want: %#v, got: %#v", want, got)
				}
			})
		}
	}
}

func TestIntListConcat(t *testing.T) {
	t.Run("the_list_argument_is_unchanged", func(t *testing.T) {
		listArray := []int{10, 20, 30}
		otherArray := []int{40, 50, 60}

		t.Run("IntList1", func(t *testing.T) {
			list := NewIntList1FromArray(listArray)
			other := NewIntList1FromArray(otherArray)

			list.Concat(other)

			got := other.ToA()
			want := otherArray
			if !reflect.DeepEqual(got, want) {
				t.Errorf("want: %#v, got: %#v", want, got)
			}
		})
		t.Run("IntList2", func(t *testing.T) {
			list := NewIntList2FromArray(listArray)
			other := NewIntList2FromArray(otherArray)

			list.Concat(other)

			got := other.ToA()
			want := otherArray
			if !reflect.DeepEqual(got, want) {
				t.Errorf("want: %#v, got: %#v", want, got)
			}
		})
	})

	cases := []struct {
		list  []int
		other []int
		want  []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{}, []int{22}, []int{22}},
		{[]int{}, []int{22, 44}, []int{22, 44}},
		{[]int{}, []int{22, 44, 33}, []int{22, 44, 33}},
		{[]int{30}, []int{}, []int{30}},
		{[]int{30}, []int{22}, []int{30, 22}},
		{[]int{30}, []int{22, 44}, []int{30, 22, 44}},
		{[]int{30}, []int{22, 44, 33}, []int{30, 22, 44, 33}},
		{[]int{30, 10}, []int{}, []int{30, 10}},
		{[]int{30, 10}, []int{22}, []int{30, 10, 22}},
		{[]int{30, 10}, []int{22, 44}, []int{30, 10, 22, 44}},
		{[]int{30, 10}, []int{22, 44, 33}, []int{30, 10, 22, 44, 33}},
		{[]int{30, 10, 50}, []int{}, []int{30, 10, 50}},
		{[]int{30, 10, 50}, []int{22}, []int{30, 10, 50, 22}},
		{[]int{30, 10, 50}, []int{22, 44}, []int{30, 10, 50, 22, 44}},
		{[]int{30, 10, 50}, []int{22, 44, 33}, []int{30, 10, 50, 22, 44, 33}},
	}
	for _, c := range cases {
		testName := fmt.Sprintf("%v+%v", c.list, c.other)
		t.Run(testName, func(t *testing.T) {
			t.Run("IntList1", func(t *testing.T) {
				list := NewIntList1FromArray(c.list)
				other := NewIntList1FromArray(c.other)

				list.Concat(other)

				got := list.ToA()
				if !reflect.DeepEqual(got, c.want) {
					t.Errorf("want: %#v, got: %#v", c.want, got)
				}
			})
			t.Run("IntList2", func(t *testing.T) {
				list := NewIntList2FromArray(c.list)
				other := NewIntList2FromArray(c.other)

				list.Concat(other)

				got := list.ToA()
				if !reflect.DeepEqual(got, c.want) {
					t.Errorf("want: %#v, got: %#v", c.want, got)
				}
			})
		})
	}
}

func BenchmarkIntList(b *testing.B) {
	for _, f := range newFromArrayFuncs {
		b.Run(f.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				list := f.body([]int{})
				for j := 0; j < 100_000; j++ {
					list.Add(j)
				}
				list.ToA()
			}
		})
	}
}
