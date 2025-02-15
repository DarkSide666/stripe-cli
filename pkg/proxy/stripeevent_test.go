package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsConnect(t *testing.T) {
	evt1 := &stripeEvent{ID: "evt_123", Type: "customer.created"}
	require.False(t, evt1.isConnect())

	evt2 := &stripeEvent{ID: "evt_123", Type: "customer.created", Account: "acct_123"}
	require.True(t, evt2.isConnect())
}

func TestURLForEventID(t *testing.T) {
	evt1 := &stripeEvent{ID: "evt_123", Type: "customer.created"}
	require.Equal(t, "https://dashboard.stripe.com/test/events/evt_123", evt1.urlForEventID())

	evt2 := &stripeEvent{ID: "evt_123", Type: "customer.created", Account: "acct_123"}
	require.Equal(t, "https://dashboard.stripe.com/acct_123/test/events/evt_123", evt2.urlForEventID())
}

func TestURLForEventType(t *testing.T) {
	evt1 := &stripeEvent{ID: "evt_123", Type: "customer.created"}
	require.Equal(t, "https://dashboard.stripe.com/test/events?type=customer.created", evt1.urlForEventType())

	evt2 := &stripeEvent{ID: "evt_123", Type: "customer.created", Account: "acct_123"}
	require.Equal(t, "https://dashboard.stripe.com/acct_123/test/events?type=customer.created", evt2.urlForEventType())
}
