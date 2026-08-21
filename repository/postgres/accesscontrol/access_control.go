package psqlaccesscontrol

import (
	"fmt"
	"github.com/lib/pq"
	"github.com/younesbeheshti/gocast_game/entity"
	"github.com/younesbeheshti/gocast_game/pkg/richerror"
	"github.com/younesbeheshti/gocast_game/pkg/slice"
	"github.com/younesbeheshti/gocast_game/repository/postgres"
	"time"
)

func (d *DB) GetUserPermissionTitles(userID uint, role entity.Role) ([]entity.PermissionTitle, error) {
	const op = "psql.GetUserPermissionTitles"

	rows, err := d.conn.Conn().Query(`select * from access_controls where actor_type=$1 and actor_id=$2`, entity.RoleActorType, role)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage("something went wrong").WithKind(richerror.KindUnexpected)
	}

	roleACL := make([]entity.AccessControl, 0)

	for rows.Next() {
		acl, err := scanAccessControl(rows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage("something went wrong").WithKind(richerror.KindUnexpected)
		}
		roleACL = append(roleACL, *acl)

	}

	if rows.Err() != nil {
		return nil, richerror.New(op).WithErr(rows.Err()).WithKind(richerror.KindUnexpected).WithMessage("something went wrong")
	}

	defer rows.Close()

	userRows, err := d.conn.Conn().Query(`select * from access_controls where actor_type = $1 and actor_id=$2`, entity.UserActorType, userID)
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithMessage("something went wrong").WithKind(richerror.KindUnexpected)
	}

	userACL := make([]entity.AccessControl, 0)

	for userRows.Next() {
		acl, err := scanAccessControl(userRows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage("something went wrong").WithKind(richerror.KindUnexpected)
		}
		userACL = append(userACL, *acl)

	}

	if userRows.Err() != nil {
		return nil, richerror.New(op).WithErr(rows.Err()).WithKind(richerror.KindUnexpected).WithMessage("something went wrong")
	}

	defer userRows.Close()

	// merge ACLs by permission id
	permissionIDs := make([]uint, 0)
	for _, acl := range roleACL {
		if !slice.DoesExist(permissionIDs, acl.PermissionID) {
			permissionIDs = append(permissionIDs, acl.PermissionID)
		}
	}

	if len(permissionIDs) == 0 {
		return nil, nil
	}

	args := make([]any, len(permissionIDs))
	for i, id := range permissionIDs {
		args[i] = id
	}

	pRows, err := d.conn.Conn().Query(`select * from permissions where id = any($1)`, pq.Array(args))
	fmt.Println("helloo")
	if err != nil {
		fmt.Println("log", err)
		return nil, richerror.New(op).WithErr(err).WithMessage(fmt.Sprint("something went wrong")).WithKind(richerror.KindUnexpected)
	}
	defer pRows.Close()

	permissionTitles := make([]entity.PermissionTitle, 0)
	for pRows.Next() {
		permission, err := scanPermission(pRows)
		if err != nil {
			return nil, richerror.New(op).WithErr(err).WithMessage(fmt.Sprint("something went wrong")).WithKind(richerror.KindUnexpected)
		}

		permissionTitles = append(permissionTitles, permission.Title)
	}

	if pRows.Err() != nil {
		return nil, richerror.New(op).WithErr(pRows.Err()).WithKind(richerror.KindUnexpected).WithMessage("something went wrong")
	}

	return permissionTitles, nil
}

func scanAccessControl(scanner postgres.Scanner) (*entity.AccessControl, error) {
	var acl entity.AccessControl
	var createdAt time.Time
	err := scanner.Scan(&acl.ID, &acl.ActorID, &acl.ActorType, &acl.PermissionID, &createdAt)
	return &acl, err
}
