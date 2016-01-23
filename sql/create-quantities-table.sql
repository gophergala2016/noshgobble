CREATE TABLE `quantities` (
    `food_id` INTEGER,
    `nutrient_id` INTEGER,
    `quantity` REAL,
	FOREIGN KEY(food_id) REFERENCES foods(id),
	FOREIGN KEY(nutrient_id) REFERENCES nutrients(id),
	PRIMARY KEY(food_id, nutrient_id)
);
