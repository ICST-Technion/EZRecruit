package helpers

import (
	"strings"

	set "github.com/deckarep/golang-set"
)

// CreateSetFromSlice returns a set contains all items in the given slice. if slice is nil, returns empty set.
func CreateSetFromSlice(slice []string) set.Set {
	if slice == nil {
		return set.NewSet()
	}

	result := set.NewSet()
	for _, item := range slice {
		result.Add(strings.ToLower(item))
	}

	return result
}

// SetContainsAll returns true if the given set contains all elements. Otherwise, returns false.
func SetContainsAll(set set.Set, elements []string) bool {
	for _, element := range elements {
		if !set.Contains(strings.ToLower(element)) {
			return false
		}
	}

	return true
}

// GetIntersectionSize returns size of intersection of two slices.
func GetIntersectionSize(slice1, slice2 []string) int {
	hitMap := make(map[string]struct{}, len(slice1))

	for _, elem := range slice1 {
		hitMap[elem] = struct{}{}
	}

	hits := 0

	for _, elem := range slice2 {
		if _, found := hitMap[elem]; found {
			hits++
		}
	}

	return hits
}
