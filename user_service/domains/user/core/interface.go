package usercore

type IRepoUser interface {
	Insert(activityCore Core) error
	Update(activityCore Core) error
	Delete(activityCore Core) error
	FindAll() ([]Core, error)
	First(activityCore Core) (Core, error)
}

type IServiceUser interface {
	Create(activityCore Core)
	Update(activityCore Core)
	Delete(activityCore Core)
	GetAll() []Core
	Get(activityCore Core) Core
}
