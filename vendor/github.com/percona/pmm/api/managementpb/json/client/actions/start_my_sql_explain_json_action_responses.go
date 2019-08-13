// Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// StartMySQLExplainJSONActionReader is a Reader for the StartMySQLExplainJSONAction structure.
type StartMySQLExplainJSONActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartMySQLExplainJSONActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewStartMySQLExplainJSONActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewStartMySQLExplainJSONActionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartMySQLExplainJSONActionOK creates a StartMySQLExplainJSONActionOK with default headers values
func NewStartMySQLExplainJSONActionOK() *StartMySQLExplainJSONActionOK {
	return &StartMySQLExplainJSONActionOK{}
}

/*StartMySQLExplainJSONActionOK handles this case with default header values.

A successful response.
*/
type StartMySQLExplainJSONActionOK struct {
	Payload *StartMySQLExplainJSONActionOKBody
}

func (o *StartMySQLExplainJSONActionOK) Error() string {
	return fmt.Sprintf("[POST /v0/management/Actions/StartMySQLExplainJSON][%d] startMySqlExplainJsonActionOk  %+v", 200, o.Payload)
}

func (o *StartMySQLExplainJSONActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMySQLExplainJSONActionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartMySQLExplainJSONActionDefault creates a StartMySQLExplainJSONActionDefault with default headers values
func NewStartMySQLExplainJSONActionDefault(code int) *StartMySQLExplainJSONActionDefault {
	return &StartMySQLExplainJSONActionDefault{
		_statusCode: code,
	}
}

/*StartMySQLExplainJSONActionDefault handles this case with default header values.

An error response.
*/
type StartMySQLExplainJSONActionDefault struct {
	_statusCode int

	Payload *StartMySQLExplainJSONActionDefaultBody
}

// Code gets the status code for the start my SQL explain JSON action default response
func (o *StartMySQLExplainJSONActionDefault) Code() int {
	return o._statusCode
}

func (o *StartMySQLExplainJSONActionDefault) Error() string {
	return fmt.Sprintf("[POST /v0/management/Actions/StartMySQLExplainJSON][%d] StartMySQLExplainJSONAction default  %+v", o._statusCode, o.Payload)
}

func (o *StartMySQLExplainJSONActionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(StartMySQLExplainJSONActionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartMySQLExplainJSONActionBody start my SQL explain JSON action body
swagger:model StartMySQLExplainJSONActionBody
*/
type StartMySQLExplainJSONActionBody struct {

	// Database name. Required if it can't be deduced from the query.
	Database string `json:"database,omitempty"`

	// pmm-agent ID where to run this Action.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// SQL query. Required.
	Query string `json:"query,omitempty"`

	// Service ID for this Action. Required.
	ServiceID string `json:"service_id,omitempty"`
}

// Validate validates this start my SQL explain JSON action body
func (o *StartMySQLExplainJSONActionBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainJSONActionBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLExplainJSONActionDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model StartMySQLExplainJSONActionDefaultBody
*/
type StartMySQLExplainJSONActionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this start my SQL explain JSON action default body
func (o *StartMySQLExplainJSONActionDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainJSONActionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartMySQLExplainJSONActionOKBody start my SQL explain JSON action OK body
swagger:model StartMySQLExplainJSONActionOKBody
*/
type StartMySQLExplainJSONActionOKBody struct {

	// Unique Action ID.
	ActionID string `json:"action_id,omitempty"`

	// pmm-agent ID where to this Action was started.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`
}

// Validate validates this start my SQL explain JSON action OK body
func (o *StartMySQLExplainJSONActionOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartMySQLExplainJSONActionOKBody) UnmarshalBinary(b []byte) error {
	var res StartMySQLExplainJSONActionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
