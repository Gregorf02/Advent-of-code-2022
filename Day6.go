package main

import(
  "fmt"
  "os"
  "bufio"
)

func main(){
  fin,_:=os.Open(os.Args[1])
  scanner:=bufio.NewScanner(fin)
  scanner.Scan()
  s:=scanner.Text()
  fmt.Println(s, "\n", trovaSec(s))
}

func trovaSec(s string) int {
  c:=14
  for i:=0; i<len(s); i++{
      sub:=s[i:i+14]
      if ripetiSub(sub){
        return c
    }
    c++
  }
  return c
}

func ripetiSub(sub string) bool {
  for i:=0; i<len(sub); i++{
    for j:=i+1; j<len(sub); j++{
      if sub[i]==sub[j]{
        return false
      }
    }
  }
  return true
}
