package leanauthz

import (
	"context"
)

// Evaluate evaluates a single Evaluation.
func Evaluate(ctx context.Context, store Store, e Evaluation) Outcome {
	outcome := NewOutcome()

	// Retrieve candidate statements that match the evaluation.
	stmts, err := store.FindCandidates(ctx, e)
	if err != nil {
		outcome.Fail(err)
		return *outcome
	}

	// If there are no statements, then the outcome is an implicit DENY.
	if len(stmts) == 0 {
		return *outcome
	}

	// Go through each statement and evaluate it.
	// The store implementation may return statements that don't match the evaluation,
	// so we still need to filter them out as we cannot trust any store implementation.
	for _, stmt := range stmts {

		// Filter out statements that don't match the evaluation
		if !stmt.MatchesEvaluation(e) {
			continue
		}

		// If the statement is an explicit DENY, then the outcome is DENY.
		if stmt.Effect == DENY {
			outcome.Deny(&stmt)
			return *outcome
		}

		// If the statement is an explicit ALLOW, then the outcome is ALLOW.
		// We don't return here because there may be other statements that are DENY.
		if stmt.Effect == ALLOW {
			outcome.Allow(&stmt)
		}
	}

	return *outcome
}

// EvaluateMany evaluates many Evaluations at once. It returns the first DENY outcome and all
func EvaluateMany(ctx context.Context, store Store, evals ...Evaluation) (Outcome, []Outcome) {
	outcomes := []Outcome{}
	if len(evals) == 0 {
		return *NewOutcome(), outcomes
	}

	for _, e := range evals {
		outcomes = append(outcomes, Evaluate(ctx, store, e))
	}

	outcome := NewOutcome()
	for _, o := range outcomes {
		if o.Effect == DENY {
			outcome.Deny(o.Decider)
			return *outcome, outcomes
		}
		if o.Effect == ALLOW {
			outcome.Allow(o.Decider)
		}
	}

	return *outcome, outcomes
}
