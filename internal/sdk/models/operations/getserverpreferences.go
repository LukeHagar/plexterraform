// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetServerPreferencesErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetServerPreferencesErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetServerPreferencesErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetServerPreferencesErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetServerPreferencesServerResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetServerPreferencesServerResponseBody struct {
	Errors []GetServerPreferencesErrors `json:"errors,omitempty"`
}

func (o *GetServerPreferencesServerResponseBody) GetErrors() []GetServerPreferencesErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type Setting struct {
	ID         *string `json:"id,omitempty"`
	Label      *string `json:"label,omitempty"`
	Summary    *string `json:"summary,omitempty"`
	Type       *string `json:"type,omitempty"`
	Default    *bool   `json:"default,omitempty"`
	Value      *bool   `json:"value,omitempty"`
	Hidden     *bool   `json:"hidden,omitempty"`
	Advanced   *bool   `json:"advanced,omitempty"`
	Group      *string `json:"group,omitempty"`
	EnumValues *string `json:"enumValues,omitempty"`
}

func (o *Setting) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Setting) GetLabel() *string {
	if o == nil {
		return nil
	}
	return o.Label
}

func (o *Setting) GetSummary() *string {
	if o == nil {
		return nil
	}
	return o.Summary
}

func (o *Setting) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *Setting) GetDefault() *bool {
	if o == nil {
		return nil
	}
	return o.Default
}

func (o *Setting) GetValue() *bool {
	if o == nil {
		return nil
	}
	return o.Value
}

func (o *Setting) GetHidden() *bool {
	if o == nil {
		return nil
	}
	return o.Hidden
}

func (o *Setting) GetAdvanced() *bool {
	if o == nil {
		return nil
	}
	return o.Advanced
}

func (o *Setting) GetGroup() *string {
	if o == nil {
		return nil
	}
	return o.Group
}

func (o *Setting) GetEnumValues() *string {
	if o == nil {
		return nil
	}
	return o.EnumValues
}

type GetServerPreferencesMediaContainer struct {
	Size    *int      `json:"size,omitempty"`
	Setting []Setting `json:"Setting,omitempty"`
}

func (o *GetServerPreferencesMediaContainer) GetSize() *int {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *GetServerPreferencesMediaContainer) GetSetting() []Setting {
	if o == nil {
		return nil
	}
	return o.Setting
}

// GetServerPreferencesResponseBody - Server Preferences
type GetServerPreferencesResponseBody struct {
	MediaContainer *GetServerPreferencesMediaContainer `json:"MediaContainer,omitempty"`
}

func (o *GetServerPreferencesResponseBody) GetMediaContainer() *GetServerPreferencesMediaContainer {
	if o == nil {
		return nil
	}
	return o.MediaContainer
}

type GetServerPreferencesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Server Preferences
	TwoHundredApplicationJSONObject *GetServerPreferencesResponseBody
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	FourHundredAndOneApplicationJSONObject *GetServerPreferencesServerResponseBody
}

func (o *GetServerPreferencesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetServerPreferencesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetServerPreferencesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetServerPreferencesResponse) GetTwoHundredApplicationJSONObject() *GetServerPreferencesResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetServerPreferencesResponse) GetFourHundredAndOneApplicationJSONObject() *GetServerPreferencesServerResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}