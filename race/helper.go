package race

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func convertLapTimeToFloat(lap Lap) float64 {
	laptime := strings.ReplaceAll(lap.LapTime, ":", "m")
	laptime = fmt.Sprintf("%ss", laptime)
	duration, _ := time.ParseDuration(laptime)
	return duration.Seconds()
}

func convertSpeedToFloat(lap Lap) float64 {
	lapSpeed := strings.ReplaceAll(lap.LapSpeedAverage, ",", ".")
	speed, _ := strconv.ParseFloat(lapSpeed, 64)
	return speed
}

func convertLapTimeToString(value float64) string {
	vString := fmt.Sprintf("%.3fs", value)
	duration, _ := time.ParseDuration(vString)
	return duration.String()
}

func convertSpeedToString(value float64) string {
	return fmt.Sprintf("%.3f", value)
}
