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
  scanner := bufio.NewScanner(fin)
  s:=""
  s2:=""
  max:=0
  sesso := true
  m := make(map[int]string)

  for scanner.Scan(){

    if sesso{
      s += scanner.Text()+"\n"
      //fmt.Println(string(scanner.Text()[1]), string(scanner.Text()[5]), string(scanner.Text()[9]))
      max = inserisci(m, scanner.Text(), max)
    }else{
      s2 += scanner.Text()+"\n"
    }

    if scanner.Text()==""{
      sesso=false
    }

  }

  fmt.Print(s)
  m=creaM(m, max)
  stampaRis(m)
  //stampaMappa(m)
  Sposta2(m, s2)
  //stampaMappa(m)
  stampaRis(m)
}

func inserisci(m map[int]string, s string, n int) int {
  cont:=0
  for i:=0; i<len(s); i++{
    if string(s[i])=="["{
      m[i]+=string(s[i+1])
      cont++
    }
  }
  if cont>n{
    n=cont
  }
  return n
}

func stampaMappa(m map[int]string) {
  for i:=1; i<len(m)+1; i++{
    fmt.Println(i, m[i])
  }
}

func creaM(m map[int]string, n int) map[int]string {

  fmt.Println()
  c:=1
  m2:= make(map[int]string)
  for k:=0; k!=4*n; k+=4{
    m2[c]=m[k]
    c++
  }
  return m2
}

func Sposta(m map[int]string, s string) {  //Per prima parte
  is := strings.Split(s[:len(s)-1], "\n")
  for i:=0; i<len(is); i++{
    div:=strings.Split(is[i], " ")
    n1, _ :=strconv.Atoi(string(div[1]))
    n2, _ :=strconv.Atoi(string(div[3]))
    n3, _ :=strconv.Atoi(string(div[5]))
    fmt.Println(n1, n2, n3)

    for j:=0; j<n1; j++ {
      temp:=string(m[n2][0])
      m[n2]=m[n2][1:]
      m[n3]=temp+m[n3]
    }
  }
}

func Sposta2(m map[int]string, s string) {
  is := strings.Split(s[:len(s)-1], "\n")
  for i:=0; i<len(is); i++{
    div:=strings.Split(is[i], " ")
    n1, _ :=strconv.Atoi(string(div[1]))
    n2, _ :=strconv.Atoi(string(div[3]))
    n3, _ :=strconv.Atoi(string(div[5]))

    temp:=string(m[n2][0:n1])
    m[n2]=m[n2][n1:]
    m[n3]=temp+m[n3]

  }
}

func stampaRis(m map[int]string){
  for i:=1; i<len(m)+1; i++{
    fmt.Print(string(m[i][0]))
  }
  fmt.Println()
}
