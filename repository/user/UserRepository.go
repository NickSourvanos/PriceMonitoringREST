package user

import (
	"PriceMonitoringService/models"
	"database/sql"
)

type userRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) *userRepositoryPostgres {
	return &userRepositoryPostgres{db}
}

func (r *userRepositoryPostgres) Save(user *models.User) error {
	var query = `insert into public.users(firstName, lastName, username, password, email, role, active) 
				values($1, $2, $3, $4, $5, $6, $7)`
	statement, err := r.db.Prepare(query)
	if err != nil {	return err	}
	defer statement.Close()

	_, err = statement.Exec(user.FirstName, user.LastName, user.Username, user.Password, user.Email, user.Role, user.Active)
	if err != nil {	return err	}

	return nil
}

func (r *userRepositoryPostgres) Update(userId string, user *models.User) error {
	query := `UPDATE public.users SET firstName=$1, lastName=$2, username=$3, password=$4, email=$5, role=$6, active=$7 WHERE id = $8`

	statement, err := r.db.Prepare(query)
	if err != nil { return err }
	defer statement.Close()

	_, err = statement.Exec(user.FirstName, user.LastName, user.Username, user.Password, user.Email, user.Role, user.Active, userId)
	if err != nil { return err }

	return nil
}

func (r *userRepositoryPostgres) Delete(userId string) error {
	query := `DELETE FROM public.users WHERE id = $1`

	statement, err := r.db.Prepare(query)
	if err != nil { return err }
	defer statement.Close()

	_, err = statement.Exec(userId)
	if err != nil { return err }

	return nil
}

func (r *userRepositoryPostgres) FindByID(userId string) (*models.ResponseUser, error) {
	query := `SELECT id, firstName, lastName, username, email, role, active FROM public.users WHERE id = $1`

	var user models.ResponseUser

	statement, err := r.db.Prepare(query)
	if err != nil { return nil, err	}

	defer statement.Close()
	err = statement.QueryRow(userId).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Email, &user.Role, &user.Active)
	if err != nil {	return nil, err	}

	return &user, nil
}

func (r *userRepositoryPostgres) FindByUsername(username string) (*models.User, error) {
	query := `SELECT id, firstName, lastName, username, password, email, role, active FROM public.users WHERE username = $1`

	var user models.User

	statement, err := r.db.Prepare(query)
	if err != nil { return nil, err	}

	defer statement.Close()
	err = statement.QueryRow(username).Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Password, &user.Email, &user.Role, &user.Active)
	if err != nil {	return nil, err	}

	return &user, nil
}


func (r *userRepositoryPostgres) FindAll() (models.ResponseUsers, error) {
	query := `SELECT id, firstName, lastName, username, email, role, active FROM public.users`

	var users models.ResponseUsers
	rows, err := r.db.Query(query)

	if err != nil {	return nil, err	}
	defer rows.Close()

	for rows.Next() {
		var user models.ResponseUser
		err = rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Username, &user.Email,
			&user.Role, &user.Active)

		if err != nil {	return nil, err	}
		users = append(users, user)
	}
	return users, nil
}
