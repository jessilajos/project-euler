package projecteuler

import (
	"fmt"
	"testing"
	"time"
)

type testrunner func(interface{}) (string, interface{})

type testcase struct {
	name      string
	input     interface{}
	want      interface{}
	benchmark interface{}
	runners   []testrunner
}

func TestProblems(t *testing.T) {
	testcases := []testcase{
		{"Problem 1", 1000, 233168, int(1e9), []testrunner{jessisolutionP1, prestonsolutionP1, jointsolutionP1}},
		{"Problem 2", 4e6, 4613732, nil, []testrunner{}},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			for _, runner := range tc.runners {
				name, got := runner(tc.input)
				if got != tc.want {
					t.Fatalf("%s failed: got %v, wanted %v", name, got, tc.want)
				}
				t.Logf("completed: %s", name)
			}
		})
		t.Run(fmt.Sprintf("%s BENCHMARK", tc.name), func(t *testing.T) {
			for _, runner := range tc.runners {
				if tc.benchmark == nil {
					return
				}
				start := time.Now()
				name, _ := runner(tc.benchmark)
				elapsed := time.Since(start)
				t.Logf("completed: %s (%v)", name, elapsed)
			}
		})
	}
}
