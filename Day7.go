package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

type file struct{
    nome  string
    val   int
}

func main(){

  m:=make(map[string][]file)
  pointer:="/"

  fin, _:=os.Open(os.Args[1])
  scanner:=bufio.NewScanner(fin)

  for scanner.Scan() {
    comando:=scanner.Text()
    fmt.Println(comando)
    pointer = leggiComando(m, comando, pointer)
    fmt.Print(pointer, "      ")
  }
  fmt.Println("Per parte 1 \nSomma delle cartelle sotto i 100000: ", controlla(m))


  fmt.Println("\n-------------------------------------------------------------")
  //70000000-41609574=28390426
  max:=(trovaMax(m))
  fmt.Println("Per parte 2\nIl mio peso totale è:", max)
  rim:=70000000-max
  fmt.Println("Quindi lo spazio rimanente è:", rim)
  serve:=30000000-rim
  fmt.Println("Mi serve ancora:", serve)
  fmt.Println("La cartella più piccola che fa abbastanza spazio pesa:", trovaCartella(m, max, serve))
}


func leggiComando(m map[string][]file, s string, pointer string) string {
  leggi:=strings.Split(s, " ")
  switch leggi[0] {
  case "$":
    if leggi[1] == "cd" {
      switch leggi[2] {
      case "/":
        return "/"
      case "..":
        return indietro(pointer)
      default:
        return avanti(pointer, leggi[2])
      }
    }else{
      //stampaMappa(m)
      stampaSottoMappa(m, pointer)
    }

  case "dir":
    var file file
    file.nome = "new"
    file.val = 0
    m[pointer]=append(m[pointer], file)

  default:
    n, _ :=strconv.Atoi(leggi[0])
    var file file
    file.nome = leggi[1]
    file.val = n
    m[pointer]=append(m[pointer], file)
  }
  return pointer
}

func stampaMappa(m map[string][]file){
  for k, v := range m {
    fmt.Println("Chiave:", k, "Val:", v)
  }
}

func stampaSottoMappa(m map[string][]file, pointer string){

  for k, _ := range m {
    if strings.Contains(k, pointer){
      fmt.Println("Punto", k, "Val:")
      for i:=0; i<len(m[k]); i++ {
        fmt.Print(m[k][i])
      }
    }
  }
}

func avanti(s1 string, s2 string) string {
  return s1+"_"+s2
}

func indietro(s string) string {
  for i:=len(s)-1; i>0; i--{
    if string(s[i])=="_"{
      return s[:i]
    }
  }
  return s
}

func controlla(m map[string][]file) int {
  fmt.Println("\n-------------------------------------------------------------")
  sum:=0
  for k, _ := range m {
    n:=valore(m, k)
    if n<=100000 {
      sum+=n
    }
  }
  return sum
}

func trovaCartella(m map[string][]file, max int, serve int) int {

  for k, _ := range m {
    n := valore(m, k)
    if n>serve && n<max {
      max=n
    }
  }
  return max
}

func trovaMax(m map[string][]file) int {
  max:=0
  for k, _ := range m {
    n := valore(m, k)
    if n>max{
      max=n
    }
  }
  return max
}

func valore(m map[string][]file, kappona string) int {
  valore := 0
  for k, _ := range m {
    if strings.Contains(k, kappona){
      for i:=0; i<len(m[k]); i++{
        valore += m[k][i].val
        }
      }
    }
  return valore
}
