package race

// Service representation
type Service struct {
	repository IRepository
	classifier IClassifier
}

// NewService constructor
func NewService(repository IRepository, classifier IClassifier) *Service {
	return &Service{
		repository: repository,
		classifier: classifier,
	}
}

// GenerateResults generate race classification and the best race time
func (s *Service) GenerateResults(inputPath, outputPath string) error {
	race, gerr := s.repository.GetRace(inputPath)
	if gerr != nil {
		return gerr
	}

	classification := s.classifier.GetClassification(race)
	finalResult := s.classifier.GetFinalResult(classification)

	serr := s.repository.SaveFinalResult(outputPath, finalResult)
	if serr != nil {
		return serr
	}
	return nil
}
