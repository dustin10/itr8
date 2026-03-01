package itr8_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dustin10/itr8"
)

func Test_All2(t *testing.T) {
	tests := map[string]struct {
		values map[string]string
	}{
		"nil":   {values: nil},
		"empty": {values: map[string]string{}},
		"non-empty": {values: map[string]string{
			"foo": "bar",
			"biz": "baz",
		}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			s := itr8.All2(test.values)

			for k, v := range s {
				e, exists := test.values[k]
				assert.True(t, exists)
				assert.Equal(t, e, v)
			}
		})
	}
}

func Test_FlatMap2(t *testing.T) {
	f := func(a int, b string) itr8.Seq[string] {
		return itr8.Of(fmt.Sprintf("%d: %s", a, b))
	}

	tests := map[string]struct {
		as []int
		bs []string
		fn func(int, string) itr8.Seq[string]
	}{
		"as empty": {as: []int{}, bs: []string{"1", "2", "3"}, fn: f},
		"as nil":   {as: nil, bs: []string{"1", "2", "3"}, fn: f},
		"bs empty": {as: []int{1, 2, 3}, bs: []string{}, fn: f},
		"bs nil":   {as: []int{1, 2, 3}, bs: nil, fn: f},
		"same len": {as: []int{1, 2, 3}, bs: []string{"1", "2", "3"}, fn: f},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			as := itr8.All(test.as)
			bs := itr8.All(test.bs)

			pairs := itr8.ZipToShortest(as, bs)

			s := itr8.FlatMap2(pairs, f)

			i := 0
			for c := range s {
				assert.Equal(t, fmt.Sprintf("%d: %s", test.as[i], test.bs[i]), c)
				i = i + 1
			}
		})
	}
}

func Test_Pull2(t *testing.T) {
	as := itr8.All([]int{1, 2, 3})
	bs := itr8.All([]string{"a", "b", "c"})

	seq := itr8.ZipToShortest(as, bs)

	next, stop := itr8.Pull2(seq)
	defer stop()

	a, b, ok := next()
	assert.True(t, ok)
	assert.Equal(t, 1, a)
	assert.Equal(t, "a", b)

	a, b, ok = next()
	assert.True(t, ok)
	assert.Equal(t, 2, a)
	assert.Equal(t, "b", b)

	stop()

	_, _, ok = next()
	assert.False(t, ok)
}

func Test_Map2(t *testing.T) {
	f := func(a int, b string) string {
		return fmt.Sprintf("%d: %s", a, b)
	}

	tests := map[string]struct {
		as []int
		bs []string
		fn func(int, string) string
	}{
		"as empty": {as: []int{}, bs: []string{"1", "2", "3"}, fn: f},
		"as nil":   {as: nil, bs: []string{"1", "2", "3"}, fn: f},
		"bs empty": {as: []int{1, 2, 3}, bs: []string{}, fn: f},
		"bs nil":   {as: []int{1, 2, 3}, bs: nil, fn: f},
		"same len": {as: []int{1, 2, 3}, bs: []string{"1", "2", "3"}, fn: f},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			as := itr8.All(test.as)
			bs := itr8.All(test.bs)

			pairs := itr8.ZipToShortest(as, bs)

			s := itr8.Map2(pairs, f)

			i := 0
			for c := range s {
				assert.Equal(t, f(test.as[i], test.bs[i]), c)
				i = i + 1
			}
		})
	}

}
