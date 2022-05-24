package race

import (
	"reflect"
	"testing"
)

func TestFileSource_GetDataFromSource(t *testing.T) {
	type args struct {
		sourceName string
	}
	tests := []struct {
		name    string
		args    args
		want    Race
		wantErr bool
	}{
		{
			"fill race with success",
			args{
				sourceName: "../tests/test_source.csv",
			},
			Race{
				Lap{
					Time:            "23:49:08.277",
					PilotCode:       "038",
					Pilot:           "F.MASSA",
					LapNumber:       "1",
					LapTime:         "1:02.852",
					LapSpeedAverage: "44,275",
				},
				Lap{
					Time:            "23:49:08.277",
					PilotCode:       "038",
					Pilot:           "F.MASSA",
					LapNumber:       "2",
					LapTime:         "1:02.852",
					LapSpeedAverage: "44,275",
				},
			},
			false,
		},
		{
			"fail to open file",
			args{
				sourceName: "",
			},
			Race{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := NewFileSource()
			got, err := f.GetDataFromSource(tt.args.sourceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileSource.GetDataFromSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileSource.GetDataFromSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFileSource_SaveDataInSource(t *testing.T) {
	type args struct {
		outputName  string
		finalResult FinalResult
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			"save data with success",
			args{
				outputName: "../tests/test_result.csv",
				finalResult: FinalResult{
					Classification: map[int]PilotResultDetailed{
						0: PilotResultDetailed{
							PilotCode:    "038",
							Pilot:        "F.MASSA",
							Laps:         4,
							Time:         "4m11.578s",
							BestLap:      3,
							BestTime:     "1m2.769s",
							AverageSpeed: "44.246",
							Difference:   "0s",
						},
						1: PilotResultDetailed{
							PilotCode:    "002",
							Pilot:        "K.RAIKKONEN",
							Laps:         4,
							Time:         "4m15.153s",
							BestLap:      4,
							BestTime:     "1m1.076s",
							AverageSpeed: "43.627",
							Difference:   "3.575s",
						},
						2: PilotResultDetailed{
							PilotCode:    "011",
							Pilot:        "S.VETTEL",
							Laps:         3,
							Time:         "6m27.276s",
							BestLap:      3,
							BestTime:     "1m18.097s",
							AverageSpeed: "25.746",
							Difference:   "2m15.698s",
						},
					},
					BestRaceTime: "1m1.076s",
				},
			},
			false,
		},
		{
			"fail to open file",
			args{
				outputName:  "",
				finalResult: FinalResult{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FileSource{}
			if err := f.SaveDataInSource(tt.args.outputName, tt.args.finalResult); (err != nil) != tt.wantErr {
				t.Errorf("FileSource.SaveDataInSource() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
