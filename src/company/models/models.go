package models

type (
	Company struct {
		CompanyID int     `json:"id"`
		Name      string  `json:"name"`
		Address   string  `json:"address"`
		City      string  `json:"city"`
		Country   string  `json:"country"`
		Email     string  `json:"email"`
		Phone     string  `json:"phone"`
		Owners    []Owner `json:"owners"`
	}
)

type (
	CompanyResume struct {
		CompanyID int    `json:"id"`
		Name      string `json:"name"`
	}
)

type (
	Owner struct {
		OwnerID int    `json:"id"`
		Name    string `json:"name"`
	}
)

type (
	CompanyOwners struct {
		CompanyID int `json:"id"`
		OwnerID   int `json:"id"`
	}
)
