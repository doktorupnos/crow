package channel

type Service struct {
	r Repo
}

func NewService(r Repo) *Service {
	return &Service{r}
}

func (s *Service) GetByName(name string) (Channel, error) {
	return s.r.GetByName(name)
}
