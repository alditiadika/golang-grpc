package role

import (
	"context"
	master_role "golang-grpc/app/master-data/role/proto"

	"go-micro.dev/v4"
)

type handler struct {
	usecase *usecase
}

func Register(service micro.Service) {
	master_role.RegisterMasterRoleServiceHandler(service.Server(), &handler{
		usecase: NewRoleUseCase(),
	})
}
func (h *handler) GetAllRole(ctx context.Context, in *master_role.GetAllRoleRequest, out *master_role.GetAllRoleResponse) error {
	return h.usecase.GetAllRole(ctx, in, out)
}
func (h *handler) GetOneUser(ctx context.Context, in *master_role.GetOneRoleRequest, out *master_role.GetOneRoleResponse) error {
	return h.usecase.GetOneRole(ctx, in, out)
}
