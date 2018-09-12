// Test Union filesystem interface
package roundrobin_test

import (
	"testing"

	_ "github.com/ncw/rclone/backend/local"
	"github.com/ncw/rclone/fstest/fstests"
)

// TestIntegration runs integration tests against the remote
func TestIntegration(t *testing.T) {
	fstests.Run(t, &fstests.Opt{
		RemoteName:  "TestRoundRobin:",
		NilObject:   nil,
		SkipFsMatch: true,
	})
}
