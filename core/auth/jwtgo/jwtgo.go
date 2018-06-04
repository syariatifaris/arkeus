package jwtgo

import (
	"github.com/dgrijalva/jwt-go"
)

var JWtMethod = jwt.SigningMethodHS256

type JwtAuth struct {
	SecretKey    string
	UnixExpireAt int64
	Issuer       string
	Data         string
}

type JwtClaims struct {
	BearerData string
	jwt.StandardClaims
}

type JwtGo struct {
	claims    JwtClaims
	secretKey string
}

func NewJwtAuth(auth JwtAuth) *JwtGo {
	return &JwtGo{
		claims: JwtClaims{
			BearerData: auth.Data,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: auth.UnixExpireAt,
				Issuer:    auth.Issuer,
			},
		},
		secretKey: auth.SecretKey,
	}
}

func (j *JwtGo) CreateToken() (string, error) {
	tokenClaim := jwt.NewWithClaims(JWtMethod, j.claims)
	return tokenClaim.SignedString([]byte(j.secretKey))
}

func (j *JwtGo) ValidateToken(token string) (bool, error) {
	authToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return false, err
	}

	return authToken.Valid, nil
}

func (j *JwtGo) ValidateTokenWithClaim(token string) (bool, error, *JwtClaims) {
	authToken, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return false, err, nil
	}

	if claims, ok := authToken.Claims.(*JwtClaims); ok && authToken.Valid {
		return true, nil, claims
	}

	return false, nil, nil
}

func (j *JwtGo) SetIssuer(newIssuer string) {
	j.claims.Issuer = newIssuer
}

func (j *JwtGo) SetBearerData(bearerData string) {
	j.claims.BearerData = bearerData
}
