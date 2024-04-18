package problems

import "testing"

func TestMultiply(t *testing.T) {
	testCases := []struct {
		num1 string
		num2 string
		want string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
		{"0", "123", "0"},
		{"123", "0", "0"},
		{"999999", "999999", "999998000001"},
		{"498828660196", "840477629533", "419254329864656431168468"},
	}

	for _, tc := range testCases {
		got := multiply(tc.num1, tc.num2)
		if got != tc.want {
			t.Errorf("multiply(%q, %q) = %q; want %q", tc.num1, tc.num2, got, tc.want)
		}
	}
}
