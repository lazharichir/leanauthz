package leanauthz_test

import (
	"context"
	"testing"

	authz "github.com/lazharichir/leanauthz"
)

func TestDevStoreImpl_SaveStatements(t *testing.T) {
	store := authz.NewDevStore()

	stmts := []authz.Statement{
		authz.NewAllowStatement(authz.NewPattern("user123"), authz.NewPattern("SaveDocuments"), "documents/*"),
		authz.NewDenyStatement(authz.NewPattern("user456"), authz.NewPattern("DeleteDocuments"), "documents/doc_123"),
	}

	err := store.SaveStatements(context.Background(), stmts)
	if err != nil {
		t.Errorf("SaveStatements() error = %v, want nil", err)
	}

	// Test that statements are saved
	got, err := store.GetStatementsByPrincipal(context.Background(), authz.NewPattern("user123"))
	if err != nil {
		t.Errorf("GetStatementsByPrincipal() error = %v, want nil", err)
	}
	if len(got) != 1 {
		t.Errorf("GetStatementsByPrincipal() len = %v, want 1", len(got))
	}
}

func TestDevStoreImpl_GetStatementsByPrincipal(t *testing.T) {
	store := authz.NewDevStore()

	stmts := []authz.Statement{
		authz.NewAllowStatement(authz.NewPattern("user123"), authz.NewPattern("SaveDocuments"), "documents/*"),
		authz.NewDenyStatement(authz.NewPattern("user456"), authz.NewPattern("DeleteDocuments"), "documents/doc_123"),
	}

	err := store.SaveStatements(context.Background(), stmts)
	if err != nil {
		t.Errorf("SaveStatements() error = %v, want nil", err)
	}

	// Test getting statements by principal
	got, err := store.GetStatementsByPrincipal(context.Background(), authz.NewPattern("user123"))
	if err != nil {
		t.Errorf("GetStatementsByPrincipal() error = %v, want nil", err)
	}
	if len(got) != 1 {
		t.Errorf("GetStatementsByPrincipal() len = %v, want 1", len(got))
	}
}

func TestDevStoreImpl_FindCandidates(t *testing.T) {
	store := authz.NewDevStore()

	stmts := []authz.Statement{
		authz.NewAllowStatement(authz.NewPattern("users/123"), authz.NewPattern("SaveDocuments"), "documents/*"),
		authz.NewDenyStatement(authz.NewPattern("users/456"), authz.NewPattern("DeleteDocuments"), "documents/doc_123"),
	}

	err := store.SaveStatements(context.Background(), stmts)
	if err != nil {
		t.Errorf("SaveStatements() error = %v, want nil", err)
	}

	// Test finding candidates
	e := authz.Evaluation{
		Principal: authz.NewPrincipal("users/123"),
		Action:    authz.Action("SaveDocuments"),
		Resource:  authz.NewResource("documents/doc_123"),
	}
	got, err := store.FindCandidates(context.Background(), e)
	if err != nil {
		t.Errorf("FindCandidates() error = %v, want nil", err)
	}
	if len(got) != 1 {
		t.Errorf("FindCandidates() len = %v, want 1", len(got))
	}
}

func TestDevStoreImpl_DeleteStatements(t *testing.T) {
	store := authz.NewDevStore()

	stmts := []authz.Statement{
		authz.NewAllowStatement(authz.NewPattern("user123"), authz.NewPattern("SaveDocuments"), "documents/*"),
		authz.NewDenyStatement(authz.NewPattern("user456"), authz.NewPattern("DeleteDocuments"), "documents/doc_123"),
	}
	err := store.SaveStatements(context.Background(), stmts)
	if err != nil {
		t.Errorf("SaveStatements() error = %v, want nil", err)
	}

	// Test deleting statements
	err = store.DeleteStatements(context.Background(), stmts)
	if err != nil {
		t.Errorf("DeleteStatements() error = %v, want nil", err)
	}

	// Test that statements are deleted
	got, err := store.GetStatementsByPrincipal(context.Background(), authz.NewPattern("user123"))
	if err != nil {
		t.Errorf("GetStatementsByPrincipal() error = %v, want nil", err)
	}
	if len(got) != 0 {
		t.Errorf("GetStatementsByPrincipal() len = %v, want 0", len(got))
	}
}
