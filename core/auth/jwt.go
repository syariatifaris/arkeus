package auth

import (
	"time"

	"github.com/syariatifaris/arkeus/core/auth/jwtgo"
)

var ExpireAt = time.Now().Add(time.Hour * 24 * 360).Unix()

type Jwt interface {
	CreateToken() (string, error)
	ValidateToken(token string) (bool, error)
	ValidateTokenWithClaim(token string) (bool, error, *jwtgo.JwtClaims)
	SetIssuer(string)
	SetBearerData(string)
}

func NewJwtAuth(secretKey string) Jwt {
	return jwtgo.NewJwtAuth(jwtgo.JwtAuth{
		SecretKey:    secretKey,
		UnixExpireAt: ExpireAt,
	})
}
