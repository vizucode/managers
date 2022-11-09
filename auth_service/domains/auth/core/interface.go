package usercore

type IRepoAuth interface {
	GetByEmail(authCore Core) (Core, error)
}

type IServiceAuth interface {
	Verify(authCore Core) Core
}
