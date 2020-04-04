package insertion

import (
	"reflect"
	"testing"
)

func Test_insertionSort(t *testing.T) {
	type args struct {
		list []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"TestOne", args{[]int{123, 12345, 1233232, 23, 5, 454, 5437}}, []int{5, 23, 123, 454, 5437, 12345, 1233232}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertionSort(tt.args.list); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertionSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
