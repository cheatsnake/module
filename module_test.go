package module

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive", a: 1, b: 2, want: 3},
		{name: "negative", a: -1, b: -2, want: -3},
		{name: "zero", a: 0, b: 0, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.a, tt.b); got != tt.want {
				t.Errorf("Add(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positive", a: 3, b: 4, want: 12},
		{name: "negative", a: -3, b: 4, want: -12},
		{name: "zero", a: 0, b: 5, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.a, tt.b); got != tt.want {
				t.Errorf("Multiply(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
