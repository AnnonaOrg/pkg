package token

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var (
	//ErrMissingHeader means the `Authorization`header was empty
	ErrMissingHeader = errors.New("the length of the `Authorization` header is zero.")
)

// Context is the context of the JSON web token
type Context struct {
	// 用户id
	ID uint64
	// 用户uuid
	UUID string
	// 用户名
	Username string
	// SafeCode string
	// 过期日期时间戳 time unix
	Expiry int64
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		//make sure the `alg`is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

// parse validates the token with the specified secret,
// and returns the context if the token was valid
func Parse(tokenString string, secret string) (*Context, error) {
	ctx := &Context{}

	//parse the token
	token, err := jwt.Parse(tokenString, secretFunc(secret))

	//parse error.
	if err != nil {
		return ctx, err
	} else if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID = uint64(claims["id"].(float64))
		ctx.UUID = claims["uuid"].(string)
		ctx.Username = claims["username"].(string)
		ctx.Expiry = int64(claims["expiry"].(float64))
		return ctx, nil
	} else {
		return ctx, err
	}
}

// parseRequest gets the token from the header and
// pass it to the parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context, error) {
	header := c.Request.Header.Get("Authorization")

	//Load the jwt secret from config
	secret := GetJwtSecret()
	if len(header) == 0 {
		return &Context{}, ErrMissingHeader
	}

	var t string
	//parse the header to get the token part
	fmt.Sscanf(header, "Bearer %s", &t)
	return Parse(t, secret)
}

// sign sings the context with the specified secret
func Sign(ctx *gin.Context, c Context, secret string) (tokenString string, err error) {
	//Load the jwt secret from the Gin config if the secret isn't specified
	if secret == "" {
		secret = GetJwtSecret()
	}
	//The token content
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       c.ID,
		"uuid":     c.UUID,
		"username": c.Username,
		"expiry":   c.Expiry,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
	})
	//sign the token with the specified secret.
	tokenString, err = token.SignedString([]byte(secret))
	return
}
