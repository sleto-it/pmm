// Code generated by go-swagger; DO NOT EDIT.

package backups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StartBackupReader is a Reader for the StartBackup structure.
type StartBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StartBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStartBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewStartBackupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewStartBackupOK creates a StartBackupOK with default headers values
func NewStartBackupOK() *StartBackupOK {
	return &StartBackupOK{}
}

/* StartBackupOK describes a response with status code 200, with default header values.

A successful response.
*/
type StartBackupOK struct {
	Payload *StartBackupOKBody
}

func (o *StartBackupOK) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/Start][%d] startBackupOk  %+v", 200, o.Payload)
}

func (o *StartBackupOK) GetPayload() *StartBackupOKBody {
	return o.Payload
}

func (o *StartBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartBackupOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStartBackupDefault creates a StartBackupDefault with default headers values
func NewStartBackupDefault(code int) *StartBackupDefault {
	return &StartBackupDefault{
		_statusCode: code,
	}
}

/* StartBackupDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type StartBackupDefault struct {
	_statusCode int

	Payload *StartBackupDefaultBody
}

// Code gets the status code for the start backup default response
func (o *StartBackupDefault) Code() int {
	return o._statusCode
}

func (o *StartBackupDefault) Error() string {
	return fmt.Sprintf("[POST /v1/management/backup/Backups/Start][%d] StartBackup default  %+v", o._statusCode, o.Payload)
}

func (o *StartBackupDefault) GetPayload() *StartBackupDefaultBody {
	return o.Payload
}

func (o *StartBackupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(StartBackupDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*StartBackupBody start backup body
swagger:model StartBackupBody
*/
type StartBackupBody struct {
	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// Machine-readable location ID.
	LocationID string `json:"location_id,omitempty"`

	// If empty then name is auto-generated.
	Name string `json:"name,omitempty"`

	// Human-readable description.
	Description string `json:"description,omitempty"`

	// Delay between each retry. Should have a suffix in JSON: 1s, 1m, 1h.
	RetryInterval string `json:"retry_interval,omitempty"`

	// How many times to retry a failed backup before giving up.
	Retries int64 `json:"retries,omitempty"`

	// DataModel is a model used for performing a backup.
	// Enum: [DATA_MODEL_INVALID PHYSICAL LOGICAL]
	DataModel *string `json:"data_model,omitempty"`
}

// Validate validates this start backup body
func (o *StartBackupBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDataModel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var startBackupBodyTypeDataModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DATA_MODEL_INVALID","PHYSICAL","LOGICAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		startBackupBodyTypeDataModelPropEnum = append(startBackupBodyTypeDataModelPropEnum, v)
	}
}

const (

	// StartBackupBodyDataModelDATAMODELINVALID captures enum value "DATA_MODEL_INVALID"
	StartBackupBodyDataModelDATAMODELINVALID string = "DATA_MODEL_INVALID"

	// StartBackupBodyDataModelPHYSICAL captures enum value "PHYSICAL"
	StartBackupBodyDataModelPHYSICAL string = "PHYSICAL"

	// StartBackupBodyDataModelLOGICAL captures enum value "LOGICAL"
	StartBackupBodyDataModelLOGICAL string = "LOGICAL"
)

// prop value enum
func (o *StartBackupBody) validateDataModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, startBackupBodyTypeDataModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *StartBackupBody) validateDataModel(formats strfmt.Registry) error {
	if swag.IsZero(o.DataModel) { // not required
		return nil
	}

	// value enum
	if err := o.validateDataModelEnum("body"+"."+"data_model", "body", *o.DataModel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this start backup body based on context it is used
func (o *StartBackupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartBackupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartBackupBody) UnmarshalBinary(b []byte) error {
	var res StartBackupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartBackupDefaultBody start backup default body
swagger:model StartBackupDefaultBody
*/
type StartBackupDefaultBody struct {
	// code
	Code int32 `json:"code,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// details
	Details []*StartBackupDefaultBodyDetailsItems0 `json:"details"`
}

// Validate validates this start backup default body
func (o *StartBackupDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartBackupDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartBackup default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartBackup default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this start backup default body based on the context it is used
func (o *StartBackupDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *StartBackupDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {
	for i := 0; i < len(o.Details); i++ {
		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("StartBackup default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("StartBackup default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *StartBackupDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartBackupDefaultBody) UnmarshalBinary(b []byte) error {
	var res StartBackupDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartBackupDefaultBodyDetailsItems0 start backup default body details items0
swagger:model StartBackupDefaultBodyDetailsItems0
*/
type StartBackupDefaultBodyDetailsItems0 struct {
	// at type
	AtType string `json:"@type,omitempty"`
}

// Validate validates this start backup default body details items0
func (o *StartBackupDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start backup default body details items0 based on context it is used
func (o *StartBackupDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartBackupDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartBackupDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res StartBackupDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*StartBackupOKBody start backup OK body
swagger:model StartBackupOKBody
*/
type StartBackupOKBody struct {
	// Unique identifier.
	ArtifactID string `json:"artifact_id,omitempty"`
}

// Validate validates this start backup OK body
func (o *StartBackupOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this start backup OK body based on context it is used
func (o *StartBackupOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *StartBackupOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *StartBackupOKBody) UnmarshalBinary(b []byte) error {
	var res StartBackupOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
