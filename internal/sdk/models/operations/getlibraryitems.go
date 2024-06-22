// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/LukeHagar/terraform-provider-PlexAPI/internal/sdk/internal/utils"
	"github.com/LukeHagar/terraform-provider-PlexAPI/internal/sdk/types"
	"net/http"
)

// Tag - A key representing a specific tag within the section.
type Tag string

const (
	TagAll            Tag = "all"
	TagUnwatched      Tag = "unwatched"
	TagNewest         Tag = "newest"
	TagRecentlyAdded  Tag = "recentlyAdded"
	TagRecentlyViewed Tag = "recentlyViewed"
	TagOnDeck         Tag = "onDeck"
	TagCollection     Tag = "collection"
	TagEdition        Tag = "edition"
	TagGenre          Tag = "genre"
	TagYear           Tag = "year"
	TagDecade         Tag = "decade"
	TagDirector       Tag = "director"
	TagActor          Tag = "actor"
	TagCountry        Tag = "country"
	TagContentRating  Tag = "contentRating"
	TagRating         Tag = "rating"
	TagResolution     Tag = "resolution"
	TagFirstCharacter Tag = "firstCharacter"
	TagFolder         Tag = "folder"
)

func (e Tag) ToPointer() *Tag {
	return &e
}
func (e *Tag) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "all":
		fallthrough
	case "unwatched":
		fallthrough
	case "newest":
		fallthrough
	case "recentlyAdded":
		fallthrough
	case "recentlyViewed":
		fallthrough
	case "onDeck":
		fallthrough
	case "collection":
		fallthrough
	case "edition":
		fallthrough
	case "genre":
		fallthrough
	case "year":
		fallthrough
	case "decade":
		fallthrough
	case "director":
		fallthrough
	case "actor":
		fallthrough
	case "country":
		fallthrough
	case "contentRating":
		fallthrough
	case "rating":
		fallthrough
	case "resolution":
		fallthrough
	case "firstCharacter":
		fallthrough
	case "folder":
		*e = Tag(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Tag: %v", v)
	}
}

type GetLibraryItemsRequest struct {
	// the Id of the library to query
	SectionID any `pathParam:"style=simple,explode=false,name=sectionId"`
	// A key representing a specific tag within the section.
	Tag Tag `pathParam:"style=simple,explode=false,name=tag"`
	// Adds the Guids object to the response
	//
	IncludeGuids *int64 `queryParam:"style=form,explode=true,name=includeGuids"`
}

func (o *GetLibraryItemsRequest) GetSectionID() any {
	if o == nil {
		return nil
	}
	return o.SectionID
}

func (o *GetLibraryItemsRequest) GetTag() Tag {
	if o == nil {
		return Tag("")
	}
	return o.Tag
}

func (o *GetLibraryItemsRequest) GetIncludeGuids() *int64 {
	if o == nil {
		return nil
	}
	return o.IncludeGuids
}

type GetLibraryItemsErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetLibraryItemsErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetLibraryItemsErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetLibraryItemsErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetLibraryItemsLibraryResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetLibraryItemsLibraryResponseBody struct {
	Errors []GetLibraryItemsErrors `json:"errors,omitempty"`
}

func (o *GetLibraryItemsLibraryResponseBody) GetErrors() []GetLibraryItemsErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type LibrarySectionIDType string

const (
	LibrarySectionIDTypeInteger LibrarySectionIDType = "integer"
	LibrarySectionIDTypeStr     LibrarySectionIDType = "str"
)

type LibrarySectionID struct {
	Integer *int64
	Str     *string

	Type LibrarySectionIDType
}

func CreateLibrarySectionIDInteger(integer int64) LibrarySectionID {
	typ := LibrarySectionIDTypeInteger

	return LibrarySectionID{
		Integer: &integer,
		Type:    typ,
	}
}

func CreateLibrarySectionIDStr(str string) LibrarySectionID {
	typ := LibrarySectionIDTypeStr

	return LibrarySectionID{
		Str:  &str,
		Type: typ,
	}
}

