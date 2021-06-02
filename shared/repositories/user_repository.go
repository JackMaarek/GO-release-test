package repositories

import (
	"database/sql"
	"errors"
)

// UserRepository handled DB operations relative to the users repositories.
type UserRepository struct {
	DB *sql.DB
}

// User contains partial informations about a user.
type User struct {
	Name      string
	Firstname string
	Email     string
}

// FindUsersByCompanyUUIDAndRoleName returns a slice of User informations by the pretty self explanatory conditions.
func (repo *Repository) FindUsersByCompanyUUIDAndRoleName(companyUUID string, roleName string) ([]*User, error) {
	var userFirstname, userLastname, userEmail string
	var users []*User

	rows, err := repo.DB.Query(`
SELECT 
	u.first_name firstname,
	u.last_name lastname,
	u.email email 
FROM
	ht_user u 
	LEFT JOIN ht_lk_users_roles ur ON ur.user_id = u.id 
	LEFT JOIN ht_role r ON r.id = ur.role_id
	LEFT JOIN ht_company c ON u.company_id = c.id
WHERE 
	c.`+"`uuid`"+`= ?
	AND r.name = ?`,
		companyUUID,
		roleName,
	)
	switch err {
	case sql.ErrNoRows:
		return nil, errors.New("no rows found")
	case nil:
		for rows.Next() {
			err := rows.Scan(&userFirstname, &userLastname, &userEmail)
			if err == nil {
				user := User{
					Name:      userLastname,
					Firstname: userFirstname,
					Email:     userEmail,
				}
				users = append(users, &user)
			}
		}
		return users, nil
	default:
		return nil, err
	}
}

// FindInitiatorInformationByUserUUID returns the initiator User informations.
func (repo *Repository) FindInitiatorInformationByUserUUID(userUUID string) (*User, error) {
	var userFirstname, userLastname, userEmail string
	row := repo.DB.QueryRow(`
SELECT 
	u.first_name firstname,
	u.last_name lastname,
	u.email email 
FROM
	ht_user u
WHERE 
	u.`+"`uuid`"+`= ?`,
		userUUID,
	)

	switch err := row.Scan(&userFirstname, &userLastname, &userEmail); err {
	case sql.ErrNoRows:
		return nil, errors.New("no rows found")
	case nil:
		user := User{
			Name:      userFirstname,
			Firstname: userLastname,
			Email:     userEmail,
		}
		return &user, nil
	default:
		return nil, err
	}
}
