package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main(){
  fin, _ := os.Open(os.Args[1])
  scanner := bufio.NewScanner(fin)
  somma:= 0
  cont := 0
  testo := ""

  for scanner.Scan(){
    testo += scanner.Text()+"\n"
    cont++
    if cont==3{
      l:=trovaLettera2(testo)
      fmt.Println(testo, string(l))
      cont=0
      testo=""
      somma+=int(calcolo(l))
    }
  }
  fmt.Println(somma)
}

func trovaLettera2(sus string) byte{
  s:=strings.Split(sus, "\n")
  for i:=0; i<len(s[0]); i++{
    if trovaIn(s[1], s[0][i]) && trovaIn(s[2], s[0][i]){
      return s[0][i]
    }
  }
  return s[0][0]
}

func trovaIn(s string, l byte) bool {
  for i:=0; i<len(s); i++{
    if s[i]==l{
      return true
    }
  }
  return false
}

func trovaLettera(s string) byte {
  s1:=s[:len(s)/2]
  s2:=s[len(s)/2:]

  for i:=0; i<len(s1); i++ {
    for j:=0; j<len(s2); j++ {
      if s1[i]==s2[j]{

        //fmt.Print(s1[i]-96, " ")
        return (s1[i])
      }
    }
  }
  return s[0]
}

func calcolo(l byte) byte {
  if l>96{
    return l-96
  }else{
    return l-38
  }
}
