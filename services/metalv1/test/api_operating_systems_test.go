/*
Metal API

Testing OperatingSystemsApiService

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

func Test_metalv1_OperatingSystemsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test OperatingSystemsApiService FindOperatingSystemVersion", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OperatingSystemsApi.FindOperatingSystemVersion(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test OperatingSystemsApiService FindOperatingSystems", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.OperatingSystemsApi.FindOperatingSystems(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
