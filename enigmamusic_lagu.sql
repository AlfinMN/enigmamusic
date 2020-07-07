-- MySQL dump 10.13  Distrib 8.0.20, for Win64 (x86_64)
--
-- Host: localhost    Database: enigmamusic
-- ------------------------------------------------------
-- Server version	8.0.20

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `lagu`
--

DROP TABLE IF EXISTS `lagu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `lagu` (
  `id_lagu` varchar(50) NOT NULL,
  `judul_lagu` varchar(100) DEFAULT NULL,
  `artist` varchar(50) DEFAULT NULL,
  `album` varchar(50) DEFAULT NULL,
  `genre` varchar(50) DEFAULT NULL,
  `tahun_rilis` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id_lagu`),
  KEY `fk_artist_idx` (`artist`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `lagu`
--

LOCK TABLES `lagu` WRITE;
/*!40000 ALTER TABLE `lagu` DISABLE KEYS */;
INSERT INTO `lagu` VALUES ('SONG001','I Love You But I\'m Letting Go','Pamungkas','I Love You But I\'m Letting Go','Pop','2016'),('SONG002','Monolog','Pamungkas','Walk The Talk','Pop','2018'),('SONG003','Kenangan Manis','Pamungkas','Walk The Talk','Pop','2018'),('SONG004','One Only','Pamungkas','Walk The Talk','Pop','2018'),('SONG005','Deeper','Pamungkas','Solipsism','Pop','2020');
/*!40000 ALTER TABLE `lagu` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-07-07 15:43:19