func (u *LibrarySectionID) UnmarshalJSON(data []byte) error {

	var integer int64 = int64(0)
	if err := utils.UnmarshalJSON(data, &integer, "", true, true); err == nil {
		u.Integer = &integer
		u.Type = LibrarySectionIDTypeInteger
		return nil
	}

	var str string = ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = LibrarySectionIDTypeStr
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for LibrarySectionID", string(data))
}

func (u LibrarySectionID) MarshalJSON() ([]byte, error) {
	if u.Integer != nil {
		return utils.MarshalJSON(u.Integer, "", true)
	}

	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	return nil, errors.New("could not marshal union type LibrarySectionID: all fields are null")
}

type GetLibraryItemsPart struct {
	ID           *int    `json:"id,omitempty"`
	Key          *string `json:"key,omitempty"`
	Duration     *int    `json:"duration,omitempty"`
	File         *string `json:"file,omitempty"`
	Size         *int64  `json:"size,omitempty"`
	Container    *string `json:"container,omitempty"`
	VideoProfile *string `json:"videoProfile,omitempty"`
}

func (o *GetLibraryItemsPart) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetLibraryItemsPart) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetLibraryItemsPart) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetLibraryItemsPart) GetFile() *string {
	if o == nil {
		return nil
	}
	return o.File
}

func (o *GetLibraryItemsPart) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetLibraryItemsPart) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *GetLibraryItemsPart) GetVideoProfile() *string {
	if o == nil {
		return nil
	}
	return o.VideoProfile
}

type GetLibraryItemsMedia struct {
	ID              *int                  `json:"id,omitempty"`
	Duration        *int                  `json:"duration,omitempty"`
	Bitrate         *int                  `json:"bitrate,omitempty"`
	Width           *int                  `json:"width,omitempty"`
	Height          *int                  `json:"height,omitempty"`
	AspectRatio     *float64              `json:"aspectRatio,omitempty"`
	AudioChannels   *int                  `json:"audioChannels,omitempty"`
	AudioCodec      *string               `json:"audioCodec,omitempty"`
	VideoCodec      *string               `json:"videoCodec,omitempty"`
	VideoResolution *string               `json:"videoResolution,omitempty"`
	Container       *string               `json:"container,omitempty"`
	VideoFrameRate  *string               `json:"videoFrameRate,omitempty"`
	VideoProfile    *string               `json:"videoProfile,omitempty"`
	Part            []GetLibraryItemsPart `json:"Part,omitempty"`
}

func (o *GetLibraryItemsMedia) GetID() *int {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetLibraryItemsMedia) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetLibraryItemsMedia) GetBitrate() *int {
	if o == nil {
		return nil
	}
	return o.Bitrate
}

func (o *GetLibraryItemsMedia) GetWidth() *int {
	if o == nil {
		return nil
	}
	return o.Width
}

func (o *GetLibraryItemsMedia) GetHeight() *int {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *GetLibraryItemsMedia) GetAspectRatio() *float64 {
	if o == nil {
		return nil
	}
	return o.AspectRatio
}

func (o *GetLibraryItemsMedia) GetAudioChannels() *int {
	if o == nil {
		return nil
	}
	return o.AudioChannels
}

func (o *GetLibraryItemsMedia) GetAudioCodec() *string {
	if o == nil {
		return nil
	}
	return o.AudioCodec
}

func (o *GetLibraryItemsMedia) GetVideoCodec() *string {
	if o == nil {
		return nil
	}
	return o.VideoCodec
}

func (o *GetLibraryItemsMedia) GetVideoResolution() *string {
	if o == nil {
		return nil
	}
	return o.VideoResolution
}

func (o *GetLibraryItemsMedia) GetContainer() *string {
	if o == nil {
		return nil
	}
	return o.Container
}

func (o *GetLibraryItemsMedia) GetVideoFrameRate() *string {
	if o == nil {
		return nil
	}
	return o.VideoFrameRate
}

func (o *GetLibraryItemsMedia) GetVideoProfile() *string {
	if o == nil {
		return nil
	}
	return o.VideoProfile
}

