package queue_test

import (
	"log"
	"os"
	"testing"

	"github.com/AhmdSkr/strctx64/queue"
)

type provider func(uint32) queue.Extended[any]

func testImplementation(name string, m *testing.M, p provider) {
	basicProvider = func(capacity uint32) queue.Basic[any] {
		return p(capacity)
	}
	extendedProvider = p

	log.Print("Running tests on ", name, " implementation")
	if code := m.Run(); code != 0 {
		os.Exit(code)
	}
}

func TestMain(m *testing.M) {
	log.Println("Starting queue tests...")
	testImplementation("static queue", m, queue.NewStatic[any])
	// Add any new implementations here...
}
