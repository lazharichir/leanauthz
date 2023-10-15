package leanauthz

// Evaluation is the input to the Evaluate function. It represents a single authorization request.
// In other words, an Evaluation asks the question "Can PRINCIPAL perform ACTION on RESOURCE?".
type Evaluation struct {
	Principal Principal
	Action    Action
	Resource  Resource
}

// NewEvaluation creates a new Evaluation.
func NewEvaluation() *Evaluation {
	return &Evaluation{}
}

// SetPrincipal sets the principal of the evaluation.
func (e *Evaluation) SetPrincipal(p Principal) *Evaluation {
	e.Principal = p
	return e
}

// SetAction sets the action of the evaluation.
func (e *Evaluation) SetAction(a Action) *Evaluation {
	e.Action = a
	return e
}

// SetResource sets the resource of the evaluation.
func (e *Evaluation) SetResource(r Resource) *Evaluation {
	e.Resource = r
	return e
}
