/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package actions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// WeaviateActionsReferencesCreateReader is a Reader for the WeaviateActionsReferencesCreate structure.
type WeaviateActionsReferencesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateActionsReferencesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateActionsReferencesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateActionsReferencesCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateActionsReferencesCreateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateActionsReferencesCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateActionsReferencesCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateActionsReferencesCreateOK creates a WeaviateActionsReferencesCreateOK with default headers values
func NewWeaviateActionsReferencesCreateOK() *WeaviateActionsReferencesCreateOK {
	return &WeaviateActionsReferencesCreateOK{}
}

/*WeaviateActionsReferencesCreateOK handles this case with default header values.

Successfully added the reference.
*/
type WeaviateActionsReferencesCreateOK struct {
}

func (o *WeaviateActionsReferencesCreateOK) Error() string {
	return fmt.Sprintf("[POST /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesCreateOK ", 200)
}

func (o *WeaviateActionsReferencesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsReferencesCreateUnauthorized creates a WeaviateActionsReferencesCreateUnauthorized with default headers values
func NewWeaviateActionsReferencesCreateUnauthorized() *WeaviateActionsReferencesCreateUnauthorized {
	return &WeaviateActionsReferencesCreateUnauthorized{}
}

/*WeaviateActionsReferencesCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateActionsReferencesCreateUnauthorized struct {
}

func (o *WeaviateActionsReferencesCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesCreateUnauthorized ", 401)
}

func (o *WeaviateActionsReferencesCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsReferencesCreateForbidden creates a WeaviateActionsReferencesCreateForbidden with default headers values
func NewWeaviateActionsReferencesCreateForbidden() *WeaviateActionsReferencesCreateForbidden {
	return &WeaviateActionsReferencesCreateForbidden{}
}

/*WeaviateActionsReferencesCreateForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateActionsReferencesCreateForbidden struct {
}

func (o *WeaviateActionsReferencesCreateForbidden) Error() string {
	return fmt.Sprintf("[POST /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesCreateForbidden ", 403)
}

func (o *WeaviateActionsReferencesCreateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateActionsReferencesCreateUnprocessableEntity creates a WeaviateActionsReferencesCreateUnprocessableEntity with default headers values
func NewWeaviateActionsReferencesCreateUnprocessableEntity() *WeaviateActionsReferencesCreateUnprocessableEntity {
	return &WeaviateActionsReferencesCreateUnprocessableEntity{}
}

/*WeaviateActionsReferencesCreateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type WeaviateActionsReferencesCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsReferencesCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateActionsReferencesCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateActionsReferencesCreateInternalServerError creates a WeaviateActionsReferencesCreateInternalServerError with default headers values
func NewWeaviateActionsReferencesCreateInternalServerError() *WeaviateActionsReferencesCreateInternalServerError {
	return &WeaviateActionsReferencesCreateInternalServerError{}
}

/*WeaviateActionsReferencesCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateActionsReferencesCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateActionsReferencesCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /actions/{id}/references/{propertyName}][%d] weaviateActionsReferencesCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateActionsReferencesCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
