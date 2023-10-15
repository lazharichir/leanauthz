package leanauthz

// Outcome is the result of an evaluation. It contains the effect of the evaluation, whether or not
// the effect was explicit, and the statement that caused the effect.
type Outcome struct {
	Effect   Effect
	Explicit bool
	Decider  *Statement
	Error    error
}

// Allow sets the effect of the outcome to ALLOW and sets the outcome to explicit.
func (outcome *Outcome) Allow(decider *Statement) *Outcome {
	outcome.Effect = ALLOW
	outcome.Explicit = true
	outcome.Decider = decider
	return outcome
}

// Deny sets the effect of the outcome to DENY and sets the outcome to explicit.
func (outcome *Outcome) Deny(decider *Statement) *Outcome {
	outcome.Effect = DENY
	outcome.Explicit = true
	outcome.Decider = decider
	return outcome
}

// Fail sets the effect of the outcome to DENY and sets the outcome to implicit. It also sets the
// error of the outcome.
func (outcome *Outcome) Fail(err error) *Outcome {
	outcome.Effect = DENY
	outcome.Explicit = false
	outcome.Error = err
	return outcome
}

// NewOutcome creates a new Outcome, which is implicitly DENY by default.
func NewOutcome() *Outcome {
	return &Outcome{
		Effect:   DENY,
		Explicit: false,
		Decider:  nil,
	}
}
