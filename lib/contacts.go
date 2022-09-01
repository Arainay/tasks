package lib

const count = 5

// Дается журнал звонков — набор записей (имя звонившего, телефон звонившего).
// Записи даны в хронологическом порядке от наиболее ранней к самой последней.
// Требуется восстановить для каждого звонившего 5 последних его номеров телефона.
// Записи могут встречаться несколько раз, то есть возможна ситуация, когда
// одна пара (имя звонившего, телефон звонившего) встречается два и более раза во входных данных.
func GetLastCallsForEveryCallMaker(list [][2]string) map[string][]string {
	contacts := map[string][]string{}

	for i := len(list) - 1; i >= 0; i-- {
		name := list[i][0]
		phone := list[i][1]
		_, isExists := contacts[name]

		if !isExists {
			value := []string{phone}
			contacts[name] = value
		} else {
			if len(contacts[name]) < count && !contains(contacts[name], phone) {
				contacts[name] = append(contacts[name], phone)
			}
		}
	}

	return contacts
}

func contains(slice []string, value string) bool {
	for _, el := range slice {
		if el == value {
			return true
		}
	}

	return false
}
