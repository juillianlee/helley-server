package app_security

type (
	TokenManager interface {
		GenerateTokenPair() (map[string]string, error)
	}
)
