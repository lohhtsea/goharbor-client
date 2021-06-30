// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v4/apiv2/model"
)

// GetAdditionReader is a Reader for the GetAddition structure.
type GetAdditionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAdditionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAdditionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAdditionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAdditionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAdditionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAdditionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAdditionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAdditionOK creates a GetAdditionOK with default headers values
func NewGetAdditionOK() *GetAdditionOK {
	return &GetAdditionOK{}
}

/* GetAdditionOK describes a response with status code 200, with default header values.

Success
*/
type GetAdditionOK struct {

	/* The content type of the addition
	 */
	ContentType string

	Payload string
}

func (o *GetAdditionOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/{addition}][%d] getAdditionOK  %+v", 200, o.Payload)
}
func (o *GetAdditionOK) GetPayload() string {
	return o.Payload
}

func (o *GetAdditionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Content-Type
	hdrContentType := response.GetHeader("Content-Type")

	if hdrContentType != "" {
		o.ContentType = hdrContentType
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdditionBadRequest creates a GetAdditionBadRequest with default headers values
func NewGetAdditionBadRequest() *GetAdditionBadRequest {
	return &GetAdditionBadRequest{}
}

/* GetAdditionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetAdditionBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetAdditionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/{addition}][%d] getAdditionBadRequest  %+v", 400, o.Payload)
}
func (o *GetAdditionBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetAdditionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdditionUnauthorized creates a GetAdditionUnauthorized with default headers values
func NewGetAdditionUnauthorized() *GetAdditionUnauthorized {
	return &GetAdditionUnauthorized{}
}

/* GetAdditionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetAdditionUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetAdditionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/{addition}][%d] getAdditionUnauthorized  %+v", 401, o.Payload)
}
func (o *GetAdditionUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetAdditionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdditionForbidden creates a GetAdditionForbidden with default headers values
func NewGetAdditionForbidden() *GetAdditionForbidden {
	return &GetAdditionForbidden{}
}

/* GetAdditionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetAdditionForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetAdditionForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/{addition}][%d] getAdditionForbidden  %+v", 403, o.Payload)
}
func (o *GetAdditionForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetAdditionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdditionNotFound creates a GetAdditionNotFound with default headers values
func NewGetAdditionNotFound() *GetAdditionNotFound {
	return &GetAdditionNotFound{}
}

/* GetAdditionNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetAdditionNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetAdditionNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/{addition}][%d] getAdditionNotFound  %+v", 404, o.Payload)
}
func (o *GetAdditionNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetAdditionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAdditionInternalServerError creates a GetAdditionInternalServerError with default headers values
func NewGetAdditionInternalServerError() *GetAdditionInternalServerError {
	return &GetAdditionInternalServerError{}
}

/* GetAdditionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetAdditionInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetAdditionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/additions/{addition}][%d] getAdditionInternalServerError  %+v", 500, o.Payload)
}
func (o *GetAdditionInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetAdditionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
