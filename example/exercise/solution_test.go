package exercise

import "testing"

func TestSolution(t *testing.T) {
	var testcase = []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "Return the string from input",
			given: "Hello world",
			want:  "Hello world",
		},
	}
	for _, tc := range testcase {
		t.Run(tc.name, func(t *testing.T) {
			got := Solution(tc.given)
			if got != tc.want {
				t.Errorf("Solution(%s) = %s, want: %s", tc.given, got, tc.want)
			}
		})
	}
}
