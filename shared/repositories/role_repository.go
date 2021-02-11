package repositories

import (
	"database/sql"
	"errors"
)

// Typology and naming of internal repositories.
const (
	ROLE_COMPANY_MANAGER = "ROLE_COMPANY_MANAGER"
	ROLE_EXPERT          = "ROLE_EXPERT"
)

// RoleRepository handled DB operatios relative to the users repositories.
type RoleRepository struct {
	DB *sql.DB
}

// UserRole contains partial informations about a users repositories. Mainly used to send a different mail depending on the role.
type UserRole struct {
	RoleID int64
	Name   string
}

// FindRoleByUserID returns the UserRole of the given user id (gathered from phishing campaign initiator.)
func (repo *RoleRepository) FindRolesByUserUUID(userUUID string) ([]*UserRole, error) {
	var roleId int64
	var roleName string
	var userRoleList []*UserRole

	rows, err := repo.DB.Query(`
SELECT 
	r.id role_id,
	r.name role_name
FROM 
	ht_lk_users_roles ur
	LEFT JOIN ht_role r ON ur.role_id = r.id
	LEFT JOIN ht_user u ON ur.user_id = u.id
WHERE
	u.`+"`uuid`"+` = ?`,
		userUUID,
	)

	switch err {
	case sql.ErrNoRows:
		return nil, errors.New("no rows found")
	case nil:
		for rows.Next() {
			err := rows.Scan(&roleId, &roleName)
			if err == nil {
				userRole := UserRole{
					RoleID: roleId,
					Name:   roleName,
				}
				userRoleList = append(userRoleList, &userRole)
			}
		}
		return userRoleList, nil
	default:
		return nil, err
	}
}

// CheckCompanyManagerRole takes the user role list and return a boolean depending on the roleName argument.
func (repo *RoleRepository) CheckCompanyManagerRole(roleList []*UserRole, roleName string) bool {
	var userRole *UserRole
	for _, userRole = range roleList {
		if userRole.Name == roleName {
			return true
		}
	}
	return false
}
