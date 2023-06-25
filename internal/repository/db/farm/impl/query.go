package impl

const (
	queryInsertFarm = `
		INSERT INTO farms(farm_name) VALUES($1)
	`

	queryGetFarmByName = `
		SELECT
			farm_id
			, farm_name
		FROM 
			farms
		WHERE
			farm_name = $1
	`

	queryGetFarm = `
		SELECT
			farm_id
			, farm_name
		FROM 
			farms
		WHERE
			status = 1
		LIMIT $1 OFFSET $2
	`

	queryGetFarmByID = `
		SELECT
			farm_id
			, farm_name
		FROM 
			farms
		WHERE
			farm_id = $1
			AND status = 1
	`
)
