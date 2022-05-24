package race

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

// ISource interface
type ISource interface {
	GetDataFromSource(sourceName string) (Race, error)
	SaveDataInSource(outputName string, finalresult FinalResult) error
}

// FileSource is file source representation
type FileSource struct{}

// NewFileSource constructor
func NewFileSource() *FileSource {
	return &FileSource{}
}

// GetDataFromSource get data from file
func (f *FileSource) GetDataFromSource(sourceName string) (Race, error) {
	race := Race{}

	file, err := os.Open(sourceName)
	if err != nil {
		return race, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		regex := regexp.MustCompile("[\\w.:,]+")
		words := regex.FindAllString(scanner.Text(), -1)
		lap := Lap{
			Time:            words[0],
			PilotCode:       words[1],
			Pilot:           words[2],
			LapNumber:       words[3],
			LapTime:         words[4],
			LapSpeedAverage: words[5],
		}
		race = append(race, lap)
	}

	return race, nil
}

// SaveDataInSource save data in a file
func (f *FileSource) SaveDataInSource(outputName string, finalResult FinalResult) error {
	file, err := os.Create(outputName)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString("Posição Chegada | Código Piloto | Nome Piloto       | Voltas Completadas | Tempo Total de Prova | Melhor volta    | Velocidade Média | Diferença  |\n")
	for position := 0; position < len(finalResult.Classification); position++ {
		file.WriteString(
			fmt.Sprintf("%-15d | %-13s | %-17s | %-18d | %-21s| %d - %-11s | %-16s | %-10s |\n",
				position+1,
				finalResult.Classification[position].PilotCode,
				finalResult.Classification[position].Pilot,
				finalResult.Classification[position].Laps,
				finalResult.Classification[position].Time,
				finalResult.Classification[position].BestLap,
				finalResult.Classification[position].BestTime,
				finalResult.Classification[position].AverageSpeed,
				finalResult.Classification[position].Difference,
			),
		)
	}
	file.WriteString(fmt.Sprintf("\nMelhor volta da corrida : %s\n", finalResult.BestRaceTime))
	return nil
}