func (o *GetLibraryItemsMedia) GetPart() []GetLibraryItemsPart {
	if o == nil {
		return nil
	}
	return o.Part
}

type GetLibraryItemsGenre struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetLibraryItemsGenre) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetLibraryItemsCountry struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetLibraryItemsCountry) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetLibraryItemsDirector struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetLibraryItemsDirector) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetLibraryItemsWriter struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetLibraryItemsWriter) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetLibraryItemsRole struct {
	Tag *string `json:"tag,omitempty"`
}

func (o *GetLibraryItemsRole) GetTag() *string {
	if o == nil {
		return nil
	}
	return o.Tag
}

type GetLibraryItemsMetadata struct {
	RatingKey              *string                   `json:"ratingKey,omitempty"`
	Key                    *string                   `json:"key,omitempty"`
	GUID                   *string                   `json:"guid,omitempty"`
	Studio                 *string                   `json:"studio,omitempty"`
	Type                   *string                   `json:"type,omitempty"`
	Title                  *string                   `json:"title,omitempty"`
	ContentRating          *string                   `json:"contentRating,omitempty"`
	Summary                *string                   `json:"summary,omitempty"`
	Rating                 *float64                  `json:"rating,omitempty"`
	AudienceRating         *float64                  `json:"audienceRating,omitempty"`
	Year                   *int                      `json:"year,omitempty"`
	Tagline                *string                   `json:"tagline,omitempty"`
	Thumb                  *string                   `json:"thumb,omitempty"`
	Art                    *string                   `json:"art,omitempty"`
	Duration               *int                      `json:"duration,omitempty"`
	OriginallyAvailableAt  *types.Date               `json:"originallyAvailableAt,omitempty"`
	AddedAt                *int                      `json:"addedAt,omitempty"`
	UpdatedAt              *int                      `json:"updatedAt,omitempty"`
	AudienceRatingImage    *string                   `json:"audienceRatingImage,omitempty"`
	ChapterSource          *string                   `json:"chapterSource,omitempty"`
	PrimaryExtraKey        *string                   `json:"primaryExtraKey,omitempty"`
	RatingImage            *string                   `json:"ratingImage,omitempty"`
	GrandparentRatingKey   *string                   `json:"grandparentRatingKey,omitempty"`
	GrandparentGUID        *string                   `json:"grandparentGuid,omitempty"`
	GrandparentKey         *string                   `json:"grandparentKey,omitempty"`
	GrandparentTitle       *string                   `json:"grandparentTitle,omitempty"`
	GrandparentThumb       *string                   `json:"grandparentThumb,omitempty"`
	GrandparentArt         *string                   `json:"grandparentArt,omitempty"`
	GrandparentTheme       *string                   `json:"grandparentTheme,omitempty"`
	Media                  []GetLibraryItemsMedia    `json:"Media,omitempty"`
	Genre                  []GetLibraryItemsGenre    `json:"Genre,omitempty"`
	Country                []GetLibraryItemsCountry  `json:"Country,omitempty"`
	Director               []GetLibraryItemsDirector `json:"Director,omitempty"`
	Writer                 []GetLibraryItemsWriter   `json:"Writer,omitempty"`
	Role                   []GetLibraryItemsRole     `json:"Role,omitempty"`
	TitleSort              *string                   `json:"titleSort,omitempty"`
	ViewCount              *int                      `json:"viewCount,omitempty"`
	LastViewedAt           *int                      `json:"lastViewedAt,omitempty"`
	OriginalTitle          *string                   `json:"originalTitle,omitempty"`
	ViewOffset             *int                      `json:"viewOffset,omitempty"`
	SkipCount              *int                      `json:"skipCount,omitempty"`
	Index                  *int                      `json:"index,omitempty"`
	Theme                  *string                   `json:"theme,omitempty"`
	LeafCount              *int                      `json:"leafCount,omitempty"`
	ViewedLeafCount        *int                      `json:"viewedLeafCount,omitempty"`
	ChildCount             *int                      `json:"childCount,omitempty"`
	HasPremiumExtras       *string                   `json:"hasPremiumExtras,omitempty"`
	HasPremiumPrimaryExtra *string                   `json:"hasPremiumPrimaryExtra,omitempty"`
	ParentRatingKey        *string                   `json:"parentRatingKey,omitempty"`
	ParentGUID             *string                   `json:"parentGuid,omitempty"`
	ParentStudio           *string                   `json:"parentStudio,omitempty"`
	ParentKey              *string                   `json:"parentKey,omitempty"`
	ParentTitle            *string                   `json:"parentTitle,omitempty"`
	ParentIndex            *int                      `json:"parentIndex,omitempty"`
	ParentYear             *int                      `json:"parentYear,omitempty"`
	ParentThumb            *string                   `json:"parentThumb,omitempty"`
	ParentTheme            *string                   `json:"parentTheme,omitempty"`
}

