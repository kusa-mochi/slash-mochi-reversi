// Function summary:
//   this function show you where to put your next stone on reversi board using current game status.
// input:
//   - current status on a reversi board
//   - next player's stone color (black/white)
//   - search tree depth
// return:
//   ・次の一手が置けるか置けないか
//   ・次の一手をどこに置くべきか
//     ・何行目
//     ・何列目
// author:
//   mochi (https://slash-mochi.net/)
// history:
//   2021/xx/xx ver.1.0

package main
import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
)

// region enum CellState
type CellState int
const (
    Empty CellState = iota
    Black
    White
    Wall
)
// endregion

type ReversiEvent struct {
    NextPlayerColor string `json:"NextPlayerColor"`
    CurrentBoard [][]CellState `json:"CurrentBoard"`
}

type ReversiPosition struct {
    Horizontal int `json:"Horizontal"`
    Vertical int `json:"Vertical"`
}

type ReversiResponse struct {
	CanPut bool `json:"CanPut"`
    Position ReversiPosition `json:"Position"`
}

func reversi_nextmove_handler(ctx context.Context, reversiEvent ReversiEvent) (ReversiResponse, error) {
    fmt.Println("----- reversi_nextmove_handler function begin. -----");
    fmt.Println("next player:", reversiEvent.NextPlayerColor);
    var ret ReversiResponse = ReversiResponse{
            CanPut: true,
            Position: ReversiPosition{
                Horizontal: 3,
                Vertical: 4,
            },
        }
    fmt.Println("----- reversi_nextmove_handler function fin. -----");
	return ret, nil
}

func main() {
    lambda.Start(reversi_nextmove_handler);
}
