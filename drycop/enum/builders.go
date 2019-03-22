/* This is free and unencumbered software released into the public domain. */

package enum

// Builder is an enumeration type for build systems
type Builder int

// The list of supported build systems
const (
	UnknownBuilder Builder = iota
	Autotools
	CMake
	DartPub
	ElixirHex
	GoBuild
	Gradle
	Make
	Maven
	OCamlDune
	PythonPIP
	RubyGems
	SwiftPackageManager
)

func (builder Builder) String() string {
	switch builder {
	case Autotools:
		return "Autotools"
	case CMake:
		return "CMake"
	case DartPub:
		return "Pub"
	case ElixirHex:
		return "Hex"
	case GoBuild:
		return "Go"
	case Gradle:
		return "Gradle"
	case Make:
		return "Make"
	case Maven:
		return "Maven"
	case OCamlDune:
		return "Dune"
	case PythonPIP:
		return "PIP"
	case RubyGems:
		return "RubyGems"
	case SwiftPackageManager:
		return "SPM"
	case UnknownBuilder:
	default:
		break
	}
	return "Unknown"
}
