// Code generated by go-swagger; DO NOT EDIT.

package pet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AddPetReader is a Reader for the AddPet structure.
type AddPetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddPetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 405:
		result := NewAddPetMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddPetMethodNotAllowed creates a AddPetMethodNotAllowed with default headers values
func NewAddPetMethodNotAllowed() *AddPetMethodNotAllowed {
	return &AddPetMethodNotAllowed{}
}

/*AddPetMethodNotAllowed handles this case with default header values.

Invalid input
*/
type AddPetMethodNotAllowed struct {
}

func (o *AddPetMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /pet][%d] addPetMethodNotAllowed ", 405)
}

func (o *AddPetMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
