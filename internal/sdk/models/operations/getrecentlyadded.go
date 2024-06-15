// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/LukeHagar/terraform-provider-PlexAPI/internal/sdk/internal/utils"
	"net/http"
	"time"
)

type GetRecentlyAddedErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetRecentlyAddedErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetRecentlyAddedErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetRecentlyAddedErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetRecentlyAddedLibraryResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetRecentlyAddedLibraryResponseBody struct {
	Errors []GetRecentlyAddedErrors `json:"errors,omitempty"`
}

func (o *GetRecentlyAddedLibraryResponseBody) GetErrors() []GetRecentlyAddedErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type Part struct {
	ID                    *float64 `json:"id,omitempty"`
	Key                   *string  `json:"key,omitempty"`
	Duration              *float64 `json:"duration,omitempty"`
	File                  *string  `json:"file,omitempty"`
	Size                  *float64 `json:"size,omitempty"`
	Container             *string  `json:"container,omitempty"`
	Has64bitOffsets       *bool    `json:"has64bitOffsets,omitempty"`
	HasThumbnail          *float64 `json:"hasThumbnail,omitempty"`
	OptimizedForStreaming *bool    `json:"optimizedForStreaming,omitempty"`
	VideoProfile          *string  `json:"videoProfile,omitempty"`
}

func (o *Part) GetID() *float64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Part) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *Part) GetDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *Part) GetFile() *string {
	if o == nil {
		return nil
	}
	return o.File
}

func (o *Part) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *Part) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *Part) GetHas64bitOffsets() *bool {
	if o == nil {
		return nil
	}
	return o.Has64bitOffsets
}

func (o *Part) GetHasThumbnail() *float64 {
	if o == nil {
		return nil
	}
	return o.HasThumbnail
}

func (o *Part) GetOptimizedForStreaming() *bool {
	if o == nil {
		return nil
	}
	return o.OptimizedForStreaming
}

func (o *Part) GetVideoProfile() *string {
	if o == nil {
		return nil
	}
	return o.VideoProfile
}

type Media struct {
	ID                    *float64 `json:"id,omitempty"`
	Duration              *float64 `json:"duration,omitempty"`
	Bitrate               *float64 `json:"bitrate,omitempty"`
	Width                 *float64 `json:"width,omitempty"`
	Height                *float64 `json:"height,omitempty"`
	AspectRatio           *float64 `json:"aspectRatio,omitempty"`
	AudioChannels         *float64 `json:"audioChannels,omitempty"`
	AudioCodec            *string  `json:"audioCodec,omitempty"`
	VideoCodec            *string  `json:"videoCodec,omitempty"`
	VideoResolution       *float64 `json:"videoResolution,omitempty"`
	Container             *string  `json:"container,omitempty"`
	VideoFrameRate        *string  `json:"videoFrameRate,omitempty"`
	OptimizedForStreaming *float64 `json:"optimizedForStreaming,omitempty"`
	Has64bitOffsets       *bool    `json:"has64bitOffsets,omitempty"`
	VideoProfile          *string  `json:"videoProfile,omitempty"`
	Part                  []Part   `json:"Part,omitempty"`
}

func (o *Media) GetID() *float64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Media) GetDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *Media) GetBitrate() *float64 {
	if o == nil {
		return nil
	}
	return o.Bitrate
}

func (o *Media) GetWidth() *float64 {
	if o == nil {
		return nil
	}
	return o.Width
}

func (o *Media) GetHeight() *float64 {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *Media) GetAspectRatio() *float64 {
	if o == nil {
		return nil
	}
	return o.AspectRatio
}

func (o *Media) GetAudioChannels() *float64 {
	if o == nil {
		return nil
	}
	return o.AudioChannels
}

func (o *Media) GetAudioCodec() *string {
	if o == nil {
		return nil
	}
	return o.AudioCodec
}

func (o *Media) GetVideoCodec() *string {
	if o == nil {
		return nil
	}
	return o.VideoCodec
}

func (o *Media) GetVideoResolution() *float64 {
	if o == nil {
		return nil
	}
	return o.VideoResolution
}

func (o *Media) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *Media) GetVideoFrameRate() *string {
	if o == nil {
		return nil
	}
	return o.VideoFrameRate
}

