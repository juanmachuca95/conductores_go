CREATE TABLE `users` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
);

CREATE TABLE `conductores` (
  `id` int NOT NULL,
  `users_id` int NOT NULL,
  `matricula` varchar(255) NOT NULL,
  `vehiculo` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
);

CREATE TABLE `roles` (
  `id` int NOT NULL,
  `role` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
);


CREATE TABLE `rolesusers` (
  `id` int NOT NULL,
  `roles_id` int NOT NULL,
  `users_id` int NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
);


CREATE TABLE `viajes` (
  `id` int NOT NULL,
  `users_id` int NOT NULL,
  `trip_status` tinyint NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL
);

INSERT INTO `roles` (`id`, `role`, `created_at`, `updated_at`) VALUES
(1, 'conductor', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(2, 'admin', '2022-05-25 17:11:13', '2022-05-25 17:11:13');


INSERT INTO `conductores` (`id`, `users_id`, `matricula`, `vehiculo`, `created_at`, `updated_at`) VALUES
(1, 1, '1088', 'Lady Shannon Keebler', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(2, 69, '1674', 'Dr. Sunny Runolfsdottir', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(3, 57, '622', 'Dr. Madalyn Roberts', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(4, 47, '359', 'Dr. Jaunita Cronin', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(5, 86, '203', 'Princess Jena White', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(6, 16, '1847', 'Lady Summer Rutherford', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(7, 96, '4', 'Ms. Nicolette Smith', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(8, 67, '722', 'Queen Thea Hermiston', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(9, 73, '613', 'Miss Kariane Klocko', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(10, 35, '283', 'Miss Desiree Terry', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(11, 31, '472', 'Prof. Emely Goodwin', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(12, 37, '1502', 'Miss Dasia Nienow', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(13, 45, '171', 'Queen Mikayla Walker', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(14, 39, '1719', 'Queen Jeanne Jacobs', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(15, 34, '1371', 'Lady Callie Labadie', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(16, 90, '1165', 'Princess Deanna Bartell', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(17, 97, '474', 'Ms. Camila Denesik', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(18, 64, '74', 'Lady Icie Altenwerth', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(19, 76, '1551', 'Mrs. Heidi Bergstrom', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(20, 12, '1319', 'Dr. Virginie Orn', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(21, 41, '1828', 'Princess Wilhelmine Daniel', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(22, 62, '745', 'Princess Marta Denesik', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(23, 93, '659', 'Queen Cora Dare', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(24, 72, '720', 'Mrs. Mariah Harber', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(25, 8, '294', 'Lady Elizabeth Mayert', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(26, 43, '838', 'Lady June Morar', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(27, 50, '1370', 'Miss Miracle Collins', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(28, 42, '1082', 'Miss Flo Sanford', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(29, 99, '108', 'Mrs. Jaunita Emmerich', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(30, 17, '1134', 'Queen Ottilie Wintheiser', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(31, 21, '784', 'Queen Florine Little', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(32, 70, '253', 'Princess Karine Hansen', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(33, 61, '180', 'Prof. Liana Littel', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(34, 38, '963', 'Miss Rhea Prosacco', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(35, 10, '1864', 'Ms. Laura Wilderman', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(36, 82, '838', 'Queen Ila Nienow', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(37, 77, '1227', 'Prof. Hettie Olson', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(38, 75, '218', 'Ms. Ayla Kerluke', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(39, 51, '3', 'Ms. Eugenia Cole', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(40, 95, '1308', 'Ms. Tara Grady', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(41, 60, '228', 'Mrs. Delfina Will', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(42, 87, '1842', 'Ms. Euna Dibbert', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(43, 78, '1721', 'Prof. Caitlyn Cartwright', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(44, 88, '651', 'Princess Marian Stanton', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(45, 59, '492', 'Queen Shyanne Heathcote', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(46, 81, '948', 'Miss Rosa Ortiz', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(47, 58, '1552', 'Mrs. Santina Cassin', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(48, 7, '649', 'Princess Verona Osinski', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(49, 100, '173', 'Queen Myrna Auer', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(50, 89, '1172', 'Miss Christa Berge', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(51, 49, '1680', 'Mrs. Myrtis Bashirian', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(52, 94, '1484', 'Mrs. Micaela Smitham', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(56, 107, '3333', 'SCANIA 220', '2022-05-25 22:05:14', '2022-05-25 22:05:14');



INSERT INTO `rolesusers` (`id`, `roles_id`, `users_id`, `created_at`, `updated_at`) VALUES
(1, 2, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(2, 2, 26, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(3, 2, 9, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(4, 1, 69, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(5, 2, 36, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(6, 2, 98, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(7, 2, 2, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(8, 1, 57, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(9, 2, 3, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(10, 2, 53, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(11, 2, 80, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(12, 2, 46, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(13, 1, 47, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(14, 1, 86, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(15, 2, 32, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(16, 1, 16, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(17, 2, 84, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(18, 2, 20, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(19, 1, 96, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(20, 1, 67, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(21, 2, 40, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(22, 1, 73, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(23, 2, 65, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(24, 1, 35, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(25, 2, 55, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(26, 1, 31, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(27, 1, 37, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(28, 2, 19, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(29, 1, 45, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(30, 2, 22, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(31, 1, 39, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(32, 2, 27, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(33, 1, 34, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(34, 1, 90, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(35, 1, 97, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(36, 2, 91, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(37, 1, 64, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(38, 2, 44, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(39, 1, 76, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(40, 1, 12, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(41, 1, 41, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(42, 1, 62, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(43, 2, 56, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(44, 1, 93, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(45, 1, 72, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(46, 1, 8, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(47, 2, 85, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(48, 1, 43, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(49, 2, 71, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(50, 2, 28, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(51, 2, 33, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(52, 2, 48, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(53, 1, 50, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(54, 1, 42, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(55, 1, 99, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(56, 2, 18, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(57, 1, 17, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(58, 2, 4, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(59, 2, 6, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(60, 2, 74, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(61, 1, 21, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(62, 2, 14, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(63, 2, 68, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(64, 1, 70, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(65, 1, 61, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(66, 1, 38, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(67, 1, 10, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(68, 1, 82, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(69, 2, 66, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(70, 1, 77, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(71, 2, 101, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(72, 1, 75, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(73, 1, 51, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(74, 1, 95, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(75, 1, 60, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(76, 1, 87, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(77, 1, 78, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(78, 2, 13, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(79, 2, 29, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(80, 2, 54, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(81, 1, 88, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(82, 2, 79, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(83, 1, 59, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(84, 2, 92, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(85, 1, 81, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(86, 2, 52, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(87, 2, 5, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(88, 2, 83, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(89, 1, 58, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(90, 2, 11, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(91, 2, 24, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(92, 2, 25, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(93, 1, 7, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(94, 1, 100, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(95, 1, 89, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(96, 2, 63, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(97, 2, 30, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(98, 1, 49, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(99, 2, 15, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(100, 1, 94, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(101, 2, 23, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(102, 1, 102, '2022-05-25 20:25:35', '2022-05-25 20:25:35'),
(103, 1, 107, '2022-05-25 22:05:14', '2022-05-25 22:05:14'),
(104, 2, 108, '2022-05-25 22:25:31', '2022-05-25 22:25:31');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--


--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `email`, `password`, `created_at`, `updated_at`) VALUES
(1, 'Administrador', 'admin@spaceguru.com', '$2y$10$WZFBKZT58ZvKd35HCv8EAuboBwld50iONiDC6StX1vsEUjOTAW3Tq', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(2, 'Queen Vivian Ritchie', 'bsobrSV@hxsWsAn.biz', 'hZJuralqQnYUJDGlvADhEZuMcZxIiRhqVQOblSjaCKSGJHDdYt', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(3, 'Mrs. Katelin Huel', 'BTJRjXO@MqqvOmJ.info', 'CQKjYDwXybnrDpEBGaaKCnbwsYybImgChpVsIRZeLWfMgJKMHQ', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(4, 'Dr. Yasmine Grimes', 'peeIkQD@dpLVTAc.ru', 'GKiCgVEvHNQyioxbXSFQatopDTDxvqXZKNMmLEtZaedTuiJARp', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(5, 'Princess Stephania McKenzie', 'WauFYTG@qmJAuyS.com', 'hfsQvvnPuRCQTxpSVmvFaLusQkEZcRdptNfZArfZxCXJmdOZtr', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(6, 'Princess Joanny Bartell', 'PkefLsq@ooGoTtF.net', 'LYUwthvLBYucYVidUwgXMatfdqvvwaoTZAYJRWJVsJPQPCOblO', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(7, 'Ms. Hailie Reinger', 'xbQJFTv@aKamwZh.ru', 'ggJUEYcwbfAqljsWFSdgQvenJTSJhuFDaNCXKIcGkCrsvygoQh', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(8, 'Princess Margaretta Schoen', 'kqwFrhu@gcHBfDl.info', 'xrMEFrANGpYRyFsHWAFMqYIqxlUDKfjqDcfWkyquHQKFHemcuC', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(9, 'Lady Kathryn White', 'aLZMqEL@kskxAEL.org', 'xOykSMMiNxQhVJOXuSsVIfVmoRbIscGhmNIcvbxCYebsjpnGCp', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(10, 'Princess Caitlyn Harber', 'SEnrbKY@kTejnMj.biz', 'OQBYmPKuupNieOsWdlSCvfAaUyghEFVcjeCaetADvgpfjlOHpu', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(11, 'Mrs. Katlynn Murazik', 'WgyuIKl@eYLfIEN.com', 'sfBjwqspAoqUHAtTrtbgTJBgnCnKcoPIvhQtZvEFcyZjkRYgis', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(12, 'Queen Kaylin Schaefer', 'JRslexZ@aFeshYF.com', 'EmAQBCgqyBkvAqGDfarOokFqZYUAYqsRxCLqEBqiSsfjiWkkUe', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(13, 'Prof. Valentine Von', 'UIcWtgv@pUCCjbk.org', 'ZhZSbtcgVTHPhNFLOMuFFRIFBZRwfeTHytMPcEOpvSBBLrTQvu', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(14, 'Dr. Joanny Swift', 'QLlCivs@NAmvoUw.biz', 'eiTkROLkHhkOLIlrPAxgwTBwgqupDEynIZCemaFiVtpcrPDhoP', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(15, 'Prof. Loyce Schmeler', 'yNOQQKp@gGdMKwf.info', 'PVsBMBGxjyfjsyclNEwIeBStoQKeaWVSPfLKEMQiWBDXenEjId', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(16, 'Ms. Vida Koelpin', 'DdFvpEJ@NLtGcTv.org', 'WawXIpcgNLYMewREoGiERFPIluZNvnQBHJZaKjBDUGcNYNYvOA', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(17, 'Princess Marge Little', 'oQEtIsu@dMDWoSZ.com', 'aCNKYRhooHFJNbxGEmoRmWVXCLPFKvwPqUQSWsJdrcVAZXQbDj', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(18, 'Mrs. Jalyn Schoen', 'OlQexwY@BcTkFvf.net', 'KqrbkCrxdgdUKtoBjFPivhgfJQBtaHpoTFqjrEgRlHRfJEYhwd', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(19, 'Princess Jermaine Schimmel', 'Hitqxru@llYahFF.info', 'QwciFVfBdAWZivVGAMJYfZUhbTGsJVUmBJITqQohOmatHCSfSQ', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(20, 'Miss Connie Lindgren', 'eFWttRK@JhCWyyQ.ru', 'gfKBZlrStIUlbZMWCjaAEFSdVpRUCmWaAVgBJYyCmhEefVXGBI', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(21, 'Mrs. Eulah Jones', 'pTHLHJu@xLZuXPx.com', 'EROaqcbfGimQSkntbhlbOMIugncNFiDXHRFpuVaYngZEltkSZN', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(22, 'Princess Hosea Goyette', 'HtAoppt@bVhccQb.org', 'CuuqKTOPOoTisBWsDKoCKVLqMuhyIkInomloGXGdoYUnxiQlta', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(23, 'Princess Aliyah Ziemann', 'YxLWgFw@ufyoHSD.ru', 'CFTXvSUTSMgxCNuASMhroupjPeheQkshFnkHRpeIpPBkPexAbW', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(24, 'Lady Pamela Wisozk', 'wJoAYyY@qMLcSbf.info', 'sisdtWWGoKdfTfRqUcCUqWNkaHJMggPFULXJRWqnJiQIfMtPeY', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(25, 'Lady Jacinthe Ondricka', 'wTLTwlh@oRhiBch.net', 'RnuidcBelEVOsEOjXcbpRtflYggLIGqIpXOrrWXJrUJYHLouOd', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(26, 'Queen Penelope Rath', 'aIOoYbS@pKlJQnB.net', 'OYnQubwyYYQnCLqYdJfWndrpYTJUyOIrgFBaPEwqHEeIysjmRJ', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(27, 'Mrs. Christa Marks', 'ICjNDfA@kgsYHdK.net', 'BUbhlhLZgIMGRsotHCnhVfSBnskPhEUJfUEjpYauKdBvCjXfDK', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(28, 'Dr. Pascale Hoppe', 'MBvDDKu@eASiYsW.com', 'hAHNoSgquKqTZqYBcNJAPbULBBktwHeqoIkkeLcHKsZZbGkIDD', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(29, 'Mrs. Lois Bernier', 'UNLZrFb@aClFECX.net', 'CmmjOkTfNeZAMUZTCwUyVDZbVfqkrGyuLBgdfVNGhymaeWBNBK', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(30, 'Princess Lauryn Langworth', 'YgEIRwT@MpTyhuH.ru', 'PYRpcDYGtHdmpJsvNhSaCEIMyIfmgfQokVIlEyGGohZsvmkgwk', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(31, 'Mrs. Clarissa McDermott', 'gCAdxcw@AxUsYny.org', 'sBbHZpqcgZEYuJhROBAsMaJKrFjAgidOMrRjCZGTSgMBsZaqqe', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(32, 'Ms. Prudence Carter', 'cvZkbvA@CxjTchR.net', 'GAtqGoPqLanZWAjRindgYLmowtDZabdomHMlAWWaaCZgvZrlpm', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(33, 'Mrs. Rachael Larkin', 'mfwVUuK@DVqIfog.biz', 'inTbgAkmymNHVFeFrRfYqTCTpGjUbFGZcAIDGhFFHOxjhfUDZF', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(34, 'Princess Hortense Gleason', 'IfKTxIY@ZrigpOe.com', 'yUkyvhblmVDLVqcHNqnUWYNwPmdMYxDiIvyUJvlbbpZqkUGIJH', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(35, 'Queen Abigayle Herman', 'FgtNCdP@jushEMO.org', 'mKysSffLSJGZuQOOsnpsYetKJeupviJZGHhCKclSBpWMQMJPqu', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(36, 'Prof. Dianna Kutch', 'BNjLfvj@XHMFbcw.net', 'FEIrDjZFfWciLHpYJnLjMoEmqbgNAfdVfGOxcGuGbwgfUNpYxx', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(37, 'Miss Zita Emard', 'ggLhBly@UXspyFv.ru', 'buxIbcGXYvFtwWkFGkcCGvfIYqVVvCVpWofoBGFSUgvWpcGOGu', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(38, 'Prof. Lexie Bartoletti', 'rxwXZbv@NtFjDhr.ru', 'SdLJDcumJNWJjKIsqdgqKBJpoigDXNiOARCBmHAqYOhyAkLZRc', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(39, 'Princess Kenyatta Mertz', 'Hygujte@hVVrYJT.net', 'fnwgBSnrmPxiDaRpYfNCETkjRtiaBIvTxfDHBZuGexWNvStRrN', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(40, 'Dr. Kitty McClure', 'eZFEaJC@AhrmnXk.biz', 'PFHKVlUVNMeZjkGLVaONXwUBZMKgvPPtUAuGgVCeepfmjteuuR', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(41, 'Queen Andreanne Treutel', 'jsaulWR@aoOssSc.ru', 'HhiAWqZFicNlXQcZhlSOsImHNkbsLMfOLCnNSZJJkGHaGuxlvM', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(42, 'Dr. Rosemarie Walter', 'NvPclBo@moLdrID.com', 'VCBayobddlgfTcNkZXHdNeKSdHUwMYrbWShGByXrobuATFMjYx', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(43, 'Prof. Lillie Satterfield', 'KUQcVey@RosPFux.ru', 'CpDBqwGXjnVGIrZKgUpBspeymyPjwwoQldjLKBhYRmRbwDjgtV', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(44, 'Prof. Carolina Douglas', 'jNKnxiD@sFiriQP.net', 'sCUOsxtQPlRZFUHxmKNoGVIQjfXaYOMnCfbmOgenIYYrjvOPXV', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(45, 'Queen Juana Schroeder', 'hrcxbwv@VkaByfn.net', 'osOvxocKVnrjhFGockNKrDhamSAxVijZeyghsISPbXeZtAxhBD', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(46, 'Dr. Aimee Romaguera', 'cMHjhbm@QdbAjKB.org', 'ySyQRSTIvCqGPslnZmMnLjABvZlKeSpXenGZtsRuhgDYFjxTdO', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(47, 'Miss Leilani Towne', 'CmnYaFt@vuVmUDr.biz', 'GUBLJSPOrvbMvxPtYtiLNcyOjiIFXGEQfVNFZQdLTSpZqVjjnB', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(48, 'Ms. Lorine Rowe', 'MJLJnit@hAnDUeP.ru', 'ssGInkoRHhTggJoRwQlKPbsZKPuSFovDXBuIEneSsBMQVawVMd', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(49, 'Miss Joannie Morissette', 'yJrFlFL@QYmsXkN.biz', 'MNyiietcZFEXOqubTCWHUIcgcfjrkIUOjtPutXXZvkBseLHmKc', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(50, 'Prof. Aracely Abshire', 'NafaLDd@HrlWMRa.biz', 'QJUaiLaXNGGOgBEgqIdNHhycNGhwADhUjNVAJKXeArPCpsXpfl', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(51, 'Mrs. Peggie Herman', 'tSLgraT@cMlxGUj.ru', 'KSQjTYNXUCWgwbfKTpdognUgIxGKKbgIuNUdgpZLbGjCSIdGqq', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(52, 'Dr. Brielle Durgan', 'VRjCWxM@WstUlHK.org', 'duAYnsnfHbGuQmELWCDrpJjMQKnmfUnQEamAgELlsUIFheumwa', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(53, 'Queen Polly Sauer', 'CCSQdRC@TftqOXO.com', 'YAdWFRAgQvhERaKfmbJvVwnWNBhQvVeStNQAHTnhxagbIOJWmv', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(54, 'Lady Fabiola Lynch', 'UPnTHfy@xVvYBvi.info', 'gBoUYoYQioyvByKsepCBZniodimnEyGenHENmnYRJafOGDHJFk', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(55, 'Miss Bulah Simonis', 'fOiOQEo@NJEJTXy.com', 'wvLFWSvQaMxnkAVKWPKOAHlgpHGHWioSSfyupPbkpYGdnifawY', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(56, 'Dr. Myriam Hamill', 'jYqlgnF@hHqUtSt.com', 'WlwOwTfsDXuLTWaxSKxFNWodGcwDrYnFMcuJDlxwCjnSoQmqVX', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(57, 'Ms. Libby O\"Keefe', 'BtHgSsT@aHTBWvN.ru', 'ikABFihjYonrJSdMLuknsmVGnmYRpnXVaQQXomeAOhPLTueWes', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(58, 'Queen Icie Bode', 'wdilPsB@dThZLRF.ru', 'ZRpGBuPUVNehdQhkRLGSoBvwLUUQyxSgpvmDDOipEcouqlFFlX', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(59, 'Queen Emmanuelle Mertz', 'UWMikaD@OuxVTIl.com', 'tTouMmGIHBIPwBJkSoeIngidtyXrOfEUVdLJmMrRjCDXbAZEiS', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(60, 'Prof. Heath Sawayn', 'tVADFJX@FtknuWU.info', 'srJkCSJHeDJAxglhBvnjjhghIIOGoVDwqDGeFbIFCQBBdObaRi', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(61, 'Prof. Abigayle Bartoletti', 'rErtnvH@pAJMRoi.net', 'VWjyUkcMyymHiijJBcVomGqSlYdRKBTEeYhlRpvOmJcClEukdb', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(62, 'Princess Heather Ankunding', 'JvcarNf@Akcclhv.ru', 'fPiMULkhmfxjFETObPfcTFrslPybCTShWpRYpDBACVkEpjybsj', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(63, 'Princess Arvilla Emmerich', 'yehWVcX@JSbHaaa.biz', 'gcNPsTWYuGimVJIWdXNyeyFOoPSoWjJaPUNMuCLRctMxkHVaRL', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(64, 'Dr. Jazmyne Haley', 'JDnhTTv@ZvmAkGF.net', 'rMNvkxQkamhDquTgMuFRnuyaZEwqJGRaYCdiwMeeFsSRVkkchm', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(65, 'Prof. Malvina Reynolds', 'FDtpEIk@EqaviBw.net', 'qpvhGFdeAxnhfvwsXqdVqnHiTyOffuwIbaWQHYnNMlMVlYDtXF', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(66, 'Ms. Daniela Leannon', 'stIrDsU@VtxJmVx.org', 'TfiEpYXpHWcWVaAZiHicMktWSZySdoCsGepsEQdwjSSIPiNrRR', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(67, 'Queen Marcella Weber', 'eMnqdlo@QVsinmE.org', 'SsOrCgBQBTckKmLmxsxSXpSCBJdircCEiZVVmYfPQEVpfDBFZJ', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(68, 'Dr. Lavonne Graham', 'qOmrQFS@ntVOsHf.info', 'KxDaXOgOBweoiJnajiXBdmkcyukHsNdKRDYsiZOCBhaHKDrans', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(69, 'Ms. Kimberly Monahan', 'bkScYfL@qNcunro.org', 'FBcgXMSHteYEAUkuFvZGcpAKlSFSgmmTXxQEWvvVLxHhOlpLeJ', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(70, 'Queen Felicity Schultz', 'QVpmsxa@PNYtLoT.com', 'UZeFIEkRGvWLBuCRPTSogbknMRrwZKwCiTxpXoQdTGpvBHYpgZ', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(71, 'Princess Alycia Farrell', 'LgwUakx@QYqWcvQ.org', 'GtPgjjegKbTvXhpFYIcFvumcoKiQLhxDfwbqcBPSsJXbOmDoKr', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(72, 'Queen Marielle Kessler', 'KBRympP@ZjQbosd.com', 'YvSeKTUknOVGxrkTeLaCaKfxLlquqDUSmvwCqdKHmtFkatlXka', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(73, 'Mrs. Lysanne O\"Keefe', 'FAKnvRJ@UPrANYZ.com', 'wuLDQiJiljCmcLhukDJHLeawsAIobnkArKmIuMKqDeEnegrIUZ', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(74, 'Miss Adell Cartwright', 'PqGiBwQ@GYaEExb.biz', 'lvcOdBhglGtXClUONhKxifqbATftiNUhalukKrkkVqFOxnpErd', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(75, 'Princess America Emmerich', 'TRjtQiu@ROpLxDe.net', 'qaWgSVrulXPtuFdGuMWobxQQrCIxPTTdOnSpWbZZLiBwxQeCJd', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(76, 'Ms. Zetta Braun', 'JqmnQhZ@XYYHrOi.ru', 'YTdJgEcNRXIKDlmyLojDUKaWTcOgkKOuwVIHnJcexbXNPJXsCB', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(77, 'Mrs. Monique Conroy', 'tHgcovb@GRbPlJE.info', 'jCeMvSsWAEHPhVtpjixNtapiQsxGKydBLOUrrlMJSJrpdGfoBJ', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(78, 'Miss Pattie Gottlieb', 'uFPHqDB@lcoMWoj.ru', 'IySMcKILHdagTbEYvXmXpNYIJMjfBLHLPZVktrYgvmbnbkJncL', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(79, 'Dr. Lenore Eichmann', 'UvvyFcS@yxgxOie.net', 'rdISNUTMPtKdyTqXBEFrfrNvNaNRDlCeeXPQRgmYadsZBcglBF', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(80, 'Prof. Lelah Wunsch', 'CfkpFlU@NQARWgL.info', 'wqaHvmOjPkNJqtxjMhltvgObkdRfSjKwQOeLoDJDHUQMlGiQim', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(81, 'Prof. Lilian Kozey', 'vorSmmr@EYvkdfT.ru', 'yTtMyOGcLFfrqDqhQdIjSOSsWkwUrIMBkZsPbiLlOagoFUQPRT', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(82, 'Dr. Herminia Lindgren', 'SHEfHUN@OdkCokP.info', 'VOTeHDcQLEGKuavQgdbVDkDLwSlSoMSJQKGRPTltUtEjbPLDjJ', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(83, 'Mrs. Eldridge Braun', 'WbilGCM@aZlJYwB.info', 'GPxqwyAJQWayunmHJkxiKgbAdMtsjFoxqoTInHDepIqbUKDYsX', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(84, 'Miss Abby Toy', 'eBtyPYt@JvWsfFv.ru', 'RDDfEicgIDBbWExNsDfIqYvhTmCKxkXPWkGiMLgJpgbOldOGEt', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(85, 'Princess Kenna Lueilwitz', 'KRWjSlh@pguIBkN.org', 'oFeeCggeJniYBPwXHgqgEbKnErHHZFCNjbrelqPjJGLgiXeJOO', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(86, 'Mrs. Birdie Mills', 'CqDslRq@nluXGUs.com', 'tdhpteBGbBYFwysupcaRHIiAmJigstgbAtSYkcTTabZVjtRhRH', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(87, 'Queen Jeanie Gleason', 'UCgHqSx@BXGWIhZ.com', 'ZVMWIFnBrOSKRfSncXXgmOKvSkCAQWUqbBIVUfCtEcjiqjmdUC', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(88, 'Ms. Abbey Nolan', 'uUXhAxH@dPjZXgR.com', 'KPEnebeMsgTcbuyhltFXFyvlZEofGoMSXCKRIHndQCeQSuKWmj', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(89, 'Ms. Vesta Carter', 'YAvnVfx@HkCyQwg.net', 'PJSnHKTiqOgTCAXFJTsIryANohVUlqcNPnLfKIJvVtZDtoVtsE', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(90, 'Dr. Meggie Lockman', 'ifOjNhK@stiLnqI.ru', 'BnkjxuwvWIFtTswMKHnnuortbnNSoXEAdeIlbDLiLfBcMjqdhh', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(91, 'Dr. Laney Robel', 'IhtKvrI@rpvIeDe.com', 'qFNIpGdknbSgJIZCHbkhrsFPRClPTFdpMsDbUcOlHJcxImXBRV', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(92, 'Lady Jazmyn Harris', 'VCZsGTH@fKqvdOf.com', 'eNaakuGJidClafEXQjVgJigPUCgEcpXHveXZyJFlucVPREgIjw', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(93, 'Dr. Eula Schowalter', 'kAiGcFU@BOhRxgX.net', 'EWdoxIxlZcabMNrBGrZobNeFviAbxNOHCOmgTjiXtRDjsAlxHu', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(94, 'Mrs. Hanna Breitenberg', 'YVqjTsY@mmpJRSd.info', 'yuicAMYKSljXOxUCZvaxbeVoEoTtjuiIiqGTohhoyCkTOOTaGk', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(95, 'Princess Annalise Howell', 'tTLpXuq@kblAPuS.info', 'CBGeKMNuHWFbdxhhwgEckxgneZUhDoViEQdSocZaeOsXwNTAto', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(96, 'Ms. Grace Donnelly', 'ekEqjvW@OSWAnAL.info', 'KmOdkAxoAstscIXkDPJFhTVZpuFGmttsZvXpNVdCfBvMIYDXeQ', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(97, 'Ms. Eulah Gerlach', 'IgeBBVw@aFBPpFj.info', 'yDIdfoHbOQuxpeqlfeAFDdbNgieGNloUJiBuKggFEmbwQNdVGq', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(98, 'Prof. Rosalyn Durgan', 'bphGcqx@xZBYBsn.net', 'ntOgmDAOmrIIyxusYDBWdRdOxLAkPHWCDccuNKLvMDrjdJvnHe', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(99, 'Prof. Mandy Pfannerstill', 'oLpmovQ@ZCgBbZl.biz', 'spKuQSwtQtPvaYPcWDrgGaNtrPoEwSKZWlQZaXifghXToeCdfD', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(100, 'Miss Linnie Rolfson', 'XNCWvCc@rBOUUlv.org', 'iQTrGlhybrVkwkSJROGSmxNwQuporvYoMjQwXKHZgTiQJwpaXb', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(101, 'Ms. Angela Jones', 'TqtUuch@fHQTIKn.net', 'PQWWMVkUapnyQtVRsBHxEqfOZufJjTqVslFItKcAlRcmEJVVnM', '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(102, 'Conductor test', 'conductor@spaceguru.com', '$2y$10$WZFBKZT58ZvKd35HCv8EAuboBwld50iONiDC6StX1vsEUjOTAW3Tq', '2022-05-25 20:23:59', '2022-05-25 20:23:59'),
(107, 'Nuevo Conductor', 'elconductor@gmail.com', '$2a$04$7gzMSYx0lZwG35nlm1zTYO26WE.SjBIcPSqoYAaBAW21sGvPLkWOy', '2022-05-25 22:05:14', '2022-05-25 22:05:14'),
(108, 'Nuevo usuario admin', 'newadmin@gmail.com', '$2a$04$eiGVfa1J86W45WH9QHB5SO7imjjQGOSQGtZO9Ous1KUAe07//WHtO', '2022-05-25 22:25:31', '2022-05-25 22:25:31');

-- --------------------------------------------------------

--
-- Table structure for table `viajes`
--


--
-- Dumping data for table `viajes`
--

INSERT INTO `viajes` (`id`, `users_id`, `trip_status`, `created_at`, `updated_at`) VALUES
(1, 1, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(2, 2, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(3, 3, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(4, 4, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(5, 5, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(6, 6, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(7, 7, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(8, 8, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(9, 9, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(10, 10, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(11, 11, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(12, 12, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(13, 13, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(14, 14, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(15, 15, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(16, 16, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(17, 17, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(18, 18, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(19, 19, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(20, 20, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(21, 21, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(22, 22, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(23, 23, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(24, 24, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(25, 25, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(26, 26, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(27, 27, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(28, 28, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(29, 29, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(30, 30, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(31, 31, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(32, 32, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(33, 33, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(34, 34, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(35, 35, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(36, 36, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(37, 37, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(38, 38, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(39, 39, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(40, 40, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(41, 41, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(42, 42, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(43, 43, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(44, 44, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(45, 45, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(46, 46, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(47, 47, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(48, 48, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(49, 49, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(50, 50, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(51, 51, 1, '2022-05-25 17:11:13', '2022-05-25 17:11:13'),
(52, 52, 0, '2022-05-25 17:11:13', '2022-05-25 17:11:13');

ALTER TABLE `conductores`
  ADD PRIMARY KEY (`id`),
  ADD KEY `foreign_usersconductor_id` (`users_id`);

ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `rolesusers`
  ADD PRIMARY KEY (`id`),
  ADD KEY `foreign_roles_id` (`roles_id`),
  ADD KEY `foreign_users_id` (`users_id`);

ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

ALTER TABLE `viajes`
  ADD PRIMARY KEY (`id`),
  ADD KEY `foreign_viajes_conductores` (`users_id`);

ALTER TABLE `conductores`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=57;

ALTER TABLE `roles`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

ALTER TABLE `rolesusers`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=105;

ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=113;

ALTER TABLE `viajes`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=53;

ALTER TABLE `conductores`
  ADD CONSTRAINT `foreign_usersconductor_id` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

ALTER TABLE `rolesusers`
  ADD CONSTRAINT `foreign_roles_id` FOREIGN KEY (`roles_id`) REFERENCES `roles` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  ADD CONSTRAINT `foreign_users_id` FOREIGN KEY (`users_id`) REFERENCES `users` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;

ALTER TABLE `viajes`
  ADD CONSTRAINT `foreign_viajes_conductores` FOREIGN KEY (`users_id`) REFERENCES `conductores` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT;
