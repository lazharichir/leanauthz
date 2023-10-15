package leanauthz_test

import (
	"testing"

	authz "github.com/lazharichir/leanauthz"
)

func TestNewStatement(t *testing.T) {
	tests := []struct {
		name      string
		effect    authz.Effect
		principal authz.Pattern
		action    authz.Pattern
		resource  authz.Pattern
	}{
		{
			name:      "allow statement",
			effect:    authz.ALLOW,
			principal: authz.Pattern("user123"),
			action:    authz.Pattern("SaveDocuments"),
			resource:  authz.Pattern("documents/*"),
		},
		{
			name:      "deny statement",
			effect:    authz.DENY,
			principal: authz.Pattern("user456"),
			action:    authz.Pattern("DeleteDocuments"),
			resource:  authz.Pattern("documents/doc_123"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := authz.NewStatement(tt.effect, tt.principal, tt.action, tt.resource)

			if got.Effect != tt.effect {
				t.Errorf("NewStatement() effect = %v, want %v", got.Effect, tt.effect)
			}

			if got.Principal != tt.principal {
				t.Errorf("NewStatement() principal = %v, want %v", got.Principal, tt.principal)
			}

			if got.Action != tt.action {
				t.Errorf("NewStatement() action = %v, want %v", got.Action, tt.action)
			}

			if got.Resource != tt.resource {
				t.Errorf("NewStatement() resource = %v, want %v", got.Resource, tt.resource)
			}
		})
	}
}

func TestStatementSetters(t *testing.T) {
	s := &authz.Statement{}

	// Test SetEffect
	wantEffect := authz.ALLOW
	got := s.SetEffect(wantEffect)
	if got.Effect != wantEffect {
		t.Errorf("SetEffect() = %v, want %v", got.Effect, wantEffect)
	}

	// Test SetPrincipal
	wantPrincipal := authz.Pattern("user123")
	got = s.SetPrincipal(wantPrincipal)
	if got.Principal != wantPrincipal {
		t.Errorf("SetPrincipal() = %v, want %v", got.Principal, wantPrincipal)
	}

	// Test SetResource
	wantResource := authz.Pattern("documents/*")
	got = s.SetResource(wantResource)
	if got.Resource != wantResource {
		t.Errorf("SetResource() = %v, want %v", got.Resource, wantResource)
	}

	// Test SetAction
	wantAction := authz.Pattern("SaveDocuments")
	got = s.SetAction(wantAction)
	if got.Action != wantAction {
		t.Errorf("SetAction() = %v, want %v", got.Action, wantAction)
	}
}

func TestStatementMatchesPrincipal(t *testing.T) {
	s := &authz.Statement{
		Principal: authz.Pattern("user*"),
	}

	tests := []struct {
		name      string
		principal authz.Principal
		want      bool
	}{
		{
			name:      "match",
			principal: authz.Principal("user123"),
			want:      true,
		},
		{
			name:      "no match",
			principal: authz.Principal("admin"),
			want:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.MatchesPrincipal(tt.principal)

			if got != tt.want {
				t.Errorf("MatchesPrincipal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatementMatchesResource(t *testing.T) {
	s := &authz.Statement{
		Resource: authz.NewPattern("documents/*"),
	}

	tests := []struct {
		name     string
		resource authz.Resource
		want     bool
	}{
		{
			name:     "match",
			resource: authz.NewResource("documents/doc_123"),
			want:     true,
		},
		{
			name:     "no match",
			resource: authz.NewResource("users/user123"),
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.MatchesResource(tt.resource)

			if got != tt.want {
				t.Errorf("MatchesResource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStatementMatchesAction(t *testing.T) {
	s := &authz.Statement{
		Action: authz.Pattern("Save*"),
	}

	tests := []struct {
		name   string
		action authz.Action
		want   bool
	}{
		{
			name:   "match",
			action: authz.Action("SaveDocuments"),
			want:   true,
		},
		{
			name:   "no match",
			action: authz.Action("DeleteDocuments"),
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := s.MatchesAction(tt.action)

			if got != tt.want {
				t.Errorf("MatchesAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
