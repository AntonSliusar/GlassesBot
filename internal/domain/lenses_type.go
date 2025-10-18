package domain

type LensesType struct {
	Code string
	Name string
}

var LensesTypes = []LensesType{
	{Code: "lenses_1", Name: "Пластик"},
	{Code: "lenses_2", Name: "Мінерал"},
	{Code: "lenses_3", Name: "полікарбонат"},
}

func GetLensesByID(code string) string {
	for _, lenses := range LensesTypes {
		if lenses.Code == code {
			return lenses.Name
		}
	}
	return ""
}
