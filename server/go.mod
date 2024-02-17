module github.com/phimtorr/phimtor/server

go 1.21.6

replace github.com/phimtorr/phimtor/common => ../common

require (
	github.com/friendsofgo/errors v0.9.2
	github.com/labstack/echo/v4 v4.11.4
	github.com/lib/pq v1.10.9
	github.com/oapi-codegen/runtime v1.1.1
	github.com/phimtorr/phimtor/common v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.32.0
	github.com/simukti/sqldb-logger v0.0.0-20230108155151-646c1a075551
	github.com/volatiletech/null/v8 v8.1.2
	github.com/volatiletech/sqlboiler/v4 v4.16.2
	github.com/volatiletech/strmangle v0.0.6
	github.com/ziflex/lecho/v3 v3.5.0
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/labstack/gommon v0.4.2 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/volatiletech/inflect v0.0.1 // indirect
	github.com/volatiletech/randomize v0.0.1 // indirect
	golang.org/x/crypto v0.17.0 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	golang.org/x/time v0.5.0 // indirect
	golang.org/x/xerrors v0.0.0-20220609144429-65e65417b02f // indirect
)
