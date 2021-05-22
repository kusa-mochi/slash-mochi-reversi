// メソッドの概要：
//   リバーシの現在のプレイ状況から、次のプレイヤーがどこに石を置くべきかを計算し返す。
// 入力：
//   ・現在の盤面の状態
//   ・次の手のプレイヤー（黒/白）
//   ・探索の深さ
// 出力：
//   ・次の一手が置けるか置けないか
//   ・次の一手をどこに置くべきか
//     ・何行目
//     ・何列目
// 作成者：
//   もち（https://slash-mochi.net/）
// 作成履歴：
//   2021/xx/xx ver.1.0

package main
import (
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

type CurrentGameStatus struct {
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

func NextReversiMove(gameStatus CurrentGameStatus) (ReversiResponse, error) {
    fmt.Println("----- NextReversiMove function begin. -----");
    fmt.Println("next player:", gameStatus.NextPlayerColor);
    var ret ReversiResponse = ReversiResponse{
            CanPut: true,
            Position: ReversiPosition{
                Horizontal: 3,
                Vertical: 4,
            },
        }
    fmt.Println("----- NextReversiMove function fin. -----");
	return ret, nil
}

func main() {
    lambda.Start(NextReversiMove);
}
