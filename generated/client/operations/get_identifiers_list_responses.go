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

// GetIdentifiersListReader is a Reader for the GetIdentifiersList structure.
type GetIdentifiersListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIdentifiersListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIdentifiersListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIdentifiersListOK creates a GetIdentifiersListOK with default headers values
func NewGetIdentifiersListOK() *GetIdentifiersListOK {
	return &GetIdentifiersListOK{}
}

/*GetIdentifiersListOK handles this case with default header values.

Get a list of all Identifiers minted across types
*/
type GetIdentifiersListOK struct {
	Payload models.Sources
}

func (o *GetIdentifiersListOK) Error() string {
	return fmt.Sprintf("[GET /identifiers/all][%d] getIdentifiersListOK  %+v", 200, o.Payload)
}

func (o *GetIdentifiersListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
