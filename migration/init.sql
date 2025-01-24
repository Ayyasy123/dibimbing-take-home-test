CREATE TABLE `users` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `email` varchar(255),
  `password` varchar(255),
  `role` user_role,
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `event` (
  `id` int PRIMARY KEY,
  `name` varchar(255),
  `description` varchar(255),
  `location` varchar(255),
  `date` timestamp,
  `category` array,
  `capacity` int,
  `price` varchar(255),
  `status` ENUM ('Active', 'Ongoing', 'Completed', 'Canceled'),
  `available_tickets` int,
  `ticket_availability` ENUM ('Available', 'Sold Out'),
  `created_at` timestamp,
  `updated_at` timestamp
);

CREATE TABLE `ticket` (
  `id` int PRIMARY KEY,
  `user_id` int,
  `event_id` int,
  `status` ENUM ('Purchased', 'Canceled'),
  `created_at` timestamp,
  `updated_at` timestamp
);

ALTER TABLE `ticket` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `ticket` ADD FOREIGN KEY (`event_id`) REFERENCES `event` (`id`);
