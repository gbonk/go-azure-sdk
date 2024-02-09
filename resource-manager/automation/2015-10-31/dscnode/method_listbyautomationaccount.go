package dscnode

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByAutomationAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DscNode
}

type ListByAutomationAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DscNode
}

type ListByAutomationAccountOperationOptions struct {
	Filter *string
}

func DefaultListByAutomationAccountOperationOptions() ListByAutomationAccountOperationOptions {
	return ListByAutomationAccountOperationOptions{}
}

func (o ListByAutomationAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByAutomationAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByAutomationAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListByAutomationAccount ...
func (c DscNodeClient) ListByAutomationAccount(ctx context.Context, id AutomationAccountId, options ListByAutomationAccountOperationOptions) (result ListByAutomationAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/nodes", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]DscNode `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByAutomationAccountComplete retrieves all the results into a single object
func (c DscNodeClient) ListByAutomationAccountComplete(ctx context.Context, id AutomationAccountId, options ListByAutomationAccountOperationOptions) (ListByAutomationAccountCompleteResult, error) {
	return c.ListByAutomationAccountCompleteMatchingPredicate(ctx, id, options, DscNodeOperationPredicate{})
}

// ListByAutomationAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DscNodeClient) ListByAutomationAccountCompleteMatchingPredicate(ctx context.Context, id AutomationAccountId, options ListByAutomationAccountOperationOptions, predicate DscNodeOperationPredicate) (result ListByAutomationAccountCompleteResult, err error) {
	items := make([]DscNode, 0)

	resp, err := c.ListByAutomationAccount(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListByAutomationAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
