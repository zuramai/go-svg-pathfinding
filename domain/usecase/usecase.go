// Package usecase defines all the interfaces for a Micro-service application.
// It is the entry point for the application's business logic. It is a top level package for a Micro-service application.
// This top level package only defines interface, the concrete implementations are defined in sub-package of it.
// It only depends on model package. No other package should dependent on it except cmd.
package usecase

type AuthUseCaseInterface interface {
	Login()
	Logout()
}
