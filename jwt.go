package nonecode

import (
	"errors"
	"fmt"
	"time"

	// "guzhengxiao/goldaigpt-serv/model"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// var (
// 	key = []byte(ConfigData.JWT.Secret)
// )

// JWT represents an authentication entry
type JWT struct {
	Secret             []byte
	Issuer             string
	DebugFlag          bool
	SessionExpiredTime time.Duration
	AuthorizationKey   string
}

// // NewJWT returns a new jwt
// func NewJWT(secret []byte) *JWT {
// 	return &JWT{Secret: secret}
// }

type Claims struct {
	ID uint `json:"id"`
	// Username string `json:"username"`
	Nonce string `json:"nonce"`
	Scope string `json:"scope"`
	jwt.StandardClaims
}
type JwtUser struct {
	// Key        string
	ID    uint
	Scope string
}

// New new a jwt token scope 判断用户类型使用
func (j *JWT) New(claims Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(j.Secret)
	return ss, err
}

// Verify  the token whether ok
func (j *JWT) Verify(authToken string) (claims *Claims, ok bool) {
	tokens := strings.Split(authToken, "Bearer ")
	if len(tokens) < 2 {
		return nil, false
	}
	tokenStr := tokens[1]

	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok1 := token.Method.(*jwt.SigningMethodHMAC); !ok1 {
			return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
		}
		return j.Secret, nil
	})
	if err != nil {
		// log.Fatalf("Error caught when parse token: %s", err)
		return nil, false
	}

	claims, ok = token.Claims.(*Claims)
	if ok && token.Valid {
		return claims, true
	}

	return claims, false
}

func (j *JWT) NewAccessToken(user *JwtUser) (string, error) {
	if user == nil {
		return "", errors.New("user is nil")
	}
	expiryTime := time.Now().Add(j.SessionExpiredTime).Unix()

	return j.New(Claims{
		ID:    user.ID,
		Nonce: RamdomAll.String(24),
		Scope: user.Scope,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime,
			Issuer:    j.Issuer,
		},
	})
}

// func (j *JWT) NewAdminUserAccessToken(user *JwtUser) (string, error) {

// 	if user == nil {
// 		return "", errors.New("UserName empty")
// 	}
// 	userScope := model.ScopeUserAdmin
// 	if user.AdminLevel == 1 {
// 		userScope = model.ScopeUserCustomer
// 	}

// 	expiryTime := time.Now().Add(ConfigData.JWT.SessionExpiredTime).Unix()
// 	claims := Claims{
// 		ID:       user.ID,
// 		Username: user.Username,
// 		Nonce:    RandomString(24),
// 		Scope:    userScope,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expiryTime,
// 			Issuer:    ConfigData.JWT.Issuer,
// 		},
// 	}

// 	return j.New(claims)
// }
// func (j *JWT) NewUserAccessToken(user *model.User) (string, error) {
// 	if user == nil {
// 		return "", errors.New("UserName empty")
// 	}
// 	expiryTime := time.Now().Add(ConfigData.JWT.SessionExpiredTime).Unix()
// 	claims := Claims{
// 		ID:       user.ID,
// 		Username: user.Username,
// 		Nonce:    RandomString(24),
// 		Scope:    model.ScopeUser,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: expiryTime,
// 			Issuer:    ConfigData.JWT.Issuer,
// 		},
// 	}

// 	return j.New(claims)
// }
