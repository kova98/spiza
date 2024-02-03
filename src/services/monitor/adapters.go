package main

type Db interface {
	GetCurrentState() (State, error)
}
