package main

import(
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main(){
  fin, _ := os.Open(os.Args[1])
  scanner := bufio.NewScanner(fin)
  x:=1
  cicli:=0
  var CRT []string
  sp:="###....................................."
  ns:=""

  for scanner.Scan(){
    comando:=scanner.Text()
    cicli, x, sp, ns, CRT = leggiComando(comando, x, cicli, sp, ns, CRT)
    //fmt.Println("", ns, "\n", sp, cicli)
  }

  stampaCRT(CRT)
}

func leggiComando(s string, x int, c int, sp string, ns string, CRT []string) (int, int, string, string, []string) {
  if s=="noop"{
    c++
    if string(sp[c-1])=="#"{
      ns+="#"
    }else{
      ns+="."
    }
    if c%40==0{
      c=0
      CRT=append(CRT, ns)
      //fmt.Println( ns)
      ns=""
    }
  }else{
    for i:=0; i<2; i++{
      c++
      if string(sp[c-1])=="#"{
        ns+="#"
      }else{
        ns+="."
      }
      if c%40==0{
        c=0
        CRT=append(CRT, ns)
        //fmt.Println( ns)
        ns=""
      }
    }
    num, _ := strconv.Atoi(s[5:])
    x+=num
    sp=spostaSp(x)
    //fmt.Println(num)
  }

  return c, x, sp, ns, CRT
}

func spostaSp(n int) string {
  s:=""
  for i:=0; i<40; i++ {
    if i+1==n || i==n || i-1==n {
      s+="#"
    }else{
      s+="."
    }
  }
  return s
}

func stampaCRT(CRT []string){
  for i:=0; i<len(CRT); i++{
    fmt.Println("->",CRT[i],"<-")
  }
}
