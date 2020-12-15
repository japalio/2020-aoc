package main

import(
  "fmt"
  "bufio"
  "log"
  "os"
)

const file_name = "puzzle_input.txt"


func main() {
  seats := readFile()
  highestSeatId := 0

  seatMap := make(map[int]bool)
  fmt.Println(len(seats))
  for _, seat := range(seats) {
    row, col := parseSeat(seat)
    seatId := row * 8 + col
    seatMap[seatId] = true


    if seatId > highestSeatId {
      highestSeatId = seatId
    }
  }
  fmt.Println("highest seat id ", highestSeatId)
  fmt.Println(len(seatMap))

  for row := 1; row < 127; row +=1 {
    for col := 0; col <= 7; col +=1 {
      id := row*8 + col
      if seatMap[id] == false {
        if seatMap[id-1] == true && seatMap[id+1] == true {
          fmt.Println("missing seat", id)
          break
        }
      }
    }
  }
}

func parseSeat(seat string) (int, int) {
  var (
    lowerRow = 0
    upperRow = 127
    lowerCol = 0
    upperCol = 7
    row = 0
    col = 0
  )
  for i, char := range(seat) {
    c := string(char)
    add := 0
    if i < 7 {
      if c == "F" {
        if i == 6 {
          row = lowerRow
        }
        // bruh why do we add 1 here
        upperRow -= (upperRow - lowerRow)/2 + 1
      } else {
        if i == 6 {
          row = upperRow
        }
        if (upperRow - lowerRow) % 2 == 1 {
          add += 1
        }

        lowerRow += (add + upperRow - lowerRow)/2
      }
    } else {
      if c == "L" {
        if i == 9 {
          col = lowerCol
        }
        upperCol -= (upperCol - lowerCol)/2  + 1
      } else {
        if i == 9 {
          col = upperCol
        }
        add = 0
        if (upperCol - lowerCol) % 2 == 1 {
          add += 1
        }
        lowerCol += (add + upperCol - lowerCol)/2
      }
    }
  }
  return row, col
}

func readFile() []string {
  f, err := os.Open(file_name)
  if err != nil {
    log.Fatalf("could not open %s", file_name)
  }
  var (
    scanner = bufio.NewScanner(f)
    seats []string
  )

  for scanner.Scan(){
    seats = append(seats, scanner.Text())
  }

  return seats
}
