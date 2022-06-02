package printer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Input struct {
	PageCount int
	Count     int
}

type Expected struct {
	Error  error
	Result []string
}
type TestCase struct {
	Name string

	Input    Input
	Expected Expected
}

var testCases = []TestCase{
	{
		Name: "Negative result, first arg < 0",
		Input: Input{
			PageCount: -5,
			Count:     10,
		},
		Expected: Expected{
			Error:  InvalidVariables,
			Result: nil,
		},
	},
	{
		Name: "Negative result, second arg < 0",
		Input: Input{
			PageCount: 10,
			Count:     -5,
		},
		Expected: Expected{
			Error:  InvalidVariables,
			Result: nil,
		},
	},

	{
		Name: "Zero values",
		Input: Input{
			PageCount: 0,
			Count:     0,
		},
		Expected: Expected{
			Error:  nil,
			Result: nil,
		},
	},

	{
		Name: "1 page",
		Input: Input{
			PageCount: 1,
			Count:     0,
		},
		Expected: Expected{
			Error:  nil,
			Result: []string{"(0, 1)"},
		},
	},
	{
		Name: "2 pages",
		Input: Input{
			PageCount: 2,
			Count:     0,
		},
		Expected: Expected{
			Error:  nil,
			Result: []string{"(0, 1)", "@", "(2, 0)"},
		},
	},
	{
		Name: "3 pages",
		Input: Input{
			PageCount: 3,
			Count:     0,
		},
		Expected: Expected{
			Error:  nil,
			Result: []string{"(0, 1)", "@", "(2, 3)"},
		},
	},
	{
		Name: "11 pages",
		Input: Input{
			PageCount: 11,
			Count:     0,
		},
		Expected: Expected{
			Error: nil,
			Result: []string{
				"(0, 1)", "(10, 3)", "(8, 5)", "@", "(2, 11)", "(4, 9)", "(6, 7)",
			},
		},
	},
	{
		Name: "16 pages",
		Input: Input{
			PageCount: 16,
			Count:     0,
		},
		Expected: Expected{
			Error: nil,
			Result: []string{
				"(16, 1)", "(14, 3)", "(12, 5)", "(10, 7)", "@", "(2, 15)", "(4, 13)", "(6, 11)", "(8, 9)",
			},
		},
	},

	{
		Name: "11 pages and 1 list count",
		Input: Input{
			PageCount: 11,
			Count:     1,
		},
		Expected: Expected{
			Error: nil,
			Result: []string{
				"(4, 1)", "@", "(2, 3)", "|", "(8, 5)", "@", "(6, 7)", "|", "(0, 9)", "@", "(10, 11)",
			},
		},
	},
	{
		Name: "11 pages and 2 list count",
		Input: Input{
			PageCount: 11,
			Count:     2,
		},
		Expected: Expected{
			Error: nil,
			Result: []string{
				"(8, 1)", "(6, 3)", "@", "(2, 7)", "(4, 5)", "|", "(0, 9)", "(0, 11)", "@", "(10, 0)", "(0, 0)",
			},
		},
	},
}

func Test_Printer(t *testing.T) {
	for _, test := range testCases {
		pages, err := Printer(test.Input.PageCount, test.Input.Count)
		assert.Equal(t, test.Expected.Error, err, "test name: ", test.Name, "expected equal errors")
		assert.Equal(t, test.Expected.Result, pages, "test name: ", test.Name, "expected equal values")
	}
}
