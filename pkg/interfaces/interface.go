package models

// TransportMethod définit l'interface commune pour tous les moyens de transport
type TransportMethod interface {
	DeliverPackage(destination string) (string, error)
	GetStatus() string
}
