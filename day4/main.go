package main
import(
  "bufio"
  "os"
  "strings"
  "fmt"
  "log"
  "strconv"
  "regexp"
)

const file_name = "puzzle_input.txt"

func main() {
  f, err := os.Open(file_name)
  if err != nil {
    log.Fatalf("could not open file %s", file_name)
  }
  var (
    scanner = bufio.NewScanner(f)
    numValid = 0
    passesValidation = 0
    id = make(map[string]string)
  )

  for scanner.Scan() {
    line := scanner.Text()
    if line == "" {
      if len(id) == 8 {

        if validate(id) == true {
          passesValidation += 1
        }
        numValid += 1
      }
      if len(id) == 7 {
        if _, ok := id["cid"]; !ok {
          if validate(id) == true {
            passesValidation += 1
          }
          numValid += 1
        }
      }
      id = make(map[string]string)
    } else {
      splitLine := strings.Split(line, " ")
      for _, attr := range(splitLine) {
        splitAttr := strings.Split(attr, ":")
        id[splitAttr[0]] = splitAttr[1]
      }
    }
  }

  fmt.Println("numValid ", numValid)
  fmt.Println("passes validation ", passesValidation)
}

/*
hi

*/

func validate(id map[string]string) bool {

  valid := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl":true, "oth": true}
  for k, v := range(id) {
    switch k {
      case "byr":
        year, _ := strconv.Atoi(v)
        if year < 1920 || year > 2002 {
          return false
        }
      case "iyr":
        year, _ := strconv.Atoi(v)
        if year < 2010 || year > 2020 {
          return false
        }
      case "eyr":
        year, _ := strconv.Atoi(v)
        if year < 2020 || year > 2030 {
          return false
        }
      case "hgt":
        cm := strings.Split(v, "cm")
        if len(cm) == 2 {
          num , _ := strconv.Atoi(cm[0])
          if num < 150 || num > 193 {
            return false
          }
        }

        in := strings.Split(v, "in")
        if len(in) == 2 {
          num , _ := strconv.Atoi(in[0])
          if num < 59 || num > 76 {
            return false
          }
        }
      case "hcl":
        if string(v[0]) != "#" {
          return false
        }
        color := v[1:]
        //if len(color) != 6 {
        //  return false
        //}
	matched, _ := regexp.MatchString("[a-fA-F0-9]{6}", color)

        if !matched {
          return false
        }
      case "ecl":
        if valid[v] != true {
          return false
        }
      case "pid":
	matched, _ := regexp.MatchString("[0-9]{9}", v)
        if !(len(v) == 9 && matched) {
          return false
        }
      default:
        continue
    }
  }
  return true
}
