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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// WeaviateSchemaThingsCreateReader is a Reader for the WeaviateSchemaThingsCreate structure.
type WeaviateSchemaThingsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateSchemaThingsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateSchemaThingsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewWeaviateSchemaThingsCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewWeaviateSchemaThingsCreateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateSchemaThingsCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateSchemaThingsCreateOK creates a WeaviateSchemaThingsCreateOK with default headers values
func NewWeaviateSchemaThingsCreateOK() *WeaviateSchemaThingsCreateOK {
	return &WeaviateSchemaThingsCreateOK{}
}

/*WeaviateSchemaThingsCreateOK handles this case with default header values.

Added the new Thing class to the ontology.
*/
type WeaviateSchemaThingsCreateOK struct {
	Payload *models.SemanticSchemaClass
}

func (o *WeaviateSchemaThingsCreateOK) Error() string {
	return fmt.Sprintf("[POST /schema/things][%d] weaviateSchemaThingsCreateOK  %+v", 200, o.Payload)
}

func (o *WeaviateSchemaThingsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SemanticSchemaClass)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaThingsCreateUnauthorized creates a WeaviateSchemaThingsCreateUnauthorized with default headers values
func NewWeaviateSchemaThingsCreateUnauthorized() *WeaviateSchemaThingsCreateUnauthorized {
	return &WeaviateSchemaThingsCreateUnauthorized{}
}

/*WeaviateSchemaThingsCreateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateSchemaThingsCreateUnauthorized struct {
}

func (o *WeaviateSchemaThingsCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schema/things][%d] weaviateSchemaThingsCreateUnauthorized ", 401)
}

func (o *WeaviateSchemaThingsCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateSchemaThingsCreateUnprocessableEntity creates a WeaviateSchemaThingsCreateUnprocessableEntity with default headers values
func NewWeaviateSchemaThingsCreateUnprocessableEntity() *WeaviateSchemaThingsCreateUnprocessableEntity {
	return &WeaviateSchemaThingsCreateUnprocessableEntity{}
}

/*WeaviateSchemaThingsCreateUnprocessableEntity handles this case with default header values.

Invalid Thing class.
*/
type WeaviateSchemaThingsCreateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsCreateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /schema/things][%d] weaviateSchemaThingsCreateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *WeaviateSchemaThingsCreateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateSchemaThingsCreateInternalServerError creates a WeaviateSchemaThingsCreateInternalServerError with default headers values
func NewWeaviateSchemaThingsCreateInternalServerError() *WeaviateSchemaThingsCreateInternalServerError {
	return &WeaviateSchemaThingsCreateInternalServerError{}
}

/*WeaviateSchemaThingsCreateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateSchemaThingsCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateSchemaThingsCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /schema/things][%d] weaviateSchemaThingsCreateInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateSchemaThingsCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
