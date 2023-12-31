// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetPlaylistRequest struct {
	// the ID of the playlist
	PlaylistID float64 `pathParam:"style=simple,explode=false,name=playlistID"`
}

func (o *GetPlaylistRequest) GetPlaylistID() float64 {
	if o == nil {
		return 0.0
	}
	return o.PlaylistID
}

type GetPlaylistErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetPlaylistErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetPlaylistErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPlaylistErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetPlaylistResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetPlaylistResponseBody struct {
	Errors []GetPlaylistErrors `json:"errors,omitempty"`
}

func (o *GetPlaylistResponseBody) GetErrors() []GetPlaylistErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type GetPlaylistResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	Object *GetPlaylistResponseBody
}

func (o *GetPlaylistResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPlaylistResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPlaylistResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPlaylistResponse) GetObject() *GetPlaylistResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
