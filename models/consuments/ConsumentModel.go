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
		insert into consuments (
			nik,
			full_name,
			legal_name,
			place_of_birth,
			date_of_birth,
			salary,
			ktp_photo,
			selfie_photo,
			created_at,
			updated_at
		)
		value (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
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

func Detail(id int) entities.Consument {
	row := config.DB.QueryRow(`select * from consuments where id = ?`, id)

	var consument entities.Consument

	err := row.Scan(
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

	return consument
}

func Update(id int, consument entities.Consument) {
	_, err := config.DB.Exec(
		`update consuments set 
		nik = ?,
		full_name = ?,
		legal_name = ?,
		place_of_birth = ?,
		date_of_birth = ?,
		salary = ?,
		ktp_photo = ?,
		selfie_photo = ?,
		updated_at = ?
		where id = ?`, 
		consument.Nik,
		consument.FullName,
		consument.LegalName,
		consument.PlaceOfBirth,
		consument.DateOfBirth,
		consument.Salary,
		consument.KtpPhoto,
		consument.SelfiePhoto,
		consument.UpdatedAt,
		id,
	)

	if err != nil {
		panic(err)
	}
}

func Delete(id int) {
	config.DB.Exec(`DELETE from consuments where id = ?`, id)
}