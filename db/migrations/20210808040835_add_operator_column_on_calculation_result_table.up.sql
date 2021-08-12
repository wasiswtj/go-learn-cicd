CREATE TYPE operator_enum AS ENUM (
	'add',
	'subtract',
	'multiply',
    'divide'
);

ALTER TABLE calculation_results ADD COLUMN operator operator_enum;