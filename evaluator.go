package leanauthz

import "context"

type Evaluator struct {
	store Store
}

func NewEvaluator(store Store) *Evaluator {
	return &Evaluator{
		store: store,
	}
}

// Evaluate evaluates a single Evaluation.
func (evaluator *Evaluator) Evaluate(ctx context.Context, e Evaluation) Outcome {
	return Evaluate(ctx, evaluator.store, e)
}

// EvaluateMany evaluates many Evaluations at once. It returns the first DENY outcome and all
func (evaluator *Evaluator) EvaluateMany(ctx context.Context, evals ...Evaluation) (Outcome, []Outcome) {
	return EvaluateMany(ctx, evaluator.store, evals...)
}
