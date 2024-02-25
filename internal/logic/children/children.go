package children

import (
	"context"
)

type sChildren struct {
}

func init() {
	//service.RegisterChildren(New())
}
func New() *sChildren {
	return &sChildren{}
}
func (s *sChildren) CheckClassMems(ctx context.Context, cid string, mid string) {

}
