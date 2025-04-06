package integration

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestAlwaysFail(t *testing.T) {
	t.Error("testing CI")
}
