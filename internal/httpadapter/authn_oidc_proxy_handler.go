package httpadapter

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"iconrepo/internal/app/security/authn"
	"iconrepo/internal/app/security/authr"
	"iconrepo/internal/app/services"
	"iconrepo/internal/logging"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const forwardedUserInfoKey = "forwarded-user-info"

func checkOIDCProxyAuthentication(authRService services.AuthorizationService) func(c *gin.Context) {
	return func(c *gin.Context) {
		logger := zerolog.Ctx(c.Request.Context()).With().Str(logging.MethodLogger, "checkOIDCProxyAuthentication").Logger()

		abort := func(details string) {
			logger.Debug().Str("details", details).Msg("Request for %v not authenticated")
			c.AbortWithStatus(401)
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			abort("missing authorization header")
			return
		}

		splitToken := strings.Split(authHeader, "Bearer ")
		reqToken := splitToken[1]

		if reqToken == "" {
			abort(fmt.Sprintf("No bearer token in auth header: %s", authHeader))
			return
		}

		tokenParts := strings.Split(reqToken, ".")
		if len(tokenParts) != 3 {
			abort(fmt.Sprintf("unexpected number of token parts (%d): %s", len(tokenParts), reqToken))
		}

		token, tokenDecodingErr := base64.RawStdEncoding.DecodeString(tokenParts[1])
		if tokenDecodingErr != nil {
			abort(fmt.Sprintf("failed to decode token: %s (%v)", tokenParts[1], tokenDecodingErr))
			return
		}

		receivedClaims := claims{}
		unmarshalErr := json.Unmarshal(token, &receivedClaims)
		if unmarshalErr != nil {
			abort(fmt.Sprintf("failed to unmarshal token: %s, %v", token, unmarshalErr))
			return
		}

		if logger.GetLevel() == zerolog.DebugLevel {
			logger.Debug().Interface("claims", receivedClaims).Msg("claims received")
		}

		groupIds := authr.GroupNamesToGroupIDs(receivedClaims.Groups)
		userInfo := authr.UserInfo{
			// FIXME: Use other than local-domain
			UserId:      authn.LocalDomain.CreateUserID(receivedClaims.Email),
			Groups:      groupIds,
			Permissions: authRService.GetPermissionsForGroups(groupIds),
		}
		c.Set(forwardedUserInfoKey, userInfo)

		r := c.Request
		ctx := context.WithValue(r.Context(), authr.UserInfoCtxKey, userInfo)
		c.Request = r.WithContext(ctx)
	}
}