func (o *Media) GetOptimizedForStreaming() *float64 {
	if o == nil {
		return nil
	}
	return o.OptimizedForStreaming
}

func (o *Media) GetHas64bitOffsets() *bool {
	if o == nil {
		return nil
	}
	return o.Has64bitOffsets
}

func (o *Media) GetVideoProfile() *string {
	if o == nil {
		return nil
	}
	return o.VideoProfile
}

func (o *Media) GetPart() []Part {
	if o == nil {
		return nil
	}
	return o.Part
}

type Genre struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *Genre) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type Director struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *Director) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type Writer struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *Writer) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type Country struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *Country) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type Role struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *Role) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetRecentlyAddedMetadata struct {
	AllowSync             *bool      `json:"allowSync,omitempty"`
	LibrarySectionID      *float64   `json:"librarySectionID,omitempty"`
	LibrarySectionTitle   *string    `json:"librarySectionTitle,omitempty"`
	LibrarySectionUUID    *string    `json:"librarySectionUUID,omitempty"`
	RatingKey             *float64   `json:"ratingKey,omitempty"`
	Key                   *string    `json:"key,omitempty"`
	GUID                  *string    `json:"guid,omitempty"`
	Studio                *string    `json:"studio,omitempty"`
	Type                  *string    `json:"type,omitempty"`
	Title                 *string    `json:"title,omitempty"`
	ContentRating         *string    `json:"contentRating,omitempty"`
	Summary               *string    `json:"summary,omitempty"`
	Rating                *float64   `json:"rating,omitempty"`
	AudienceRating        *float64   `json:"audienceRating,omitempty"`
	Year                  *float64   `json:"year,omitempty"`
	Tagline               *string    `json:"tagline,omitempty"`
	Thumb                 *string    `json:"thumb,omitempty"`
	Art                   *string    `json:"art,omitempty"`
	Duration              *float64   `json:"duration,omitempty"`
	OriginallyAvailableAt *time.Time `json:"originallyAvailableAt,omitempty"`
	AddedAt               *float64   `json:"addedAt,omitempty"`
	UpdatedAt             *float64   `json:"updatedAt,omitempty"`
	AudienceRatingImage   *string    `json:"audienceRatingImage,omitempty"`
	ChapterSource         *string    `json:"chapterSource,omitempty"`
	PrimaryExtraKey       *string    `json:"primaryExtraKey,omitempty"`
	RatingImage           *string    `json:"ratingImage,omitempty"`
	Media                 []Media    `json:"Media,omitempty"`
	Genre                 []Genre    `json:"Genre,omitempty"`
	Director              []Director `json:"Director,omitempty"`
	Writer                []Writer   `json:"Writer,omitempty"`
	Country               []Country  `json:"Country,omitempty"`
	Role                  []Role     `json:"Role,omitempty"`
}

func (g GetRecentlyAddedMetadata) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetRecentlyAddedMetadata) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetRecentlyAddedMetadata) GetAllowSync() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSync
}

func (o *GetRecentlyAddedMetadata) GetLibrarySectionID() *float64 {
	if o == nil {
		return nil
	}
	return o.LibrarySectionID
}

func (o *GetRecentlyAddedMetadata) GetLibrarySectionTitle() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionTitle
}

func (o *GetRecentlyAddedMetadata) GetLibrarySectionUUID() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionUUID
}

func (o *GetRecentlyAddedMetadata) GetRatingKey() *float64 {
	if o == nil {
		return nil
	}
	return o.RatingKey
}

func (o *GetRecentlyAddedMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetRecentlyAddedMetadata) GetGUID() *string {
	if o == nil {
		return nil
	}
	return o.GUID
}

func (o *GetRecentlyAddedMetadata) GetStudio() *string {
	if o == nil {
		return nil
	}
	return o.Studio
}

func (o *GetRecentlyAddedMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetRecentlyAddedMetadata) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetRecentlyAddedMetadata) GetContentRating() *string {
	if o == nil {
		return nil
	}
	return o.ContentRating
}

func (o *GetRecentlyAddedMetadata) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *GetRecentlyAddedMetadata) GetRating() *float64 {
	if o == nil {
		return nil
	}
	return o.Rating
}

