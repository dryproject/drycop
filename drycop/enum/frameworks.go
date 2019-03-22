/* This is free and unencumbered software released into the public domain. */

package enum

// Framework is an enumeration type for development frameworks
type Framework int

// The list of supported development frameworks
const (
	UnknownFramework Framework = iota
	Arduino
	Flutter
)

func (framework Framework) String() string {
	switch framework {
	case Arduino:
		return "Arduino"
	case Flutter:
		return "Flutter"
	case UnknownFramework:
	default:
		break
	}
	return "Unknown"
}
