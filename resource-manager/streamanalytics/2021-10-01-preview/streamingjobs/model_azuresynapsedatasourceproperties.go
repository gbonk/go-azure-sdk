package streamingjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureSynapseDataSourceProperties struct {
	AuthenticationMode *AuthenticationMode `json:"authenticationMode,omitempty"`
	Database           *string             `json:"database,omitempty"`
	Password           *string             `json:"password,omitempty"`
	Server             *string             `json:"server,omitempty"`
	Table              *string             `json:"table,omitempty"`
	User               *string             `json:"user,omitempty"`
}
