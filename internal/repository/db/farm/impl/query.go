package impl

const (
	queryInsertFarm = `
		INSERT INTO farms(farm_name, description) VALUES($1,$2)
	`

	queryGetFarmByName = `
		SELECT
			farm_id
			, farm_name
		FROM 
			farms
		WHERE
			farm_name = $1
			AND status = 1
	`

	queryGetFarm = `
		SELECT
			farm_id
			, farm_name
			, description
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
			, description
		FROM 
			farms
		WHERE
			farm_id = $1
			AND status = 1
	`

	queryDeleteFarm = `
		UPDATE 
			farms 
		SET 
			status = 2 
		WHERE 
			farm_id = $1
	`

	queryUpdateFarm = `
		UPDATE
			farms
		SET
			description = $2
		WHERE
			farm_id = $1
			AND status = 1
	`
)
