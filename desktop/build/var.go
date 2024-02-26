package build

var (
	Version = "dev" // Version of the app, set by build flags. It can be "m-<commit hash>" or "v<semver>"
	AppName = "PhimTor"

	ServerAddr = "http://localhost:8080"
	IsLocal    = "true"
)
