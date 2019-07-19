package validParentheses

import "testing"

func Test_isValid(t *testing.T) {
	tests := []struct {
		name string
		tc   string
		want bool
	}{
		{name: "Simple ()", tc: "()", want: true},
		{name: "Simple {}", tc: "{}", want: true},
		{name: "Simple []", tc: "[]", want: true},
		{name: "Fails", tc: "(()", want: false},
		{name: "Combined", tc: "(){}[]", want: true},
		{name: "Mixed", tc: "({[]})", want: true},
		{name: "Single closer", tc: ")", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.tc); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
