CREATE TABLE `consuments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `nik` varchar(100) NOT NULL,
  `full_name` varchar(255) NOT NULL,
  `legal_name` varchar(255) NOT NULL,
  `place_of_birth` varchar(100) NOT NULL,
  `date_of_birth` date NOT NULL,
  `salary` int NOT NULL,
  `ktp_photo` varchar(255) NOT NULL,
  `selfie_photo` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `transactions` (
  `id` int NOT NULL AUTO_INCREMENT,
  `consument_id` int NOT NULL,
  `contract_number` varchar(255) NOT NULL,
  `otr` int NOT NULL, -- PRICE
  `admin_fee` int NOT NULL,
  `installment` int NOT NULL, -- CICILAN
  `interest` int NOT NULL, -- BUNGA
  `asset_name` varchar(255) NOT NULL, -- BUNGA
  `created_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci