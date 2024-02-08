package domain

type Db interface {
	GetCurrentState() (State, error)
}
