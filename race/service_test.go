package race

import (
	"errors"
	"testing"
)

func TestService_GenerateResults(t *testing.T) {
	type fields struct {
		repository IRepository
		classifier IClassifier
	}
	type args struct {
		inputPath  string
		outputPath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"generate result with success",
			fields{
				repository: repositoryMock{
					getRace: func(sourceName string) (Race, error) {
						return Race{}, nil
					},
					saveFinalResult: func(outputName string, finalresult FinalResult) error {
						return nil
					},
				},
				classifier: classifierMock{
					getClassification: func(Race) Classification {
						return Classification{}
					},
					getFinalResult: func(Classification) FinalResult {
						return FinalResult{}
					},
				},
			},
			args{
				inputPath:  "input.file",
				outputPath: "output.file",
			},
			false,
		},
		{
			"error on getRace",
			fields{
				repository: repositoryMock{
					getRace: func(sourceName string) (Race, error) {
						return Race{}, errors.New("failed to get data")
					},
				},
			},
			args{
				inputPath:  "input.file",
				outputPath: "output.file",
			},
			true,
		},
		{
			"error on save result",
			fields{
				repository: repositoryMock{
					getRace: func(sourceName string) (Race, error) {
						return Race{}, nil
					},
					saveFinalResult: func(outputName string, finalresult FinalResult) error {
						return errors.New("failed to save data")
					},
				},
				classifier: classifierMock{
					getClassification: func(Race) Classification {
						return Classification{}
					},
					getFinalResult: func(Classification) FinalResult {
						return FinalResult{}
					},
				},
			},
			args{
				inputPath:  "input.file",
				outputPath: "output.file",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(
				tt.fields.repository,
				tt.fields.classifier,
			)
			if err := s.GenerateResults(tt.args.inputPath, tt.args.outputPath); (err != nil) != tt.wantErr {
				t.Errorf("Service.GenerateResults() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type repositoryMock struct {
	getRace         func(sourceName string) (Race, error)
	saveFinalResult func(outputName string, finalresult FinalResult) error
}

func (r repositoryMock) GetRace(sourceName string) (Race, error) {
	return r.getRace(sourceName)
}
func (r repositoryMock) SaveFinalResult(outputName string, finalresult FinalResult) error {
	return r.saveFinalResult(outputName, finalresult)
}

type classifierMock struct {
	getClassification func(Race) Classification
	getFinalResult    func(Classification) FinalResult
}

func (c classifierMock) GetClassification(race Race) Classification {
	return c.getClassification(race)
}
func (c classifierMock) GetFinalResult(classification Classification) FinalResult {
	return c.getFinalResult(classification)
}
