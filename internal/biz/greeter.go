package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound("ErrUserNotFound", "user not found")
)

// Greeter is a Greeter model.
type Mess struct {
	Info string
}

// DataController is a Greater repo.
type DataController interface {
	FindByID(context.Context, int64) (*Mess, error)
}

// DataControllerUsecase is a Greeter usecase.
type DataControllerUsecase struct {
	repo DataController
	log  *log.Helper
}

// NewDataControllerUsecase new a Greeter usecase.
func NewDataControllerUsecase(repo DataController, logger log.Logger) *DataControllerUsecase {
	return &DataControllerUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *DataControllerUsecase) CreateMess(ctx context.Context, id int64) (*Mess, error) {
	me, err := uc.repo.FindByID(ctx, id)
	return me, err
	// return &Mess{Info: "wyq"}, nil
}
