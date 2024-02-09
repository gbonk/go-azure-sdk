package dscnode

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DscNodeClient struct {
	Client *resourcemanager.Client
}

func NewDscNodeClientWithBaseURI(sdkApi sdkEnv.Api) (*DscNodeClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "dscnode", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DscNodeClient: %+v", err)
	}

	return &DscNodeClient{
		Client: client,
	}, nil
}
