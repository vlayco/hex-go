package ports

// This file contains driven/right-side adapters; connects to framework layer

type DBPort interface {
	CloseDBConnection()
	AddToHistory(answer int32, operation string) error
}
