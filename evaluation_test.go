package leanauthz_test

import (
	"testing"

	authz "github.com/lazharichir/leanauthz"
)

func TestEvaluation(t *testing.T) {

	got := authz.NewEvaluation().
		SetPrincipal(authz.NewUserPrincipal("john")).
		SetAction(authz.NewAction("SaveDocuments")).
		SetResource(authz.NewResource("google/documents/abcdef"))

	want := authz.Evaluation{
		Principal: authz.Principal("users/john"),
		Action:    authz.Action("SaveDocuments"),
		Resource:  authz.Resource("google/documents/abcdef"),
	}

	if got.Principal != want.Principal {
		t.Errorf("Expected Principal %v, got %v", want.Principal, got.Principal)
	}

	if got.Action != want.Action {
		t.Errorf("Expected Action %v, got %v", want.Action, got.Action)
	}

	if got.Resource != want.Resource {
		t.Errorf("Expected Resource %v, got %v", want.Resource, got.Resource)
	}

}
