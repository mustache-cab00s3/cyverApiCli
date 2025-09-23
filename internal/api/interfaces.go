package api

// ClientInterface defines the common client operations
type ClientInterface interface {
}

// PentesterInterface defines the common pentester operations
type PentesterInterface interface {
	GetPentesterInfo(id string) (*PentesterInfo, error)
	ListPentesters() ([]*PentesterInfo, error)
	AssignPentester(clientID, pentesterID string) error
	RemovePentester(clientID, pentesterID string) error
}

// TokenAuthInterface defines the common token authentication operations
type TokenAuthInterface interface {
	GenerateToken() (*TokenInfo, error)
	ValidateToken(token string) error
	RevokeToken(token string) error
	RefreshToken(token string) (*TokenInfo, error)
}

// FindingsInterface defines the common findings operations
type FindingsInterface interface {
	GetFindings(params map[string]interface{}) (interface{}, error)
	GetFindingByID(id string) (interface{}, error)
}

// Common types used across interfaces
type ClientInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type PentesterInfo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type TokenInfo struct {
	Token     string `json:"token"`
	ExpiresAt string `json:"expires_at"`
	Type      string `json:"type"`
}
