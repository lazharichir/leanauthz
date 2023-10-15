package leanauthz

import (
	"context"
)

// Store is an interface for storing and retrieving Statements.
// It is up to the implementer to decide how to store the Statements.
// Recommended implementations are in-memory for dev and testing, and SQL/NoSQL in production.
type Store interface {

	// SaveStatements saves the given Statements to the store.
	SaveStatements(ctx context.Context, stmts []Statement) error

	// GetStatementsByPrincipal returns all Statements that match the given principal Pattern.
	GetStatementsByPrincipal(ctx context.Context, p Pattern) ([]Statement, error)

	// DeleteStatements hard-deletes the given Statements from the store.
	DeleteStatements(ctx context.Context, stmts []Statement) error

	// FindCandidates returns all Statements that match the given Evaluation. It's used to help narrow down the Statements that need to be evaluated.
	// The returned Statements are not guaranteed to be a match for the given Evaluation (this is why they are called "Candidates"),
	// and will be filtered again within in our Evaluate function. This is to allow for more efficient querying of the store.
	// If you implement your own SQL store, you can use this function to narrow down the Statements that need to be evaluated using a like "%" search.
	FindCandidates(ctx context.Context, e Evaluation) ([]Statement, error)
}
