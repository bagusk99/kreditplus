package transactions

import (
	"kreditplus/config"
	"kreditplus/entities"
)

func GetAll() []entities.Transaction {
	rows, err := config.DB.Query("select * FROM transactions");

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var transactions []entities.Transaction

	for rows.Next() {
		var transaction entities.Transaction
		err := rows.Scan(
			&transaction.Id, 
			&transaction.ConsumentId,
			&transaction.ContractNumber,
			&transaction.Otr,
			&transaction.AdminFee,
			&transaction.Installment,
			&transaction.Interest,
			&transaction.AssetName,
			&transaction.CreatedAt, 
			&transaction.UpdatedAt,
		)

		if err != nil {
			panic(err)
		}

		transactions = append(transactions, transaction)
	}

	return transactions
}

func Create(transaction entities.Transaction) {
	_, err := config.DB.Exec(`
		insert into transactions (
			consument_id,
			contract_number,
			otr,
			admin_fee,
			installment,
			interest,
			asset_name,
			created_at,
			updated_at
		)
		value (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		transaction.ConsumentId,
		transaction.ContractNumber,
		transaction.Otr,
		transaction.AdminFee,
		transaction.Installment,
		transaction.Interest,
		transaction.AssetName,
		transaction.CreatedAt, 
		transaction.UpdatedAt,
	)	

	if err != nil {
		panic(err)
	}
}

func Detail(id int) entities.Transaction {
	row := config.DB.QueryRow(`select * from transactions where id = ?`, id)

	var transaction entities.Transaction

	err := row.Scan(
		&transaction.Id, 
		&transaction.ConsumentId,
		&transaction.ContractNumber,
		&transaction.Otr,
		&transaction.AdminFee,
		&transaction.Installment,
		&transaction.Interest,
		&transaction.AssetName,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
	)

	if err != nil {
		panic(err)
	}

	return transaction
}

func Update(id int, transaction entities.Transaction) {
	_, err := config.DB.Exec(
		`update transactions set 
		consument_id = ?,
		contract_number = ?,
		otr = ?,
		admin_fee = ?,
		installment = ?,
		interest = ?,
		asset_name = ?,
		updated_at = ?
		where id = ?`, 
		transaction.ConsumentId,
		transaction.ContractNumber,
		transaction.Otr,
		transaction.AdminFee,
		transaction.Installment,
		transaction.Interest,
		transaction.AssetName,
		transaction.UpdatedAt,
		id,
	)

	if err != nil {
		panic(err)
	}
}

func Delete(id int) {
	config.DB.Exec(`DELETE from transactions where id = ?`, id)
}