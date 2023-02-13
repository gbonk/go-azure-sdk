package job

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetOutputOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]byte
}

// GetOutput ...
func (c JobClient) GetOutput(ctx context.Context, id JobId) (result GetOutputOperationResponse, err error) {
	req, err := c.preparerForGetOutput(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "GetOutput", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "GetOutput", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetOutput(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "job.JobClient", "GetOutput", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetOutput prepares the GetOutput request.
func (c JobClient) preparerForGetOutput(ctx context.Context, id JobId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/output", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetOutput handles the response to the GetOutput request. The method always
// closes the http.Response Body.
func (c JobClient) responderForGetOutput(resp *http.Response) (result GetOutputOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}