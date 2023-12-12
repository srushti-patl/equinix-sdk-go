/*
Metal API

Testing TwoFactorAuthApiService

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

func Test_metalv1_TwoFactorAuthApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TwoFactorAuthApiService DisableTfaApp", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TwoFactorAuthApi.DisableTfaApp(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TwoFactorAuthApiService DisableTfaSms", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TwoFactorAuthApi.DisableTfaSms(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TwoFactorAuthApiService EnableTfaApp", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TwoFactorAuthApi.EnableTfaApp(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test TwoFactorAuthApiService EnableTfaSms", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.TwoFactorAuthApi.EnableTfaSms(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
