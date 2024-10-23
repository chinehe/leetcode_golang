package problemset

import (
	"math/rand"
	"testing"
	"time"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				nums: func(num int) []int {
					nums := make([]int, num)
					rand.Seed(time.Now().Unix())
					for i := 0; i < num; i++ {
						nums[i] = rand.Intn(2)
					}
					return nums
				}(100000),
				val: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			if got := removeElement2(tt.args.nums, tt.args.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, time.Now().Sub(start).Milliseconds())
			}
		})
	}
}
