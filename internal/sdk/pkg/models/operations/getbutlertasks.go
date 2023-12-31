// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
)

type GetButlerTasksErrors struct {
	Code    *float64 `json:"code,omitempty"`
	Message *string  `json:"message,omitempty"`
	Status  *float64 `json:"status,omitempty"`
}

func (o *GetButlerTasksErrors) GetCode() *float64 {
	if o == nil {
		return nil
	}
	return o.Code
}

func (o *GetButlerTasksErrors) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *GetButlerTasksErrors) GetStatus() *float64 {
	if o == nil {
		return nil
	}
	return o.Status
}

// GetButlerTasksButlerResponseBody - Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
type GetButlerTasksButlerResponseBody struct {
	Errors []GetButlerTasksErrors `json:"errors,omitempty"`
}

func (o *GetButlerTasksButlerResponseBody) GetErrors() []GetButlerTasksErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

type ButlerTask struct {
	Name               *string  `json:"name,omitempty"`
	Interval           *float64 `json:"interval,omitempty"`
	ScheduleRandomized *bool    `json:"scheduleRandomized,omitempty"`
	Enabled            *bool    `json:"enabled,omitempty"`
	Title              *string  `json:"title,omitempty"`
	Description        *string  `json:"description,omitempty"`
}

func (o *ButlerTask) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ButlerTask) GetInterval() *float64 {
	if o == nil {
		return nil
	}
	return o.Interval
}

func (o *ButlerTask) GetScheduleRandomized() *bool {
	if o == nil {
		return nil
	}
	return o.ScheduleRandomized
}

func (o *ButlerTask) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *ButlerTask) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *ButlerTask) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

type ButlerTasks struct {
	ButlerTask []ButlerTask `json:"ButlerTask,omitempty"`
}

func (o *ButlerTasks) GetButlerTask() []ButlerTask {
	if o == nil {
		return nil
	}
	return o.ButlerTask
}

// GetButlerTasksResponseBody - All butler tasks
type GetButlerTasksResponseBody struct {
	ButlerTasks *ButlerTasks `json:"ButlerTasks,omitempty"`
}

func (o *GetButlerTasksResponseBody) GetButlerTasks() *ButlerTasks {
	if o == nil {
		return nil
	}
	return o.ButlerTasks
}

type GetButlerTasksResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// All butler tasks
	TwoHundredApplicationJSONObject *GetButlerTasksResponseBody
	// Unauthorized - Returned if the X-Plex-Token is missing from the header or query.
	FourHundredAndOneApplicationJSONObject *GetButlerTasksButlerResponseBody
}

func (o *GetButlerTasksResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetButlerTasksResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetButlerTasksResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetButlerTasksResponse) GetTwoHundredApplicationJSONObject() *GetButlerTasksResponseBody {
	if o == nil {
		return nil
	}
	return o.TwoHundredApplicationJSONObject
}

func (o *GetButlerTasksResponse) GetFourHundredAndOneApplicationJSONObject() *GetButlerTasksButlerResponseBody {
	if o == nil {
		return nil
	}
	return o.FourHundredAndOneApplicationJSONObject
}
