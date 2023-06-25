package model

import (
	"aquafarm/common/constant"
	"errors"
	"net/http"
	"strconv"
)

type InsertFarmReq struct {
	FarmName string `json:"farm_name"`
}

func (m *InsertFarmReq) Validate() error {
	if m.FarmName == "" {
		return errors.New("farm name cannot be empty")
	}

	return nil
}

type GetFarmByNameReq struct {
	FarmName string `json:"farm_name"`
}

type GetFarmByNameRes struct {
	FarmID   int64  `json:"farm_id" db:"farm_id"`
	FarmName string `json:"farm_name" db:"farm_name"`
}

type GetFarmReq struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"per_page"`
}

func (m *GetFarmReq) ParseAndValidate(r *http.Request) error {
	var (
		err error
	)
	page := r.FormValue("page")
	if page == "" {
		m.Page = constant.DefaultPaginationPage
	}

	m.Page, err = strconv.ParseInt(page, 10, 64)
	if err != nil {
		return err
	}

	perPage := r.FormValue("per_page")
	if perPage == "" {
		m.PerPage = constant.DefaultPaginationPerPage
	}

	m.PerPage, err = strconv.ParseInt(perPage, 10, 64)
	if err != nil {
		return err
	}

	return err
}

type GetFarmResult struct {
	FarmID   int64  `json:"farm_id" db:"farm_id"`
	FarmName string `json:"farm_name" db:"farm_name"`
}

type GetFarmByIDReq struct {
	FarmID int64 `json:"farm_id"`
}

type GetFarmByIDRes struct {
	FarmID   int64  `json:"farm_id" db:"farm_id"`
	FarmName string `json:"farm_name" db:"farm_name"`
}
