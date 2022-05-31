package common

import (
	"math/rand"
	"strings"
	"time"
)

// Generate random float value
func GenerateFloat64Vals(valRange float64, minVal float64) float64 {
	newRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Return randomly generated float64
	return (newRand.Float64() * valRange) + minVal
}

// Build formatted time string
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

// Utility to build a slice of strings
func BuildStrSlice(orgStr string, delimiter string) []string {
	newStrList := []string{}

	strList := strings.Split(orgStr, delimiter)
	for _, value := range strList {
		newStrList = append(newStrList, value)
	}

	return newStrList
}

// Utility to build a slice of strings
func BuildDelimitedStr(strs []string, delimiter string) string {
	newStr := ""

	strSize := len(strs)
	for idx := 0; idx < strSize-1; idx++ {
		newStr = newStr + strs[idx] + delimiter
	}

	newStr = newStr + strs[strSize-1]

	return newStr
}

// Utility to move string item to a different position within the list
func ShuffleList(strList []string) []string {
	rand.Shuffle(len(strList), func(idx1, idx2 int) {
		strList[idx1], strList[idx2] = strList[idx2], strList[idx1]
	})

	return strList
}
