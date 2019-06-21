/* This is free and unencumbered software released into the public domain. */

package enum

import "fmt"

// Framework is an enumeration type for development frameworks
type Framework int

// The list of supported development frameworks
const (
	UnknownFramework Framework = iota
	NoFramework
	Android
	Arduino
	Flutter
)

func (framework *Framework) Type() string {
	return "framework"
}

func (framework *Framework) String() string {
	switch *framework {
	case NoFramework:
		return "none"
	case Android:
		return "android"
	case Arduino:
		return "arduino"
	case Flutter:
		return "flutter"
	case UnknownFramework:
	}
	return "unknown"
}

func (framework *Framework) Set(input string) error {
	switch input {
	case "none":
		*framework = NoFramework
	case "android":
		*framework = Android
	case "arduino":
		*framework = Arduino
	case "flutter":
		*framework = Flutter
	default:
		*framework = UnknownFramework
		return fmt.Errorf("unknown framework: %s", input)
	}
	return nil
}
