package azblob

// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"fmt"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io"
	"net/http"
	"net/url"
	"time"
)

// appendBlobsClient is the client for the AppendBlobs methods of the Azblob service.
type appendBlobsClient struct {
	managementClient
}

// newAppendBlobsClient creates an instance of the appendBlobsClient client.
func newAppendBlobsClient(url url.URL, p pipeline.Pipeline) appendBlobsClient {
	return appendBlobsClient{newManagementClient(url, p)}
}

// AppendBlock the Append Block operation commits a new block of data to the end of an existing append blob. The Append
// Block operation is permitted only if the blob was created with x-ms-blob-type set to AppendBlob. Append Block is
// supported only on version 2015-02-21 version or later.
//
// body is initial data body will be closed upon successful return. Callers should ensure closure when receiving an
// error.timeout is the timeout parameter is expressed in seconds. For more information, see <a
// href="https://docs.microsoft.com/en-us/rest/api/storageservices/fileservices/setting-timeouts-for-blob-service-operations">Setting
// Timeouts for Blob Service Operations.</a> leaseID is if specified, the operation only succeeds if the container's
// lease is active and matches this ID. maxSize is optional conditional header. The max length in bytes permitted for
// the append blob. If the Append Block operation would cause the blob to exceed that limit or if the blob size is
// already greater than the value specified in this header, the request will fail with MaxBlobSizeConditionNotMet error
// (HTTP status code 412 - Precondition Failed). appendPosition is optional conditional header, used only for the
// Append Block operation. A number indicating the byte offset to compare. Append Block will succeed only if the append
// position is equal to this number. If it is not, the request will fail with the AppendPositionConditionNotMet error
// (HTTP status code 412 - Precondition Failed). ifModifiedSince is specify this header value to operate only on a blob
// if it has been modified since the specified date/time. ifUnmodifiedSince is specify this header value to operate
// only on a blob if it has not been modified since the specified date/time. ifMatches is specify an ETag value to
// operate only on blobs with a matching value. ifNoneMatch is specify an ETag value to operate only on blobs without a
// matching value. requestID is provides a client-generated, opaque value with a 1 KB character limit that is recorded
// in the analytics logs when storage analytics logging is enabled.
func (client appendBlobsClient) AppendBlock(ctx context.Context, body io.ReadSeeker, timeout *int32, leaseID *string, maxSize *int32, appendPosition *int32, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatches *ETag, ifNoneMatch *ETag, requestID *string) (*AppendBlobsAppendBlockResponse, error) {
	if err := validate([]validation{
		{targetValue: body,
			constraints: []constraint{{target: "body", name: null, rule: true, chain: nil}}},
		{targetValue: timeout,
			constraints: []constraint{{target: "timeout", name: null, rule: false,
				chain: []constraint{{target: "timeout", name: inclusiveMinimum, rule: 0, chain: nil}}}}}}); err != nil {
		return nil, err
	}
	req, err := client.appendBlockPreparer(body, timeout, leaseID, maxSize, appendPosition, ifModifiedSince, ifUnmodifiedSince, ifMatches, ifNoneMatch, requestID)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.appendBlockResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*AppendBlobsAppendBlockResponse), err
}

// appendBlockPreparer prepares the AppendBlock request.
func (client appendBlobsClient) appendBlockPreparer(body io.ReadSeeker, timeout *int32, leaseID *string, maxSize *int32, appendPosition *int32, ifModifiedSince *time.Time, ifUnmodifiedSince *time.Time, ifMatches *ETag, ifNoneMatch *ETag, requestID *string) (pipeline.Request, error) {
	req, err := pipeline.NewRequest("PUT", client.url, body)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if timeout != nil {
		params.Set("timeout", fmt.Sprintf("%v", *timeout))
	}
	params.Set("comp", "appendblock")
	req.URL.RawQuery = params.Encode()
	if leaseID != nil {
		req.Header.Set("x-ms-lease-id", *leaseID)
	}
	if maxSize != nil {
		req.Header.Set("x-ms-blob-condition-maxsize", fmt.Sprintf("%v", *maxSize))
	}
	if appendPosition != nil {
		req.Header.Set("x-ms-blob-condition-appendpos", fmt.Sprintf("%v", *appendPosition))
	}
	if ifModifiedSince != nil {
		req.Header.Set("If-Modified-Since", (*ifModifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifUnmodifiedSince != nil {
		req.Header.Set("If-Unmodified-Since", (*ifUnmodifiedSince).In(gmt).Format(time.RFC1123))
	}
	if ifMatches != nil {
		req.Header.Set("If-Match", string(*ifMatches))
	}
	if ifNoneMatch != nil {
		req.Header.Set("If-None-Match", string(*ifNoneMatch))
	}
	req.Header.Set("x-ms-version", ServiceVersion)
	if requestID != nil {
		req.Header.Set("x-ms-client-request-id", *requestID)
	}
	return req, nil
}

// appendBlockResponder handles the response to the AppendBlock request.
func (client appendBlobsClient) appendBlockResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK, http.StatusCreated)
	if resp == nil {
		return nil, err
	}
	return &AppendBlobsAppendBlockResponse{rawResponse: resp.Response()}, err
}