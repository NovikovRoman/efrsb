package efrsb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReferenceBookMessageTypes(t *testing.T) {
	ctx := context.Background()
	m, err := testClient.ReferenceBookMessageTypes(ctx)
	if err != nil {
		t.Errorf("ReferenceBookMessageTypes() error = %v", err)
	}
	found := false
	for _, i := range m {
		if i.Code == MessageArbitralDecree {
			found = true
			break
		}
	}
	assert.True(t, found)
	assert.Greater(t, len(m), 0)
}

func TestReferenceBookCourtDecisionTypes(t *testing.T) {
	ctx := context.Background()
	m, err := testClient.ReferenceBookCourtDecisionTypes(ctx)
	if err != nil {
		t.Errorf("ReferenceBookCourtDecisionTypes() error = %v", err)
	}
	found := false
	for _, i := range m {
		if i.Code == ActObservation {
			found = true
			break
		}
	}
	assert.True(t, found)
	assert.Greater(t, len(m), 0)
}
