package maxprodsubarray

import "testing"

func Test_maxProduct(t *testing.T) {
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
			want:6,
			args: args{
				nums: []int{2,3,-2,4},
			},
		},
		{
			name:"tc2",
			want:2,
			args: args{
				nums: []int{0,2},
			},
		},
		{
			name:"tc3",
			want:0,
			args: args{
				nums: []int{-2,0,-1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.args.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
