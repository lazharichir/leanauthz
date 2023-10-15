package leanauthz

// Action is a non-empty string that represents an action that can be performed on a resource.
// We recommend using a capitalized verb in the imperative mood and a capitalized noun for the object.
// e.g., "ReadDocuments", "SaveDocuments", "DeleteDocuments", "ReadUsers", "SaveUsers", "DeleteUsers"
type Action string

// NewAction returns a new Action.
func NewAction(str string) Action {
	if str == "" {
		panic("authz: action string cannot be empty")
	}
	return Action(str)
}
