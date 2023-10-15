package leanauthz

import (
	"context"
)

// devStoreImpl is an in-memory implementation of Store.
// The below line is a compile-time check that devStoreImpl implements Store.
var _ Store = &devStoreImpl{}

type devStoreImpl struct {
	stmts []Statement
}

func NewDevStore() *devStoreImpl {
	impl := &devStoreImpl{}
	return impl
}

func (store *devStoreImpl) statementExists(s Statement) bool {
	for _, stmt := range store.stmts {
		if stmt.ID == s.ID {
			return true
		}
	}
	return false
}

func (store *devStoreImpl) SaveStatements(ctx context.Context, stmts []Statement) error {
	for _, stmt := range stmts {
		if store.statementExists(stmt) {
			continue
		}
		store.stmts = append(store.stmts, stmt)
	}
	return nil
}

func (store *devStoreImpl) GetStatementsByPrincipal(ctx context.Context, p Pattern) ([]Statement, error) {
	items := []Statement{}
	for _, stmt := range store.stmts {
		if stmt.Principal == p {
			items = append(items, stmt)
		}
	}
	return items, nil
}

func (store *devStoreImpl) FindCandidates(ctx context.Context, e Evaluation) ([]Statement, error) {
	candidates := []Statement{}
	for _, stmt := range store.stmts {
		matchesPrincipal := stmt.MatchesPrincipal(e.Principal)
		matchesResource := stmt.MatchesResource(e.Resource)
		matchesAction := stmt.MatchesAction(e.Action)

		// fmt.Println("--------")
		// fmt.Printf("stmt %+#v \n", stmt)
		// fmt.Printf("eval %+#v \n", e)
		// fmt.Println("matchesPrincipal:", matchesPrincipal)
		// fmt.Println("matchesResource:", matchesResource)
		// fmt.Println("matchesAction:", matchesAction)
		// fmt.Println("--------")

		if !matchesAction || !matchesResource || !matchesPrincipal {
			continue
		}
		candidates = append(candidates, stmt)
	}
	return candidates, nil
}

func (store *devStoreImpl) DeleteStatements(ctx context.Context, stmts []Statement) error {
	for _, stmt := range stmts {
		for i, s := range store.stmts {
			if s.ID == stmt.ID {
				store.stmts = append(store.stmts[:i], store.stmts[i+1:]...)
				break
			}
		}
	}
	return nil
}
