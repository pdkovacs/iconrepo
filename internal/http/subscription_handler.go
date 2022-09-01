package http

import (
	"context"
	"errors"
	"igo-repo/internal/app/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"nhooyr.io/websocket"
)

type socketAdapter struct {
	wsConn *websocket.Conn
}

func (sa *socketAdapter) CloseRead(ctx context.Context) context.Context {
	return sa.wsConn.CloseRead(ctx)
}

func (sa *socketAdapter) Close() error {
	return sa.wsConn.Close(websocket.StatusPolicyViolation, "connection too slow to keep up with messages")
}

func (sa *socketAdapter) Write(ctx context.Context, msg string) error {
	return sa.wsConn.Write(ctx, websocket.MessageText, []byte(msg))
}

func subscriptionHandler(ns *services.Notification, logger zerolog.Logger) gin.HandlerFunc {
	return func(g *gin.Context) {
		wsConn, subsErr := websocket.Accept(g.Writer, g.Request, nil)
		if subsErr != nil {
			logger.Error().Msgf("Failed to accept WS connection request: %v", subsErr)
			g.Error(subsErr)
			g.AbortWithStatus(500)
			return
		}

		defer wsConn.Close(websocket.StatusInternalError, "")

		session := mustGetUserSession(g)

		curriedContext := wsConn.CloseRead(g.Request.Context())                                            // Clients wan't write to the WS.(?)
		subscriptionError := ns.Subscribe(curriedContext, &socketAdapter{wsConn}, session.UserInfo.UserId) // we block here until Error or Done

		if errors.Is(subscriptionError, context.Canceled) {
			return // Done
		}

		if websocket.CloseStatus(subscriptionError) == websocket.StatusNormalClosure ||
			websocket.CloseStatus(subscriptionError) == websocket.StatusGoingAway {
			return
		}
		if subscriptionError != nil {
			logger.Error().Msgf("%v", subscriptionError)
			return
		}
	}
}
