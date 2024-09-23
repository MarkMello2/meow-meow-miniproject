insert
	into
	public.categories
("name",
	description,
	image,
	created_at,
	updated_at,
	deleted_at)
values(
	'Pet supplies', 
	'Pet supplies', 
	'category/pet_supplies.png',
	CURRENT_TIMESTAMP(3),
	CURRENT_TIMESTAMP(3),
	NULL
);
