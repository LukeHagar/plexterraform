// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/LukeHagar/terraform-provider-PlexAPI/internal/sdk/internal/utils"
	"net/http"
	"time"
)

var GetPinServerList = []string{
	"https://plex.tv/api/v2",
}

type GetPinRequest struct {
	// Determines the kind of code returned by the API call
	// Strong codes are used for Pin authentication flows
	// Non-Strong codes are used for `Plex.tv/link`
	//
	Strong *bool `default:"false" queryParam:"style=form,explode=true,name=strong"`
	// The unique identifier for the client application
	// This is used to track the client application and its usage
	// (UUID, serial number, or other number unique per device)
	//
	XPlexClientIdentifier *string `header:"style=simple,explode=false,name=X-Plex-Client-Identifier"`
}

func (g GetPinRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetPinRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetPinRequest) GetStrong() *bool {
	if o == nil {
		return nil
	}
	return o.Strong
}

func (o *GetPinRequest) GetXPlexClientIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.XPlexClientIdentifier
}

type GetPinErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetPinErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetPinErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetPinErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetPinPlexResponseBody - X-Plex-Client-Identifier is missing
type GetPinPlexResponseBody struct {
	Errors []GetPinErrors `json:"errors,omitempty"`
}

func (o *GetPinPlexResponseBody) GetErrors() []GetPinErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type Location struct {
	Code                       *string  `json:"code,omitempty"`
	EuropeanUnionMember        *bool    `json:"european_union_member,omitempty"`
	ContinentCode              *string  `json:"continent_code,omitempty"`
	Country                    *string  `json:"country,omitempty"`
	City                       *string  `json:"city,omitempty"`
	TimeZone                   *string  `json:"time_zone,omitempty"`
	PostalCode                 *float64 `json:"postal_code,omitempty"`
	InPrivacyRestrictedCountry *bool    `json:"in_privacy_restricted_country,omitempty"`
	Subdivisions               *string  `json:"subdivisions,omitempty"`
	Coordinates                *string  `json:"coordinates,omitempty"`
}

func (o *Location) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *Location) GetEuropeanUnionMember() *bool {
	if o == nil {
		return nil
	}
	return o.EuropeanUnionMember
}

func (o *Location) GetContinentCode() *string {
	if o == nil {
		return nil
	}
	return o.ContinentCode
}

func (o *Location) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *Location) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *Location) GetTimeZone() *string {
	if o == nil {
		return nil
	}
	return o.TimeZone
}

func (o *Location) GetPostalCode() *float64 {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *Location) GetInPrivacyRestrictedCountry() *bool {
	if o == nil {
		return nil
	}
	return o.InPrivacyRestrictedCountry
}

func (o *Location) GetSubdivisions() *string {
	if o == nil {
		return nil
	}
	return o.Subdivisions
}

func (o *Location) GetCoordinates() *string {
	if o == nil {
		return nil
	}
	return o.Coordinates
}

// GetPinResponseBody - The Pin
type GetPinResponseBody struct {
	// PinID for use with authentication
	ID      *float64 `json:"id,omitempty"`
	Code    *string  `json:"code,omitempty"`
	Product *string  `json:"product,omitempty"`
	Trusted *bool    `json:"trusted,omitempty"`
	// a link to a QR code hosted on plex.tv
	// The QR code redirects to the relevant `plex.tv/link` authentication page
	// Which then prompts the user for the 4 Digit Link Pin
	//
	Qr               *string    `json:"qr,omitempty"`
	ClientIdentifier *string    `json:"clientIdentifier,omitempty"`
	Location         *Location  `json:"location,omitempty"`
	ExpiresIn        *float64   `json:"expiresIn,omitempty"`
	CreatedAt        *time.Time `json:"createdAt,omitempty"`
	ExpiresAt        *time.Time `json:"expiresAt,omitempty"`
	AuthToken        *string    `json:"authToken,omitempty"`
	NewRegistration  *string    `json:"newRegistration,omitempty"`
}

func (g GetPinResponseBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetPinResponseBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetPinResponseBody) GetID() *float64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GetPinResponseBody) GetCode() *string {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetPinResponseBody) GetProduct() *string {
	if o == nil {
		return nil
	}
	return o.Product
}

func (o *GetPinResponseBody) GetTrusted() *bool {
	if o == nil {
		return nil
	}
	return o.Trusted
}

func (o *GetPinResponseBody) GetQr() *string {
	if o == nil {
		return nil
	}
	return o.Qr
}

func (o *GetPinResponseBody) GetClientIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.ClientIdentifier
}

func (o *GetPinResponseBody) GetLocation() *Location {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *GetPinResponseBody) GetExpiresIn() *float64 {
	if o == nil {
		return nil
	}
	return o.ExpiresIn
}

func (o *GetPinResponseBody) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GetPinResponseBody) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *GetPinResponseBody) GetAuthToken() *string {
	if o == nil {
		return nil
	}
	return o.AuthToken
}

func (o *GetPinResponseBody) GetNewRegistration() *string {
	if o == nil {
		return nil
	}
	return o.NewRegistration
}

type GetPinResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The Pin
	TwoHundredApplicationJSONObject *GetPinResponseBody
	// X-Plex-Client-Identifier is missing
	FourHundredApplicationJSONObject *GetPinPlexResponseBody
}

func (o *GetPinResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPinResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPinResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetPinResponse) GetTwoHundredApplicationJSONObject() *GetPinResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetPinResponse) GetFourHundredApplicationJSONObject() *GetPinPlexResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredApplicationJSONObject
}