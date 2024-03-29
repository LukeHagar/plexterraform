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

// GetPlaylistPlaylistsResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetPlaylistPlaylistsResponseBody struct {
	Errors []GetPlaylistErrors `json:"errors,omitempty"`
}

func (o *GetPlaylistPlaylistsResponseBody) GetErrors() []GetPlaylistErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type GetPlaylistMetadata struct {
	Content      *string `json:"content,omitempty"`
	RatingKey    *string `json:"ratingKey,omitempty"`
	Key          *string `json:"key,omitempty"`
	GUID         *string `json:"guid,omitempty"`
	Type         *string `json:"type,omitempty"`
	Title        *string `json:"title,omitempty"`
	Summary      *string `json:"summary,omitempty"`
	Smart        *bool   `json:"smart,omitempty"`
	PlaylistType *string `json:"playlistType,omitempty"`
	Composite    *string `json:"composite,omitempty"`
	Icon         *string `json:"icon,omitempty"`
	Duration     *int    `json:"duration,omitempty"`
	LeafCount    *int    `json:"leafCount,omitempty"`
	AddedAt      *int    `json:"addedAt,omitempty"`
	UpdatedAt    *int    `json:"updatedAt,omitempty"`
}

func (o *GetPlaylistMetadata) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *GetPlaylistMetadata) GetRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.RatingKey
}

func (o *GetPlaylistMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetPlaylistMetadata) GetGUID() *string {
	if o == nil {
		return nil
	}
	return o.GUID
}

func (o *GetPlaylistMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetPlaylistMetadata) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetPlaylistMetadata) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *GetPlaylistMetadata) GetSmart() *bool {
	if o == nil {
		return nil
	}
	return o.Smart
}

func (o *GetPlaylistMetadata) GetPlaylistType() *string {
	if o == nil {
		return nil
	}
	return o.PlaylistType
}

func (o *GetPlaylistMetadata) GetComposite() *string {
	if o == nil {
		return nil
	}
	return o.Composite
}

func (o *GetPlaylistMetadata) GetIcon() *string {
	if o == nil {
		return nil
	}
	return o.Icon
}

func (o *GetPlaylistMetadata) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetPlaylistMetadata) GetLeafCount() *int {
	if o == nil {
		return nil
	}
	return o.LeafCount
}

func (o *GetPlaylistMetadata) GetAddedAt() *int {
	if o == nil {
		return nil
	}
	return o.AddedAt
}

func (o *GetPlaylistMetadata) GetUpdatedAt() *int {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

type GetPlaylistMediaContainer struct {
	Size     *int                  `json:"size,omitempty"`
	Metadata []GetPlaylistMetadata `json:"Metadata,omitempty"`
}

func (o *GetPlaylistMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetPlaylistMediaContainer) GetMetadata() []GetPlaylistMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// GetPlaylistResponseBody - The playlist
type GetPlaylistResponseBody struct {
	MediaContainer *GetPlaylistMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetPlaylistResponseBody) GetMediaContainer() *GetPlaylistMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetPlaylistResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The playlist
	TwoHundredApplicationJSONObject *GetPlaylistResponseBody
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	FourHundredAndOneApplicationJSONObject *GetPlaylistPlaylistsResponseBody
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

func (o *GetPlaylistResponse) GetTwoHundredApplicationJSONObject() *GetPlaylistResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetPlaylistResponse) GetFourHundredAndOneApplicationJSONObject() *GetPlaylistPlaylistsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}
