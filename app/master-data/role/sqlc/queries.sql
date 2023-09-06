-- name: SelectAllRole :many
select * from master_role;
-- name: SelectOneRole :one
select * from master_role where role_name = $1;