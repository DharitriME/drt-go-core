package mock

import "github.com/DharitriME/drt-go-core/data"

// EmptyBlockCreatorStub -
type EmptyBlockCreatorStub struct {
	CreateNewHeaderCalled func() data.HeaderHandler
}

// CreateNewHeader -
func (e *EmptyBlockCreatorStub) CreateNewHeader() data.HeaderHandler {
	if e.CreateNewHeaderCalled != nil {
		return e.CreateNewHeaderCalled()
	}
	return nil
}

// IsInterfaceNil -
func (e *EmptyBlockCreatorStub) IsInterfaceNil() bool {
	return e == nil
}
