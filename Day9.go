// 3973 troppo alto

package main

import(
  "fmt"
  "bufio"
  "os"
  "strconv"
)

type coordinate struct {
  x int
  y int
}

func main(){

  fin, _ := os.Open(os.Args[1])
  scanner := bufio.NewScanner(fin)
  mappa := make(map[coordinate]bool)
  var corda [10]coordinate

  for scanner.Scan(){
    comando := scanner.Text()
    fmt.Println(comando)
    corda = eseguicomando(mappa, corda, comando)
  }
  fmt.Println(len(mappa))
}

func eseguicomando(m map[coordinate]bool, corda [10]coordinate, s string) [10]coordinate {
  com := string(s[0])
  n, _ := strconv.Atoi(s[2:])

  switch com {
  case "R":
    for i:=0; i<n; i++ {
      corda[0].x++
      for j:=0; j<len(corda)-1; j++{
        if !controllaRaggio(corda[j].x, corda[j].y, corda[j+1].x, corda[j+1].y){
          corda[j+1] = aggiustaCoda(corda[j], corda[j+1])
        }
      }
      aggiornaMappa(m, corda)
    }

  case "L":
    for i:=0; i<n; i++ {
      corda[0].x--
      for j:=0; j<len(corda)-1; j++{
        if !controllaRaggio(corda[j].x, corda[j].y, corda[j+1].x, corda[j+1].y){
          corda[j+1] = aggiustaCoda(corda[j], corda[j+1])
        }
      }
      aggiornaMappa(m, corda)
    }

  case "U":
    for i:=0; i<n; i++ {
      corda[0].y++
      for j:=0; j<len(corda)-1; j++{
        if !controllaRaggio(corda[j].x, corda[j].y, corda[j+1].x, corda[j+1].y){
          corda[j+1] = aggiustaCoda(corda[j], corda[j+1])
        }
      }
      aggiornaMappa(m, corda)
    }

  case "D":
    for i:=0; i<n; i++ {
      corda[0].y--
      for j:=0; j<len(corda)-1; j++{
        if !controllaRaggio(corda[j].x, corda[j].y, corda[j+1].x, corda[j+1].y){
          corda[j+1] = aggiustaCoda(corda[j], corda[j+1])
        }
      }
      aggiornaMappa(m, corda)
    }
  }
  return corda
}

func aggiustaCoda (testa coordinate, coda coordinate) coordinate {


    if testa.x<coda.x{
      coda.x--
    }else if testa.x>coda.x{
      coda.x++
    }

    if testa.y<coda.y{
      coda.y--
    }else if testa.y>coda.y{
      coda.y++
    }

  return coda
}

func stampaCorda(corda [10]coordinate){
  for i:=0; i<len(corda); i++{
    fmt.Println("posizione", i, " x", corda[i].x, " y", corda[i].y)
  }
}

func aggiornaMappa(m map[coordinate]bool, corda [10]coordinate) {
  m[corda[9]]=true
}

func controllaRaggio(hx int, hy int, tx int, ty int) bool {
  if (hy==ty && (hx==tx+1 || hx==tx-1 || hx==tx)) || (hx==tx && (hy==ty+1 || hy==ty-1)){ // spaghetto per vedere se sono vicini in croce
    return true
  }else if (hx==tx+1 || hx==tx-1) && (hy==ty+1 || hy==ty-1) {
    return true
  }
  return false
}
