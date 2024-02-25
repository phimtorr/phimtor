module github.com/phimtorr/phimtor/server

go 1.21.6

replace github.com/phimtorr/phimtor/common => ../common

require (
	github.com/friendsofgo/errors v0.9.2
	github.com/go-chi/chi/v5 v5.0.12
	github.com/go-chi/render v1.0.3
	github.com/go-sql-driver/mysql v1.6.0
	github.com/golang-migrate/migrate/v4 v4.17.0
	github.com/oapi-codegen/runtime v1.1.1
	github.com/phimtorr/phimtor/common v0.0.0-00010101000000-000000000000
	github.com/rs/cors v1.10.1
	github.com/rs/zerolog v1.32.0
	github.com/simukti/sqldb-logger v0.0.0-20230108155151-646c1a075551
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.16.2
	github.com/volatiletech/strmangle v0.0.6
)

require (
	github.com/ajg/form v1.5.1 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
)
