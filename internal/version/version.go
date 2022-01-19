// Package version records versioning information about this module.
package version

import (
	"fmt"
)

const (
	Major = 0
	Minor = 0
	Patch = 0
)

// String formats the semantic version string for this module.
func String() string {
	return fmt.Sprintf("v%d.%d.%d", Major, Minor, Patch)
}
