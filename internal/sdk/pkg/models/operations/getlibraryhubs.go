// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/LukeHagar/terraform-provider-PlexAPI/internal/sdk/pkg/types"
	"github.com/LukeHagar/terraform-provider-PlexAPI/internal/sdk/pkg/utils"
	"net/http"
)

// QueryParamOnlyTransient - Only return hubs which are "transient", meaning those which are prone to changing after media playback or addition (e.g. On Deck, or Recently Added).
type QueryParamOnlyTransient int64

const (
	QueryParamOnlyTransientZero QueryParamOnlyTransient = 0
	QueryParamOnlyTransientOne  QueryParamOnlyTransient = 1
)

func (e QueryParamOnlyTransient) ToPointer() *QueryParamOnlyTransient {
	return &e
}

func (e *QueryParamOnlyTransient) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = QueryParamOnlyTransient(v)
		return nil
	default:
		return fmt.Errorf("invalid value for QueryParamOnlyTransient: %v", v)
	}
}

type GetLibraryHubsRequest struct {
	// the Id of the library to query
	SectionID float64 `pathParam:"style=simple,explode=false,name=sectionId"`
	// The number of items to return with each hub.
	Count *float64 `queryParam:"style=form,explode=true,name=count"`
	// Only return hubs which are "transient", meaning those which are prone to changing after media playback or addition (e.g. On Deck, or Recently Added).
	OnlyTransient *QueryParamOnlyTransient `queryParam:"style=form,explode=true,name=onlyTransient"`
}

func (o *GetLibraryHubsRequest) GetSectionID() float64 {
	if o == nil {
		return 0.0
	}
	return o.SectionID
}

func (o *GetLibraryHubsRequest) GetCount() *float64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *GetLibraryHubsRequest) GetOnlyTransient() *QueryParamOnlyTransient {
	if o == nil {
		return nil
	}
	return o.OnlyTransient
}

type GetLibraryHubsErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetLibraryHubsErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetLibraryHubsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetLibraryHubsErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetLibraryHubsHubsResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetLibraryHubsHubsResponseBody struct {
	Errors []GetLibraryHubsErrors `json:"errors,omitempty"`
}

func (o *GetLibraryHubsHubsResponseBody) GetErrors() []GetLibraryHubsErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type GetLibraryHubsPart struct {
	ID                    *int    `json:"id,omitempty"`
	Key                   *string `json:"key,omitempty"`
	Duration              *int    `json:"duration,omitempty"`
	File                  *string `json:"file,omitempty"`
	Size                  *int    `json:"size,omitempty"`
	AudioProfile          *string `json:"audioProfile,omitempty"`
	Container             *string `json:"container,omitempty"`
	Has64bitOffsets       *bool   `json:"has64bitOffsets,omitempty"`
	OptimizedForStreaming *bool   `json:"optimizedForStreaming,omitempty"`
	VideoProfile          *string `json:"videoProfile,omitempty"`
}

func (o *GetLibraryHubsPart) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetLibraryHubsPart) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetLibraryHubsPart) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetLibraryHubsPart) GetFile() *string {
	if o == nil {
		return nil
	}
	return o.File
}

func (o *GetLibraryHubsPart) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetLibraryHubsPart) GetAudioProfile() *string {
	if o == nil {
		return nil
	}
	return o.AudioProfile
}

func (o *GetLibraryHubsPart) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *GetLibraryHubsPart) GetHas64bitOffsets() *bool {
	if o == nil {
		return nil
	}
	return o.Has64bitOffsets
}

func (o *GetLibraryHubsPart) GetOptimizedForStreaming() *bool {
	if o == nil {
		return nil
	}
	return o.OptimizedForStreaming
}

func (o *GetLibraryHubsPart) GetVideoProfile() *string {
	if o == nil {
		return nil
	}
	return o.VideoProfile
}

