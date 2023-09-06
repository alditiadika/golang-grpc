package role

import (
	"context"
	master_role "golang-grpc/app/master-data/role/proto"
	role_sql "golang-grpc/app/master-data/role/sqlc"
	"golang-grpc/config"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type usecase struct {
	repository *role_sql.Queries
}

func NewRoleUseCase() *usecase {
	return &usecase{
		repository: role_sql.New(config.Application.DB),
	}
}
func (u *usecase) GetAllRole(ctx context.Context, req *master_role.GetAllRoleRequest, res *master_role.GetAllRoleResponse) error {
	var err error
	var data []*master_role.RoleItem
	dataSelect, err := u.repository.SelectAllRole(ctx)
	if err != nil {
		res.IsError = true
		res.ErrMessage = err.Error()
		return err
	}
	for _, item := range dataSelect {
		dataitem := &master_role.RoleItem{}
		dataitem.ID = item.ID
		dataitem.RoleCode = item.RoleCode
		dataitem.RoleName = item.RoleName
		dataitem.IsActive = item.IsActive
		dataitem.CreatedBy = item.CreatedBy
		dataitem.CreatedDate = timestamppb.New(item.CreatedDate)
		dataitem.ModifiedBy = item.ModifiedBy
		dataitem.ModifiedDate = timestamppb.New(item.ModifiedDate)
		data = append(data, dataitem)
	}
	res.Data = data
	res.IsError = false
	return nil
}
func (u *usecase) GetOneRole(ctx context.Context, req *master_role.GetOneRoleRequest, res *master_role.GetOneRoleResponse) error {
	var err error
	// var data role_sql.MasterRole
	dataRepo, err := u.repository.SelectOneRole(ctx, req.RoleName)
	if err != nil {
		res.IsError = true
		res.ErrMessage = err.Error()
		return err
	}
	data := &master_role.RoleItem{}
	data.ID = dataRepo.ID
	data.RoleCode = dataRepo.RoleCode
	data.RoleName = dataRepo.RoleName
	data.IsActive = dataRepo.IsActive
	data.CreatedBy = dataRepo.CreatedBy
	data.CreatedDate = timestamppb.New(dataRepo.CreatedDate)
	data.ModifiedBy = dataRepo.ModifiedBy
	data.ModifiedDate = timestamppb.New(dataRepo.ModifiedDate)
	res.Data = data
	res.IsError = false
	return nil
}
