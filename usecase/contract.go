package usecase

import (
	"crud_multi_transport/helpers/jwe"
	"crud_multi_transport/helpers/jwt"
	"crud_multi_transport/usecase/viewmodel"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/middleware"
	"time"
)

const (
	defaultLimit    = 10
	maxLimit        = 50
	defaultOrderBy  = "id"
	defaultSort     = "asc"
	defaultLastPage = 0
)

type UcContract struct {
	DB *sql.DB
	Redis *redis.Client
	Jwe jwe.Credential
	JwtConfig middleware.JWTConfig
	JwtCred jwt.JwtCredential
}


func (uc UcContract) setPaginationParameter(page, limit int, order, sort string) (int, int, int, string, string) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 || limit > maxLimit {
		limit = defaultLimit
	}

	if order == "" {
		order = defaultOrderBy
	}

	if sort == "" {
		sort = defaultSort
	}

	offset := (page - 1) * limit

	return offset, limit, page, order, sort
}

func (uc UcContract) setPaginationResponse(page, limit, total int) (paginationResponse viewmodel.PaginationVm) {
	var lastPage int

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = defaultLastPage
	}

	paginationResponse = viewmodel.PaginationVm{
		CurrentPage: page,
		LastPage:    lastPage,
		Total:       total,
		PerPage:     limit,
	}

	return paginationResponse
}

func (uc UcContract) StoreToRedistWithExpired(key string, val interface{}, duration string) error {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}

	b, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = uc.Redis.Set(key, string(b), dur).Err()

	return err
}

func (uc UcContract) StoreToRedis(key string, val interface{}) error {
	b, err := json.Marshal(val)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = uc.Redis.Set(key, string(b), 0).Err()

	return err
}

func (uc UcContract) GetFromRedis(key string, cb interface{}) error {
	res, err := uc.Redis.Get(key).Result()
	if err != nil {
		return err
	}

	if res == "" {
		return errors.New("[Redis] Value of " + key + " is empty.")
	}

	err = json.Unmarshal([]byte(res), &cb)
	if err != nil {
		return err
	}

	return err
}

func (uc UcContract) RemoveFromRedis(key string) error {
	return uc.Redis.Del(key).Err()
}