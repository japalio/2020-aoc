package main

import(
  "bufio"
  "fmt"
  "log"
  "os"
)

const file_name = "puzzle_input.txt"

type slope struct {
  x int
  y int
}

func main() {
  input := readInput()
  // all the slopes for part 2
  // part 1 was only 3,1
  slopes := []slope{
    slope {
      x: 3,
      y: 1,
    },
    slope {
      x: 1,
      y: 1,
    },
    slope {
      x: 5,
      y: 1,
    },
    slope {
      x: 1,
      y: 2,
    },
    slope {
      x: 7,
      y: 1,
    },
  }

  product := 1

  for _, s := range(slopes) {
    right := s.x
    down := s.y
    y := 0
    x := 0
    numTrees := 0

    for y < len(input) {
      line := input[y]
      for x >= len(line) {
        line += line
      }
      if string(line[x]) == "#" {
        numTrees += 1
      }
      y += down
      x += right
    }
    product *= numTrees
  }

  fmt.Println("Product ", product)

}

func readInput() []string {
  f, err := os.Open(file_name)
  if err != nil {
    log.Fatalf("failed to open file %s", file_name)
  }

  var (
    scanner = bufio.NewScanner(f)
    lines []string
  )

  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }

  return lines
}
