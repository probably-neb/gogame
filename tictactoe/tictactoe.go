package tictactoe

import (
	"encoding/json"
	"gogame/htmx"
	"gogame/player"
	"log"
    "strconv"
)

type Player = player.Player

type Symbol rune

const (
	X Symbol = 'X'
	O Symbol = 'O'
)

type Move struct {
	cell   string
	player *Player
}

type GameBoard struct {
    cells [9]Symbol
}

func (g *GameBoard) place(cellIdStr string, sym Symbol) {
    cellId, err := strconv.Atoi(cellIdStr[1:])
    if err != nil {
        log.Println("failed to parse cellId:", cellIdStr, err.Error())
    } else {
        log.Println("parsed cellId:", cellId)
    }
    g.cells[cellId] = sym
}

func (g *GameBoard) checkForWins() (*Symbol, int, int, int) {
    eq := func(x int, y int, z int) bool {
        return g.cells[x] == g.cells[y] && g.cells[y] == g.cells[z] && g.cells[x] != 0
    }
    // top row
    if eq(0, 1, 2) {
        return &g.cells[0], 0, 1, 2
    }
    // middle row
    if eq(3, 4, 5) {
        return &g.cells[3], 3, 4, 5
    }
    // bottom row
    if eq(6, 7, 8) {
        return &g.cells[6], 6, 7, 8
    }
    // left column
    if eq(0, 3, 6) {
        return &g.cells[0], 0, 3, 6
    }
    // middle column
    if eq(1, 4, 7) {
        return &g.cells[1], 1, 4, 7
    }
    // right column
    if eq(2, 5, 8) {
        return &g.cells[2], 2, 5, 8
    }
    // tl->br diagonal
    if eq(0, 4, 8) {
        return &g.cells[0], 0, 4, 8
    }
    // tr->bl diagonal
    if eq(2, 4, 6) {
        return &g.cells[2], 2, 4, 6
    }
    return nil, 0, 0, 0
}

func listenForMoves(msgs chan player.PlayerMsg, moves chan Move) {
	for {
		// TODO: check if channel closed
		msg := <-msgs

		// TODO: assert/check msg.Type
		var req struct {
            Headers htmx.Headers `json:"HEADERS"`
        }

		if err := json.Unmarshal(msg.Message.Data, &req); err != nil {
			log.Println(err)
			return
		}
		cell_id := msg.Message.Headers.HXTarget
		log.Println(msg.Player, "played", cell_id)
		moves <- Move{cell: cell_id, player: msg.Player}
	}
}



func Start(host *Player, guest *Player, msgs chan player.PlayerMsg, exit chan struct{}) {
	defer func() { exit <- struct{}{} }()
	moves := make(chan Move)
	symbols := map[*Player]Symbol{host: X, guest: O}
	host.Send <- Game()
	guest.Send <- Game()
    board := GameBoard{}
	go listenForMoves(msgs, moves)
	turn := X
    done := false
	for {
        if (done) {
            continue
        }
		move := <-moves
		symbol := symbols[move.player]
		if symbol != turn {
			log.Println("wrong turn it's", turn, "turn")
			// TODO: consider send not turn warning?
			continue
		}
        if turn == symbols[host] {
			log.Println("guest's turn")
			turn = symbols[guest]
		} else {
			log.Println("host's turn")
			turn = symbols[host]
		}
		newBox := Box(move.cell, symbol)
		host.Send <- newBox
		guest.Send <- newBox

        board.place(move.cell, symbol)
        winner, a, b, c := board.checkForWins()
        if winner == nil {
           continue
        }
        winBoard := WinBoard(board.cells, [3]int{a,b,c})
        host.Send <- winBoard
        guest.Send <- winBoard
        if *winner == symbols[host] {
            log.Println("host won", board.cells)
        } else {
            log.Println("guest won", board.cells)
        }
	}
}
