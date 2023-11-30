package integration_test

import (
	"testing"
)

func TestPostsWorkflowIntegration(t *testing.T) {
	if testing.Short() {
		t.SkipNow()
	}
}
