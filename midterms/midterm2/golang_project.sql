-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 24, 2023 at 07:47 PM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `golang_project`
--

-- --------------------------------------------------------

--
-- Table structure for table `comments`
--

CREATE TABLE `comments` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `username` varchar(40) NOT NULL DEFAULT 'Some Username',
  `item_id` int(11) NOT NULL,
  `content` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `comments`
--

INSERT INTO `comments` (`id`, `user_id`, `username`, `item_id`, `content`) VALUES
(1, 1, 'Medeubekov Sergazy', 23, 'TEST COMMENT'),
(2, 2, 'Askhat Alaziz', 23, 'Comment'),
(3, 2, 'Askhat Alaziz', 23, 'IS is the comment'),
(4, 1, 'Medeubekov Sergazy', 23, 'I\'m here'),
(5, 2, 'Askhat Alaziz', 23, 'IS is the comment'),
(6, 2, '', 23, 'fefwefwef'),
(7, 2, '', 23, 'fwfwefwfe'),
(8, 2, '', 23, 'efefwef'),
(9, 2, '', 23, 'aaa'),
(10, 2, '', 23, '222'),
(11, 2, '', 23, '333'),
(12, 2, 'Medoka', 23, 'mmmmmmmm'),
(13, 3, 'admin', 23, '33333'),
(14, 2, 'Medoka', 23, 'Test.exe'),
(15, 2, 'Medoka', 23, 'Test.exe'),
(16, 2, 'Medoka', 23, 'Test.exe'),
(17, 2, 'Medoka', 23, 'Test.exe'),
(18, 2, 'Medoka', 23, 'ffwrefef'),
(19, 3, 'admin', 23, 'grgrge'),
(20, 2, 'Medoka', 17, 'efwefwe'),
(21, 2, 'Medoka', 15, 'erihbgfkrfbefk'),
(22, 3, 'admin', 15, 'rgekrbgjetrg r'),
(23, 2, 'Medoka', 16, 'erghbkfnek krferf'),
(24, 3, 'admin', 16, 'rgeifbhqefjbef');

-- --------------------------------------------------------

--
-- Table structure for table `items`
--