func (g GetLibraryItemsMetadata) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetLibraryItemsMetadata) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetLibraryItemsMetadata) GetRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.RatingKey
}

func (o *GetLibraryItemsMetadata) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetLibraryItemsMetadata) GetGUID() *string {
	if o == nil {
		return nil
	}
	return o.GUID
}

func (o *GetLibraryItemsMetadata) GetStudio() *string {
	if o == nil {
		return nil
	}
	return o.Studio
}

func (o *GetLibraryItemsMetadata) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetLibraryItemsMetadata) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetLibraryItemsMetadata) GetContentRating() *string {
	if o == nil {
		return nil
	}
	return o.ContentRating
}

func (o *GetLibraryItemsMetadata) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *GetLibraryItemsMetadata) GetRating() *float64 {
	if o == nil {
		return nil
	}
	return o.Rating
}

func (o *GetLibraryItemsMetadata) GetAudienceRating() *float64 {
	if o == nil {
		return nil
	}
	return o.AudienceRating
}

func (o *GetLibraryItemsMetadata) GetYear() *int {
	if o == nil {
		return nil
	}
	return o.Year
}

func (o *GetLibraryItemsMetadata) GetTagline() *string {
	if o == nil {
		return nil
	}
	return o.Tagline
}

func (o *GetLibraryItemsMetadata) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetLibraryItemsMetadata) GetArt() *string {
	if o == nil {
		return nil
	}
	return o.Art
}

func (o *GetLibraryItemsMetadata) GetDuration() *int {
	if o == nil {
		return nil
	}
	return o.Duration
}

func (o *GetLibraryItemsMetadata) GetOriginallyAvailableAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.OriginallyAvailableAt
}

func (o *GetLibraryItemsMetadata) GetAddedAt() *int {
	if o == nil {
		return nil
	}
	return o.AddedAt
}

func (o *GetLibraryItemsMetadata) GetUpdatedAt() *int {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GetLibraryItemsMetadata) GetAudienceRatingImage() *string {
	if o == nil {
		return nil
	}
	return o.AudienceRatingImage
}

func (o *GetLibraryItemsMetadata) GetChapterSource() *string {
	if o == nil {
		return nil
	}
	return o.ChapterSource
}

func (o *GetLibraryItemsMetadata) GetPrimaryExtraKey() *string {
	if o == nil {
		return nil
	}
	return o.PrimaryExtraKey
}

func (o *GetLibraryItemsMetadata) GetRatingImage() *string {
	if o == nil {
		return nil
	}
	return o.RatingImage
}

func (o *GetLibraryItemsMetadata) GetGrandparentRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentRatingKey
}

func (o *GetLibraryItemsMetadata) GetGrandparentGUID() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentGUID
}

func (o *GetLibraryItemsMetadata) GetGrandparentKey() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentKey
}

func (o *GetLibraryItemsMetadata) GetGrandparentTitle() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentTitle
}

func (o *GetLibraryItemsMetadata) GetGrandparentThumb() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentThumb
}

func (o *GetLibraryItemsMetadata) GetGrandparentArt() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentArt
}

func (o *GetLibraryItemsMetadata) GetGrandparentTheme() *string {
	if o == nil {
		return nil
	}
	return o.GrandparentTheme
}

