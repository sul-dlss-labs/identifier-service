// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/sul-dlss-labs/identifier-service/generated/models"
)

// GetIdentifiersInfoReader is a Reader for the GetIdentifiersInfo structure.
type GetIdentifiersInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentifiersInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIdentifiersInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetIdentifiersInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIdentifiersInfoOK creates a GetIdentifiersInfoOK with default headers values
func NewGetIdentifiersInfoOK() *GetIdentifiersInfoOK {
	return &GetIdentifiersInfoOK{}
}

/*GetIdentifiersInfoOK handles this case with default header values.

Get a list of identifier minters supported by this API.
*/
type GetIdentifiersInfoOK struct {
	Payload models.Sources
}

func (o *GetIdentifiersInfoOK) Error() string {
	return fmt.Sprintf("[GET /identifiers][%d] getIdentifiersInfoOK  %+v", 200, o.Payload)
}

func (o *GetIdentifiersInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIdentifiersInfoInternalServerError creates a GetIdentifiersInfoInternalServerError with default headers values
func NewGetIdentifiersInfoInternalServerError() *GetIdentifiersInfoInternalServerError {
	return &GetIdentifiersInfoInternalServerError{}
}

/*GetIdentifiersInfoInternalServerError handles this case with default header values.

Your request could not be processed at this time.
*/
type GetIdentifiersInfoInternalServerError struct {
}

func (o *GetIdentifiersInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /identifiers][%d] getIdentifiersInfoInternalServerError ", 500)
}

func (o *GetIdentifiersInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
