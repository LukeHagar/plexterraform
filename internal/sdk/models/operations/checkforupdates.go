// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Download - Indicate that you want to start download any updates found.
type Download int64

const (
	DownloadZero Download = 0
	DownloadOne  Download = 1
)

func (e Download) ToPointer() *Download {
	return &e
}
func (e *Download) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 0:
		fallthrough
	case 1:
		*e = Download(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Download: %v", v)
	}
}

type CheckForUpdatesRequest struct {
	// Indicate that you want to start download any updates found.
	Download *Download `queryParam:"style=form,explode=true,name=download"`
}

func (o *CheckForUpdatesRequest) GetDownload() *Download {
	if o == nil {
		return nil
	}
	return o.Download
}

type CheckForUpdatesErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *CheckForUpdatesErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *CheckForUpdatesErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CheckForUpdatesErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// CheckForUpdatesResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type CheckForUpdatesResponseBody struct {
	Errors []CheckForUpdatesErrors `json:"errors,omitempty"`
}

func (o *CheckForUpdatesResponseBody) GetErrors() []CheckForUpdatesErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type CheckForUpdatesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	Object *CheckForUpdatesResponseBody
}

func (o *CheckForUpdatesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CheckForUpdatesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CheckForUpdatesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CheckForUpdatesResponse) GetObject() *CheckForUpdatesResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
