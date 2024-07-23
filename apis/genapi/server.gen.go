// Package genapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package genapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

// AuthRes defines model for AuthRes.
type AuthRes struct {
	ExpireAt int64  `json:"expireAt"`
	Id       int32  `json:"id"`
	Username string `json:"username"`
}

// LoginReq defines model for LoginReq.
type LoginReq struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

// LoginRes defines model for LoginRes.
type LoginRes struct {
	Token string `json:"token"`
}

// AuthAuthParams defines parameters for AuthAuth.
type AuthAuthParams struct {
	Authorization string `json:"authorization"`
}

// AuthLoginJSONRequestBody defines body for AuthLogin for application/json ContentType.
type AuthLoginJSONRequestBody = LoginReq

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (GET /api/auth)
	AuthAuth(ctx echo.Context, params AuthAuthParams) error

	// (POST /api/auth/login)
	AuthLogin(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AuthAuth converts echo context to params.
func (w *ServerInterfaceWrapper) AuthAuth(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params AuthAuthParams

	headers := ctx.Request().Header
	// ------------- Required header parameter "authorization" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("authorization")]; found {
		var Authorization string
		n := len(valueList)
		if n != 1 {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Expected one value for authorization, got %d", n))
		}

		err = runtime.BindStyledParameterWithOptions("simple", "authorization", valueList[0], &Authorization, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: true})
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter authorization: %s", err))
		}

		params.Authorization = Authorization
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Header parameter authorization is required, but not found"))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AuthAuth(ctx, params)
	return err
}

// AuthLogin converts echo context to params.
func (w *ServerInterfaceWrapper) AuthLogin(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AuthLogin(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET(baseURL+"/api/auth", wrapper.AuthAuth)
	router.POST(baseURL+"/api/auth/login", wrapper.AuthLogin)

}

type AuthAuthRequestObject struct {
	Params AuthAuthParams
}

type AuthAuthResponseObject interface {
	VisitAuthAuthResponse(w http.ResponseWriter) error
}

type AuthAuth200JSONResponse AuthRes

func (response AuthAuth200JSONResponse) VisitAuthAuthResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AuthLoginRequestObject struct {
	Body *AuthLoginJSONRequestBody
}

type AuthLoginResponseObject interface {
	VisitAuthLoginResponse(w http.ResponseWriter) error
}

type AuthLogin200JSONResponse LoginRes

func (response AuthLogin200JSONResponse) VisitAuthLoginResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {

	// (GET /api/auth)
	AuthAuth(ctx context.Context, request AuthAuthRequestObject) (AuthAuthResponseObject, error)

	// (POST /api/auth/login)
	AuthLogin(ctx context.Context, request AuthLoginRequestObject) (AuthLoginResponseObject, error)
}

type StrictHandlerFunc = strictecho.StrictEchoHandlerFunc
type StrictMiddlewareFunc = strictecho.StrictEchoMiddlewareFunc

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
}

// AuthAuth operation middleware
func (sh *strictHandler) AuthAuth(ctx echo.Context, params AuthAuthParams) error {
	var request AuthAuthRequestObject

	request.Params = params

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AuthAuth(ctx.Request().Context(), request.(AuthAuthRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AuthAuth")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AuthAuthResponseObject); ok {
		return validResponse.VisitAuthAuthResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// AuthLogin operation middleware
func (sh *strictHandler) AuthLogin(ctx echo.Context) error {
	var request AuthLoginRequestObject

	var body AuthLoginJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AuthLogin(ctx.Request().Context(), request.(AuthLoginRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AuthLogin")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AuthLoginResponseObject); ok {
		return validResponse.VisitAuthLoginResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}
