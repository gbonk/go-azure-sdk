package communitygalleries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGallery struct {
	Identifier *CommunityGalleryIdentifier `json:"identifier,omitempty"`
	Location   *string                     `json:"location,omitempty"`
	Name       *string                     `json:"name,omitempty"`
	Properties *CommunityGalleryProperties `json:"properties,omitempty"`
	Type       *string                     `json:"type,omitempty"`
}
