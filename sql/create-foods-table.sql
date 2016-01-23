CREATE TABLE `foods` (
    `id` INTEGER PRIMARY KEY,
    `food_group_id` INTEGER,
    `description` VARCHAR(200) NULL,
    `short_description` VARCHAR(60) NULL,
    `common_name` VARCHAR(100) NULL,
    `manufacturer_name` VARCHAR(65) NULL,
    `refuse_description` VARCHAR(135) NULL,
    `refuse` INTEGER NULL,
    `scientific_name` VARCHAR(65) NULL,
    `nitrogen_factor` REAL NULL,
    `protein_factor` REAL NULL,
    `fat_factor` REAL NULL,
    `carbohydrate_factor` REAL NULL
);
