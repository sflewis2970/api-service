package common

import (
	"strings"
	"time"
)

func GetFormattedTime(timeNow time.Time, timeFormat string) string {
	return timeNow.Format(timeFormat)
}

// Build UUID string
func BuildUUID(uuid string, delimiter string, nbrOfGroups int) string {
	newUUID := ""

	uuidList := strings.Split(uuid, delimiter)
	for key, value := range uuidList {
		if key < nbrOfGroups {
			newUUID = newUUID + value
		}
	}

	return newUUID
}
