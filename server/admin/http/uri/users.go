package uri

func ListUsers(nextPageToken string) string {
	if nextPageToken == "" {
		return "/users"
	}
	return "/users?nextPageToken=" + nextPageToken
}
