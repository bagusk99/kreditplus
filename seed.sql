-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 23, 2025 at 04:39 PM
-- Server version: 8.0.33
-- PHP Version: 8.3.21

SET FOREIGN_KEY_CHECKS=0;
SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kreditplus`
--

--
-- Truncate table before insert `consuments`
--

TRUNCATE TABLE `consuments`;
--
-- Dumping data for table `consuments`
--

INSERT INTO `consuments` (`id`, `nik`, `full_name`, `legal_name`, `place_of_birth`, `date_of_birth`, `salary`, `ktp_photo`, `selfie_photo`, `created_at`, `updated_at`) VALUES
(7, '064000020', 'Yvon Keningham', 'Keningham', 'Indonesia', '1993-04-23', 5607819, 'box-be8f2207-16eb-443a-bd13-d9cae9bcc578.png', 'box-b5c02fc5-0f61-4e72-9518-315a6b3ec6e4.png', '2024-10-18 17:00:00', '2024-07-25 17:00:00'),
(8, '091302597', 'Goddart Tribbeck', 'Tribbeck', 'Indonesia', '1990-10-13', 6820560, 'box-be8f2207-16eb-443a-bd13-d9cae9bcc578.png', 'box-b5c02fc5-0f61-4e72-9518-315a6b3ec6e4.png', '2025-01-01 17:00:00', '2024-07-12 17:00:00'),
(9, '074001048', 'Maddy Heining', 'Heining', 'Afghanistan', '1994-01-12', 4997271, 'box-be8f2207-16eb-443a-bd13-d9cae9bcc578.png', 'box-b5c02fc5-0f61-4e72-9518-315a6b3ec6e4.png', '2024-06-11 17:00:00', '2024-10-29 17:00:00'),
(11, '121141864', 'Salvatore Cotherill', 'Cotherill', 'Indonesia', '1996-05-28', 7659122, 'box-be8f2207-16eb-443a-bd13-d9cae9bcc578.png', 'box-b5c02fc5-0f61-4e72-9518-315a6b3ec6e4.png', '2024-09-13 17:00:00', '2024-06-29 17:00:00');

--
-- Truncate table before insert `transactions`
--

TRUNCATE TABLE `transactions`;
--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `consument_id`, `contract_number`, `otr`, `admin_fee`, `installment`, `interest`, `asset_name`, `created_at`, `updated_at`) VALUES
(8, 8, '1231', 1231, 1231, 1231, 12312, 'wew asset', '2025-05-23 07:42:01', '2025-05-23 07:59:07');
SET FOREIGN_KEY_CHECKS=1;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
