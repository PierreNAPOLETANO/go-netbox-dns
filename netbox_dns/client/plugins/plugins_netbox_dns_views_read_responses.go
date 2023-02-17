// Code generated by go-swagger; DO NOT EDIT.

package plugins

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ProtonMail/go-netbox-dns/netbox_dns/models"
)

// PluginsNetboxDNSViewsReadReader is a Reader for the PluginsNetboxDNSViewsRead structure.
type PluginsNetboxDNSViewsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginsNetboxDNSViewsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginsNetboxDNSViewsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPluginsNetboxDNSViewsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPluginsNetboxDNSViewsReadOK creates a PluginsNetboxDNSViewsReadOK with default headers values
func NewPluginsNetboxDNSViewsReadOK() *PluginsNetboxDNSViewsReadOK {
	return &PluginsNetboxDNSViewsReadOK{}
}

/*
PluginsNetboxDNSViewsReadOK describes a response with status code 200, with default header values.

PluginsNetboxDNSViewsReadOK plugins netbox Dns views read o k
*/
type PluginsNetboxDNSViewsReadOK struct {
	Payload *models.View
}

// IsSuccess returns true when this plugins netbox Dns views read o k response has a 2xx status code
func (o *PluginsNetboxDNSViewsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plugins netbox Dns views read o k response has a 3xx status code
func (o *PluginsNetboxDNSViewsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plugins netbox Dns views read o k response has a 4xx status code
func (o *PluginsNetboxDNSViewsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plugins netbox Dns views read o k response has a 5xx status code
func (o *PluginsNetboxDNSViewsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plugins netbox Dns views read o k response a status code equal to that given
func (o *PluginsNetboxDNSViewsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plugins netbox Dns views read o k response
func (o *PluginsNetboxDNSViewsReadOK) Code() int {
	return 200
}

func (o *PluginsNetboxDNSViewsReadOK) Error() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/views/{id}/][%d] pluginsNetboxDnsViewsReadOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSViewsReadOK) String() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/views/{id}/][%d] pluginsNetboxDnsViewsReadOK  %+v", 200, o.Payload)
}

func (o *PluginsNetboxDNSViewsReadOK) GetPayload() *models.View {
	return o.Payload
}

func (o *PluginsNetboxDNSViewsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.View)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginsNetboxDNSViewsReadDefault creates a PluginsNetboxDNSViewsReadDefault with default headers values
func NewPluginsNetboxDNSViewsReadDefault(code int) *PluginsNetboxDNSViewsReadDefault {
	return &PluginsNetboxDNSViewsReadDefault{
		_statusCode: code,
	}
}

/*
PluginsNetboxDNSViewsReadDefault describes a response with status code -1, with default header values.

PluginsNetboxDNSViewsReadDefault plugins netbox dns views read default
*/
type PluginsNetboxDNSViewsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this plugins netbox dns views read default response has a 2xx status code
func (o *PluginsNetboxDNSViewsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plugins netbox dns views read default response has a 3xx status code
func (o *PluginsNetboxDNSViewsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plugins netbox dns views read default response has a 4xx status code
func (o *PluginsNetboxDNSViewsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plugins netbox dns views read default response has a 5xx status code
func (o *PluginsNetboxDNSViewsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plugins netbox dns views read default response a status code equal to that given
func (o *PluginsNetboxDNSViewsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plugins netbox dns views read default response
func (o *PluginsNetboxDNSViewsReadDefault) Code() int {
	return o._statusCode
}

func (o *PluginsNetboxDNSViewsReadDefault) Error() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/views/{id}/][%d] plugins_netbox-dns_views_read default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSViewsReadDefault) String() string {
	return fmt.Sprintf("[GET /plugins/netbox-dns/views/{id}/][%d] plugins_netbox-dns_views_read default  %+v", o._statusCode, o.Payload)
}

func (o *PluginsNetboxDNSViewsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *PluginsNetboxDNSViewsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}