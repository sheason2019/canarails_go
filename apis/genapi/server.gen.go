// Package genapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package genapi

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"github.com/oapi-codegen/runtime"
	strictecho "github.com/oapi-codegen/runtime/strictmiddleware/echo"
)

const (
	BasicAuthScopes = "BasicAuth.Scopes"
)

// App defines model for App.
type App struct {
	Description string   `json:"description"`
	Hostnames   []string `json:"hostnames"`
	Id          int32    `json:"id"`
	Title       string   `json:"title"`
}

// AuthRes defines model for AuthRes.
type AuthRes struct {
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

// AppsDeleteJSONBody defines parameters for AppsDelete.
type AppsDeleteJSONBody struct {
	Id int32 `json:"id"`
}

// AppsDeleteParams defines parameters for AppsDelete.
type AppsDeleteParams struct {
	Authorization string `json:"authorization"`
}

// AppsCreateParams defines parameters for AppsCreate.
type AppsCreateParams struct {
	Authorization string `json:"authorization"`
}

// AppsPutParams defines parameters for AppsPut.
type AppsPutParams struct {
	Authorization string `json:"authorization"`
}

// AuthAuthParams defines parameters for AuthAuth.
type AuthAuthParams struct {
	Authorization string `json:"authorization"`
}

// AppsDeleteJSONRequestBody defines body for AppsDelete for application/json ContentType.
type AppsDeleteJSONRequestBody AppsDeleteJSONBody

// AppsCreateJSONRequestBody defines body for AppsCreate for application/json ContentType.
type AppsCreateJSONRequestBody = App

// AppsPutJSONRequestBody defines body for AppsPut for application/json ContentType.
type AppsPutJSONRequestBody = App

// AuthLoginJSONRequestBody defines body for AuthLogin for application/json ContentType.
type AuthLoginJSONRequestBody = LoginReq

// ServerInterface represents all server handlers.
type ServerInterface interface {

	// (DELETE /api/app)
	AppsDelete(ctx echo.Context, params AppsDeleteParams) error

	// (GET /api/app)
	AppsList(ctx echo.Context) error

	// (POST /api/app)
	AppsCreate(ctx echo.Context, params AppsCreateParams) error

	// (PUT /api/app)
	AppsPut(ctx echo.Context, params AppsPutParams) error

	// (GET /api/app/{id})
	AppsFindById(ctx echo.Context, id int32) error

	// (GET /api/auth)
	AuthAuth(ctx echo.Context, params AuthAuthParams) error

	// (POST /api/auth/login)
	AuthLogin(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AppsDelete converts echo context to params.
func (w *ServerInterfaceWrapper) AppsDelete(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params AppsDeleteParams

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
	err = w.Handler.AppsDelete(ctx, params)
	return err
}

// AppsList converts echo context to params.
func (w *ServerInterfaceWrapper) AppsList(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AppsList(ctx)
	return err
}

// AppsCreate converts echo context to params.
func (w *ServerInterfaceWrapper) AppsCreate(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params AppsCreateParams

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
	err = w.Handler.AppsCreate(ctx, params)
	return err
}

// AppsPut converts echo context to params.
func (w *ServerInterfaceWrapper) AppsPut(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params AppsPutParams

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
	err = w.Handler.AppsPut(ctx, params)
	return err
}

// AppsFindById converts echo context to params.
func (w *ServerInterfaceWrapper) AppsFindById(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "id" -------------
	var id int32

	err = runtime.BindStyledParameterWithOptions("simple", "id", ctx.Param("id"), &id, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter id: %s", err))
	}

	// Invoke the callback with all the unmarshaled arguments
	err = w.Handler.AppsFindById(ctx, id)
	return err
}

// AuthAuth converts echo context to params.
func (w *ServerInterfaceWrapper) AuthAuth(ctx echo.Context) error {
	var err error

	ctx.Set(BasicAuthScopes, []string{})

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

	router.DELETE(baseURL+"/api/app", wrapper.AppsDelete)
	router.GET(baseURL+"/api/app", wrapper.AppsList)
	router.POST(baseURL+"/api/app", wrapper.AppsCreate)
	router.PUT(baseURL+"/api/app", wrapper.AppsPut)
	router.GET(baseURL+"/api/app/:id", wrapper.AppsFindById)
	router.GET(baseURL+"/api/auth", wrapper.AuthAuth)
	router.POST(baseURL+"/api/auth/login", wrapper.AuthLogin)

}

type AppsDeleteRequestObject struct {
	Params AppsDeleteParams
	Body   *AppsDeleteJSONRequestBody
}

type AppsDeleteResponseObject interface {
	VisitAppsDeleteResponse(w http.ResponseWriter) error
}

type AppsDelete200JSONResponse int32

func (response AppsDelete200JSONResponse) VisitAppsDeleteResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AppsListRequestObject struct {
}

type AppsListResponseObject interface {
	VisitAppsListResponse(w http.ResponseWriter) error
}

type AppsList200JSONResponse []App

func (response AppsList200JSONResponse) VisitAppsListResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AppsCreateRequestObject struct {
	Params AppsCreateParams
	Body   *AppsCreateJSONRequestBody
}

type AppsCreateResponseObject interface {
	VisitAppsCreateResponse(w http.ResponseWriter) error
}

type AppsCreate200JSONResponse int32

func (response AppsCreate200JSONResponse) VisitAppsCreateResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AppsPutRequestObject struct {
	Params AppsPutParams
	Body   *AppsPutJSONRequestBody
}

type AppsPutResponseObject interface {
	VisitAppsPutResponse(w http.ResponseWriter) error
}

type AppsPut200JSONResponse int32

func (response AppsPut200JSONResponse) VisitAppsPutResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type AppsFindByIdRequestObject struct {
	Id int32 `json:"id"`
}

type AppsFindByIdResponseObject interface {
	VisitAppsFindByIdResponse(w http.ResponseWriter) error
}

type AppsFindById200JSONResponse App

func (response AppsFindById200JSONResponse) VisitAppsFindByIdResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
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

	// (DELETE /api/app)
	AppsDelete(ctx context.Context, request AppsDeleteRequestObject) (AppsDeleteResponseObject, error)

	// (GET /api/app)
	AppsList(ctx context.Context, request AppsListRequestObject) (AppsListResponseObject, error)

	// (POST /api/app)
	AppsCreate(ctx context.Context, request AppsCreateRequestObject) (AppsCreateResponseObject, error)

	// (PUT /api/app)
	AppsPut(ctx context.Context, request AppsPutRequestObject) (AppsPutResponseObject, error)

	// (GET /api/app/{id})
	AppsFindById(ctx context.Context, request AppsFindByIdRequestObject) (AppsFindByIdResponseObject, error)

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

// AppsDelete operation middleware
func (sh *strictHandler) AppsDelete(ctx echo.Context, params AppsDeleteParams) error {
	var request AppsDeleteRequestObject

	request.Params = params

	var body AppsDeleteJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AppsDelete(ctx.Request().Context(), request.(AppsDeleteRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AppsDelete")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AppsDeleteResponseObject); ok {
		return validResponse.VisitAppsDeleteResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// AppsList operation middleware
func (sh *strictHandler) AppsList(ctx echo.Context) error {
	var request AppsListRequestObject

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AppsList(ctx.Request().Context(), request.(AppsListRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AppsList")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AppsListResponseObject); ok {
		return validResponse.VisitAppsListResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// AppsCreate operation middleware
func (sh *strictHandler) AppsCreate(ctx echo.Context, params AppsCreateParams) error {
	var request AppsCreateRequestObject

	request.Params = params

	var body AppsCreateJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AppsCreate(ctx.Request().Context(), request.(AppsCreateRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AppsCreate")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AppsCreateResponseObject); ok {
		return validResponse.VisitAppsCreateResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// AppsPut operation middleware
func (sh *strictHandler) AppsPut(ctx echo.Context, params AppsPutParams) error {
	var request AppsPutRequestObject

	request.Params = params

	var body AppsPutJSONRequestBody
	if err := ctx.Bind(&body); err != nil {
		return err
	}
	request.Body = &body

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AppsPut(ctx.Request().Context(), request.(AppsPutRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AppsPut")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AppsPutResponseObject); ok {
		return validResponse.VisitAppsPutResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
}

// AppsFindById operation middleware
func (sh *strictHandler) AppsFindById(ctx echo.Context, id int32) error {
	var request AppsFindByIdRequestObject

	request.Id = id

	handler := func(ctx echo.Context, request interface{}) (interface{}, error) {
		return sh.ssi.AppsFindById(ctx.Request().Context(), request.(AppsFindByIdRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "AppsFindById")
	}

	response, err := handler(ctx, request)

	if err != nil {
		return err
	} else if validResponse, ok := response.(AppsFindByIdResponseObject); ok {
		return validResponse.VisitAppsFindByIdResponse(ctx.Response())
	} else if response != nil {
		return fmt.Errorf("unexpected response type: %T", response)
	}
	return nil
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

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"H4sIAAAAAAAC/+xWwW7bMAz9lYLbYQOMOmhvvrUdBhToYeh2K3JQLTZSl0iqRG3IAv37QNm1m8Ze0mUZ",
	"imGnODJNPb73RHEFtV04a9BQgGoFoVa4EPnxzDn+cd469KQxL0oMtdeOtDX8l5YOoYJAXpsZpAKUDWTE",
	"ognWhIswGNYuCO/Fkv9ryWF31i8EQQXa0OkJdGHaEM7Q5w81zXEgZSrA40PUHiVUN5zvMbZYw/wU4bTb",
	"wN7eY02c/yySum7Qrxe+M8IY0HP+HUF24UNoruxMm2t82ITjRAjfrZeD7O4OoYss+oy/ADLAC9mvaLZv",
	"1IRtpk4FBKyj17T8zNZrkp6LoGuWorMkf3PLqz3nishB4gza3NkMoTEHvMsP76GAb+hDtipMjifHEy7F",
	"OjTCaajgNC9x5aTytqVwuhSN7SXOkTKDXK1g81xKqPhUhA/NO/7SiwUS+gDVzQo0b6RQSPRQQKMAiEjK",
	"ev1DtP7rWSEfsWhP3BCD0yYYA51bueSI2hpCQ/wonJvrOict70NzGvtUv+XdTXsO6JWeV5AXgrMmNJud",
	"TCYvgrodVlo/whV8UXjUMnOkRDgKsa4RJcrjNUNlTZ5Y6WbKlJKYsVos5IU15O18jh6mqYAZ0rDeVzoQ",
	"7Flm1wvferyDCt6UfeMt265bcsvd6I4vrn+8QmfDSIkXHsXrtvRW3v5la7o4otunSP9Fe5WipaK7T8qV",
	"lomxjXaYj9rI8+WlHNGSb6heyTw3jMu3nYDpnhTvpOofaVkdie0sMExgJJUV+SsH4VC8tYPnwewZSY1S",
	"W855vsuDw/ANEUnlERAO0wq6OfcA/WCHfcM+jt3gtX+5erTcs6BU9G/WDJ+m6WcAAAD//+2WhJCRDQAA",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
