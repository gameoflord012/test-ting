package Logic

import (
	"wan-api-kol-event/DTO"
	"wan-api-kol-event/Initializers"
)

// * Get Kols from the database based on the range of pageIndex and pageSize
// ! USE GORM TO QUERY THE DATABASE
// ? There are some support function that can be access in Utils folder (/BE/Utils)
// --------------------------------------------------------------------------------
// @params: pageIndex
// @params: pageSize
// @return: List of KOLs and error message

func GetKolLogic() ([]*DTO.KolDTO, error) {
	var kolDTOs []*DTO.KolDTO

	if err := Initializers.DB.Raw("SELECT * FROM kol").Scan(&kolDTOs).Error; err != nil {
		return nil, err
	}

	return kolDTOs, nil
}
