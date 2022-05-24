package race

import (
	"reflect"
	"testing"
)

func TestClassifier_GetClassification(t *testing.T) {
	type args struct {
		race Race
	}
	tests := []struct {
		name string
		args args
		want Classification
	}{
		{

			"GetClassification with success",
			args{
				Race{
					Lap{
						Time:            "23:49:08.277",
						PilotCode:       "002",
						Pilot:           "K.RAIKKONEN",
						LapNumber:       "1",
						LapTime:         "1:04.108",
						LapSpeedAverage: "43,408",
					},
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
						LapTime:         "1:03.170",
						LapSpeedAverage: "44,053",
					},
					Lap{
						Time:            "23:49:08.277",
						PilotCode:       "002",
						Pilot:           "K.RAIKKONEN",
						LapNumber:       "2",
						LapTime:         "1:03.982",
						LapSpeedAverage: "43,493",
					},
					Lap{
						Time:            "23:49:08.277",
						PilotCode:       "038",
						Pilot:           "F.MASSA",
						LapNumber:       "3",
						LapTime:         "1:02.769",
						LapSpeedAverage: "44,334",
					},
					Lap{
						Time:            "23:49:08.277",
						PilotCode:       "002",
						Pilot:           "K.RAIKKONEN",
						LapNumber:       "3",
						LapTime:         "1:03.987",
						LapSpeedAverage: "43,49",
					},
					Lap{
						Time:            "23:49:08.277",
						PilotCode:       "011",
						Pilot:           "S.VETTEL",
						LapNumber:       "1",
						LapTime:         "3:31.315",
						LapSpeedAverage: "13,169",
					},
					Lap{
						Time:            "23:49:08.277",
						PilotCode:       "038",
						Pilot:           "F.MASSA",
						LapNumber:       "4",
						LapTime:         "1:02.787",
						LapSpeedAverage: "44,321",
					},
					Lap{
						Time:            "23:49:08.277",
						PilotCode:       "002",
						Pilot:           "K.RAIKKONEN",
						LapNumber:       "4",
						LapTime:         "1:03.076",
						LapSpeedAverage: "44,118",
					},
					Lap{
						Time:            "23:49:08.277",
						PilotCode:       "011",
						Pilot:           "S.VETTEL",
						LapNumber:       "2",
						LapTime:         "1:37.864",
						LapSpeedAverage: "28,435",
					},
					Lap{
						Time:            "23:49:08.277",
						PilotCode:       "011",
						Pilot:           "S.VETTEL",
						LapNumber:       "3",
						LapTime:         "1:18.097",
						LapSpeedAverage: "35,633",
					},
				},
			},
			Classification{
				0: PilotResult{
					PilotCode:  "038",
					Pilot:      "F.MASSA",
					Laps:       4,
					Time:       251.578,
					BestLap:    3,
					BestTime:   62.769,
					TotalSpeed: 176.983,
				},
				1: PilotResult{
					PilotCode:  "002",
					Pilot:      "K.RAIKKONEN",
					Laps:       4,
					Time:       255.153,
					BestLap:    4,
					BestTime:   63.076,
					TotalSpeed: 174.50900000000001,
				},
				2: PilotResult{
					PilotCode:  "011",
					Pilot:      "S.VETTEL",
					Laps:       3,
					Time:       387.27599999999995,
					BestLap:    3,
					BestTime:   78.097,
					TotalSpeed: 77.237,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewClassifier()
			if got := c.GetClassification(tt.args.race); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Classifier.GetClassification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClassifier_GetFinalResult(t *testing.T) {
	type args struct {
		classification Classification
	}
	tests := []struct {
		name string
		args args
		want FinalResult
	}{
		{
			"get final result with success",
			args{
				Classification{
					0: PilotResult{
						PilotCode:  "038",
						Pilot:      "F.MASSA",
						Laps:       4,
						Time:       251.578,
						BestLap:    3,
						BestTime:   62.769,
						TotalSpeed: 176.983,
					},
					1: PilotResult{
						PilotCode:  "002",
						Pilot:      "K.RAIKKONEN",
						Laps:       4,
						Time:       255.153,
						BestLap:    4,
						BestTime:   61.076,
						TotalSpeed: 174.50900000000001,
					},
					2: PilotResult{
						PilotCode:  "011",
						Pilot:      "S.VETTEL",
						Laps:       3,
						Time:       387.27599999999995,
						BestLap:    3,
						BestTime:   78.097,
						TotalSpeed: 77.237,
					},
				},
			},
			FinalResult{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Classifier{}
			if got := c.GetFinalResult(tt.args.classification); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Classifier.GetFinalResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
