package helpers

import (
	set "github.com/deckarep/golang-set"
	"strings"
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
