package race

// Race represents a race
type Race []Lap

// Lap represents a lap
type Lap struct {
	Time            string
	PilotCode       string
	Pilot           string
	LapNumber       string
	LapTime         string
	LapSpeedAverage string
}

// PilotResult represents a pilot result from race
type PilotResult struct {
	PilotCode  string
	Pilot      string
	Laps       int
	Time       float64
	BestLap    int
	BestTime   float64
	TotalSpeed float64
}

// PilotsResults many PilotResult
type PilotsResults map[string]PilotResult

// Classification represents a race classification
type Classification map[int]PilotResult

// FinalResult represents all data from race
type FinalResult struct {
	Classification map[int]PilotResultDetailed
	BestRaceTime   string
}

// PilotResultDetailed represents PilotResult with more info and formatted to output
type PilotResultDetailed struct {
	PilotCode    string
	Pilot        string
	Laps         int
	Time         string
	BestLap      int
	AverageSpeed string
	BestTime     string
	Difference   string
}
