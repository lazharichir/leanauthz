package leanauthz_test

import (
	"testing"

	authz "github.com/lazharichir/leanauthz"
)

func TestOutcome(t *testing.T) {
	decider := &authz.Statement{
		Action:   "read",
		Resource: "document",
	}

	outcome := authz.NewOutcome()
	if outcome.Effect != authz.DENY {
		t.Errorf("expected effect to be %v, but got %v", authz.DENY, outcome.Effect)
	}
	if outcome.Explicit != false {
		t.Errorf("expected explicit flag to be %v, but got %v", false, outcome.Explicit)
	}
	if outcome.Decider != nil {
		t.Errorf("expected decider to be %v, but got %v", nil, outcome.Decider)
	}

	outcome.Allow(decider)
	if outcome.Effect != authz.ALLOW {
		t.Errorf("expected effect to be %v, but got %v", authz.ALLOW, outcome.Effect)
	}
	if outcome.Explicit != true {
		t.Errorf("expected explicit flag to be %v, but got %v", true, outcome.Explicit)
	}
	if outcome.Decider != decider {
		t.Errorf("expected decider to be %v, but got %v", decider, outcome.Decider)
	}

	outcome.Deny(decider)
	if outcome.Effect != authz.DENY {
		t.Errorf("expected effect to be %v, but got %v", authz.DENY, outcome.Effect)
	}
	if outcome.Explicit != true {
		t.Errorf("expected explicit flag to be %v, but got %v", true, outcome.Explicit)
	}
	if outcome.Decider != decider {
		t.Errorf("expected decider to be %v, but got %v", decider, outcome.Decider)
	}
}
