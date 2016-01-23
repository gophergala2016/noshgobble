CREATE TABLE `food` (
    `id` INTEGER PRIMARY KEY AUTOINCREMENT,
    `food_group_id` INTEGER,
    `description` VARCHAR(200) NULL,
    `short_description` VARCHAR(60) NULL,
    `common_name` VARCHAR(100) NULL,
    `manufacturer_name` VARCHAR(65) NULL,
    `refuse_description` VARCHAR(135) NULL,
    `refuse` INTEGER NULL,
    `scientif_name` VARCHAR(65) NULL,
    `nitrogen_factor` REAL NULL,
    `protein_factor` REAL NULL,
    `fat_factor` REAL NULL,
    `carbohydrate_factor` REAL NULL
);
