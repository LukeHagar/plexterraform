// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetSessionHistoryErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetSessionHistoryErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetSessionHistoryErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetSessionHistoryErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetSessionHistoryResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetSessionHistoryResponseBody struct {
	Errors []GetSessionHistoryErrors `json:"errors,omitempty"`
}

func (o *GetSessionHistoryResponseBody) GetErrors() []GetSessionHistoryErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type GetSessionHistoryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	Object *GetSessionHistoryResponseBody
}

func (o *GetSessionHistoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetSessionHistoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetSessionHistoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetSessionHistoryResponse) GetObject() *GetSessionHistoryResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
