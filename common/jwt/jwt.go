package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
	"uminer/common/errors"
)

type TokenClaims struct {
	UserId    string `json:"userId"`
	CreatedAt int64  `json:"createdAt"`
	jwt.StandardClaims
}

func CreateToken(uid, secret string, expiration time.Duration) (string, error) {
	nowTime := time.Now()
	nowTimeUnix := nowTime.Unix()
	claims := TokenClaims{
		UserId:    uid,
		CreatedAt: nowTimeUnix,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  nowTimeUnix,
			ExpiresAt: nowTime.Add(expiration).Unix(),
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", errors.Errorf(err, errors.ErrorCreateTokenFailed)
	}
	return token, nil
}

func ParseToken(tokenStr string, secret string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, errors.Errorf(err, errors.ErrorParseTokenFailed)
	}
	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		if time.Now().Unix() > claims.ExpiresAt {
			return nil, errors.Errorf(err, errors.ErrorParseTokenFailed)
		}
		return claims, nil
	} else {
		return nil, errors.Errorf(err, errors.ErrorTokenInvalid)
	}
}

func FormatToken(prefix, token string) string {
	if prefix == "" {
		return token
	}
	return fmt.Sprintf("%s_%s", prefix, token)
}
