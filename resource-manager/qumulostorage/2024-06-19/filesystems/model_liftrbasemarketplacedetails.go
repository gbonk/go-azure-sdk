package filesystems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LiftrBaseMarketplaceDetails struct {
	MarketplaceSubscriptionId     *string                        `json:"marketplaceSubscriptionId,omitempty"`
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus `json:"marketplaceSubscriptionStatus,omitempty"`
	OfferId                       string                         `json:"offerId"`
	PlanId                        string                         `json:"planId"`
	PublisherId                   *string                        `json:"publisherId,omitempty"`
	TermUnit                      *string                        `json:"termUnit,omitempty"`
}
