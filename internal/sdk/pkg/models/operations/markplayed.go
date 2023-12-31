// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type MarkPlayedRequest struct {
	// The media key to mark as played
	Key float64 `queryParam:"style=form,explode=true,name=key"`
}

func (o *MarkPlayedRequest) GetKey() float64 {
	if o == nil {
		return 0.0
	}
	return o.Key
}

type MarkPlayedErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *MarkPlayedErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *MarkPlayedErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *MarkPlayedErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// MarkPlayedResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type MarkPlayedResponseBody struct {
	Errors []MarkPlayedErrors `json:"errors,omitempty"`
}

func (o *MarkPlayedResponseBody) GetErrors() []MarkPlayedErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type MarkPlayedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	Object *MarkPlayedResponseBody
}

func (o *MarkPlayedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *MarkPlayedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *MarkPlayedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *MarkPlayedResponse) GetObject() *MarkPlayedResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
