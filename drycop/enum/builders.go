/* This is free and unencumbered software released into the public domain. */

package enum

// Builder is an enumeration type for build systems
type Builder int

// The list of supported build systems
const (
	UnknownBuilder Builder = iota
	Autoconf
	Automake
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
	case Autoconf:
		return "autoconf"
	case Automake:
		return "automake"
	case CMake:
		return "cmake"
	case DartPub:
		return "pub"
	case ElixirHex:
		return "hex"
	case GoBuild:
		return "go"
	case Gradle:
		return "gradle"
	case Make:
		return "make"
	case Maven:
		return "mvn"
	case OCamlDune:
		return "dune"
	case PythonPIP:
		return "pip"
	case RubyGems:
		return "gem"
	case SwiftPackageManager:
		return "spm"
	case UnknownBuilder:
	}
	return "unknown"
}
