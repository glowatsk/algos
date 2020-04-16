package substring

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test One", args{"aabcdeff"}, 6},
		{"Test Two", args{"lkasjdlkj"}, 6},
		{"Test Three", args{"uuuuuuuuuuuuuuuuuuuuuuwueuuuuuuuuuuuu"}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_lengthOfLongestSubstringRedux(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test One", args{"aabcdeff"}, 6},
		{"Test Two", args{"lkasjdlkj"}, 6},
		{"Test Three", args{"uuuuuuuuuuuuuuuuuuuuuuwueuuuuuuuuuuuu"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringRedux(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringRedux() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxInt(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxInt(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("maxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
