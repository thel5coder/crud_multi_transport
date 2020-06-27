package actions

import (
	"crud_multi_transport/db/models"
	"crud_multi_transport/db/repositories/contracts"
	"crud_multi_transport/helpers/str"
	"crud_multi_transport/usecase/viewmodel"
	"database/sql"
	"fmt"
	"strings"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) contracts.IUserRepository {
	return UserRepository{DB: DB}
}

func (repository UserRepository) Browse(search, order, sort string, limit, offset int) (res []models.Users, count int, err error) {
	statement := `select * from "users" where (lower("full_name") like $1 or lower("email") like $1 or "mobile_phone" like $1) and "deleted_at" is null order by ` + order + ` ` + sort + ` limit $2 offset $3`
	rows, err := repository.DB.Query(statement, "%"+strings.ToLower(search)+"%", limit, offset)
	if err != nil {
		return res, count, err
	}

	for rows.Next() {
		temp := models.Users{}
		err = rows.Scan(&temp.ID, &temp.FullName, &temp.Email, &temp.Password, &temp.MobilePhone, &temp.CreatedAt, &temp.UpdatedAt, &temp.DeletedAt)
		if err != nil {
			return res, count, err
		}
		res = append(res, temp)
	}

	statement = `select count("id") from "users" where (lower("full_name") like $1 or lower("email") like $1 or "mobile_phone" like $1) and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, "%"+strings.ToLower(search)+"%").Scan(&count)
	if err != nil {
		return res, count, err
	}

	return res, count, err
}

func (repository UserRepository) ReadBy(column,value string) (res models.Users, err error) {
	statement := `select * from "users" where `+ column+ `=$1 and "deleted_at" is null`
	err = repository.DB.QueryRow(statement, value).Scan(&res.ID, &res.FullName, &res.Email, &res.Password, &res.MobilePhone, &res.CreatedAt, &res.UpdatedAt, &res.DeletedAt)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository UserRepository) Edit(body viewmodel.UserVm, password string) (res string, err error) {
	if password == "" {
		statement := `update "users" set "full_name"=$1, "email"=$2, "mobile_phone"=$3, "updated_at"=$4 where "id"=$5 returning "id"`
		err = repository.DB.QueryRow(statement, body.FullName, body.Email, body.MobilePhone, str.StrParseToTime(body.UpdatedAt, time.RFC3339), body.ID).Scan(&res)
	} else {
		statement := `update "users" set "full_name"=$1, "email"=$2, "password"=$3, "mobile_phone"=$4, "updated_at"=$5 where "id"=$6 returning "id"`
		err = repository.DB.QueryRow(statement, body.FullName, body.Email, password, body.MobilePhone, str.StrParseToTime(body.UpdatedAt, time.RFC3339), body.ID).Scan(&res)
	}

	if err != nil {
		return res, err
	}

	return res, err
}

func (repository UserRepository) Add(body viewmodel.UserVm, password string) (res string, err error) {
	statement := `insert into "users" ("full_name","email","password","mobile_phone","created_at","updated_at") values($1,$2,$3,$4,$5,$6) returning "id"`
	err = repository.DB.QueryRow(statement, body.FullName, body.Email, password, body.MobilePhone, str.StrParseToTime(body.CreatedAt, time.RFC3339),str.StrParseToTime(body.UpdatedAt,time.RFC3339)).Scan(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (repository UserRepository) Delete(ID ,updatedAt, deletedAt string) (res string, err error) {
	fmt.Println(updatedAt)
	statement := `update "users" set "updated_at"=$1, "deleted_at"=$2 where "id"=$3 returning "id"`
	err = repository.DB.QueryRow(statement, str.StrParseToTime(updatedAt, time.RFC3339), str.StrParseToTime(deletedAt, time.RFC3339),ID).Scan(&res)
	if err != nil {
		fmt.Println(err)
		return res, err
	}

	return res, err
}

func (repository UserRepository) IsExist(ID, email,mobilePhone string) (res bool, err error) {
	var count int
	if ID == "" {
		statement := `select count("id") from "users" where ("email"=$1 or "mobile_phone"=$2) and "deleted_at" is null`
		err = repository.DB.QueryRow(statement,email,mobilePhone).Scan(&count)
	}else{
		statement := `select count("id") from "users" where ("email"=$1 or "mobile_phone"=$2) and "deleted_at" is null and "id" <> $3`
		err = repository.DB.QueryRow(statement,email,mobilePhone,ID).Scan(&count)
	}

	return count > 0, err
}
