package race

// IRepository repository interface
type IRepository interface {
	GetRace(sourceName string) (Race, error)
	SaveFinalResult(outputName string, finalresult FinalResult) error
}

// Repository representation
type Repository struct {
	source ISource
}

// NewRepository constructor
func NewRepository(source ISource) *Repository {
	return &Repository{
		source: source,
	}
}

// GetRace get race data
func (r *Repository) GetRace(sourceName string) (Race, error) {
	return r.source.GetDataFromSource(sourceName)
}

// SaveFinalResult saves the final result from a race
func (r *Repository) SaveFinalResult(outputName string, finalResult FinalResult) error {
	return r.source.SaveDataInSource(outputName, finalResult)
}
