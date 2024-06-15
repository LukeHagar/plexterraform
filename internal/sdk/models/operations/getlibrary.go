// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/LukeHagar/terraform-provider-PlexAPI/internal/sdk/internal/utils"
	"net/http"
)

// IncludeDetails - Whether or not to include details for a section (types, filters, and sorts).
// Only exists for backwards compatibility, media providers other than the server libraries have it on always.
type IncludeDetails int64

const (
	IncludeDetailsZero IncludeDetails = 0
	IncludeDetailsOne  IncludeDetails = 1
)

func (e IncludeDetails) ToPointer() *IncludeDetails {
	return &e
}
func (e *IncludeDetails) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = IncludeDetails(v)
		return nil
	default:
		return fmt.Errorf("invalid value for IncludeDetails: %v", v)
	}
}

type GetLibraryRequest struct {
	// the Id of the library to query
	SectionID float64 `pathParam:"style=simple,explode=false,name=sectionId"`
	// Whether or not to include details for a section (types, filters, and sorts).
	// Only exists for backwards compatibility, media providers other than the server libraries have it on always.
	//
	IncludeDetails *IncludeDetails `default:"0" queryParam:"style=form,explode=true,name=includeDetails"`
}

func (g GetLibraryRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetLibraryRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetLibraryRequest) GetSectionID() float64 {
	if o == nil {
		return 0.0
	}
	return o.SectionID
}

func (o *GetLibraryRequest) GetIncludeDetails() *IncludeDetails {
	if o == nil {
		return nil
	}
	return o.IncludeDetails
}

type GetLibraryErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetLibraryErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetLibraryErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetLibraryErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetLibraryLibraryResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetLibraryLibraryResponseBody struct {
	Errors []GetLibraryErrors `json:"errors,omitempty"`
}

func (o *GetLibraryLibraryResponseBody) GetErrors() []GetLibraryErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type GetLibraryDirectory struct {
	Key       *string `json:"key,omitempty"`
	Title     *string `json:"title,omitempty"`
	Secondary *bool   `json:"secondary,omitempty"`
	Prompt    *string `json:"prompt,omitempty"`
	Search    *bool   `json:"search,omitempty"`
}

func (o *GetLibraryDirectory) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetLibraryDirectory) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetLibraryDirectory) GetSecondary() *bool {
	if o == nil {
		return nil
	}
	return o.Secondary
}

func (o *GetLibraryDirectory) GetPrompt() *string {
	if o == nil {
		return nil
	}
	return o.Prompt
}

func (o *GetLibraryDirectory) GetSearch() *bool {
	if o == nil {
		return nil
	}
	return o.Search
}

type GetLibraryFilter struct {
	Filter     *string `json:"filter,omitempty"`
	FilterType *string `json:"filterType,omitempty"`
	Key        *string `json:"key,omitempty"`
	Title      *string `json:"title,omitempty"`
	Type       *string `json:"type,omitempty"`
}

func (o *GetLibraryFilter) GetFilter() *string {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetLibraryFilter) GetFilterType() *string {
	if o == nil {
		return nil
	}
	return o.FilterType
}

func (o *GetLibraryFilter) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetLibraryFilter) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetLibraryFilter) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

type Sort struct {
	Default           *string `json:"default,omitempty"`
	DefaultDirection  *string `json:"defaultDirection,omitempty"`
	DescKey           *string `json:"descKey,omitempty"`
	FirstCharacterKey *string `json:"firstCharacterKey,omitempty"`
	Key               *string `json:"key,omitempty"`
	Title             *string `json:"title,omitempty"`
}

func (o *Sort) GetDefault() *string {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *Sort) GetDefaultDirection() *string {
	if o == nil {
		return nil
	}
	return o.DefaultDirection
}

func (o *Sort) GetDescKey() *string {
	if o == nil {
		return nil
	}
	return o.DescKey
}

func (o *Sort) GetFirstCharacterKey() *string {
	if o == nil {
		return nil
	}
	return o.FirstCharacterKey
}

func (o *Sort) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *Sort) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

type Field struct {
	Key     *string `json:"key,omitempty"`
	Title   *string `json:"title,omitempty"`
	Type    *string `json:"type,omitempty"`
	SubType *string `json:"subType,omitempty"`
}

func (o *Field) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *Field) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Field) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Field) GetSubType() *string {
	if o == nil {
		return nil
	}
	return o.SubType
}

