// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetMetadataChildrenRequest struct {
	// the id of the library item to return the children of.
	RatingKey float64 `pathParam:"style=simple,explode=false,name=ratingKey"`
}

func (o *GetMetadataChildrenRequest) GetRatingKey() float64 {
	if o == nil {
		return 0.0
	}
	return o.RatingKey
}

type GetMetadataChildrenErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetMetadataChildrenErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetMetadataChildrenErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetMetadataChildrenErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetMetadataChildrenResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetMetadataChildrenResponseBody struct {
	Errors []GetMetadataChildrenErrors `json:"errors,omitempty"`
}

func (o *GetMetadataChildrenResponseBody) GetErrors() []GetMetadataChildrenErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type GetMetadataChildrenResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	Object *GetMetadataChildrenResponseBody
}

func (o *GetMetadataChildrenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMetadataChildrenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMetadataChildrenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMetadataChildrenResponse) GetObject() *GetMetadataChildrenResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
