package race

import (
	"errors"
	"reflect"
	"testing"
)

func TestRepository_GetRace(t *testing.T) {
	type fields struct {
		source ISource
	}
	tests := []struct {
		name       string
		fields     fields
		sourceName string
		want       Race
		wantErr    bool
	}{
		{
			"return race with success",
			fields{
				source: sourceMock{
					getDataFromSource: func(sourceName string) (Race, error) {
						return Race{}, nil
					},
				},
			},
			"test.csv",
			Race{},
			false,
		},
		{
			"error on get race data",
			fields{
				source: sourceMock{
					getDataFromSource: func(sourceName string) (Race, error) {
						return Race{}, errors.New("error to open file")
					},
				},
			},
			"invalid.file",
			Race{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(tt.fields.source)

			got, err := r.GetRace(tt.sourceName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetRace() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetRace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_SaveFinalResult(t *testing.T) {
	type fields struct {
		source ISource
	}
	type args struct {
		outputName  string
		finalresult FinalResult
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"save data with success",
			fields{
				sourceMock{
					saveDataInSource: func(outputName string, finalresult FinalResult) error {
						return nil
					},
				},
			},
			args{
				outputName:  "result.csv",
				finalresult: FinalResult{},
			},
			false,
		},
		{
			"fail to save data",
			fields{
				sourceMock{
					saveDataInSource: func(outputName string, finalresult FinalResult) error {
						return errors.New("failed to save data")
					},
				},
			},
			args{
				outputName:  "invalid.file",
				finalresult: FinalResult{},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRepository(tt.fields.source)
			if err := r.SaveFinalResult(tt.args.outputName, tt.args.finalresult); (err != nil) != tt.wantErr {
				t.Errorf("Repository.SaveFinalResult() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type sourceMock struct {
	getDataFromSource func(sourceName string) (Race, error)
	saveDataInSource  func(outputName string, finalresult FinalResult) error
}

func (s sourceMock) GetDataFromSource(sourceName string) (Race, error) {
	return s.getDataFromSource(sourceName)
}

func (s sourceMock) SaveDataInSource(outputName string, finalresult FinalResult) error {
	return s.saveDataInSource(outputName, finalresult)
}
