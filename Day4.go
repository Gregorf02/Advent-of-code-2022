package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func main(){
  fin, _:=os.Open(os.Args[1])
  scanner:=bufio.NewScanner(fin)
  cont:=0

  for scanner.Scan(){
    s:=scanner.Text()
    if overlap(s){
      cont++
    }
  }
  fmt.Println(cont)
}

func fullyContain(s string) bool {
  sez := strings.Split(s, ",")
  sez1 := strings.Split(sez[0], "-")
  sez2 := strings.Split(sez[1], "-")

  min1, _ := strconv.Atoi(sez1[0])
  max1, _ := strconv.Atoi(sez1[1])
  min2, _ := strconv.Atoi(sez2[0])
  max2, _ := strconv.Atoi(sez2[1])

  if min1<=min2 && max1>=max2 || min2<=min1 && max2>=max1{
    return true
  }else{
    return false
  }
}

func overlap(s string) bool {
  sez := strings.Split(s, ",")
  sez1 := strings.Split(sez[0], "-")
  sez2 := strings.Split(sez[1], "-")

  min1, _ := strconv.Atoi(sez1[0])
  max1, _ := strconv.Atoi(sez1[1])
  min2, _ := strconv.Atoi(sez2[0])
  max2, _ := strconv.Atoi(sez2[1])

  if min1>max2 || min2>max1{
    return false
  }else{
    return true
  }
}
