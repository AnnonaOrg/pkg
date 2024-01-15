package token

var jwt_secret string

func SetJwtSecret(jwtSecret string) {
	if len(jwtSecret) > 0 {
		jwt_secret = jwtSecret
	} else {
		jwt_secret = "jwt_secret"
	}
}
func GetJwtSecret() string {
	if len(jwt_secret) > 0 {
		return jwt_secret
	}
	return "jwt_secret"
}
