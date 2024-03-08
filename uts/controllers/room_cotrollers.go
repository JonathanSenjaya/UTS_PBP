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
	db, err := connectGorm()
	if err != nil {
		log.Fatal(err)
		sendErrorResponse(w, "Failed to Connect Database!")
	}
	var participants int64
	// var max_player int64
	id_room := r.Form.Get("id_room")
	strconv.Atoi(id_room)
	// room_name := r.Form.Get("room_name")
	// id_games, _ := strconv.Atoi(r.Form.Get("id_game"))
	// if id_games > max_players
	db.Model(&m.Participant{}).Where("room_id = ?", id_room).Count(&participants)
	fmt.Println(participants)
	fmt.Println(id_room)
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
