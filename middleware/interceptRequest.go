package middleware

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
)

var (
	clientID     = "example-app"
	clientSecret = "ZXhhbXBsZS1hcHAtc2VjcmV0"
	issuerURL    = "http://127.0.0.1:5556/dex"
	redirectURL  = "http://127.0.0.1:5555/callback"
)

func InterceptBearerToken(ctx *gin.Context) {

	tokenString := ctx.GetHeader("Authorization")
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
		ctx.Abort()
		return
	}
	err := authorize(ctx, strings.Split(tokenString, " ")[1])
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Not Authorized"})
		ctx.Abort()
		return
	}
	ctx.Next()
}

func authorize(ctx *gin.Context, bearerToken string) error {

	provider, err := oidc.NewProvider(ctx, issuerURL)
	if err != nil {
		log.Fatal("errDzp", err)
		return err
	}
	accessTokenVerifier := provider.Verifier(&oidc.Config{ClientID: clientID})
	accessToken, err := accessTokenVerifier.Verify(ctx, bearerToken)
	if err != nil {
		return err
	}

	var claims struct {
		Email    string   `json:"email"`
		Sub      string   `json:"sub"`
		Verified bool     `json:"email_verified"`
		Name     string   `json:"name"`
		Groups   []string `json:"groups"`
	}
	if err := accessToken.Claims(&claims); err != nil {
		return fmt.Errorf("failed to parse claims: %v", err)
	}
	if !claims.Verified {
		return fmt.Errorf("email (%q) in returned claims was not verified", claims.Email)
	}

	ctx.Set("sub", claims.Sub)

	return nil
}
