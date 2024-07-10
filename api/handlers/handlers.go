package handler

import (
	auth "api_gateway/genproto/authentication_service"
)

func NewAuthHandler(authSer auth.AuthenticationServiceClient) authHandler {
	return &authHandlerImpl{authService: authSer}
}

type Handlers struct {
	authHandler
}

func NewHandlers(authSer auth.AuthenticationServiceClient) *Handlers {
	auth := NewAuthHandler(authSer)
	return &Handlers{authHandler: auth}
}
