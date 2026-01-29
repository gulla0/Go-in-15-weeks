DAY 1 — Go Fundamentals

• Go is a compiled language that produces native executables (binaries).
• Compilation time checks code correctness; runtime is when the program executes.
• A binary/executable runs directly on the OS without an interpreter.
• `go run` compiles and runs immediately (no saved binary).
• `go build` compiles and produces a reusable executable.
• A Go program must use `package main` to be executable.
• `func main()` is the single entry point; it takes no arguments and returns nothing.
• Programs communicate via side effects (printing, files, network), not return values.
• Go code is organized into packages; all `.go` files in a directory form one package.
• Variables are statically typed; types are fixed at compile time.
• `var` declares explicitly; `:=` declares + assigns inside functions only.
• Basic types learned: `int`, `string`, `bool`.
• `if` requires a boolean condition; no truthy/falsy values.
• `for` is the only loop construct (classic, while-style, infinite).
• `for` init/condition/post each allow exactly one expression.
• Go enforces formatting via `gofmt` / `go fmt`.
• `gofmt` enforces structure, not personal spacing preferences.
• A Go module is a directory with a `go.mod` file.
• `go mod init` is run once per project; optional for throwaway snippets.

