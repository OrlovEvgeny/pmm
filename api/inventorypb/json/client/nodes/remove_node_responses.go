// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// RemoveNodeReader is a Reader for the RemoveNode structure.
type RemoveNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveNodeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveNodeOK creates a RemoveNodeOK with default headers values
func NewRemoveNodeOK() *RemoveNodeOK {
	return &RemoveNodeOK{}
}

/*RemoveNodeOK handles this case with default header values.

A successful response.
*/
type RemoveNodeOK struct {
	Payload interface{}
}

func (o *RemoveNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/Remove][%d] removeNodeOk  %+v", 200, o.Payload)
}

func (o *RemoveNodeOK) GetPayload() interface{} {
	return o.Payload
}

func (o *RemoveNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveNodeDefault creates a RemoveNodeDefault with default headers values
func NewRemoveNodeDefault(code int) *RemoveNodeDefault {
	return &RemoveNodeDefault{
		_statusCode: code,
	}
}

/*RemoveNodeDefault handles this case with default header values.

An error response.
*/
type RemoveNodeDefault struct {
	_statusCode int

	Payload *RemoveNodeDefaultBody
}

// Code gets the status code for the remove node default response
func (o *RemoveNodeDefault) Code() int {
	return o._statusCode
}

func (o *RemoveNodeDefault) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/Remove][%d] RemoveNode default  %+v", o._statusCode, o.Payload)
}

func (o *RemoveNodeDefault) GetPayload() *RemoveNodeDefaultBody {
	return o.Payload
}

func (o *RemoveNodeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(RemoveNodeDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RemoveNodeBody remove node body
swagger:model RemoveNodeBody
*/
type RemoveNodeBody struct {

	// Unique randomly generated instance identifier.
	NodeID string `json:"node_id,omitempty"`

	// Remove node with all dependencies.
	Force bool `json:"force,omitempty"`
}

// Validate validates this remove node body
func (o *RemoveNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveNodeBody) UnmarshalBinary(b []byte) error {
	var res RemoveNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*RemoveNodeDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model RemoveNodeDefaultBody
*/
type RemoveNodeDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this remove node default body
func (o *RemoveNodeDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RemoveNodeDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RemoveNodeDefaultBody) UnmarshalBinary(b []byte) error {
	var res RemoveNodeDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
