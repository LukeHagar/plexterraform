// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetFileHashRequest struct {
	// This is the path to the local file, must be prefixed by `file://`
	URL string `queryParam:"style=form,explode=true,name=url"`
	// Item type
	Type *float64 `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetFileHashRequest) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *GetFileHashRequest) GetType() *float64 {
	if o == nil {
		return nil
	}
	return o.Type
}

type GetFileHashErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetFileHashErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetFileHashErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetFileHashErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetFileHashResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetFileHashResponseBody struct {
	Errors []GetFileHashErrors `json:"errors,omitempty"`
}

func (o *GetFileHashResponseBody) GetErrors() []GetFileHashErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type GetFileHashResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	Object *GetFileHashResponseBody
}

func (o *GetFileHashResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFileHashResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFileHashResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetFileHashResponse) GetObject() *GetFileHashResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}