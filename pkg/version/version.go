// Package version is used to access the current CLI version.
package version

import "fmt"

// major is the major version number
const major = 0

// minor is the minor version number
const minor = 1

// patch is the patch version number
const patch = 0

// GetVersion returns the full version string for the current GROOT software
func GetVersion() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

// GetBaseVersion returns the major minor version string for the current GROOT software
func GetBaseVersion() string {
	return fmt.Sprintf("%d.%d", major, minor)
}
