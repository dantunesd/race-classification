package race

// IClassifier interface
type IClassifier interface {
	GetClassification(Race) Classification
	GetFinalResult(Classification) FinalResult
}

// Classifier representation
type Classifier struct{}

// NewClassifier constructor
func NewClassifier() *Classifier {
	return &Classifier{}
}

// GetClassification classifies a race
func (c *Classifier) GetClassification(race Race) Classification {
	return generateClassification(
		aggragateResults(race),
	)
}

// GetFinalResult gets the final results
func (c *Classifier) GetFinalResult(classification Classification) FinalResult {
	return generateFinalResult(classification)
}

func aggragateResults(race Race) PilotsResults {
	results := PilotsResults{}
	for _, lap := range race {
		results = aggregate(results, lap)
	}
	return results
}

func generateClassification(results PilotsResults) Classification {
	classification := Classification{}
	for _, pilot := range results {
		classification = addPilotToClassification(classification, pilot)
	}
	return classification
}

func aggregate(results PilotsResults, lap Lap) PilotsResults {
	if isAggregated(results, lap) {
		return aggregateExistent(results, lap)
	}
	return aggregateNew(results, lap)
}

func isAggregated(results PilotsResults, lap Lap) bool {
	_, ok := results[lap.PilotCode]
	return ok
}

func aggregateNew(results PilotsResults, lap Lap) PilotsResults {
	lapTime := convertLapTimeToFloat(lap)
	pilotResult, _ := results[lap.PilotCode]
	pilotResult.Pilot = lap.Pilot
	pilotResult.PilotCode = lap.PilotCode
	pilotResult.Laps = 1
	pilotResult.BestLap = 1
	pilotResult.Time = lapTime
	pilotResult.BestTime = lapTime
	pilotResult.TotalSpeed = convertSpeedToFloat(lap)
	results[lap.PilotCode] = pilotResult
	return results
}

func aggregateExistent(results PilotsResults, lap Lap) PilotsResults {
	lapTime := convertLapTimeToFloat(lap)
	pilotResult, _ := results[lap.PilotCode]
	pilotResult.Laps++
	pilotResult.Time = pilotResult.Time + lapTime
	if lapTime < pilotResult.BestTime {
		pilotResult.BestLap = pilotResult.Laps
		pilotResult.BestTime = lapTime
	}
	pilotResult.TotalSpeed = pilotResult.TotalSpeed + convertSpeedToFloat(lap)
	results[lap.PilotCode] = pilotResult
	return results
}

func addPilotToClassification(classification Classification, pilot PilotResult) Classification {
	totalClassified := len(classification)
	if isClassificationEmpty(totalClassified) {
		return initializeClassification(classification, pilot)
	}
	return updateClassification(classification, pilot, totalClassified)
}

func isClassificationEmpty(totalClassified int) bool {
	return totalClassified == 0
}

func initializeClassification(classification Classification, pilot PilotResult) Classification {
	classification[0] = pilot
	return classification
}

func updateClassification(actualClassification Classification, pilot PilotResult, lastPosition int) Classification {
	newClassification := Classification{}
	for position := 0; position < lastPosition; position++ {
		newClassification[position] = actualClassification[position]
		if hasBetterResults(pilot, actualClassification[position]) {
			return updatePositions(actualClassification, newClassification, pilot, position, lastPosition)
		}
	}
	newClassification[lastPosition] = pilot
	return newClassification
}
func hasBetterResults(pilot PilotResult, classificatedPilot PilotResult) bool {
	return hasMoreLapsCompleted(pilot.Laps, classificatedPilot.Laps) &&
		hasBetterTime(pilot.Time, classificatedPilot.Time)
}

func hasMoreLapsCompleted(new, actual int) bool {
	return new >= actual
}

func hasBetterTime(new, actual float64) bool {
	return new < actual
}

func updatePositions(actualClassification, newClassification Classification, pilot PilotResult, actualPosition, lastPosition int) Classification {
	newClassification[actualPosition] = pilot
	for position := actualPosition; position < lastPosition; position++ {
		newClassification[position+1] = actualClassification[position]
	}
	return newClassification
}

func generateFinalResult(classification Classification) FinalResult {
	classificationDetailed := map[int]PilotResultDetailed{}
	bestTime := 999999.9

	for position := 0; position < len(classification); position++ {
		pd := classificationDetailed[position]
		pd.Pilot = classification[position].Pilot
		pd.PilotCode = classification[position].PilotCode
		pd.Laps = classification[position].Laps
		pd.Time = convertLapTimeToString(classification[position].Time)
		pd.BestLap = classification[position].BestLap
		pd.BestTime = convertLapTimeToString(classification[position].BestTime)
		pd.AverageSpeed = convertSpeedToString(classification[position].TotalSpeed / float64(classification[position].Laps))
		pd.Difference = convertLapTimeToString(classification[position].Time - classification[0].Time)
		bestTime = getBestRaceTime(classification[position].BestTime, bestTime)
		classificationDetailed[position] = pd
	}

	return FinalResult{
		Classification: classificationDetailed,
		BestRaceTime:   convertLapTimeToString(bestTime),
	}
}

func getBestRaceTime(new, actual float64) float64 {
	if new < actual {
		return new
	}
	return actual
}
