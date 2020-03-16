// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/innovationnorway/go-databricks/models"
)

// RestartReader is a Reader for the Restart structure.
type RestartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestartOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRestartBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewRestartDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestartOK creates a RestartOK with default headers values
func NewRestartOK() *RestartOK {
	return &RestartOK{}
}

/*RestartOK handles this case with default header values.

OK
*/
type RestartOK struct {
}

func (o *RestartOK) Error() string {
	return fmt.Sprintf("[POST /clusters/restart][%d] restartOK ", 200)
}

func (o *RestartOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRestartBadRequest creates a RestartBadRequest with default headers values
func NewRestartBadRequest() *RestartBadRequest {
	return &RestartBadRequest{}
}

/*RestartBadRequest handles this case with default header values.

Error
*/
type RestartBadRequest struct {
	Payload *models.Error
}

func (o *RestartBadRequest) Error() string {
	return fmt.Sprintf("[POST /clusters/restart][%d] restartBadRequest  %+v", 400, o.Payload)
}

func (o *RestartBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RestartBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestartDefault creates a RestartDefault with default headers values
func NewRestartDefault(code int) *RestartDefault {
	return &RestartDefault{
		_statusCode: code,
	}
}

/*RestartDefault handles this case with default header values.

Default
*/
type RestartDefault struct {
	_statusCode int

	Payload string
}

// Code gets the status code for the restart default response
func (o *RestartDefault) Code() int {
	return o._statusCode
}

func (o *RestartDefault) Error() string {
	return fmt.Sprintf("[POST /clusters/restart][%d] restart default  %+v", o._statusCode, o.Payload)
}

func (o *RestartDefault) GetPayload() string {
	return o.Payload
}

func (o *RestartDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*RestartBody restart body
swagger:model RestartBody
*/
type RestartBody struct {

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`
}

// Validate validates this restart body
func (o *RestartBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *RestartBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *RestartBody) UnmarshalBinary(b []byte) error {
	var res RestartBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}