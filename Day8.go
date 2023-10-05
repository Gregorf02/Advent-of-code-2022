package main

import(
  "fmt"
  "os"
  "bufio"
  "strconv"
)

type tree struct{
  altezza int
  visto   bool
}

func main(){
  fin, _ := os.Open(os.Args[1])
  scanner := bufio.NewScanner(fin)
  var tabella [][]tree

  for scanner.Scan(){
    s:=scanner.Text()
    tabella=append(tabella, inserisci(s))
  }

  tabella = chekkaRighe(tabella)
  tabella = chekkaColonne(tabella)
  fmt.Println(contaTrue(tabella))
  fmt.Println(trovaSpazio(tabella))

}

func spazioOrizontale(tab [][]tree, i int, j int) int{
  spDestra:=0
  spSinistra:=0
  for cont:=j+1; cont<len(tab[i]); cont++{ //ciclo che guarda a destra che funziona
    //fmt.Println("confronto:", tab[i][j].altezza, tab[i][cont].altezza)
    spDestra++
    if tab[i][j].altezza <= tab[i][cont].altezza{
      break
    }
  }
  //fmt.Println("spDestra di ", i, j, "=", spDestra)

  for cont:=j-1; cont>=0; cont--{ //ciclo che guarda a sinistra
    //fmt.Println("confronto:", tab[i][j].altezza, tab[i][cont].altezza)
    spSinistra++
    if tab[i][j].altezza <= tab[i][cont].altezza{
      break
    }
  }
  //fmt.Println("spSinistra di ", i, j, "=", spSinistra)
  //fmt.Println(spSinistra, spDestra)
  return spDestra*spSinistra
}

func spazioVerticale(tab [][]tree, i int, j int) int{
  spSu:=0
  spGiu:=0
  for cont:=i-1; cont>=0; cont--{ //ciclo che guarda su
    //fmt.Println("confronto:", tab[i][j].altezza, tab[i][cont].altezza)
    spSu++
    if tab[i][j].altezza <= tab[cont][j].altezza{
      break
    }
  }

  for cont:=i+1; cont<len(tab); cont++{ //ciclo che guarda giu
    //fmt.Println("confronto:", tab[i][j].altezza, tab[i][cont].altezza)
    spGiu++
    if tab[i][j].altezza <= tab[cont][j].altezza{
      break
    }
  }
  fmt.Println(spSu, spGiu)
  return spSu*spGiu
}

func trovaSpazio(tab [][]tree) int {
  max:=0
  for i:=1; i<len(tab)-1; i++{
    for j:=1; j<len(tab[i])-1; j++{
      n1:=spazioOrizontale(tab, i, j)
      n2:=spazioVerticale(tab, i, j)
      n3:=n1*n2
      //fmt.Println(i, j, "=", n1, "*", n2, "=", n3)
      if n3>max{
        max=n3
      }
    }
  }
  return max
}

func contaTrue(tab [][]tree) int {
  cont:=0
  for i:=0; i<len(tab); i++{
    for j:=0; j<len(tab[0]); j++{
      if tab[i][j].visto==true{
        cont++
      }
    }
  }
  return cont
}

func chekkaColonne(tab [][]tree) [][]tree{
  for i:=0; i<len(tab[0]); i++{
    n:=0
    for j:=0; j<len(tab); j++{
      if tab[j][i].altezza>=n{
        tab[j][i].visto=true
        n=tab[j][i].altezza+1
      }
    }
    n=0
    for j:=len(tab)-1; j>0; j--{
      if tab[j][i].altezza>=n{
        tab[j][i].visto=true
        n=tab[j][i].altezza+1
      }
    }
  }
  return tab
}

func chekkaRighe(tab [][]tree) [][]tree{
  for i:=0; i<len(tab); i++{
    n:=0
    for j:=0; j<len(tab[i]); j++{
      if tab[i][j].altezza>=n{
        tab[i][j].visto=true
        n=tab[i][j].altezza+1
      }
    }
    n=0
    for j:=len(tab[i])-1; j>0; j--{
      if tab[i][j].altezza>=n{
        tab[i][j].visto=true
        n=tab[i][j].altezza+1
      }
    }
  }
  return tab
}

func inserisci(s string) (riga []tree) {
  var temp tree
  for i:=0; i<len(s); i++{
    val, _ := strconv.Atoi(string(s[i]))
    temp.altezza = val
    riga=append(riga, temp)
    //fmt.Print(riga[i].altezza)
  }
  //fmt.Println()
  return
}

func stampaTab(tab [][]tree){
  for i:=0; i<len(tab); i++{
    for j:=0; j<len(tab[i]); j++{
      fmt.Print(tab[i][j].altezza, tab[i][j].visto, "  ")
    }
    fmt.Println()
  }
}
