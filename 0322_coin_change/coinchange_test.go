package coinchange

import "testing"

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "125,11",
			args: args{
				coins:  []int{1,2,5},
				amount: 11,
			},
			want:3,
		},
		{
			name: "2,3",
			args: args{
				coins:  []int{2},
				amount: 3,
			},
			want:-1,
		},
		{
			name:"1,0",
			args: args{
				coins:  []int{1},
				amount: 0,
			},
			want: 0,
		},
		{
			name: "1,1",
			args: args{
				coins:  []int{1},
				amount: 1,
			},
			want: 1,
		},
		{
			name: "1,2",
			args: args{
				coins:  []int{1},
				amount: 2,
			},
			want: 2,
		},
		{
			name: "big numbers",
			args: args{
				coins:  []int{186,419,83,408},
				amount: 6249,
			},
			want: 20,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := coinChange(tt.args.coins, tt.args.amount); got != tt.want {
				t.Errorf("coinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
