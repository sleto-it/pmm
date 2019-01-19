// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// AddRemoteNodeReader is a Reader for the AddRemoteNode structure.
type AddRemoteNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRemoteNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddRemoteNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddRemoteNodeOK creates a AddRemoteNodeOK with default headers values
func NewAddRemoteNodeOK() *AddRemoteNodeOK {
	return &AddRemoteNodeOK{}
}

/*AddRemoteNodeOK handles this case with default header values.

(empty)
*/
type AddRemoteNodeOK struct {
	Payload *AddRemoteNodeOKBody
}

func (o *AddRemoteNodeOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Nodes/AddRemote][%d] addRemoteNodeOK  %+v", 200, o.Payload)
}

func (o *AddRemoteNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddRemoteNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddRemoteNodeBody add remote node body
swagger:model AddRemoteNodeBody
*/
type AddRemoteNodeBody struct {

	// Unique Node identifier. Will be generated if empty.
	ID string `json:"id,omitempty"`

	// Unique user-defined Node name.
	Name string `json:"name,omitempty"`
}

// Validate validates this add remote node body
func (o *AddRemoteNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteNodeBody) UnmarshalBinary(b []byte) error {
	var res AddRemoteNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRemoteNodeOKBody add remote node o k body
swagger:model AddRemoteNodeOKBody
*/
type AddRemoteNodeOKBody struct {

	// remote
	Remote *AddRemoteNodeOKBodyRemote `json:"remote,omitempty"`
}

// Validate validates this add remote node o k body
func (o *AddRemoteNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRemote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddRemoteNodeOKBody) validateRemote(formats strfmt.Registry) error {

	if swag.IsZero(o.Remote) { // not required
		return nil
	}

	if o.Remote != nil {
		if err := o.Remote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addRemoteNodeOK" + "." + "remote")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteNodeOKBody) UnmarshalBinary(b []byte) error {
	var res AddRemoteNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddRemoteNodeOKBodyRemote RemoteNode represents a generic remote Node.
// Agents can't be run on remote Nodes.
swagger:model AddRemoteNodeOKBodyRemote
*/
type AddRemoteNodeOKBodyRemote struct {

	// Linux distribution. May be empty.
	Distro string `json:"distro,omitempty"`

	// Linux distribution version. May be empty.
	DistroVersion string `json:"distro_version,omitempty"`

	// Hostname. Is not unique. May be empty.
	Hostname string `json:"hostname,omitempty"`

	// Unique Node identifier.
	ID string `json:"id,omitempty"`

	// Unique user-defined Node name.
	Name string `json:"name,omitempty"`
}

// Validate validates this add remote node o k body remote
func (o *AddRemoteNodeOKBodyRemote) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddRemoteNodeOKBodyRemote) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddRemoteNodeOKBodyRemote) UnmarshalBinary(b []byte) error {
	var res AddRemoteNodeOKBodyRemote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
