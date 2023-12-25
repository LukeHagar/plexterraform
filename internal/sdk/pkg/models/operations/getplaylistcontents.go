// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetPlaylistContentsRequest struct {
	// the ID of the playlist
	PlaylistID float64 `pathParam:"style=simple,explode=false,name=playlistID"`
	// the metadata type of the item to return
	Type float64 `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetPlaylistContentsRequest) GetPlaylistID() float64 {
	if o == nil {
		return 0.0
	}
	return o.PlaylistID
}

func (o *GetPlaylistContentsRequest) GetType() float64 {
	if o == nil {
		return 0.0
	}
	return o.Type
}

type GetPlaylistContentsErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetPlaylistContentsErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetPlaylistContentsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPlaylistContentsErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetPlaylistContentsResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetPlaylistContentsResponseBody struct {
	Errors []GetPlaylistContentsErrors `json:"errors,omitempty"`
}

func (o *GetPlaylistContentsResponseBody) GetErrors() []GetPlaylistContentsErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type GetPlaylistContentsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	Object *GetPlaylistContentsResponseBody
}

func (o *GetPlaylistContentsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPlaylistContentsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPlaylistContentsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPlaylistContentsResponse) GetObject() *GetPlaylistContentsResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
