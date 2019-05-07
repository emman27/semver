package semver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	assert := assert.New(t)
	type result struct {
		version *Semver
		wantErr bool
	}
	tests := map[string]result{
		"1.0.0":              result{&Semver{Major: 1, Minor: 0, Patch: 0, PreRelease: "", BuildMetadata: ""}, false},
		"1.2.3":              result{&Semver{Major: 1, Minor: 2, Patch: 3, PreRelease: "", BuildMetadata: ""}, false},
		"1.2.3-rc.1":         result{&Semver{Major: 1, Minor: 2, Patch: 3, PreRelease: "rc.1", BuildMetadata: ""}, false},
		"1.2.3-rc.1+0abcdef": result{&Semver{Major: 1, Minor: 2, Patch: 3, PreRelease: "rc.1", BuildMetadata: "0abcdef"}, false},
		"1.2.3+0abcdef":      result{&Semver{Major: 1, Minor: 2, Patch: 3, PreRelease: "", BuildMetadata: "0abcdef"}, false},
	}
	for semverString, expectedResult := range tests {
		v, err := FromString(semverString)
		if expectedResult.wantErr {
			assert.NotNil(err)
		} else {
			assert.Equal(expectedResult.version, v)
		}
	}
}
