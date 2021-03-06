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

package contextionary_api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/entities/models"
)

// WeaviateC11yWordsReader is a Reader for the WeaviateC11yWords structure.
type WeaviateC11yWordsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateC11yWordsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewWeaviateC11yWordsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewWeaviateC11yWordsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewWeaviateC11yWordsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewWeaviateC11yWordsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewWeaviateC11yWordsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 501:
		result := NewWeaviateC11yWordsNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateC11yWordsOK creates a WeaviateC11yWordsOK with default headers values
func NewWeaviateC11yWordsOK() *WeaviateC11yWordsOK {
	return &WeaviateC11yWordsOK{}
}

/*WeaviateC11yWordsOK handles this case with default header values.

Successful response.
*/
type WeaviateC11yWordsOK struct {
	Payload *models.C11yWordsResponse
}

func (o *WeaviateC11yWordsOK) Error() string {
	return fmt.Sprintf("[GET /c11y/words/{words}][%d] weaviateC11yWordsOK  %+v", 200, o.Payload)
}

func (o *WeaviateC11yWordsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.C11yWordsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateC11yWordsBadRequest creates a WeaviateC11yWordsBadRequest with default headers values
func NewWeaviateC11yWordsBadRequest() *WeaviateC11yWordsBadRequest {
	return &WeaviateC11yWordsBadRequest{}
}

/*WeaviateC11yWordsBadRequest handles this case with default header values.

Incorrect request
*/
type WeaviateC11yWordsBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateC11yWordsBadRequest) Error() string {
	return fmt.Sprintf("[GET /c11y/words/{words}][%d] weaviateC11yWordsBadRequest  %+v", 400, o.Payload)
}

func (o *WeaviateC11yWordsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateC11yWordsUnauthorized creates a WeaviateC11yWordsUnauthorized with default headers values
func NewWeaviateC11yWordsUnauthorized() *WeaviateC11yWordsUnauthorized {
	return &WeaviateC11yWordsUnauthorized{}
}

/*WeaviateC11yWordsUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type WeaviateC11yWordsUnauthorized struct {
}

func (o *WeaviateC11yWordsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /c11y/words/{words}][%d] weaviateC11yWordsUnauthorized ", 401)
}

func (o *WeaviateC11yWordsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateC11yWordsForbidden creates a WeaviateC11yWordsForbidden with default headers values
func NewWeaviateC11yWordsForbidden() *WeaviateC11yWordsForbidden {
	return &WeaviateC11yWordsForbidden{}
}

/*WeaviateC11yWordsForbidden handles this case with default header values.

Insufficient permissions.
*/
type WeaviateC11yWordsForbidden struct {
}

func (o *WeaviateC11yWordsForbidden) Error() string {
	return fmt.Sprintf("[GET /c11y/words/{words}][%d] weaviateC11yWordsForbidden ", 403)
}

func (o *WeaviateC11yWordsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateC11yWordsInternalServerError creates a WeaviateC11yWordsInternalServerError with default headers values
func NewWeaviateC11yWordsInternalServerError() *WeaviateC11yWordsInternalServerError {
	return &WeaviateC11yWordsInternalServerError{}
}

/*WeaviateC11yWordsInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type WeaviateC11yWordsInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *WeaviateC11yWordsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /c11y/words/{words}][%d] weaviateC11yWordsInternalServerError  %+v", 500, o.Payload)
}

func (o *WeaviateC11yWordsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWeaviateC11yWordsNotImplemented creates a WeaviateC11yWordsNotImplemented with default headers values
func NewWeaviateC11yWordsNotImplemented() *WeaviateC11yWordsNotImplemented {
	return &WeaviateC11yWordsNotImplemented{}
}

/*WeaviateC11yWordsNotImplemented handles this case with default header values.

Not (yet) implemented.
*/
type WeaviateC11yWordsNotImplemented struct {
}

func (o *WeaviateC11yWordsNotImplemented) Error() string {
	return fmt.Sprintf("[GET /c11y/words/{words}][%d] weaviateC11yWordsNotImplemented ", 501)
}

func (o *WeaviateC11yWordsNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
