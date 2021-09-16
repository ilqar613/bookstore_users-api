package users

import (
	"fmt"

	usersdb "github.com/ilqar613/bookstore_users-api/datasources/mysql/users_db"
	"github.com/ilqar613/bookstore_users-api/logger"
	"github.com/ilqar613/bookstore_users-api/utils/errors"
)

const (
	//indexUniqueEmail = "email_UNIQUE"
	//errorNoRows      = "no rows in result set"
	queryInsertUser       = "INSERT INTO users(first_name,last_name,email,date_created,status,password) VALUES (?, ?, ?, ?, ?, ?);"
	queryGetUser          = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	queryUpdateUser       = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
	queryDeleteUser       = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?;"
)

//get data
func (user *User) Get() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)
	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); getErr != nil {
		logger.Error("error when trying to  get user by id", getErr)
		return errors.NewInternalServerError("database error")
		//return mysqlutils.ParseError(getErr)
	}

	return nil
}

// save data
func (user *User) Save() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare save user statement", err)
		return errors.NewInternalServerError("database error")
		//return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to  save user", saveErr)
		return errors.NewInternalServerError("database error")
		//return mysqlutils.ParseError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last id after creating a new user", err)
		return errors.NewInternalServerError("database error")
		//return mysqlutils.ParseError(err)
	}
	user.Id = userId
	return nil
}

//update data
func (user *User) Update() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
		//return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		logger.Error("error when trying to update user", err)
		return errors.NewInternalServerError("database error")
		//return mysqlutils.ParseError(err)
	}
	return nil
}

// delete data
func (user *User) Delete() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
		//return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		logger.Error("error when trying to delete user", err)
		return errors.NewInternalServerError("database error")
		//return mysqlutils.ParseError(err)
	}
	return nil
}

// find data
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, errors.NewInternalServerError("database error")
		//return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to find user by status", err)
		return nil, errors.NewInternalServerError("database error")
		//return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to scan user into user struct", err)
			return nil, errors.NewInternalServerError("database error")
			//return nil, mysqlutils.ParseError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NewNotFoundRequestError(fmt.Sprintf("no users matching status %s", status))
	}

	return results, nil
}
