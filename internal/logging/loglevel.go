package logging

import "strings"

var LevelPriority = map[string]int{
	"info":    1,
	"warning": 2,
	"error":   3,
}

func ShouldLog(entryLevel string, minLevel string) bool {
	entryLevel = strings.ToLower(entryLevel)
	minLevel = strings.ToLower(minLevel)

	entryPriority, ok1 := LevelPriority[entryLevel]
	minPriority, ok2 := LevelPriority[minLevel]

	if !ok1 || !ok2 {
		return false
	}

	return entryPriority >= minPriority
}
