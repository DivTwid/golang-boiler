package service

type DummyService interface {
	GetVal() string
	AddVal() string
}
type dummyService struct {
}

func NewDummyService() DummyService {
	return &dummyService{}
}

func (ds dummyService) GetVal() string {
	return "pong"
}

func (ds dummyService) AddVal() string {
	return "something"
}
