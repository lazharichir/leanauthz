package leanauthz_test

import (
	"context"
	"fmt"
	"testing"

	authz "github.com/lazharichir/leanauthz"
)

func checkOutcome(t *testing.T, label string, want authz.Outcome, got authz.Outcome) {
	t.Run(label, func(t *testing.T) {
		if want.Effect != got.Effect {
			t.Errorf("Expected Effect %v, got %v", want.Effect, got.Effect)
		}
		if want.Explicit != got.Explicit {
			t.Errorf("Expected Explicit %v, got %v", want.Explicit, got.Explicit)
		}
		if want.Decider == nil && got.Decider != nil {
			t.Errorf("Expected Decider %v, got %v", want.Decider, got.Decider)
		} else if want.Decider != nil && got.Decider == nil {
			t.Errorf("Expected Decider %v, got %v", want.Decider, got.Decider)
		} else if want.Decider != nil && got.Decider != nil {
			if want.Decider.ID != got.Decider.ID {
				t.Errorf("Expected Decider.Principal %v, got %v", want.Decider.Principal, got.Decider.Principal)
			}
		}
	})
}

func TestEvaluateMany(t *testing.T) {
	ctx := context.Background()
	store := authz.NewDevStore()

	// allow john to save documents in google
	stmt1 := authz.NewAllowStatement(authz.NewPattern("users/john"), authz.NewPattern("SaveDocuments"), authz.NewPattern("accounts/google/documents/*"))
	// deny john to delete documents in google
	stmt2 := authz.NewDenyStatement(authz.NewPattern("users/john"), authz.NewPattern("DeleteDocuments"), authz.NewPattern("accounts/google/documents/*"))
	// allow anybody to read anything in any account
	stmt3 := authz.NewAllowStatement(authz.NewPattern("users/*"), authz.NewPattern("Read*"), authz.NewPattern("accounts/*"))

	// save the statements
	_ = store.SaveStatements(ctx, []authz.Statement{stmt1, stmt2, stmt3})

	{
		// evaluate many statements
		e1 := authz.NewEvaluation()
		e1.SetPrincipal(authz.NewUserPrincipal("john"))
		e1.SetAction(authz.NewAction("SaveDocuments"))
		e1.SetResource(authz.NewResource("accounts/google/documents/abcdef"))

		e2 := authz.NewEvaluation()
		e2.SetPrincipal(authz.NewUserPrincipal("john"))
		e2.SetAction(authz.NewAction("DeleteDocuments"))
		e2.SetResource(authz.NewResource("accounts/google/documents/ghejkl"))

		gotReduced, gotSlice := authz.EvaluateMany(ctx, store, *e1, *e2)

		want := []authz.Outcome{
			*authz.NewOutcome().Allow(&stmt1),
			*authz.NewOutcome().Deny(&stmt2),
		}

		if len(want) != len(gotSlice) {
			t.Errorf("Expected %d outcome items, got %d", len(want), len(gotSlice))
		}

		for i, wantOutcome := range want {
			checkOutcome(t, fmt.Sprintf("EvaluateMany %d", i), wantOutcome, gotSlice[i])
		}

		checkOutcome(t, "EvaluateMany reduced", *authz.NewOutcome().Deny(&stmt2), gotReduced)
	}
}

func TestEvaluate(t *testing.T) {
	ctx := context.Background()
	store := authz.NewDevStore()

	// allow john to save documents in google
	stmt1 := authz.NewAllowStatement(authz.NewPattern("users/john"), authz.NewPattern("SaveDocuments"), authz.NewPattern("accounts/google/documents/*"))
	// deny john to delete documents in google
	stmt2 := authz.NewDenyStatement(authz.NewPattern("users/john"), authz.NewPattern("DeleteDocuments"), authz.NewPattern("accounts/google/documents/*"))
	// allow anybody to read anything in any account
	stmt3 := authz.NewAllowStatement(authz.NewPattern("users/*"), authz.NewPattern("Read*"), authz.NewPattern("accounts/*"))

	// save the statements
	_ = store.SaveStatements(ctx, []authz.Statement{stmt1, stmt2, stmt3})

	{
		e := authz.NewEvaluation()
		e.SetPrincipal(authz.NewUserPrincipal("john"))
		e.SetAction(authz.NewAction("SaveDocuments"))
		e.SetResource(authz.NewResource("accounts/google/documents/abcdef"))
		checkOutcome(t, "john can save documents", authz.Outcome{
			Effect:   authz.ALLOW,
			Explicit: true,
			Decider:  &stmt1,
		}, authz.Evaluate(ctx, store, *e))

	}

	{
		e := authz.NewEvaluation()
		e.SetPrincipal(authz.NewUserPrincipal("john"))
		e.SetAction(authz.NewAction("DeleteDocuments"))
		e.SetResource(authz.NewResource("accounts/google/documents/abcdef"))
		checkOutcome(t, "john can't delete documents", authz.Outcome{
			Effect:   authz.DENY,
			Explicit: true,
			Decider:  &stmt2,
		}, authz.Evaluate(ctx, store, *e))
	}

	{
		e := authz.NewEvaluation()
		e.SetPrincipal(authz.NewUserPrincipal("anonymous"))
		e.SetAction(authz.NewAction("ReadInvoices"))
		e.SetResource(authz.NewResource("accounts/microsoft/invoices/inv-00120124/versions/4"))
		checkOutcome(t, "anonymous can read invoices", authz.Outcome{
			Effect:   authz.ALLOW,
			Explicit: true,
			Decider:  &stmt3,
		}, authz.Evaluate(ctx, store, *e))
	}
}