type GetLibraryType struct {
	Key    *string            `json:"key,omitempty"`
	Type   *string            `json:"type,omitempty"`
	Title  *string            `json:"title,omitempty"`
	Active *bool              `json:"active,omitempty"`
	Filter []GetLibraryFilter `json:"Filter,omitempty"`
	Sort   []Sort             `json:"Sort,omitempty"`
	Field  []Field            `json:"Field,omitempty"`
}

func (o *GetLibraryType) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *GetLibraryType) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetLibraryType) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *GetLibraryType) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *GetLibraryType) GetFilter() []GetLibraryFilter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *GetLibraryType) GetSort() []Sort {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *GetLibraryType) GetField() []Field {
	if o == nil {
		return nil
	}
	return o.Field
}

type Operator struct {
	Key   *string `json:"key,omitempty"`
	Title *string `json:"title,omitempty"`
}

func (o *Operator) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *Operator) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

type FieldType struct {
	Type     *string    `json:"type,omitempty"`
	Operator []Operator `json:"Operator,omitempty"`
}

func (o *FieldType) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *FieldType) GetOperator() []Operator {
	if o == nil {
		return nil
	}
	return o.Operator
}

type GetLibraryMediaContainer struct {
	Size             *int                  `json:"size,omitempty"`
	AllowSync        *bool                 `json:"allowSync,omitempty"`
	Art              *string               `json:"art,omitempty"`
	Content          *string               `json:"content,omitempty"`
	Identifier       *string               `json:"identifier,omitempty"`
	LibrarySectionID *int                  `json:"librarySectionID,omitempty"`
	MediaTagPrefix   *string               `json:"mediaTagPrefix,omitempty"`
	MediaTagVersion  *int                  `json:"mediaTagVersion,omitempty"`
	Thumb            *string               `json:"thumb,omitempty"`
	Title1           *string               `json:"title1,omitempty"`
	ViewGroup        *string               `json:"viewGroup,omitempty"`
	ViewMode         *int                  `json:"viewMode,omitempty"`
	Directory        []GetLibraryDirectory `json:"Directory,omitempty"`
	Type             []GetLibraryType      `json:"Type,omitempty"`
	FieldType        []FieldType           `json:"FieldType,omitempty"`
}

func (o *GetLibraryMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetLibraryMediaContainer) GetAllowSync() *bool {
	if o == nil {
		return nil
	}
	return o.AllowSync
}

func (o *GetLibraryMediaContainer) GetArt() *string {
	if o == nil {
		return nil
	}
	return o.Art
}

func (o *GetLibraryMediaContainer) GetContent() *string {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *GetLibraryMediaContainer) GetIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.Identifier
}

func (o *GetLibraryMediaContainer) GetLibrarySectionID() *int {
	if o == nil {
		return nil
	}
	return o.LibrarySectionID
}

func (o *GetLibraryMediaContainer) GetMediaTagPrefix() *string {
	if o == nil {
		return nil
	}
	return o.MediaTagPrefix
}

func (o *GetLibraryMediaContainer) GetMediaTagVersion() *int {
	if o == nil {
		return nil
	}
	return o.MediaTagVersion
}

func (o *GetLibraryMediaContainer) GetThumb() *string {
	if o == nil {
		return nil
	}
	return o.Thumb
}

func (o *GetLibraryMediaContainer) GetTitle1() *string {
	if o == nil {
		return nil
	}
	return o.Title1
}

func (o *GetLibraryMediaContainer) GetViewGroup() *string {
	if o == nil {
		return nil
	}
	return o.ViewGroup
}

func (o *GetLibraryMediaContainer) GetViewMode() *int {
	if o == nil {
		return nil
	}
	return o.ViewMode
}

func (o *GetLibraryMediaContainer) GetDirectory() []GetLibraryDirectory {
	if o == nil {
		return nil
	}
	return o.Directory
}

func (o *GetLibraryMediaContainer) GetType() []GetLibraryType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *GetLibraryMediaContainer) GetFieldType() []FieldType {
	if o == nil {
		return nil
	}
	return o.FieldType
}

// GetLibraryResponseBody - The details of the library
type GetLibraryResponseBody struct {
	MediaContainer *GetLibraryMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetLibraryResponseBody) GetMediaContainer() *GetLibraryMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetLibraryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The details of the library
	TwoHundredApplicationJSONObject *GetLibraryResponseBody
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	FourHundredAndOneApplicationJSONObject *GetLibraryLibraryResponseBody
}

func (o *GetLibraryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLibraryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLibraryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetLibraryResponse) GetTwoHundredApplicationJSONObject() *GetLibraryResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetLibraryResponse) GetFourHundredAndOneApplicationJSONObject() *GetLibraryLibraryResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}