type GetLibraryHubsMedia struct {
	ID                    *int                 `json:"id,omitempty"`
	Duration              *int                 `json:"duration,omitempty"`
	Bitrate               *int                 `json:"bitrate,omitempty"`
	Width                 *int                 `json:"width,omitempty"`
	Height                *int                 `json:"height,omitempty"`
	AspectRatio           *float64             `json:"aspectRatio,omitempty"`
	AudioChannels         *int                 `json:"audioChannels,omitempty"`
	AudioCodec            *string              `json:"audioCodec,omitempty"`
	VideoCodec            *string              `json:"videoCodec,omitempty"`
	VideoResolution       *string              `json:"videoResolution,omitempty"`
	Container             *string              `json:"container,omitempty"`
	VideoFrameRate        *string              `json:"videoFrameRate,omitempty"`
	OptimizedForStreaming *int                 `json:"optimizedForStreaming,omitempty"`
	AudioProfile          *string              `json:"audioProfile,omitempty"`
	Has64bitOffsets       *bool                `json:"has64bitOffsets,omitempty"`
	VideoProfile          *string              `json:"videoProfile,omitempty"`
	Part                  []GetLibraryHubsPart `json:"Part,omitempty"`
}

func (o *GetLibraryHubsMedia) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetLibraryHubsMedia) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetLibraryHubsMedia) GetBitrate() *int {
	if o == nil {
		return nil
	}
	return o.Bitrate
}

func (o *GetLibraryHubsMedia) GetWidth() *int {
	if o == nil {
		return nil
	}
	return o.Width
}

func (o *GetLibraryHubsMedia) GetHeight() *int {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *GetLibraryHubsMedia) GetAspectRatio() *float64 {
	if o == nil {
		return nil
	}
	return o.AspectRatio
}

func (o *GetLibraryHubsMedia) GetAudioChannels() *int {
	if o == nil {
		return nil
	}
	return o.AudioChannels
}

func (o *GetLibraryHubsMedia) GetAudioCodec() *string {
	if o == nil {
		return nil
	}
	return o.AudioCodec
}

func (o *GetLibraryHubsMedia) GetVideoCodec() *string {
	if o == nil {
		return nil
	}
	return o.VideoCodec
}

func (o *GetLibraryHubsMedia) GetVideoResolution() *string {
	if o == nil {
		return nil
	}
	return o.VideoResolution
}

func (o *GetLibraryHubsMedia) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *GetLibraryHubsMedia) GetVideoFrameRate() *string {
	if o == nil {
		return nil
	}
	return o.VideoFrameRate
}

func (o *GetLibraryHubsMedia) GetOptimizedForStreaming() *int {
	if o == nil {
		return nil
	}
	return o.OptimizedForStreaming
}

func (o *GetLibraryHubsMedia) GetAudioProfile() *string {
	if o == nil {
		return nil
	}
	return o.AudioProfile
}

func (o *GetLibraryHubsMedia) GetHas64bitOffsets() *bool {
	if o == nil {
		return nil
	}
	return o.Has64bitOffsets
}

func (o *GetLibraryHubsMedia) GetVideoProfile() *string {
	if o == nil {
		return nil
	}
	return o.VideoProfile
}

func (o *GetLibraryHubsMedia) GetPart() []GetLibraryHubsPart {
	if o == nil {
		return nil
	}
	return o.Part
}

type GetLibraryHubsGenre struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetLibraryHubsGenre) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetLibraryHubsCountry struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetLibraryHubsCountry) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetLibraryHubsDirector struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetLibraryHubsDirector) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetLibraryHubsRole struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetLibraryHubsRole) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetLibraryHubsWriter struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetLibraryHubsWriter) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetLibraryHubsMetadata struct {
	RatingKey             *string                  `json:"ratingKey,omitempty"`
	Key                   *string                  `json:"key,omitempty"`
	GUID                  *string                  `json:"guid,omitempty"`
	Studio                *string                  `json:"studio,omitempty"`
	Type                  *string                  `json:"type,omitempty"`
	Title                 *string                  `json:"title,omitempty"`
	LibrarySectionTitle   *string                  `json:"librarySectionTitle,omitempty"`
	LibrarySectionID      *int                     `json:"librarySectionID,omitempty"`
	LibrarySectionKey     *string                  `json:"librarySectionKey,omitempty"`
	ContentRating         *string                  `json:"contentRating,omitempty"`
	Summary               *string                  `json:"summary,omitempty"`
	Rating                *float64                 `json:"rating,omitempty"`
	AudienceRating        *float64                 `json:"audienceRating,omitempty"`
	ViewCount             *int                     `json:"viewCount,omitempty"`
	LastViewedAt          *int                     `json:"lastViewedAt,omitempty"`
	Year                  *int                     `json:"year,omitempty"`
	Tagline               *string                  `json:"tagline,omitempty"`
	Thumb                 *string                  `json:"thumb,omitempty"`
	Art                   *string                  `json:"art,omitempty"`
	Duration              *int                     `json:"duration,omitempty"`
	OriginallyAvailableAt *types.Date              `json:"originallyAvailableAt,omitempty"`
	AddedAt               *int                     `json:"addedAt,omitempty"`
	UpdatedAt             *int                     `json:"updatedAt,omitempty"`
	AudienceRatingImage   *string                  `json:"audienceRatingImage,omitempty"`
	PrimaryExtraKey       *string                  `json:"primaryExtraKey,omitempty"`
	RatingImage           *string                  `json:"ratingImage,omitempty"`
	Media                 []GetLibraryHubsMedia    `json:"Media,omitempty"`
	Genre                 []GetLibraryHubsGenre    `json:"Genre,omitempty"`
	Country               []GetLibraryHubsCountry  `json:"Country,omitempty"`
	Director              []GetLibraryHubsDirector `json:"Director,omitempty"`
	Role                  []GetLibraryHubsRole     `json:"Role,omitempty"`
	Writer                []GetLibraryHubsWriter   `json:"Writer,omitempty"`
	SkipCount             *int                     `json:"skipCount,omitempty"`
	ChapterSource         *string                  `json:"chapterSource,omitempty"`
}