func (o *GetRecentlyAddedMetadata) GetAudienceRating() *float64 {
	if o == nil {
		return nil
	}
	return o.AudienceRating
}

func (o *GetRecentlyAddedMetadata) GetYear() *float64 {
	if o == nil {
		return nil
	}
	return o.Year
}

func (o *GetRecentlyAddedMetadata) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}

func (o *GetRecentlyAddedMetadata) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetRecentlyAddedMetadata) GetArt() *string {
	if o == nil {
		return nil
	}
	return o.Art
}

func (o *GetRecentlyAddedMetadata) GetDuration() *float64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetRecentlyAddedMetadata) GetOriginallyAvailableAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.OriginallyAvailableAt
}

func (o *GetRecentlyAddedMetadata) GetAddedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.AddedAt
}

func (o *GetRecentlyAddedMetadata) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetRecentlyAddedMetadata) GetAudienceRatingImage() *string {
	if o == nil {
		return nil
	}
	return o.AudienceRatingImage
}

func (o *GetRecentlyAddedMetadata) GetChapterSource() *string {
	if o == nil {
		return nil
	}
	return o.ChapterSource
}

func (o *GetRecentlyAddedMetadata) GetPrimaryExtraKey() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryExtraKey
}

func (o *GetRecentlyAddedMetadata) GetRatingImage() *string {
	if o == nil {
		return nil
	}
	return o.RatingImage
}

func (o *GetRecentlyAddedMetadata) GetMedia() []Media {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *GetRecentlyAddedMetadata) GetGenre() []Genre {
	if o == nil {
		return nil
	}
	return o.Genre
}

func (o *GetRecentlyAddedMetadata) GetDirector() []Director {
	if o == nil {
		return nil
	}
	return o.Director
}

func (o *GetRecentlyAddedMetadata) GetWriter() []Writer {
	if o == nil {
		return nil
	}
	return o.Writer
}

func (o *GetRecentlyAddedMetadata) GetCountry() []Country {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *GetRecentlyAddedMetadata) GetRole() []Role {
	if o == nil {
		return nil
	}
	return o.Role
}

type GetRecentlyAddedMediaContainer struct {
	Size            *float64                   `json:"size,omitempty"`
	AllowSync       *bool                      `json:"allowSync,omitempty"`
	Identifier      *string                    `json:"identifier,omitempty"`
	MediaTagPrefix  *string                    `json:"mediaTagPrefix,omitempty"`
	MediaTagVersion *float64                   `json:"mediaTagVersion,omitempty"`
	MixedParents    *bool                      `json:"mixedParents,omitempty"`
	Metadata        []GetRecentlyAddedMetadata `json:"Metadata,omitempty"`
}

func (o *GetRecentlyAddedMediaContainer) GetSize() *float64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetRecentlyAddedMediaContainer) GetAllowSync() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSync
}

func (o *GetRecentlyAddedMediaContainer) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *GetRecentlyAddedMediaContainer) GetMediaTagPrefix() *string {
	if o == nil {
		return nil
	}
	return o.MediaTagPrefix
}

func (o *GetRecentlyAddedMediaContainer) GetMediaTagVersion() *float64 {
	if o == nil {
		return nil
	}
	return o.MediaTagVersion
}

func (o *GetRecentlyAddedMediaContainer) GetMixedParents() *bool {
	if o == nil {
		return nil
	}
	return o.MixedParents
}

func (o *GetRecentlyAddedMediaContainer) GetMetadata() []GetRecentlyAddedMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// GetRecentlyAddedResponseBody - The recently added content
type GetRecentlyAddedResponseBody struct {
	MediaContainer *GetRecentlyAddedMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetRecentlyAddedResponseBody) GetMediaContainer() *GetRecentlyAddedMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetRecentlyAddedResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The recently added content
	TwoHundredApplicationJSONObject *GetRecentlyAddedResponseBody
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	FourHundredAndOneApplicationJSONObject *GetRecentlyAddedLibraryResponseBody
}

func (o *GetRecentlyAddedResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRecentlyAddedResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRecentlyAddedResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetRecentlyAddedResponse) GetTwoHundredApplicationJSONObject() *GetRecentlyAddedResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetRecentlyAddedResponse) GetFourHundredAndOneApplicationJSONObject() *GetRecentlyAddedLibraryResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}
