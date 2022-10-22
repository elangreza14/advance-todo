package domain

type (
	// ContextValue is type context
	// that used for naming context value
	ContextValue string
)

const (
	// ContextValueIP is type context for holding ip user
	ContextValueIP ContextValue = "context-ip"

	// ContextValueUserID is type context for holding user id
	ContextValueUserID ContextValue = "context-user-id"
)