func (g GetLibraryHubsMetadata) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetLibraryHubsMetadata) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetLibraryHubsMetadata) GetRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.RatingKey
}

func (o *GetLibraryHubsMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetLibraryHubsMetadata) GetGUID() *string {
	if o == nil {
		return nil
	}
	return o.GUID
}

func (o *GetLibraryHubsMetadata) GetStudio() *string {
	if o == nil {
		return nil
	}
	return o.Studio
}

func (o *GetLibraryHubsMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetLibraryHubsMetadata) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetLibraryHubsMetadata) GetLibrarySectionTitle() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionTitle
}

func (o *GetLibraryHubsMetadata) GetLibrarySectionID() *int {
	if o == nil {
		return nil
	}
	return o.LibrarySectionID
}

func (o *GetLibraryHubsMetadata) GetLibrarySectionKey() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionKey
}

func (o *GetLibraryHubsMetadata) GetContentRating() *string {
	if o == nil {
		return nil
	}
	return o.ContentRating
}

func (o *GetLibraryHubsMetadata) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *GetLibraryHubsMetadata) GetRating() *float64 {
	if o == nil {
		return nil
	}
	return o.Rating
}

func (o *GetLibraryHubsMetadata) GetAudienceRating() *float64 {
	if o == nil {
		return nil
	}
	return o.AudienceRating
}

func (o *GetLibraryHubsMetadata) GetViewCount() *int {
	if o == nil {
		return nil
	}
	return o.ViewCount
}

func (o *GetLibraryHubsMetadata) GetLastViewedAt() *int {
	if o == nil {
		return nil
	}
	return o.LastViewedAt
}

func (o *GetLibraryHubsMetadata) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

func (o *GetLibraryHubsMetadata) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}

func (o *GetLibraryHubsMetadata) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetLibraryHubsMetadata) GetArt() *string {
	if o == nil {
		return nil
	}
	return o.Art
}

func (o *GetLibraryHubsMetadata) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetLibraryHubsMetadata) GetOriginallyAvailableAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.OriginallyAvailableAt
}

func (o *GetLibraryHubsMetadata) GetAddedAt() *int {
	if o == nil {
		return nil
	}
	return o.AddedAt
}

func (o *GetLibraryHubsMetadata) GetUpdatedAt() *int {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetLibraryHubsMetadata) GetAudienceRatingImage() *string {
	if o == nil {
		return nil
	}
	return o.AudienceRatingImage
}

func (o *GetLibraryHubsMetadata) GetPrimaryExtraKey() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryExtraKey
}

func (o *GetLibraryHubsMetadata) GetRatingImage() *string {
	if o == nil {
		return nil
	}
	return o.RatingImage
}

func (o *GetLibraryHubsMetadata) GetMedia() []GetLibraryHubsMedia {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *GetLibraryHubsMetadata) GetGenre() []GetLibraryHubsGenre {
	if o == nil {
		return nil
	}
	return o.Genre
}

func (o *GetLibraryHubsMetadata) GetCountry() []GetLibraryHubsCountry {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *GetLibraryHubsMetadata) GetDirector() []GetLibraryHubsDirector {
	if o == nil {
		return nil
	}
	return o.Director
}

