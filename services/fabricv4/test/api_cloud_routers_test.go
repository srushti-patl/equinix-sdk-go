/*
Equinix Fabric API v4

Testing CloudRoutersApiService

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

func Test_fabricv4_CloudRoutersApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CloudRoutersApiService CreateCloudRouter", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CloudRoutersApi.CreateCloudRouter(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService CreateCloudRouterAction", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routerId string

		resp, httpRes, err := apiClient.CloudRoutersApi.CreateCloudRouterAction(context.Background(), routerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService DeleteCloudRouterByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routerId string

		httpRes, err := apiClient.CloudRoutersApi.DeleteCloudRouterByUuid(context.Background(), routerId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService GetCloudRouterActions", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routerId string

		resp, httpRes, err := apiClient.CloudRoutersApi.GetCloudRouterActions(context.Background(), routerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService GetCloudRouterActionsByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routerId string
		var actionId string

		resp, httpRes, err := apiClient.CloudRoutersApi.GetCloudRouterActionsByUuid(context.Background(), routerId, actionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService GetCloudRouterByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routerId string

		resp, httpRes, err := apiClient.CloudRoutersApi.GetCloudRouterByUuid(context.Background(), routerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService GetCloudRouterPackageByCode", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routerPackageCode openapiclient.RouterPackageCode

		resp, httpRes, err := apiClient.CloudRoutersApi.GetCloudRouterPackageByCode(context.Background(), routerPackageCode).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService GetCloudRouterPackages", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CloudRoutersApi.GetCloudRouterPackages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService SearchCloudRouterRoutes", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routerId string

		resp, httpRes, err := apiClient.CloudRoutersApi.SearchCloudRouterRoutes(context.Background(), routerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService SearchCloudRouters", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.CloudRoutersApi.SearchCloudRouters(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService SearchConnectionAdvertisedRoutes", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var connectionId string

		resp, httpRes, err := apiClient.CloudRoutersApi.SearchConnectionAdvertisedRoutes(context.Background(), connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService SearchConnectionReceivedRoutes", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var connectionId string

		resp, httpRes, err := apiClient.CloudRoutersApi.SearchConnectionReceivedRoutes(context.Background(), connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService SearchRouterActions", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routerId string

		resp, httpRes, err := apiClient.CloudRoutersApi.SearchRouterActions(context.Background(), routerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test CloudRoutersApiService UpdateCloudRouterByUuid", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var routerId string

		resp, httpRes, err := apiClient.CloudRoutersApi.UpdateCloudRouterByUuid(context.Background(), routerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
