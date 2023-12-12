/*
Metal API

Testing InvitationsApiService

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

func Test_metalv1_InvitationsApiService(t *testing.T) {
	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test InvitationsApiService AcceptInvitation", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.InvitationsApi.AcceptInvitation(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InvitationsApiService DeclineInvitation", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		httpRes, err := apiClient.InvitationsApi.DeclineInvitation(context.Background(), id).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)
	})

	t.Run("Test InvitationsApiService FindInvitationById", func(t *testing.T) {
		t.Skip("skip test") // remove to run test

		var id string

		resp, httpRes, err := apiClient.InvitationsApi.FindInvitationById(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	})
}
