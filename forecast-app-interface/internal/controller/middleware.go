package controller

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/valyala/fasthttp"
)

var (
	secretKey = []byte(os.Getenv("SECRET_KEY"))
)

func AuthMiddleware(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		tokenString := string(ctx.Request.Header.Cookie("token"))
		if err := verifyToken(tokenString); err != nil {
			ctx.Redirect("/", fasthttp.StatusSeeOther)
			return
		}

		next(ctx)
	}
}

func GetUserPredictionsMiddleware(next fasthttp.RequestHandler, r *Router) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		username := string(ctx.Request.Header.Cookie("username"))
		userPredictsNames, err := r.useCase.GetForecasts(username)
		if err != nil {
			log.Println("cannot get user predicts")
		} else {
			ctx.SetUserValue("user_predicts_names", userPredictsNames)
		}

		next(ctx)
	}
}

func createToken(username string) (string, time.Time, error) {
	expiration := time.Now().Add(time.Hour * 24)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      expiration.Unix(),
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiration, nil
}

func verifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}