func (o *GetLibraryHubsMetadata) GetRole() []GetLibraryHubsRole {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *GetLibraryHubsMetadata) GetWriter() []GetLibraryHubsWriter {
	if o == nil {
		return nil
	}
	return o.Writer
}

func (o *GetLibraryHubsMetadata) GetSkipCount() *int {
	if o == nil {
		return nil
	}
	return o.SkipCount
}

func (o *GetLibraryHubsMetadata) GetChapterSource() *string {
	if o == nil {
		return nil
	}
	return o.ChapterSource
}

type GetLibraryHubsHub struct {
	Key           *string                  `json:"key,omitempty"`
	Title         *string                  `json:"title,omitempty"`
	Type          *string                  `json:"type,omitempty"`
	HubIdentifier *string                  `json:"hubIdentifier,omitempty"`
	Context       *string                  `json:"context,omitempty"`
	Size          *int                     `json:"size,omitempty"`
	More          *bool                    `json:"more,omitempty"`
	Style         *string                  `json:"style,omitempty"`
	HubKey        *string                  `json:"hubKey,omitempty"`
	Metadata      []GetLibraryHubsMetadata `json:"Metadata,omitempty"`
	Promoted      *bool                    `json:"promoted,omitempty"`
	Random        *bool                    `json:"random,omitempty"`
}

func (o *GetLibraryHubsHub) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetLibraryHubsHub) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetLibraryHubsHub) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetLibraryHubsHub) GetHubIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.HubIdentifier
}

func (o *GetLibraryHubsHub) GetContext() *string {
	if o == nil {
		return nil
	}
	return o.Context
}

func (o *GetLibraryHubsHub) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetLibraryHubsHub) GetMore() *bool {
	if o == nil {
		return nil
	}
	return o.More
}

func (o *GetLibraryHubsHub) GetStyle() *string {
	if o == nil {
		return nil
	}
	return o.Style
}

func (o *GetLibraryHubsHub) GetHubKey() *string {
	if o == nil {
		return nil
	}
	return o.HubKey
}

func (o *GetLibraryHubsHub) GetMetadata() []GetLibraryHubsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *GetLibraryHubsHub) GetPromoted() *bool {
	if o == nil {
		return nil
	}
	return o.Promoted
}

func (o *GetLibraryHubsHub) GetRandom() *bool {
	if o == nil {
		return nil
	}
	return o.Random
}

type GetLibraryHubsMediaContainer struct {
	Size                *int                `json:"size,omitempty"`
	AllowSync           *bool               `json:"allowSync,omitempty"`
	Identifier          *string             `json:"identifier,omitempty"`
	LibrarySectionID    *int                `json:"librarySectionID,omitempty"`
	LibrarySectionTitle *string             `json:"librarySectionTitle,omitempty"`
	LibrarySectionUUID  *string             `json:"librarySectionUUID,omitempty"`
	Hub                 []GetLibraryHubsHub `json:"Hub,omitempty"`
}

func (o *GetLibraryHubsMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetLibraryHubsMediaContainer) GetAllowSync() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSync
}

func (o *GetLibraryHubsMediaContainer) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *GetLibraryHubsMediaContainer) GetLibrarySectionID() *int {
	if o == nil {
		return nil
	}
	return o.LibrarySectionID
}

func (o *GetLibraryHubsMediaContainer) GetLibrarySectionTitle() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionTitle
}

func (o *GetLibraryHubsMediaContainer) GetLibrarySectionUUID() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionUUID
}

func (o *GetLibraryHubsMediaContainer) GetHub() []GetLibraryHubsHub {
	if o == nil {
		return nil
	}
	return o.Hub
}

// GetLibraryHubsResponseBody - The hubs specific to the library
type GetLibraryHubsResponseBody struct {
	MediaContainer *GetLibraryHubsMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetLibraryHubsResponseBody) GetMediaContainer() *GetLibraryHubsMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetLibraryHubsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The hubs specific to the library
	TwoHundredApplicationJSONObject *GetLibraryHubsResponseBody
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	FourHundredAndOneApplicationJSONObject *GetLibraryHubsHubsResponseBody
}

func (o *GetLibraryHubsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLibraryHubsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLibraryHubsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLibraryHubsResponse) GetTwoHundredApplicationJSONObject() *GetLibraryHubsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetLibraryHubsResponse) GetFourHundredAndOneApplicationJSONObject() *GetLibraryHubsHubsResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}
