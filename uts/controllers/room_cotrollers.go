package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	m "uts/models"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := "SELECT id, room_name FROM rooms"

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Failed to query!")
		return
	}
	var room m.RoomDataResponse
	var rooms []m.RoomDataResponse
	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.Room_name); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}
	w.Header().Set("Content-Type", "application/json")

	var response m.RoomsResponse
	response.Status = 200
	response.Message = "Success"
	response.Data.Rooms = rooms
	json.NewEncoder(w).Encode(response)
}

func GetAllDetailRooms(w http.ResponseWriter, r *http.Request) {
	db := connect()
	defer db.Close()

	query := `SELECT r.id, r.room_name, p.id, p.id_account, a.username FROM rooms r
			  JOIN participants p ON r.id = p.id_room
			  JOIN accounts a ON p.id_account = a.id`

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		sendErrorResponse(w, "Failed to query!")
		return
	}
	var room m.RoomDetailResponse
	var rooms []m.RoomDetailResponse
	for rows.Next() {
		if err := rows.Scan(&room.Id, &room.Room_name, &room.Participant.Id, &room.Participant.Id_account, &room.Participant.Username); err != nil {
			log.Println(err)
			return
		} else {
			rooms = append(rooms, room)
		}
	}
	w.Header().Set("Content-Type", "application/json")

	var response m.RoomsDetailResponse
	response.Status = 200
	response.Message = "Success"
	response.Data.Rooms = rooms
	json.NewEncoder(w).Encode(response)
}

func InsertToRoom(w http.ResponseWriter, r *http.Request) {
	db, err := connectGorm()
	if err != nil {
		log.Fatal(err)
		sendErrorResponse(w, "Failed to connect Database!")
		return
	}
	er := r.ParseForm()
	if er != nil {
		return
	}
	id_room, errIdRoom := strconv.Atoi(r.Form.Get("id_room"))
	id_account, errIdAccount := strconv.Atoi(r.Form.Get("id_account"))
	var max_player int
	var participants int
	var id int
	var ids []int
	rows, err := db.Raw("SELECT g.max_player, COUNT(p.id) AS participants, p.id FROM rooms r JOIN participants p ON r.id = p.id_room JOIN games g ON g.id = r.id_games WHERE r.id = ?", id_room).Rows()

	for rows.Next() { //mendapat nilai value max_player dan jumlah participants
		if err := rows.Scan(&max_player, &participants, &id); err != nil {
			return
		} else {
			ids = append(ids, id)
		}
	}

	if errIdRoom != nil { //Error kesalahan ID room
		sendErrorResponse(w, "invalid id room!")
		return
	}

	if errIdAccount != nil { //Error kesalahan ID account
		sendErrorResponse(w, "invalid id account!")
		return
	} /*else {
		for i := 0; i < len(ids); i++ { //Error jika akun sudah ada di dalam room game
			if ids[i] == id_account {
				fmt.Println(ids[i])
				sendErrorResponse(w, "Account Already in room game")
				return
			}
		}
	}*/

	if participants == max_player { //Error ketika room sudah penuh
		sendErrorResponse(w, "Game Room already Full!")
		return
	}

	participant := m.Participant{Id_room: id_room, Id_account: id_account}
	result := db.Create(&participant)
	if result.Error != nil {
		sendErrorResponse(w, "Failed to Insert!")
	} else {
		sendSuccessResponse(w, "Success Insert!")
	}
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response m.RoomsResponse
	response.Status = http.StatusNotAcceptable
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter, message string) {
	var response m.RoomsResponse
	response.Status = http.StatusAccepted
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
