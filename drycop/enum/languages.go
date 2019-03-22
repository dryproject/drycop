/* This is free and unencumbered software released into the public domain. */

package enum

// Language is an enumeration type for programming languages
type Language int

// The list of supported programming languages
const (
	UnknownLanguage Language = iota
	C
	Csharp
	Cxx
	CommonLisp
	D
	Dart
	DRY
	Elixir
	Erlang
	Go
	Java
	JavaScript
	Julia
	Kotlin
	Lua
	ObjectiveC
	OCaml
	PHP
	Python
	RestructuredText
	Ruby
	Rust
	Swift
	TypeScript
	YAML
	Zig
)

func (language Language) String() string {
	switch language {
	case C:
		return "C"
	case Csharp:
		return "C#"
	case Cxx:
		return "C++"
	case CommonLisp:
		return "Common Lisp"
	case D:
		return "D"
	case Dart:
		return "Dart"
	case DRY:
		return "DRY"
	case Elixir:
		return "Elixir"
	case Erlang:
		return "Erlang"
	case Go:
		return "Go"
	case Java:
		return "Java"
	case JavaScript:
		return "JavaScript"
	case Julia:
		return "Julia"
	case Kotlin:
		return "Kotlin"
	case Lua:
		return "Lua"
	case ObjectiveC:
		return "Objective-C"
	case OCaml:
		return "OCaml"
	case PHP:
		return "PHP"
	case Python:
		return "Python"
	case RestructuredText:
		return "reStructuredText"
	case Ruby:
		return "Ruby"
	case Rust:
		return "Rust"
	case Swift:
		return "Swift"
	case TypeScript:
		return "TypeScript"
	case YAML:
		return "YAML"
	case Zig:
		return "Zig"
	case UnknownLanguage:
	default:
		break
	}
	return "Unknown"
}
