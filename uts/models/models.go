package models

//account
type Account struct {
	Id       int    `json:"id_account"`
	Username string `json:"username"`
}

//game
type Game struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Max_player int    `json:"max_player"`
}

//room
type Room struct {
	Id        int    `json:"id"`
	Room_name string `json:"room_name"`
	Id_game   int    `json:"id_game"`
}

//participant
type Participant struct {
	Id         int `json:"id"`
	Id_room    int `json:"id_room"`
	Id_account int `json:"id_account"`
}

//response
type RoomDataResponse struct {
	Id        int    `json:"id"`
	Room_name string `json:"room_name"`
}

type RoomResponse struct {
	Rooms []RoomDataResponse `json:"rooms"`
}

type RoomsResponse struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Data    RoomResponse `json:"data"`
}

type DetailParticipant struct {
	Id         int    `json:"id"`
	Id_account int    `json:"id_account"`
	Username   string `json:"username"`
}

type RoomDetailResponse struct {
	Id          int               `json:"id"`
	Room_name   string            `json:"room_name"`
	Participant DetailParticipant `json:"participants"`
}

type RoomDetailsResponse struct {
	Rooms []RoomDetailResponse `json:"rooms"`
}

type RoomsDetailResponse struct {
	Status  int                 `json:"status"`
	Message string              `json:"message"`
	Data    RoomDetailsResponse `json:"data"`
}
