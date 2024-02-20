module github.com/phimtorr/phimtor/desktop

go 1.21.6

replace github.com/phimtorr/phimtor/common => ../common

require (
	github.com/a-h/templ v0.2.543
	github.com/go-chi/chi/v5 v5.0.12
	github.com/oapi-codegen/runtime v1.1.1
	github.com/phimtorr/phimtor/common v0.0.0-00010101000000-000000000000
	github.com/rs/cors v1.8.3
	github.com/rs/zerolog v1.32.0
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/sys v0.15.0 // indirect
)
