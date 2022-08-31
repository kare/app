package application

import "fmt"

// Build contains application build and version information.
type Build struct {
	Time     string `json:"time"`
	Version  string `json:"version"`
	Revision string `json:"revision"`
}

func require(name, value string) error {
	if value == "" {
		return fmt.Errorf("application: Build.%s is a required value", name)
	}
	return nil
}

// NewBuild creates a new build version.
func NewBuild(time, version, revision string) (*Build, error) {
	if err := require("Time", time); err != nil {
		return nil, err
	}
	if err := require("Version", version); err != nil {
		return nil, err
	}
	if err := require("Revision", revision); err != nil {
		return nil, err
	}
	return &Build{
		Time:     time,
		Version:  version,
		Revision: revision,
	}, nil
}