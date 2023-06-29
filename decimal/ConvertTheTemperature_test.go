package decimal

import (
	"reflect"
	"testing"
)

func Test_convertTemperature(t *testing.T) {
	type args struct {
		celsius float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{
			name: "test1",
			args: args{
				celsius: 36.50,
			},
			want: []float64{309.65000, 97.70000},
		},
		{
			name: "test2",
			args: args{
				celsius: 122.11,
			},
			want: []float64{395.26000, 251.79800},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertTemperature(tt.args.celsius); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertTemperature() = %v, want %v", got, tt.want)
			}
		})
	}
}
