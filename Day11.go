package main

import(
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

type scimmia struct{
  num int
  items []int
  operation string //simbolo + num senza spazio
  test []int // pos 0 vede se è divisibile, 1 lancia a true, 2 lancia a false
  cont int
}

func main(){
  fin, _ := os.Open(os.Args[1])
  scanner:=bufio.NewScanner(fin)
  text:=""
  var ugabuga []scimmia

  for scanner.Scan(){
    text+=scanner.Text()+"\n"
    if scanner.Text()==""{
      ugabuga = append(ugabuga, scimmialo(text))
      text=""
    }
  }
  ugabuga = append(ugabuga, scimmialo(text))
  //stampaScimmie(ugabuga)

  mcm:=1
  for i:=0; i<len(ugabuga); i++{
    mcm*=ugabuga[i].test[0]
  }
  //fmt.Println(mcm)

  for i:=1; i<10001; i++ {
    for j:=0; j<len(ugabuga); j++ {
      ugabuga=gioca(ugabuga, j, mcm)
    }
    if i==1 || i==20 || i==1000 || i==2000 || i==3000 || i==4000 || i==8000 || i==10000{
      fmt.Println("== After round", i, "==")
      stampaCont(ugabuga)
      fmt.Println()
    }
  }


  trovaris(ugabuga)
}

func stampaCont(ugabuga []scimmia){
  for i:=0; i<len(ugabuga); i++{
    fmt.Println("Monkey", i, "inspected items", ugabuga[i].cont, "times.")
  }
}

func trovaris(ugabuga []scimmia) {
  max1:=0
  max2:=0
  for i:=0; i<len(ugabuga); i++{
    if ugabuga[i].cont>max1{
      max2=max1
      max1=ugabuga[i].cont
    }else if ugabuga[i].cont>max2{
      max2=ugabuga[i].cont
    }
  }
  fmt.Println(max1, max2, max1*max2)
  fmt.Println(ugabuga[0].items)
}

func gioca (ugabuga []scimmia, j int, mcm int) []scimmia{
  //fmt.Println("Scimmia", j)
  for i:=0; i<len(ugabuga[j].items); i++{
    //fmt.Println(len(ugabuga[j].items))
    //fmt.Println(" La schimmia guarda l'oggetto", ugabuga[j].items[i])
    ugabuga[j].cont++
    n:=operazione(ugabuga[j], ugabuga[j].items[i], mcm)
    //n/=3
    //fmt.Println("La schimmia si stufa, livello di preoccupazione:", n)
    if n%ugabuga[j].test[0]==0{
      //fmt.Println("Il livello di preoccupazione è divisibile per", ugabuga[j].test[0])
      ugabuga[ugabuga[j].test[1]].items=append(ugabuga[ugabuga[j].test[1]].items, n)
      //fmt.Println("L'item con danger", n, "viene lanciato alla scimmia", ugabuga[j].test[1])
      //ugabuga[j].items=ugabuga[j].items[1:]
    }else{

      //fmt.Println("Il livello di preoccupazione NON è divisibile per", ugabuga[j].test[0])
      ugabuga[ugabuga[j].test[2]].items=append(ugabuga[ugabuga[j].test[2]].items, n)
      //ugabuga[j].items=ugabuga[j].items[1:]
      //fmt.Println("L'item con danger", n, "viene lanciato alla scimmia", ugabuga[j].test[2])
    }
  }
  ugabuga[j].items=ugabuga[j].items[:0]
  return ugabuga
}

func operazione(scimmia scimmia, item int, mcm int) int {
  n, _ := strconv.Atoi(scimmia.operation[1:])
  if n==0{
    n=item
  }
  //fmt.Println("Livello di preoccupazione",string(scimmia.operation[0]),n )
  if string(scimmia.operation[0])=="*" {
    return (item*n)%mcm
  }else{
    return (item+n)%mcm
  }
}

func stampaScimmie(ugabuga []scimmia) {
  for i:=0; i<len(ugabuga); i++{
    fmt.Println("Numero:", ugabuga[i].num)
    fmt.Print("Items:")
    for j:=0; j<len(ugabuga[i].items); j++{
      fmt.Print(ugabuga[i].items[j], " ")
    }
    fmt.Println("\nOperation:", ugabuga[i].operation)
    fmt.Print("Test:")
    for j:=0; j<len(ugabuga[i].test); j++{
      fmt.Print(ugabuga[i].test[j], " ")
    }
    fmt.Println("\n")
  }
}

func scimmialo (s string) scimmia {
  cose := strings.Split(s, "\n")
  var bubu scimmia

  s = string(cose[0][7])
  bubu.num, _ = strconv.Atoi(s)

  cose[1]=strings.ReplaceAll(cose[1], ",", "")
  temp:=strings.Split(cose[1], " ")
  for i:=4; i<len(temp); i++{
    num, _ := strconv.Atoi(temp[i])
    bubu.items=append(bubu.items, num)
  }

  temp=strings.Split(cose[2], " ")
  bubu.operation=temp[len(temp)-2]+temp[len(temp)-1]


  temp=strings.Split(cose[3], " ")
  s=temp[len(temp)-1]
  num, _ :=strconv.Atoi(s)
  bubu.test=append(bubu.test, num)

  s=string(cose[4][len(cose[4])-1])
  num, _ =strconv.Atoi(s)
  bubu.test=append(bubu.test, num)

  s=string(cose[5][len(cose[5])-1])
  num, _ =strconv.Atoi(s)
  bubu.test=append(bubu.test, num)

  return bubu
}
