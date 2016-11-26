package dbsql

import (
	"database/sql"
	"log"

	"company/models"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./Company.db")
	if err != nil {
		log.Fatal(err)
	}

	sqlStmt := `
	create table if not exists Owners (
		OwnerID INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		Name CHAR(50) NOT NULL
	);
	create table if not exists Companies (
		CompanyID INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
		Name CHAR(50) NOT NULL,
		Address CHAR(50) NOT NULL,
		City CHAR(50) NOT NULL,
		Country CHAR(50) NOT NULL,
		Email CHAR(50) NULL,
		Phone CHAR(50) NULL
	);
	create table if not exists CompanyOwners(
		Id INTEGER PRIMARY KEY AUTOINCREMENT,
		CompanyID INTEGER REFERENCES Companies(CompanyID),
		OwnerID INTEGER REFERENCES Owners(OwnerID)
	);
	insert or ignore into Companies (CompanyID, Name, Address, City, Country, Email, Phone) values (1, 'Buzzshare', '9479 Oakridge Court', 'Oakland', 'United States', 'lharris0@rakuten.co.jp', '1-(415)603-9558');
	insert or ignore into Companies (CompanyID, Name, Address, City, Country, Email, Phone) values (2, 'Linkbridge', '976 Melrose Court', 'Cotmon', 'Philippines', 'dmiller1@cnet.com', '63-(264)803-3205');
	insert or ignore into Companies (CompanyID, Name, Address, City, Country, Email, Phone) values (3, 'Vitz', '54 Twin Pines Terrace', 'Mollebamba', 'Peru', 'kmorales2@soundcloud.com', '51-(112)547-3580');
	insert or ignore into Companies (CompanyID, Name, Address, City, Country, Email, Phone) values (4, 'Bluezoom', '8 Pond Street', 'Goianésia', 'Brazil', 'kallen3@surveymonkey.com', '55-(100)903-7431');
	insert or ignore into Companies (CompanyID, Name, Address, City, Country, Email, Phone) values (5, 'Jetwire', '27 Mariners Cove Parkway', 'Abha', 'Saudi Arabia', 'sperkins4@cmu.edu', '966-(806)822-9186');
	insert or ignore into Companies (CompanyID, Name, Address, City, Country, Email, Phone) values (6, 'Youspan', '722 Swallow Hill', 'Périgueux', 'France', 'bberry5@multiply.com', '33-(952)485-8378');
	insert or ignore into Companies (CompanyID, Name, Address, City, Country, Email, Phone) values (7, 'Viva', '21996 Bultman Alley', 'Putinci', 'Serbia', 'sowens6@businessinsider.com', '381-(125)749-4462');
	insert or ignore into Companies (CompanyID, Name, Address, City, Country, Email, Phone) values (8, 'Feednation', '6091 Ohio Street', 'Sykiés', 'Greece', 'nvasquez7@wordpress.com', '30-(506)783-2157');
	insert or ignore into Companies (CompanyID, Name, Address, City, Country, Email, Phone) values (9, 'Meedoo', '59 Green Ridge Way', 'Ribeira', 'Portugal', null, null);
	insert or ignore into Companies (CompanyID, Name, Address, City, Country, Email, Phone) values (10, 'Omba', '29221 Banding Alley', 'Guhuai', 'China', 'ahawkins9@boston.com', '86-(378)697-2261');
	insert or ignore into Owners (OwnerID, Name) values (1, 'Amanda');
	insert or ignore into Owners (OwnerID, Name) values (2, 'Richard');
	insert or ignore into Owners (OwnerID, Name) values (3, 'Howard');
	insert or ignore into Owners (OwnerID, Name) values (4, 'Cheryl');
	insert or ignore into Owners (OwnerID, Name) values (5, 'Denise');
	insert or ignore into CompanyOwners (Id, CompanyID, OwnerID) values (1, 1, 1);
	insert or ignore into CompanyOwners (Id, CompanyID, OwnerID) values (2, 1, 2);
	insert or ignore into CompanyOwners (Id, CompanyID, OwnerID) values (3, 4, 4);
	insert or ignore into CompanyOwners (Id, CompanyID, OwnerID) values (4, 8, 3);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return nil
	}

	return db
}

func addCompanyOwners(db *sql.DB, company *models.Company) error {
	sqlStmt := `
		delete from CompanyOwners where CompanyID = ?
		`
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, stmtErr := stmt.Exec(company.CompanyID)
	if stmtErr != nil {
		return stmtErr
	}

	sqlStmtInsert := `
		insert into CompanyOwners (CompanyID, OwnerID) values (?, ?);
		`
	stmtInsert, err := db.Prepare(sqlStmtInsert)
	if err != nil {
		return err
	}
	defer stmtInsert.Close()

	for i := 0; i < len(company.Owners); i++ {
		_, stmtErr := stmtInsert.Exec(company.CompanyID, company.Owners[i].OwnerID)
		if stmtErr != nil {
			return stmtErr
		}
	}

	return nil
}

func UpdateCompany(db *sql.DB, company models.Company) error {
	sqlStmt := `
		REPLACE INTO Companies(
			CompanyID,
			Name,
			Address,
			City,
			Country,
			Email,
			Phone
		) values (?, ?, ?, ?, ?, ?, ?)
		`

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, stmtErr := stmt.Exec(
		company.CompanyID,
		company.Name,
		company.Address,
		company.City,
		company.Country,
		company.Email,
		company.Phone)
	if stmtErr != nil {
		return stmtErr
	}

	errCompanyOwners := addCompanyOwners(db, &company)
	if errCompanyOwners != nil {
		return errCompanyOwners
	}

	return nil
}

func AddCompany(db *sql.DB, company models.Company) error {
	sqlStmt := `
		INSERT INTO Companies(
			Name,
			Address,
			City,
			Country,
			Email,
			Phone
		) values (?, ?, ?, ?, ?, ?)
		`

	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, stmtErr := stmt.Exec(
		company.Name,
		company.Address,
		company.City,
		company.Country,
		company.Email,
		company.Phone)
	if stmtErr != nil {
		return stmtErr
	}

	errCompanyOwners := addCompanyOwners(db, &company)
	if errCompanyOwners != nil {
		return errCompanyOwners
	}

	return nil
}

func GetAll(db *sql.DB) ([]models.Company, error) {
	sqlStmt := `
		select 
			CompanyID,
			Name,
			Address,
			City,
			Country,
			coalesce(Email,''),
			coalesce(Phone,'')
		from Companies
		`

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Company
	for rows.Next() {
		item := models.Company{}
		rowsErr := rows.Scan(
			&item.CompanyID,
			&item.Name,
			&item.Address,
			&item.City,
			&item.Country,
			&item.Email,
			&item.Phone)
		if rowsErr != nil {
			return nil, rowsErr
		}
		result = append(result, item)
	}

	return result, nil
}

func GetAllResume(db *sql.DB) ([]models.Company, error) {
	sqlStmt := `
		select 
			CompanyID,
			Name
		from Companies
		`

	rows, err := db.Query(sqlStmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Company
	for rows.Next() {
		item := models.Company{}
		rowsErr := rows.Scan(
			&item.CompanyID,
			&item.Name)
		if rowsErr != nil {
			return nil, rowsErr
		}
		result = append(result, item)
	}

	return result, nil
}

func GetById(db *sql.DB, id int) (*models.Company, error) {
	result := models.Company{}
	err := db.QueryRow(`
		select 
			CompanyID,
			Name,
			Address,
			City,
			Country,
			coalesce(Email,''),
			coalesce(Phone,'')
		from Companies
		where CompanyID = ?
		`, id).Scan(
		&result.CompanyID,
		&result.Name,
		&result.Address,
		&result.City,
		&result.Country,
		&result.Email,
		&result.Phone)
	if err != nil {
		return nil, err
	}

	sqlStmt := `
		select 
			o.OwnerID,
			o.Name
		from CompanyOwners co
		inner join Owners o on co.OwnerID = o.OwnerID 
		where co.CompanyID = ?
		`
	rows, err := db.Query(sqlStmt, result.CompanyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var owners []models.Owner
	for rows.Next() {
		item := models.Owner{}
		rowsErr := rows.Scan(
			&item.OwnerID,
			&item.Name)
		if rowsErr != nil {
			return nil, rowsErr
		}
		owners = append(owners, item)
	}

	result.Owners = owners

	return &result, nil
}

func AddCompanyOwners(db *sql.DB, co models.CompanyOwners) error {
	sqlStmt := `
		INSERT OR REPLACE INTO CompanyOwners(
			CompanyID,
			OwnerID
		) values (?, ?)
		`
	stmt, err := db.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, stmtErr := stmt.Exec(
		co.CompanyID,
		co.OwnerID)
	if stmtErr != nil {
		return stmtErr
	}

	return nil
}

func GetAvailableOwners(db *sql.DB, id int) ([]models.Owner, error) {
	sqlStmt := `
		select
			o.OwnerID,
			o.Name
		from Owners o
		where o.OwnerID not in (
			select co.OwnerID 
			from CompanyOwners co
			where co.CompanyID = ?)
		`

	rows, err := db.Query(sqlStmt, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []models.Owner
	for rows.Next() {
		item := models.Owner{}
		rowsErr := rows.Scan(
			&item.OwnerID,
			&item.Name)
		if rowsErr != nil {
			return nil, rowsErr
		}
		result = append(result, item)
	}

	return result, nil
}
