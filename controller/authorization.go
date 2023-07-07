package controller

import (
	"diana/model/request"
	"diana/model/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/go-querystring/query"
	"net/http"
	"reflect"
	"strings"
)

func requestValidate(ctx *gin.Context) (*request.AuthorizationRequest, *response.OAuthError) {
	req := &request.AuthorizationRequest{}
	if err := ctx.ShouldBindQuery(&req); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		var errorFields []string
		if ok {
			for _, err := range errs {
				fieldName := err.Field()
				field, ok := reflect.TypeOf(req).FieldByName(fieldName)
				if ok {
					jsonName := field.Tag.Get("json")
					errorFields = append(errorFields, jsonName)
				}
			}
			msg := fmt.Sprintf("The request is missing %s", strings.Join(errorFields, ","))
			oauthError := response.NewOAuthError(response.InvalidRequest, msg, req.RedirectUri, req.State)
			return nil, &oauthError
		}
	}
	return req, nil
}

// Auth is Authorization Endpoint.
func Auth(ctx *gin.Context) {
	req, err := requestValidate(ctx)
	if err != nil {
		queryString, _ := query.Values(err)
		ctx.Redirect(http.StatusFound, req.RedirectUri+"?"+queryString.Encode())
		return
	}
}
