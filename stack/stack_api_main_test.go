package stack_test

import (
	"log"
	"os"
	"testing"

	"github.com/AhmdSkr/strctx64/stack"
)

type provider func(uint32) stack.Extended[any]

func testImplementation(name string, m *testing.M, p provider) {
	basicProvider = func(capacity uint32) stack.Basic[any] {
		return p(capacity)
	}
	extendedProvider = p

	log.Print("Running tests on ", name, " implementation")
	if code := m.Run(); code != 0 {
		os.Exit(code)
	}
}

func TestMain(m *testing.M) {
	log.Print("Starting stack tests...")
	testImplementation("static stack", m, stack.NewStatic[any])
	// Add any new implementations here...
}
