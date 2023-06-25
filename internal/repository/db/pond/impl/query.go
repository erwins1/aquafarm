package impl

const (
	queryGetPondByFarmID = `
		SELECT
			pond_id
			, farm_id
			, pond_name
			, description
		FROM
			ponds
		WHERE
			farm_id = $1
			AND status = 1
	`

	queryInsertPond = `
		INSERT INTO ponds(farm_id,pond_name,description) VALUES($1,$2,$3)
	`

	queryGetPondByName = `
		SELECT
			pond_id
			, farm_id
			, pond_name
			, description
		FROM
			ponds
		WHERE
			pond_name = $1
			AND status = 1
	`

	queryGetPond = `
		SELECT
			pond_id
			, farm_id
			, pond_name
			, description
		FROM 
			ponds
		WHERE
			status = 1
		LIMIT $1 OFFSET $2
	`

	queryGetPondByID = `
		SELECT
			pond_id
			, farm_id
			, pond_name
			, description
		FROM 
			ponds
		WHERE
			pond_id = $1
			AND status = 1
	`

	queryDeletePond = `
		UPDATE 
			ponds 
		SET 
			status = 2 
		WHERE 
			pond_id = $1
	`
)
