package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
)

const MapSize int = 40
const Turns int = 1000000

const should_log bool = false

var chance_card_squares = [...]int {7, 22, 36}
var chance_card_targets = [...]int {39, 5, 11, 24}
var community_card_squares = [...]int {2, 17, 33}
const go_to_jail_square = 30
const jail_square = 10

type Player struct {
    position int
}

func(p *Player) move_by(n int){
    p.position += n
    p.position %= MapSize
}

func roll() (int, int){
    return rand.IntN(6) + 1, rand.IntN(6) + 1
}

func handle_special_fields(p *Player){
    if p.position == go_to_jail_square{
        p.position = jail_square
        return
    }
    if !slices.Contains(community_card_squares[:], p.position){
        // There are 16 cards and one of them moves the player to the first square
        r := rand.IntN(16)
        if r == 0 {
            p.position = 0
        }
        return
    }
    if !slices.Contains(chance_card_squares[:], p.position){
        // half of the chance cards don't move the player
        if rand.IntN(2) == 0{
            return
        }

        r := rand.IntN(8);
        // 4 cards have specific target squares, rest depends on your position
        if r < 4{
            p.position = chance_card_targets[r];
            return
        }
        // 2 "go to the nearest train station cards" 
        if r < 6{
            var target int = 10 * (p.position / 10) + 5
            p.position = target
            return
        }
        if r == 6{
            p.position -= 3
            if p.position < 0 {
                p.position += MapSize
            }
            return
        }
        // I won't be bothered to do the last one ("go to nearest shop")


        return
    }
}


func main(){
    var square_visits [MapSize]int
    var players [4]Player

    fmt.Printf("Simmulating %d turns on a %d square map \n\n", Turns, MapSize)
    for turn := range Turns{
        if should_log{
            fmt.Println("Start of turn ", turn + 1)
        }
        for i := range players{
            r1, r2 := 0, 0
            roll_sum := 0
            for r1 == r2{
                r1, r2 = roll()
                if should_log{
                    fmt.Println("Player ", i+1, " rolled ", r1, " ", r2)
                }
                roll_sum += r1 + r2
            }
            players[i].move_by(roll_sum)

            if should_log{
                fmt.Println("Player ", i+1, " moved to square ", players[i].position)
            }
            handle_special_fields(&players[i]);
            square_visits[players[i].position] += 1
        }
    }
    
    fmt.Printf(" --- Summary --- \n\n")

    for i, v := range(square_visits){
        fmt.Println("square ", i, "\t visits: ", v)
    }
        

    f, _ := os.Create("res.csv")
    defer f.Close()
    w := bufio.NewWriter(f)
    
    var exclude_squares = [...]int{0, 2, 7, 10, 17, 22, 30, 33, 36}

    for i, v := range(square_visits){
        if slices.Contains(exclude_squares[:], i){
            continue
        }
        fmt.Fprintf(w, "%d, %d\n", i, v)
    }
    w.Flush()
}
