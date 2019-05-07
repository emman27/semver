package semver

import (
	"strconv"
	"strings"
)

// Semver contains all information about a semantic version
type Semver struct {
	Major         uint64
	Minor         uint64
	Patch         uint64
	PreRelease    string
	BuildMetadata string
}

// IsPreRelease returns true if the given version is a pre-release version
func (s *Semver) IsPreRelease() bool {
	return s.PreRelease == ""
}

// HasBuildMetadata returns true if there is any build metadata associated with the given version
func (s *Semver) HasBuildMetadata() bool {
	return s.BuildMetadata != ""
}

// FromString initializes a Semver object from a version string
// Will accept any string that conforms to the semver standard
// Will also accept strings prefixed with 'v', for convenience purposes
func FromString(s string) (*Semver, error) {
	onlyVersion := strings.FieldsFunc(s, isHyphen)[0]
	dotSeparated := strings.FieldsFunc(onlyVersion, isFullStop)
	major, err := strconv.Atoi(dotSeparated[0])
	if err != nil {
		return nil, err
	}
	minor, err := strconv.Atoi(dotSeparated[1])
	if err != nil {
		return nil, err
	}
	patch, err := strconv.Atoi(dotSeparated[2])
	if err != nil {
		return nil, err
	}
	return &Semver{
		Major: uint64(major),
		Minor: uint64(minor),
		Patch: uint64(patch),
	}, nil
}

func isFullStop(r rune) bool {
	return r == '.'
}

func isHyphen(r rune) bool {
	return r == '-'
}
