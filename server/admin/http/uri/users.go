package uri

func ListUsers(nextPageToken string) string {
	if nextPageToken == "" {
		return "/users"
	}
	return "/users?nextPageToken=" + nextPageToken
}

func ViewUser(uid string) string {
	return "/users/" + uid
}

func UpdatePremium(uid string) string {
	return "/users/" + uid + "/update-premium"
}
