package usercore

type IRepoUser interface {
	Insert(activityCore Core) error
	Update(activityCore Core) error
	Delete(activityCore Core) error
	FindAll(activityCore Core) ([]Core, error)
	First(activityCore Core) (Core, error)
}

type IServiceUser interface {
	Create(activityCore Core)
	Update(activityCore Core)
	Delete(activityCore Core)
	FindAll(activityCore Core) []Core
	First(activityCore Core) Core
}
