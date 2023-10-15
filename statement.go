package leanauthz

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/lazharichir/leanauthz/matching"
)

func NewAllowStatement(principal, action, resource Pattern) Statement {
	return NewStatement(ALLOW, principal, action, resource)
}

func NewDenyStatement(principal, action, resource Pattern) Statement {
	return NewStatement(DENY, principal, action, resource)
}

func generateStatementID() (string, error) {
	// Generate 16 random bytes
	b := make([]byte, 22)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Encode the bytes as a base64 string
	id := base64.URLEncoding.EncodeToString(b)

	return id, nil
}

func NewStatement(effect Effect, principal, action, resource Pattern) Statement {
	id, err := generateStatementID()
	if err != nil {
		panic(err)
	}

	return Statement{
		ID:        id,
		Effect:    effect,
		Principal: principal,
		Resource:  resource,
		Action:    action,
	}
}

type Statement struct {
	ID        string
	Effect    Effect
	Principal Pattern
	Resource  Pattern
	Action    Pattern
}

func (s *Statement) SetEffect(e Effect) *Statement {
	s.Effect = e
	return s
}

func (s *Statement) SetPrincipal(p Pattern) *Statement {
	s.Principal = p
	return s
}

func (s *Statement) SetResource(r Pattern) *Statement {
	s.Resource = r
	return s
}

func (s *Statement) SetAction(a Pattern) *Statement {
	s.Action = a
	return s
}

func (s *Statement) MatchesPrincipal(p Principal) bool {
	return matching.Match(string(s.Principal), string(p))
}

func (s *Statement) MatchesResource(r Resource) bool {
	return matching.Match(s.Resource.String(), string(r))
}

func (s *Statement) MatchesAction(a Action) bool {
	return matching.Match(s.Action.String(), string(a))
}

func (s *Statement) MatchesEvaluation(eval Evaluation) bool {
	return s.MatchesPrincipal(eval.Principal) && s.MatchesResource(eval.Resource) && s.MatchesAction(eval.Action)
}
