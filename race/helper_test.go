package race

import (
	"testing"
)

func Test_convertLapTimeToFloat(t *testing.T) {
	type args struct {
		lap Lap
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"convert with success",
			args{
				lap: Lap{
					LapTime: "1:02.852",
				},
			},
			62.852,
		},
		{
			"fail to convert",
			args{
				lap: Lap{
					LapTime: "xpto",
				},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertLapTimeToFloat(tt.args.lap); got != tt.want {
				t.Errorf("convertLapTimeToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertSpeedToFloat(t *testing.T) {
	type args struct {
		lap Lap
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"convert with success",
			args{
				lap: Lap{
					LapSpeedAverage: "43,500",
				},
			},
			43.500,
		},
		{
			"fail to convert",
			args{
				lap: Lap{
					LapSpeedAverage: "xpto",
				},
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertSpeedToFloat(tt.args.lap); got != tt.want {
				t.Errorf("convertSpeedToFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertLapTimeToString(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"convert with success",
			args{
				value: 62.852,
			},
			"1m2.852s",
		},
		{
			"convert with success 2",
			args{
				value: 62.0,
			},
			"1m2s",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertLapTimeToString(tt.args.value); got != tt.want {
				t.Errorf("convertLapTimeToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertSpeedToString(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"convert with success",
			args{
				value: 62.0,
			},
			"62.000",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertSpeedToString(tt.args.value); got != tt.want {
				t.Errorf("convertSpeedToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
