package card

func TranslateMCC(code string) string {
	// представим, что mcc читается из файла (научимся позже)
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5912": "Аптеки",
	}
	badStatus := "Категория не указана"

	for i, value := range mcc {
		if i == code {
			return value
		}
	}
	return badStatus
}
