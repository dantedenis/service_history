package contract

import "context"

type IService interface {
	Run(ctx context.Context) error
	Stop() error
}
