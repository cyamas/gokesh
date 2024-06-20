package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/cyamas/gokesh/board"
	"github.com/cyamas/gokesh/game"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	BLACK = "BLACK"
	WHITE = "WHITE"
)

var Game *game.Game

func RunServer() {
	router := chi.NewRouter()
	fileServer := http.FileServer(http.Dir("static"))
	router.Use(middleware.Logger)
	router.Get("/", home)
	router.Get("/play", play)
	router.Get("/botmove", botMove)
	router.Post("/usermove", userMove)
	router.Handle("/static/*", http.StripPrefix("/static/", fileServer))
	http.ListenAndServe(":6969", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := "OK"
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func play(w http.ResponseWriter, r *http.Request) {
	board := board.New()
	board.SetupPieces()
	Game = game.New(board)
	var data map[string]interface{}
	if Game.Bot.Color == BLACK {
		data = map[string]interface{}{
			"color": "white",
			"from":  "none",
			"to":    "none",
		}
	} else {
		move := Game.Bot.CreateMove(Game.Board)
		Game.ExecuteTurn(move)
		data = map[string]interface{}{
			"color": "black",
			"from":  []int{move.From.Row, move.From.Column},
			"to":    []int{move.To.Row, move.To.Column},
		}
	}
	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

type ClientMove struct {
	From      []int  `json:"from"`
	To        []int  `json:"to"`
	Promotion string `json:"promotion"`
}

func userMove(w http.ResponseWriter, r *http.Request) {
	var move ClientMove
	err := json.NewDecoder(r.Body).Decode(&move)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println("Error decoding client request data: ", err)
		return
	}
	fromRow := move.From[0]
	fromCol := move.From[1]
	toRow := move.To[0]
	toCol := move.To[1]

	fromSq := Game.Board.Squares[fromRow][fromCol]
	toSq := Game.Board.Squares[toRow][toCol]
	userMove := &board.Move{
		Turn:  Game.Turn,
		Piece: fromSq.Piece,
		From:  fromSq,
		To:    toSq,
	}
	if move.Promotion == "QUEEN" {
		promotedPiece := Game.Board.CreatePiece(Game.Turn, move.Promotion)
		userMove.Promotion = promotedPiece
	}
	var data map[string]interface{}
	receipt, gameError := Game.ExecuteTurn(userMove)
	log.Println(receipt)
	if gameError != nil {
		data = map[string]interface{}{
			"valid":   false,
			"receipt": receipt,
		}
	} else {
		data = map[string]interface{}{
			"color":     userMove.Turn,
			"type":      userMove.Type,
			"valid":     true,
			"from":      move.From,
			"to":        move.To,
			"promotion": move.Promotion,
			"receipt":   receipt,
		}
	}

	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func botMove(w http.ResponseWriter, r *http.Request) {
	if Game.Board.Checkmate {
		handleCheckmate(w)
		return
	}
	move := Game.Bot.CreateMove(Game.Board)
	receipt, _ := Game.ExecuteTurn(move)
	data := map[string]interface{}{
		"type":    move.Type,
		"color":   move.Turn,
		"from":    []int{move.From.Row, move.From.Column},
		"to":      []int{move.To.Row, move.To.Column},
		"receipt": receipt,
	}
	if move.Promotion != nil {
		data["promotion"] = move.Promotion.Type()
	}
	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func handleCheckmate(w http.ResponseWriter) {
	data := map[string]interface{}{
		"type":   "CHECKMATE",
		"winner": Game.Turn,
	}
	json, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
	return
}