CREATE TABLE `items` (
  `id` int(11) NOT NULL,
  `name` varchar(20) NOT NULL DEFAULT 'item_name',
  `content` varchar(100) NOT NULL DEFAULT 'item_content',
  `picture` varchar(40) DEFAULT '.',
  `price` int(20) NOT NULL DEFAULT 1488,
  `tags` varchar(100) NOT NULL DEFAULT 'product',
  `rating` float NOT NULL DEFAULT 10
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `items`
--

INSERT INTO `items` (`id`, `name`, `content`, `picture`, `price`, `tags`, `rating`) VALUES
(1, 'item_name', 'item_content', '.', 1000, 'product', 0),
(2, 'item_name1', 'item_content1', 'pic1.jpg', 1100, 'product', 0),
(3, 'Project', 'aaaa', './pictures/ешеду.png', 1600, 'product', 0),
(4, 'Project', 'aaaa', './pictures/ешеду.png', 100, 'product', 0),
(5, 'Project', 'aaaa', './pictures/ешеду.png', 500, 'product', 0),
(6, 'Project', 'aaaa', 'pictures/ешеду.png', 1488, 'product', 0),
(7, 'Project', 'aaaa', 'pictures/ешеду.png', 700, 'product', 0),
(8, 'Project', 'aaaa', 'pictures/ешеду.png', 1488, 'product', 0),
(9, 'test', 'test', 'pictures/pixc2.png', 1488, 'product', 0),
(10, 'test', 'test', 'pictures/pixc2.png', 1488, 'product', 0),
(11, 'test', 'test', 'pictures/image001_1525420555097b.png', 5000, 'product', 0),
(12, 'test', 'test', 'pictures/image001_1525420555097b.png', 1488, 'product', 0),
(13, 'Test2', 'tets2', 'pictures/pic1.png', 1488, 'product', 0),
(14, 'Name', 'content', 'pictures/3.png', 1488, 'product', 0),
(15, 'admin', 'admin', 'pictures/stp ethtrunk.png', 1200, 'product', 0),
(16, 'Name of post', 'some content', 'pictures/4.png', 1488, 'product', 0),
(17, 'Ajaja', 'Jojojo', 'pictures/Без имени-1.png', 325, 'product', 0),
(18, 'admin', 'admin', 'pictures/stp ethtrunk.png', 1400, 'product', 0),
(19, 'Phone', 'The phone', '.', 120000, 'phone', 0),
(20, 'TV', 'the tv', '.', 200000, 'tv samsung ', 0),
(21, 'TV', 'the tv', '.', 200000, 'tv samsung ', 0),
(22, 'TV', 'the tv', '.', 200000, 'tv samsung ', 0),
(23, 'tv', 'the tv', '.', 200000, 'tv samsung ', 4);

-- --------------------------------------------------------

--
-- Table structure for table `ratings`
--

CREATE TABLE `ratings` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `item_id` int(11) NOT NULL,
  `value` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `ratings`
--

INSERT INTO `ratings` (`id`, `user_id`, `item_id`, `value`) VALUES
(1, 2, 23, 5),
(2, 2, 23, 1),
(3, 2, 23, 5),
(4, 2, 23, 5);

-- --------------------------------------------------------

--
-- Table structure for table `tags`
--

CREATE TABLE `tags` (
  `id` int(11) NOT NULL,
  `name` varchar(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `tags`
--

INSERT INTO `tags` (`id`, `name`) VALUES
(1, 'tv'),
(2, 'phone'),
(3, 'product');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `full_name` varchar(40) NOT NULL,
  `email` varchar(40) NOT NULL,
  `username` varchar(40) NOT NULL,
  `password` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `full_name`, `email`, `username`, `password`) VALUES
(1, 'Medeubekov Sergazy', 'mdbkvs@mail.ru', 'mdbkvS', '$2a$10$NdkYGJ.hm/Gu/'),
(2, 'Askhat Alaziz', 'alaziz.200223@gmail.com', 'Medoka', '$2a$10$jUMdIzyLyyCZ1vI0vZcpNushTnFv48XwOSDEhFakmqp7QwVdCixj.'),
(3, 'Askhat Alaziz', 'admin@mail.com', 'admin', '$2a$10$qgXmcK.DbDMs8MH0a/8p/.lAq61JX7TVj8PqpV6Bbus2QZjjT9uNC');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `comments`
--
ALTER TABLE `comments`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user` (`user_id`),
  ADD KEY `item` (`item_id`);

--
-- Indexes for table `items`
--
ALTER TABLE `items`
  ADD PRIMARY KEY (`id`);
ALTER TABLE `items` ADD FULLTEXT KEY `full_text_index` (`name`,`content`,`picture`);
ALTER TABLE `items` ADD FULLTEXT KEY `name` (`name`,`content`);

--
-- Indexes for table `ratings`
--
ALTER TABLE `ratings`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `tags`
--
ALTER TABLE `tags`
  ADD UNIQUE KEY `unique name` (`name`),
  ADD KEY `id` (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);
ALTER TABLE `users` ADD FULLTEXT KEY `fulltext_index_name` (`full_name`,`email`,`username`,`password`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `comments`
--
ALTER TABLE `comments`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=25;

--
-- AUTO_INCREMENT for table `items`
--
ALTER TABLE `items`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=24;

--
-- AUTO_INCREMENT for table `ratings`
--
ALTER TABLE `ratings`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `tags`
--
ALTER TABLE `tags`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `comments`
--
ALTER TABLE `comments`
  ADD CONSTRAINT `item` FOREIGN KEY (`item_id`) REFERENCES `items` (`id`),
  ADD CONSTRAINT `user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
