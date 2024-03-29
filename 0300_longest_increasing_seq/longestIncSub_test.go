package longestincsub

import "testing"

func Test_lengthOfLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"tc1",
			want: 4,
			args: args{
				nums: []int{10,9,2,5,3,7,101,18},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLIS(tt.args.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}


func Test_max(t *testing.T) {
	v := max([]int{10,9,2,5,3,7,101,18})
	if v != 101 {
		t.Errorf("expected 101 but got %v", v)
	}
}