func (o *GetLibraryItemsMetadata) GetMedia() []GetLibraryItemsMedia {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *GetLibraryItemsMetadata) GetGenre() []GetLibraryItemsGenre {
	if o == nil {
		return nil
	}
	return o.Genre
}

func (o *GetLibraryItemsMetadata) GetCountry() []GetLibraryItemsCountry {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *GetLibraryItemsMetadata) GetDirector() []GetLibraryItemsDirector {
	if o == nil {
		return nil
	}
	return o.Director
}

func (o *GetLibraryItemsMetadata) GetWriter() []GetLibraryItemsWriter {
	if o == nil {
		return nil
	}
	return o.Writer
}

func (o *GetLibraryItemsMetadata) GetRole() []GetLibraryItemsRole {
	if o == nil {
		return nil
	}
	return o.Role
}

func (o *GetLibraryItemsMetadata) GetTitleSort() *string {
	if o == nil {
		return nil
	}
	return o.TitleSort
}

func (o *GetLibraryItemsMetadata) GetViewCount() *int {
	if o == nil {
		return nil
	}
	return o.ViewCount
}

func (o *GetLibraryItemsMetadata) GetLastViewedAt() *int {
	if o == nil {
		return nil
	}
	return o.LastViewedAt
}

func (o *GetLibraryItemsMetadata) GetOriginalTitle() *string {
	if o == nil {
		return nil
	}
	return o.OriginalTitle
}

func (o *GetLibraryItemsMetadata) GetViewOffset() *int {
	if o == nil {
		return nil
	}
	return o.ViewOffset
}

func (o *GetLibraryItemsMetadata) GetSkipCount() *int {
	if o == nil {
		return nil
	}
	return o.SkipCount
}

func (o *GetLibraryItemsMetadata) GetIndex() *int {
	if o == nil {
		return nil
	}
	return o.Index
}

func (o *GetLibraryItemsMetadata) GetTheme() *string {
	if o == nil {
		return nil
	}
	return o.Theme
}

func (o *GetLibraryItemsMetadata) GetLeafCount() *int {
	if o == nil {
		return nil
	}
	return o.LeafCount
}

func (o *GetLibraryItemsMetadata) GetViewedLeafCount() *int {
	if o == nil {
		return nil
	}
	return o.ViewedLeafCount
}

func (o *GetLibraryItemsMetadata) GetChildCount() *int {
	if o == nil {
		return nil
	}
	return o.ChildCount
}

func (o *GetLibraryItemsMetadata) GetHasPremiumExtras() *string {
	if o == nil {
		return nil
	}
	return o.HasPremiumExtras
}

func (o *GetLibraryItemsMetadata) GetHasPremiumPrimaryExtra() *string {
	if o == nil {
		return nil
	}
	return o.HasPremiumPrimaryExtra
}

func (o *GetLibraryItemsMetadata) GetParentRatingKey() *string {
	if o == nil {
		return nil
	}
	return o.ParentRatingKey
}

func (o *GetLibraryItemsMetadata) GetParentGUID() *string {
	if o == nil {
		return nil
	}
	return o.ParentGUID
}

func (o *GetLibraryItemsMetadata) GetParentStudio() *string {
	if o == nil {
		return nil
	}
	return o.ParentStudio
}

func (o *GetLibraryItemsMetadata) GetParentKey() *string {
	if o == nil {
		return nil
	}
	return o.ParentKey
}

func (o *GetLibraryItemsMetadata) GetParentTitle() *string {
	if o == nil {
		return nil
	}
	return o.ParentTitle
}

func (o *GetLibraryItemsMetadata) GetParentIndex() *int {
	if o == nil {
		return nil
	}
	return o.ParentIndex
}

func (o *GetLibraryItemsMetadata) GetParentYear() *int {
	if o == nil {
		return nil
	}
	return o.ParentYear
}

func (o *GetLibraryItemsMetadata) GetParentThumb() *string {
	if o == nil {
		return nil
	}
	return o.ParentThumb
}

func (o *GetLibraryItemsMetadata) GetParentTheme() *string {
	if o == nil {
		return nil
	}
	return o.ParentTheme
}

