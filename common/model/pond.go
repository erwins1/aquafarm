package model

import (
	"aquafarm/common/constant"
	"errors"
	"net/http"
	"strconv"
)

type Pond struct {
	PondID      int64  `json:"pond_id" db:"pond_id"`
	FarmID      int64  `json:"farm_id" db:"farm_id"`
	PondName    string `json:"pond_name" db:"pond_name"`
	Description string `json:"description" db:"description"`
}

type GetPondByFarmIDReq struct {
	FarmID int64
}

type InsertPondReq struct {
	FarmID      int64  `json:"farm_id"`
	PondName    string `json:"pond_name"`
	Description string `json:"description"`
}

func (m *InsertPondReq) Validate() error {
	if m.PondName == "" {
		return errors.New("pond name cannot be empty")
	}

	if m.FarmID == 0 {
		return errors.New("farm id cannot be empty")
	}

	return nil
}

type GetPondByNameReq struct {
	PondName string `json:"pond_name"`
}

type GetPondByNameRes struct {
	PondID      int64  `json:"pond_id" db:"pond_id"`
	FarmID      int64  `json:"farm_id" db:"farm_id"`
	PondName    string `json:"pond_name" db:"pond_name"`
	Description string `json:"description" db:"description"`
}

type GetPondReq struct {
	Page    int64 `json:"page"`
	PerPage int64 `json:"per_page"`
}

func (m *GetPondReq) ParseAndValidate(r *http.Request) error {
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

type GetPondByID struct {
	PondID int64 `json:"pond_id"`
}

type DeletePondByID struct {
	PondID int64 `json:"pond_id"`
}

func (m *DeletePondByID) Validate() error {
	if m.PondID == 0 {
		return errors.New("pond id cannot be empty")
	}

	return nil
}
