/* This is free and unencumbered software released into the public domain. */

package enum

import "fmt"

// Language is an enumeration type for programming languages
type Language int

// The list of supported programming languages
const (
	UnknownLanguage Language = iota
	C
	Csharp
	Cpp
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

func (language *Language) Type() string {
	return "language"
}

func (language *Language) String() string {
	switch *language {
	case C:
		return "c"
	case Csharp:
		return "csharp"
	case Cpp:
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

func (language *Language) Set(input string) error {
	switch input {
	case "c":
		*language = C
	case "csharp":
		*language = Csharp
	case "cpp":
		*language = Cpp
	case "lisp":
		*language = CommonLisp
	case "d":
		*language = D
	case "dart":
		*language = Dart
	case "dry":
		*language = DRY
	case "elixir":
		*language = Elixir
	case "erlang":
		*language = Erlang
	case "go":
		*language = Go
	case "java":
		*language = Java
	case "javascript":
		*language = JavaScript
	case "julia":
		*language = Julia
	case "kotlin":
		*language = Kotlin
	case "lua":
		*language = Lua
	case "markdown":
		*language = Markdown
	case "objectivec":
		*language = ObjectiveC
	case "ocaml":
		*language = OCaml
	case "php":
		*language = PHP
	case "python":
		*language = Python
	case "restructuredtext":
		*language = RestructuredText
	case "ruby":
		*language = Ruby
	case "rust":
		*language = Rust
	case "swift":
		*language = Swift
	case "typescript":
		*language = TypeScript
	case "yaml":
		*language = YAML
	case "zig":
		*language = Zig
	default:
		*language = UnknownLanguage
		return fmt.Errorf("unknown language: %s", input)
	}
	return nil
}

func (language *Language) TestCommand() string {
	return "" // TODO
}
