package admin

import (
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/service"
	"github.com/stretchr/testify/require"
)

func TestDeviceTokenForExistingKeyRedactsSecret(t *testing.T) {
	t.Parallel()

	out := deviceTokenForExistingKey(&service.APIKey{
		ID:     7,
		UserID: 42,
		Key:    "sk-live-secret",
		Name:   "Simulator MacBook Pro",
		Status: service.StatusAPIKeyActive,
	})

	require.NotNil(t, out)
	require.Equal(t, int64(7), out.ID)
	require.Equal(t, "Simulator MacBook Pro", out.Name)
	require.Empty(t, out.Key)
}
