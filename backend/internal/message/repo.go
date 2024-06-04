package message

type Repo interface {
	Create(*Message) error
}
