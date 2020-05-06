package _0983

import "testing"

func Test_mincostTickets(t *testing.T) {
	type args struct {
		days  []int
		costs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test1", args: args{
			days:  []int{1, 4, 6, 7, 8, 20},
			costs: []int{2, 7, 15},
		}, want: 11},
		{name: "test2", args: args{
			days:  []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 30, 31},
			costs: []int{2, 7, 15},
		}, want: 17},
		{name: "test3", args: args{
			days:  []int{1, 4, 6, 7, 8, 364},
			costs: []int{2, 7, 15},
		}, want: 11},
		{name: "test_1", args: args{
			days:  []int{1, 4, 6, 7, 8, 365},
			costs: []int{2, 7, 15},
		}, want: 11},
		{name: "test_2", args: args{
			days:  []int{6, 8, 9, 18, 20, 21, 23, 25},
			costs: []int{2, 10, 41},
		}, want: 16},
		{name: "test_3", args: args{
			days:  []int{6, 9, 10, 14, 15, 16, 17, 18, 20, 22, 23, 24, 29, 30, 31, 33, 35, 37, 38, 40, 41, 46, 47, 51, 54, 57, 59, 65, 70, 76, 77, 81, 85, 87, 90, 91, 93, 94, 95, 97, 98, 100, 103, 104, 105, 106, 107, 111, 112, 113, 114, 116, 117, 118, 120, 124, 128, 129, 135, 137, 139, 145, 146, 151, 152, 153, 157, 165, 166, 173, 174, 179, 181, 182, 185, 187, 188, 190, 191, 192, 195, 196, 204, 205, 206, 208, 210, 214, 218, 219, 221, 225, 229, 231, 233, 235, 239, 240, 245, 247, 249, 251, 252, 258, 261, 263, 268, 270, 273, 274, 275, 276, 280, 283, 285, 286, 288, 289, 290, 291, 292, 293, 296, 298, 299, 301, 303, 307, 313, 314, 319, 323, 325, 327, 329, 334, 339, 340, 341, 342, 344, 346, 349, 352, 354, 355, 356, 357, 358, 359, 363, 364},
			costs: []int{21, 115, 345},
		}, want: 3040},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mincostTickets(tt.args.days, tt.args.costs); got != tt.want {
				t.Errorf("mincostTickets() = %v, want %v", got, tt.want)
			}
		})
	}
}
