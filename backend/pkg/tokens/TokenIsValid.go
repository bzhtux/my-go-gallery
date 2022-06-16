package tokens

func TokenIsValid(email_token, db_token string) bool {
	if email_token == db_token {
		return true
	} else {
		return false
	}
}
