package db

func GetTeamIDByName(name string) (int64, error) {
	var id int64
	err := db.QueryRow("SELECT id FROM Team WHERE nome = $1", name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetCampionatoIDByName(name string) (int64, error) {
	var id int64
	err := db.QueryRow("SELECT id FROM Campionato WHERE nome = $1", name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetGiocatoreIDByUsername(username string) (int64, error) {
	var id int64
	err := db.QueryRow("SELECT id FROM Giocatore WHERE username = $1", username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetEventoIDByName(name string) (int64, error) {
	var id int64
	err := db.QueryRow("SELECT id FROM Evento WHERE nome = $1", name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetCartaIDByName(name string) (int64, error) {
	var id int64
	err := db.QueryRow("SELECT id FROM Carta WHERE nome = $1", name).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
