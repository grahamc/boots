package vmware

import (
	"os"
	"testing"

	"github.com/packethost/boots/installers"
	"github.com/packethost/boots/job"
	l "github.com/packethost/pkg/log"
)

func TestMain(m *testing.M) {
	os.Setenv("PACKET_ENV", "test")
	os.Setenv("PACKET_VERSION", "0")
	os.Setenv("ROLLBAR_DISABLE", "1")
	os.Setenv("ROLLBAR_TOKEN", "1")

	logger, _ := l.Init("github.com/packethost/boots")
	installers.Init(logger)
	job.Init(logger)
	os.Exit(m.Run())
}