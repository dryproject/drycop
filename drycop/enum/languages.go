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
	Markdown
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
		return "c"
	case Csharp:
		return "csharp"
	case Cxx:
		return "cpp"
	case CommonLisp:
		return "lisp"
	case D:
		return "d"
	case Dart:
		return "dart"
	case DRY:
		return "dry"
	case Elixir:
		return "elixir"
	case Erlang:
		return "erlang"
	case Go:
		return "go"
	case Java:
		return "java"
	case JavaScript:
		return "javascript"
	case Julia:
		return "julia"
	case Kotlin:
		return "kotlin"
	case Lua:
		return "lua"
	case Markdown:
		return "markdown"
	case ObjectiveC:
		return "objectivec"
	case OCaml:
		return "ocaml"
	case PHP:
		return "php"
	case Python:
		return "python"
	case RestructuredText:
		return "restructuredtext"
	case Ruby:
		return "ruby"
	case Rust:
		return "rust"
	case Swift:
		return "swift"
	case TypeScript:
		return "typescript"
	case YAML:
		return "yaml"
	case Zig:
		return "zig"
	case UnknownLanguage:
	}
	return "unknown"
}
