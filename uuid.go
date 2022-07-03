package nulls

import "github.com/gofrs/uuid"

// NewUUID creates a new valid uuid.NullUUID.
func NewUUID(id uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{
		UUID:  id,
		Valid: true,
	}
}
