/*
Metal API

Testing DevicesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package metalv1

import (
	"context"
	"testing"

	openapiclient "github.com/equinix/equinix-sdk-go/services/metalv1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_metalv1_DevicesApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DevicesApiService CreateBgpSession", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.CreateBgpSession(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService CreateDevice", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.CreateDevice(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService CreateIPAssignment", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.CreateIPAssignment(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService DeleteDevice", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.DevicesApi.DeleteDevice(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService FindBgpSessions", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.FindBgpSessions(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService FindDeviceById", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.FindDeviceById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService FindDeviceCustomdata", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.DevicesApi.FindDeviceCustomdata(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService FindDeviceMetadataByID", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.FindDeviceMetadataByID(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService FindDeviceUserdataByID", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.FindDeviceUserdataByID(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService FindIPAssignmentCustomdata", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var instanceId string
		var id string

		httpRes, err := apiClient.DevicesApi.FindIPAssignmentCustomdata(context.Background(), instanceId, id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService FindIPAssignments", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.FindIPAssignments(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService FindInstanceBandwidth", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.DevicesApi.FindInstanceBandwidth(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService FindOrganizationDevices", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.FindOrganizationDevices(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

		resp, err = apiClient.DevicesApi.FindOrganizationDevices(context.Background(), id).ExecuteWithPagination()
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("Test DevicesApiService FindProjectDevices", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.FindProjectDevices(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

		resp, err = apiClient.DevicesApi.FindProjectDevices(context.Background(), id).ExecuteWithPagination()
		require.Nil(t, err)
		require.NotNil(t, resp)
	})

	t.Run("Test DevicesApiService FindTraffic", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.DevicesApi.FindTraffic(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService GetBgpNeighborData", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.GetBgpNeighborData(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService GetDeviceFirmwareSets", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.GetDeviceFirmwareSets(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService GetDeviceHealthRollup", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.GetDeviceHealthRollup(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService PerformAction", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.DevicesApi.PerformAction(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test DevicesApiService UpdateDevice", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.DevicesApi.UpdateDevice(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
