package handler

import (
	"context"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xylonx/icc-core/internal/model"
	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

type auth struct {
}

func ctxWithAuthToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, auth{}, token)
}

func mustGetTokenFromCtx(ctx context.Context) string {
	return ctx.Value(auth{}).(string) //nolint:forcetypeassert
}

func HttpAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			zapx.Error("auth header missing", zap.String("Authorization", authHeader))
			respAbortWithUnauthError(c, ErrAuthHeaderNotFound)
			return
		}

		token := authHeader[7:]

		if err := model.CheckTokenExists(c.Request.Context(), token); err != nil {
			zapx.Error("check token exists failed", zap.Error(err), zap.String("token", token))
			respAbortWithUnauthError(c, err)
			return
		}

		ctx := ctxWithAuthToken(c.Request.Context(), token)
		c.Request = c.Request.Clone(ctx)
	}
}
