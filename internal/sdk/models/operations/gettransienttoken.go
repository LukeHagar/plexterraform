// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTransientTokenQueryParamType - `delegation` - This is the only supported `type` parameter.
type GetTransientTokenQueryParamType string

const (
	GetTransientTokenQueryParamTypeDelegation GetTransientTokenQueryParamType = "delegation"
)

func (e GetTransientTokenQueryParamType) ToPointer() *GetTransientTokenQueryParamType {
	return &e
}
func (e *GetTransientTokenQueryParamType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "delegation":
		*e = GetTransientTokenQueryParamType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetTransientTokenQueryParamType: %v", v)
	}
}

// Scope - `all` - This is the only supported `scope` parameter.
type Scope string

const (
	ScopeAll Scope = "all"
)

func (e Scope) ToPointer() *Scope {
	return &e
}
func (e *Scope) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "all":
		*e = Scope(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Scope: %v", v)
	}
}

type GetTransientTokenRequest struct {
	// `delegation` - This is the only supported `type` parameter.
	Type GetTransientTokenQueryParamType `queryParam:"style=form,explode=true,name=type"`
	// `all` - This is the only supported `scope` parameter.
	Scope Scope `queryParam:"style=form,explode=true,name=scope"`
}

func (o *GetTransientTokenRequest) GetType() GetTransientTokenQueryParamType {
	if o == nil {
		return GetTransientTokenQueryParamType("")
	}
	return o.Type
}

func (o *GetTransientTokenRequest) GetScope() Scope {
	if o == nil {
		return Scope("")
	}
	return o.Scope
}

type GetTransientTokenErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetTransientTokenErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetTransientTokenErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetTransientTokenErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetTransientTokenResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetTransientTokenResponseBody struct {
	Errors []GetTransientTokenErrors `json:"errors,omitempty"`
}

func (o *GetTransientTokenResponseBody) GetErrors() []GetTransientTokenErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type GetTransientTokenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	Object *GetTransientTokenResponseBody
}

func (o *GetTransientTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransientTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransientTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTransientTokenResponse) GetObject() *GetTransientTokenResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
