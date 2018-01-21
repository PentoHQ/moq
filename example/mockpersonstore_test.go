package example

const PersonStoreErrorMessage = "PersonStore-error"

type PersonStoreOpts struct {
	IsCreateError bool
	
	IsGetError bool
	GetResponse *Person
}

type MockPersonStore struct {
	MockCreate func(ctx context.Context, person *Person, confirm bool) error
	createCalls int
	MockGet func(ctx context.Context, id string) (*Person, error)
	getCalls int
}


func (mock *MockPersonStore) Create(ctx context.Context, person *Person, confirm bool) error {
	mock.createCalls++
	return mock.MockCreate(ctx, person, confirm)
}

func (mock *MockPersonStore) CreateCalls() int {
	return mock.createCalls
}

func (mock *MockPersonStore) Get(ctx context.Context, id string) (*Person, error) {
	mock.getCalls++
	return mock.MockGet(ctx, id)
}

func (mock *MockPersonStore) GetCalls() int {
	return mock.getCalls
}

func (mock *MockPersonStore) SetOpts (opts PersonStoreOpts) {
	mock.MockCreate = func(ctx context.Context, person *Person, confirm bool) error {
		if opts.IsCreateError {
			return errors.New(PersonStoreErrorMessage)
		}
		if opts.CreateResponse != nil {
			return nil
		}
		return nil
	}
	mock.MockGet = func(ctx context.Context, id string) (*Person, error) {
		if opts.IsGetError {
			return nil, errors.New(PersonStoreErrorMessage)
		}
		if opts.GetResponse != nil {
			return nil, nil
		}
		return nil, nil
	}
}

func NewMockPersonStore(opts PersonStoreOpts) *MockPersonStore {
	mock := new(MockPersonStore)
	mock.SetOpts(opts)
	return mock
}
