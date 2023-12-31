// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetMyPlexAccountErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetMyPlexAccountErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetMyPlexAccountErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetMyPlexAccountErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetMyPlexAccountServerResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetMyPlexAccountServerResponseBody struct {
	Errors []GetMyPlexAccountErrors `json:"errors,omitempty"`
}

func (o *GetMyPlexAccountServerResponseBody) GetErrors() []GetMyPlexAccountErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type MyPlex struct {
	AuthToken            *string  `json:"authToken,omitempty"`
	Username             *string  `json:"username,omitempty"`
	MappingState         *string  `json:"mappingState,omitempty"`
	MappingError         *string  `json:"mappingError,omitempty"`
	SignInState          *string  `json:"signInState,omitempty"`
	PublicAddress        *string  `json:"publicAddress,omitempty"`
	PublicPort           *float64 `json:"publicPort,omitempty"`
	PrivateAddress       *string  `json:"privateAddress,omitempty"`
	PrivatePort          *float64 `json:"privatePort,omitempty"`
	SubscriptionFeatures *string  `json:"subscriptionFeatures,omitempty"`
	SubscriptionActive   *bool    `json:"subscriptionActive,omitempty"`
	SubscriptionState    *string  `json:"subscriptionState,omitempty"`
}

func (o *MyPlex) GetAuthToken() *string {
	if o == nil {
		return nil
	}
	return o.AuthToken
}

func (o *MyPlex) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *MyPlex) GetMappingState() *string {
	if o == nil {
		return nil
	}
	return o.MappingState
}

func (o *MyPlex) GetMappingError() *string {
	if o == nil {
		return nil
	}
	return o.MappingError
}

func (o *MyPlex) GetSignInState() *string {
	if o == nil {
		return nil
	}
	return o.SignInState
}

func (o *MyPlex) GetPublicAddress() *string {
	if o == nil {
		return nil
	}
	return o.PublicAddress
}

func (o *MyPlex) GetPublicPort() *float64 {
	if o == nil {
		return nil
	}
	return o.PublicPort
}

func (o *MyPlex) GetPrivateAddress() *string {
	if o == nil {
		return nil
	}
	return o.PrivateAddress
}

func (o *MyPlex) GetPrivatePort() *float64 {
	if o == nil {
		return nil
	}
	return o.PrivatePort
}

func (o *MyPlex) GetSubscriptionFeatures() *string {
	if o == nil {
		return nil
	}
	return o.SubscriptionFeatures
}

func (o *MyPlex) GetSubscriptionActive() *bool {
	if o == nil {
		return nil
	}
	return o.SubscriptionActive
}

func (o *MyPlex) GetSubscriptionState() *string {
	if o == nil {
		return nil
	}
	return o.SubscriptionState
}

// GetMyPlexAccountResponseBody - MyPlex Account
type GetMyPlexAccountResponseBody struct {
	MyPlex *MyPlex `json:"MyPlex,omitempty"`
}

func (o *GetMyPlexAccountResponseBody) GetMyPlex() *MyPlex {
	if o == nil {
		return nil
	}
	return o.MyPlex
}

type GetMyPlexAccountResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// MyPlex Account
	TwoHundredApplicationJSONObject *GetMyPlexAccountResponseBody
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	FourHundredAndOneApplicationJSONObject *GetMyPlexAccountServerResponseBody
}

func (o *GetMyPlexAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetMyPlexAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetMyPlexAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetMyPlexAccountResponse) GetTwoHundredApplicationJSONObject() *GetMyPlexAccountResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetMyPlexAccountResponse) GetFourHundredAndOneApplicationJSONObject() *GetMyPlexAccountServerResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}
