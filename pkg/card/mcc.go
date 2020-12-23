package card

// TranslateMCC ....
func TranslateMCC(code string) string {
	// представим, что mcc читается из файла (научимся позже)
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5533": "Автоуслуги",
		"5912": "Аптеки",
	}
	badStatus := "Категория не указана"
	value, ok := mcc[code]
	if !ok {
		return badStatus
	}
	return value
}
