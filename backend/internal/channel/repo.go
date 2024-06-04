package channel

type Repo interface {
	GetByName(string) (Channel, error)
}
