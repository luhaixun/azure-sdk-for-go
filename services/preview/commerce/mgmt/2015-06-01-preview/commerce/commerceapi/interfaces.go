package commerceapi

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/commerce/mgmt/2015-06-01-preview/commerce"
	"github.com/Azure/go-autorest/autorest/date"
)

// UsageAggregatesClientAPI contains the set of methods on the UsageAggregatesClient type.
type UsageAggregatesClientAPI interface {
	List(ctx context.Context, reportedStartTime date.Time, reportedEndTime date.Time, showDetails *bool, aggregationGranularity commerce.AggregationGranularity, continuationToken string) (result commerce.UsageAggregationListResultPage, err error)
}

var _ UsageAggregatesClientAPI = (*commerce.UsageAggregatesClient)(nil)

// RateCardClientAPI contains the set of methods on the RateCardClient type.
type RateCardClientAPI interface {
	Get(ctx context.Context, filter string) (result commerce.ResourceRateCardInfo, err error)
}

var _ RateCardClientAPI = (*commerce.RateCardClient)(nil)
