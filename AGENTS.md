# Agent Guidelines for Tournament Winner

## Build/Test Commands
- **Build**: `go build -o tournament-winner`
- **Run**: `go run main.go` (serves on :3000)
- **Test**: `go test ./...` (no tests exist yet)
- **Single test**: `go test -run TestName ./package`
- **Format**: `go fmt ./...`
- **Vet**: `go vet ./...`
- **Generate**: `templ generate` (for .templ files)

## Project Structure
- Go 1.24.2 with chi router, templ templates, JWT auth
- `handlers/` - HTTP handlers with Register() method
- `middleware/` - HTTP middleware (JWT auth)
- `models/` - Data models and enums
- `views/` - Templ templates (*.templ files, *_templ.go generated)

## Code Style
- Standard Go formatting with `go fmt`
- Package names: lowercase, single word
- Struct names: PascalCase
- Interface names: PascalCase ending with -er when appropriate
- Constants: PascalCase or UPPER_CASE for enums
- Use `chi.Router` for routing, `http.Handler` for middleware
- JWT claims use json tags
- Error handling: return errors, use http.Error() for HTTP responses