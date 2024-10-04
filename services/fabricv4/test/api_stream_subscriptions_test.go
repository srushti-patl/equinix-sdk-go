/*
Equinix Fabric API v4

Testing StreamSubscriptionsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package fabricv4

import (
	"context"
	"testing"

	openapiclient "github.com/equinix/equinix-sdk-go/services/fabricv4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_fabricv4_StreamSubscriptionsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test StreamSubscriptionsApiService CreateStreamSubscriptions", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StreamSubscriptionsApi.CreateStreamSubscriptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test StreamSubscriptionsApiService DeleteStreamSubscriptionByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamSubscriptionId string

		httpRes, err := apiClient.StreamSubscriptionsApi.DeleteStreamSubscriptionByUuid(context.Background(), streamSubscriptionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test StreamSubscriptionsApiService GetStreamSubscriptionByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamSubscriptionId string

		resp, httpRes, err := apiClient.StreamSubscriptionsApi.GetStreamSubscriptionByUuid(context.Background(), streamSubscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test StreamSubscriptionsApiService GetStreamSubscriptions", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.StreamSubscriptionsApi.GetStreamSubscriptions(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test StreamSubscriptionsApiService UpdateStreamSubscriptionByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var streamSubscriptionId string

		resp, httpRes, err := apiClient.StreamSubscriptionsApi.UpdateStreamSubscriptionByUuid(context.Background(), streamSubscriptionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
