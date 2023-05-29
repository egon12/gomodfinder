package gomodfinder

import (
	"fmt"
	"testing"
)

func TestFind(t *testing.T) {
	_, err := Find()
	if err != nil {
		t.Error("It should not be error, because we have go.mod in parent ancestor")
	}
}

func TestFindRecursive(t *testing.T) {
	tests := []struct {
		inputDir       string
		inputRecursive int
		outputError    string
	}{

		{"fake_folder/../", 0, "can't find go.mod in parent ancestor: cannot read in 'fake_folder/../'"},
		{"/", 0, "can't find go.mod in parent ancestor"},
		{".", 101, "can't find go.mod in parent ancestor: '.' nested more than 100 level"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("input %s, %d", tt.inputDir, tt.inputRecursive), func(t *testing.T) {
			_, err := findRecursiveGoMod(tt.inputDir, tt.inputRecursive)
			if err == nil {
				t.Errorf("It should be error with error message: %v", tt.outputError)
				return
			}

			if err.Error() != tt.outputError {
				t.Errorf("It should be error with error message: %v\ngot: %v", tt.outputError, err.Error())
			}
		})
	}
}

func BenchmarkFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = Find()
	}
}
