// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// ListVolumeAttachmentsRequest wrapper for the ListVolumeAttachments operation
type ListVolumeAttachmentsRequest struct {

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" contributesTo:"query" name:"compartmentId"`

	// The name of the Availability Domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" contributesTo:"query" name:"availabilityDomain"`

	// The maximum number of items to return in a paginated "List" call.
	// Example: `500`
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The value of the `opc-next-page` response header from the previous "List" call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The OCID of the instance.
	InstanceId *string `mandatory:"false" contributesTo:"query" name:"instanceId"`

	// The OCID of the volume.
	VolumeId *string `mandatory:"false" contributesTo:"query" name:"volumeId"`
}

func (request ListVolumeAttachmentsRequest) String() string {
	return common.PointerString(request)
}

// ListVolumeAttachmentsResponse wrapper for the ListVolumeAttachments operation
type ListVolumeAttachmentsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The []VolumeAttachment instance
	Items []VolumeAttachment `presentIn:"body"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then a partial list might have been returned. Include this value as the `page` parameter for the
	// subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response ListVolumeAttachmentsResponse) String() string {
	return common.PointerString(response)
}
