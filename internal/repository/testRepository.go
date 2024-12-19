package repository

type TestRepository interface {
	Test() (string, error)
}

type TestRepositoryImpl struct {
}

func TestRepositoryInit() *TestRepositoryImpl {
	return &TestRepositoryImpl{}
}

func (u *TestRepositoryImpl) Test() (string, error) {

	return "test", nil
}
