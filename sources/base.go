package sources

import (
	"avaandmed/utils"
)

type Ettevote struct {
	ID   int64
	Name string
}

type Isik struct {
	ID           string
	FirstName    *string
	LastName     *string
	BirthDate    string
	BirthDateInt int64
}

func CreateIsik(id *string, firstName *string, lastName *string) *Isik {
	if id == nil || len(*id) != 11 {
		return nil
	}
	
	year := (*id)[1:3]
	if (*id)[0] == '3' || (*id)[0] == '4' {
		year = "19" + year
	} else {
		year = "20" + year
	}
	month := (*id)[3:5]
	day := (*id)[5:7]
	birthDate := day + "." + month + "." + year
	return &Isik{
		ID:           *id,
		FirstName:    firstName,
		LastName:     lastName,
		BirthDate:    birthDate,
		BirthDateInt: utils.Date(birthDate),
	}
}
