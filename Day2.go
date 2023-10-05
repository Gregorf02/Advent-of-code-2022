package main

import(
  "fmt"
  "os"
  "bufio"
)

func main(){
  fin, _ := os.Open(os.Args[1])
  var somma int

  scanner := bufio.NewScanner(fin)
  for scanner.Scan(){
    testo := scanner.Text()
    //fmt.Println(testo)
    somma += calcolo2(testo)
  }
  fmt.Println(somma)
}


func calcolo1(s string) (n int) {

  if string(s[2])=="X"{
    n=1
  }else if string(s[2])=="Y"{
    n=2
  }else if string(s[2])=="Z"{
    n=3
  }

  if s=="A Y" || s=="B Z" || s=="C X"{
    n+=6
  }else if s=="A X" || s=="B Y" || s=="C Z"{
    n+=3
  }

  return
}

func calcolo2(s string) (n int) {
  /*A roccia B carta C forbici*/
  /*X roccia Y carta Z forbici*/
  if string(s[2])=="Z"{
    n=6
    if string(s[0])=="A"{
      n+=2
    }else if string(s[0])=="B"{
      n+=3
    }else {
      n+=1
    }
  }else if string(s[2])=="Y"{
    n=3
    if string(s[0])=="A"{
      n+=1
    }else if string(s[0])=="B"{
      n+=2
    }else {
      n+=3
    }
  }else{
    if string(s[0])=="A"{
      n+=3
    }else if string(s[0])=="B"{
      n+=1
    }else {
      n+=2
    }
  }

  return
}
