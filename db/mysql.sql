CREATE DATABASE /*!32312 IF NOT EXISTS*/ spacesguru /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE spacesguru;

DROP TABLE IF EXISTS users;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `rol` enum('admin','conductor') DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO users(id,name,email,password,rol,created_at,updated_at) VALUES(1,'Juan Gabriel Machuca','admin@spaceguru.com','$2y$10$uyUdoTnJdkIKxh7g4t.uFeOrncRnOBaSGtSzkjnyBDVRfVyKqK6ha',NULL,'2022-05-24 09:03:07','2022-05-24 09:03:07');


