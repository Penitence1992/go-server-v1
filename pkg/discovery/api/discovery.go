package api

type Discovery interface {
	IsAppExists() (bool, error)
	CreateInstance() (bool, error)
	Heartbeat() (bool, error)
	RemoveInstance() (bool, error)
}
