/* This is free and unencumbered software released into the public domain. */

package enum

import "fmt"

// Builder is an enumeration type for build systems
type Builder int

// The list of supported build systems
const (
	UnknownBuilder Builder = iota
	NoBuilder
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

func (builder *Builder) Type() string {
	return "builder"
}

func (builder *Builder) String() string {
	switch *builder {
	case NoBuilder:
		return "none"
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

func (builder *Builder) Set(input string) error {
	switch input {
	case "none":
		*builder = NoBuilder
	case "autoconf":
		*builder = Autoconf
	case "automake":
		*builder = Automake
	case "cmake":
		*builder = CMake
	case "pub":
		*builder = DartPub
	case "hex":
		*builder = ElixirHex
	case "go":
		*builder = GoBuild
	case "gradle":
		*builder = Gradle
	case "make":
		*builder = Make
	case "mvn":
		*builder = Maven
	case "dune":
		*builder = OCamlDune
	case "pip":
		*builder = PythonPIP
	case "gem":
		*builder = RubyGems
	case "spm":
		*builder = SwiftPackageManager
	default:
		*builder = UnknownBuilder
		return fmt.Errorf("unknown builder: %s", input)
	}
	return nil
}
