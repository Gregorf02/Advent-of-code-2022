package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func main(){
  fin, _ := os.Open(os.Args[1])
  testo:=""

  scanner := bufio.NewScanner(fin)
  for scanner.Scan(){
    testo += scanner.Text()+"\n"
  }
  numerelli:=strings.Split(testo, "\n\n")
  //fmt.Println(numerelli)

  max1:=0
  max2:=0
  max3:=0

  for i:=0; i<len(numerelli); i++{
    n:=somma(numerelli[i])
    if n>max1{
      max3=max2
      max2=max1
      max1=n
    }else if n>max2{
      max3=max2
      max2=n
    }else if n>max3{
      max3=n
    }
  }
  fmt.Println(max1, max2, max3)
  fmt.Println(max1+ max2+ max3)
}

func somma (s string) (somma int) {
  sus:=strings.Split(s, "\n")
  for i:=0; i<len(sus); i++{
    n, _:=strconv.Atoi(sus[i])
    somma+=n
  }
  return
}
