package consuments

import (
	"kreditplus/config"
	"kreditplus/entities"
)

func GetAll() []entities.Consument {
	rows, err := config.DB.Query("select * FROM consuments");

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var consuments []entities.Consument

	for rows.Next() {
		var consument entities.Consument
		err := rows.Scan(
			&consument.Id, 
			&consument.Nik,
			&consument.FullName,
			&consument.LegalName,
			&consument.PlaceOfBirth,
			&consument.DateOfBirth,
			&consument.Salary,
			&consument.KtpPhoto,
			&consument.SelfiePhoto,
			&consument.CreatedAt, 
			&consument.UpdatedAt,
		)

		if err != nil {
			panic(err)
		}

		consuments = append(consuments, consument)
	}

	return consuments
}

func Create(consument entities.Consument) {
	_, err := config.DB.Exec(`
		insert into categories (name, created_at, updated_at)
		value (?, ?, ?, ?, ?, ?, ?, ?)`,
		consument.Nik,
		consument.FullName,
		consument.LegalName,
		consument.PlaceOfBirth,
		consument.DateOfBirth,
		consument.Salary,
		consument.KtpPhoto,
		consument.SelfiePhoto,

		consument.CreatedAt, 
		consument.UpdatedAt,
	)	

	if err != nil {
		panic(err)
	}
}