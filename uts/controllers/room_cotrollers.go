package controllers

import (
	"encoding/json"
	"fmt"
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
	db := connect()
	defer db.Close()

	// Read request from body
	err := r.ParseForm()
	if err != nil {
		return
	}
	id_room, _ := strconv.Atoi(r.Form.Get("id_room"))
	id_account, _ := strconv.Atoi(r.Form.Get("id_account"))

	_, errQuery := db.Exec("INSERT INTO participants (id_room, id_account) VALUES (?,?)",
		id_room,
		id_account,
	)

	if errQuery != nil {
		sendErrorResponse(w, "Failed to Insert!")
	} else {
		sendSuccessResponse(w, "Success Insert!")
	}
}

func sendErrorResponse(w http.ResponseWriter, message string) {
	var response m.RoomsResponse
	response.Status = 400
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func sendSuccessResponse(w http.ResponseWriter, message string) {
	var response m.RoomsResponse
	response.Status = 200
	response.Message = message
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