type GetLibraryItemsMediaContainer struct {
	Size                *int                      `json:"size,omitempty"`
	AllowSync           *bool                     `json:"allowSync,omitempty"`
	Art                 *string                   `json:"art,omitempty"`
	Identifier          *string                   `json:"identifier,omitempty"`
	LibrarySectionID    *LibrarySectionID         `json:"librarySectionID,omitempty"`
	LibrarySectionTitle *string                   `json:"librarySectionTitle,omitempty"`
	LibrarySectionUUID  *string                   `json:"librarySectionUUID,omitempty"`
	MediaTagPrefix      *string                   `json:"mediaTagPrefix,omitempty"`
	MediaTagVersion     *int                      `json:"mediaTagVersion,omitempty"`
	Thumb               *string                   `json:"thumb,omitempty"`
	Title1              *string                   `json:"title1,omitempty"`
	Title2              *string                   `json:"title2,omitempty"`
	ViewGroup           *string                   `json:"viewGroup,omitempty"`
	ViewMode            *int                      `json:"viewMode,omitempty"`
	MixedParents        *bool                     `json:"mixedParents,omitempty"`
	Metadata            []GetLibraryItemsMetadata `json:"Metadata,omitempty"`
}

func (o *GetLibraryItemsMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetLibraryItemsMediaContainer) GetAllowSync() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSync
}

func (o *GetLibraryItemsMediaContainer) GetArt() *string {
	if o == nil {
		return nil
	}
	return o.Art
}

func (o *GetLibraryItemsMediaContainer) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *GetLibraryItemsMediaContainer) GetLibrarySectionID() *LibrarySectionID {
	if o == nil {
		return nil
	}
	return o.LibrarySectionID
}

func (o *GetLibraryItemsMediaContainer) GetLibrarySectionTitle() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionTitle
}

func (o *GetLibraryItemsMediaContainer) GetLibrarySectionUUID() *string {
	if o == nil {
		return nil
	}
	return o.LibrarySectionUUID
}

func (o *GetLibraryItemsMediaContainer) GetMediaTagPrefix() *string {
	if o == nil {
		return nil
	}
	return o.MediaTagPrefix
}

func (o *GetLibraryItemsMediaContainer) GetMediaTagVersion() *int {
	if o == nil {
		return nil
	}
	return o.MediaTagVersion
}

func (o *GetLibraryItemsMediaContainer) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetLibraryItemsMediaContainer) GetTitle1() *string {
	if o == nil {
		return nil
	}
	return o.Title1
}

func (o *GetLibraryItemsMediaContainer) GetTitle2() *string {
	if o == nil {
		return nil
	}
	return o.Title2
}

func (o *GetLibraryItemsMediaContainer) GetViewGroup() *string {
	if o == nil {
		return nil
	}
	return o.ViewGroup
}

func (o *GetLibraryItemsMediaContainer) GetViewMode() *int {
	if o == nil {
		return nil
	}
	return o.ViewMode
}

func (o *GetLibraryItemsMediaContainer) GetMixedParents() *bool {
	if o == nil {
		return nil
	}
	return o.MixedParents
}

func (o *GetLibraryItemsMediaContainer) GetMetadata() []GetLibraryItemsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

// GetLibraryItemsResponseBody - The contents of the library by section and tag
type GetLibraryItemsResponseBody struct {
	MediaContainer *GetLibraryItemsMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetLibraryItemsResponseBody) GetMediaContainer() *GetLibraryItemsMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetLibraryItemsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The contents of the library by section and tag
	TwoHundredApplicationJSONObject *GetLibraryItemsResponseBody
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	FourHundredAndOneApplicationJSONObject *GetLibraryItemsLibraryResponseBody
}

func (o *GetLibraryItemsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLibraryItemsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLibraryItemsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLibraryItemsResponse) GetTwoHundredApplicationJSONObject() *GetLibraryItemsResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetLibraryItemsResponse) GetFourHundredAndOneApplicationJSONObject() *GetLibraryItemsLibraryResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}
