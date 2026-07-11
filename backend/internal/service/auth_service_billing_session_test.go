package service

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	dbent "github.com/Wei-Shaw/sub2api/ent"
	"github.com/Wei-Shaw/sub2api/ent/enttest"
	"github.com/Wei-Shaw/sub2api/internal/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func TestValidateBillingSessionReturnURL(t *testing.T) {
	t.Parallel()

	for _, valid := range []string{
		"https://simulator.college/settings/subscription",
		"https://staging.simulator.college/billing/return",
	} {
		require.NoError(t, validateBillingSessionReturnURL(valid))
	}

	for _, invalid := range []string{
		"",
		"/settings/subscription",
		"https://simulator.college.evil.example/return",
		"https://evil.example/return",
		"javascript:alert(1)",
	} {
		require.ErrorIs(t, validateBillingSessionReturnURL(invalid), ErrBillingSessionInvalid)
	}
}

func TestBillingSessionCanOnlyBeConsumedOnce(t *testing.T) {
	db, err := sql.Open("sqlite", "file:billing-session-consume?mode=memory&cache=shared")
	require.NoError(t, err)
	t.Cleanup(func() { require.NoError(t, db.Close()) })
	_, err = db.Exec("PRAGMA foreign_keys = ON")
	require.NoError(t, err)
	drv := entsql.OpenDB(dialect.SQLite, db)
	client := enttest.NewClient(t, enttest.WithOptions(dbent.Driver(drv)))
	t.Cleanup(func() { require.NoError(t, client.Close()) })

	svc := &AuthService{entClient: client}
	claims := &BillingSessionClaims{}
	claims.ID = "billing-session-jti"
	claims.ExpiresAt = jwtNumericDate(time.Now().Add(5 * time.Minute))

	require.NoError(t, svc.consumeBillingSession(context.Background(), claims))
	require.ErrorIs(t, svc.consumeBillingSession(context.Background(), claims), ErrBillingSessionConsumed)
}

func TestGenerateBillingSessionTokenIncludesUniqueID(t *testing.T) {
	svc := &AuthService{cfg: &config.Config{JWT: config.JWTConfig{Secret: "test-secret"}}}
	user := &User{ID: 42, Email: "user@simulator.college", Status: StatusActive}

	first, err := svc.GenerateBillingSessionToken(user, BillingSessionInput{ReturnURL: "https://simulator.college/settings/subscription"})
	require.NoError(t, err)
	second, err := svc.GenerateBillingSessionToken(user, BillingSessionInput{ReturnURL: "https://simulator.college/settings/subscription"})
	require.NoError(t, err)
	require.NotEqual(t, first.Token, second.Token)
}

func jwtNumericDate(value time.Time) *jwt.NumericDate {
	return jwt.NewNumericDate(value)
}
