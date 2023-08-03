
## `github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/workflowrunactions` Documentation

The `workflowrunactions` SDK allows for interaction with the Azure Resource Manager Service `web` (API Version `2022-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/hashicorp/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/hashicorp/go-azure-sdk/resource-manager/web/2022-09-01/workflowrunactions"
```


### Client Initialization

```go
client := workflowrunactions.NewWorkflowRunActionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkflowRunActionsClient.CopeRepetitionsGet`

```go
ctx := context.TODO()
id := workflowrunactions.NewScopeRepetitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "runValue", "actionValue", "scopeRepetitionValue")

read, err := client.CopeRepetitionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowRunActionsClient.CopeRepetitionsList`

```go
ctx := context.TODO()
id := workflowrunactions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "runValue", "actionValue")

// alternatively `client.CopeRepetitionsList(ctx, id)` can be used to do batched pagination
items, err := client.CopeRepetitionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkflowRunActionsClient.Get`

```go
ctx := context.TODO()
id := workflowrunactions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "runValue", "actionValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowRunActionsClient.List`

```go
ctx := context.TODO()
id := workflowrunactions.NewRunID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "runValue")

// alternatively `client.List(ctx, id, workflowrunactions.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, workflowrunactions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkflowRunActionsClient.ListExpressionTraces`

```go
ctx := context.TODO()
id := workflowrunactions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "runValue", "actionValue")

// alternatively `client.ListExpressionTraces(ctx, id)` can be used to do batched pagination
items, err := client.ListExpressionTracesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkflowRunActionsClient.WorkflowRunActionRepetitionsGet`

```go
ctx := context.TODO()
id := workflowrunactions.NewRepetitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "runValue", "actionValue", "repetitionValue")

read, err := client.WorkflowRunActionRepetitionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowRunActionsClient.WorkflowRunActionRepetitionsList`

```go
ctx := context.TODO()
id := workflowrunactions.NewActionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "runValue", "actionValue")

// alternatively `client.WorkflowRunActionRepetitionsList(ctx, id)` can be used to do batched pagination
items, err := client.WorkflowRunActionRepetitionsListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkflowRunActionsClient.WorkflowRunActionRepetitionsListExpressionTraces`

```go
ctx := context.TODO()
id := workflowrunactions.NewRepetitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "runValue", "actionValue", "repetitionValue")

// alternatively `client.WorkflowRunActionRepetitionsListExpressionTraces(ctx, id)` can be used to do batched pagination
items, err := client.WorkflowRunActionRepetitionsListExpressionTracesComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkflowRunActionsClient.WorkflowRunActionRepetitionsRequestHistoriesGet`

```go
ctx := context.TODO()
id := workflowrunactions.NewRequestHistoryID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "runValue", "actionValue", "repetitionValue", "requestHistoryValue")

read, err := client.WorkflowRunActionRepetitionsRequestHistoriesGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkflowRunActionsClient.WorkflowRunActionRepetitionsRequestHistoriesList`

```go
ctx := context.TODO()
id := workflowrunactions.NewRepetitionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "siteValue", "workflowValue", "runValue", "actionValue", "repetitionValue")

// alternatively `client.WorkflowRunActionRepetitionsRequestHistoriesList(ctx, id)` can be used to do batched pagination
items, err := client.WorkflowRunActionRepetitionsRequestHistoriesListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```