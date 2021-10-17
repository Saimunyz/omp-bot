package receipt

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Receipt {
	return allEntities
}

func (s *Service) Get(idx int) (*Receipt, error) {
	return &allEntities[idx], nil
}
