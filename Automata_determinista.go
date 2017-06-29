package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
  "strconv"
  "strings"
  //"strings"
)
var m = 0
var estado = 0
var i = 0


//Estas dos pequeñas funciones son de las mas importantes del programa
//error maneja errores que salgan de la verificación de los Automatas
//mientras revisar es la funcion que se encarga de revisar que el estado pertenezca
// a ese automata.
func error() {
  fmt.Println("Se quedo en el estado: error en la cadena",estado)
  os.Exit(0)
}
func revisar(now int, cad string)int{
  if strconv.Itoa(now) == cad {
    return 1
  }else{
    return 0
  }
}

func check_n(now int, e string) int{
//En esta parte del codigo se verifica que no se haya sobrepasado el tamaño de la cadena
//Que ingreso el usuario.
//Si es igual a la longitud, entonces se verifica que el estado actual sea el estado final del automata.
  if now > len(e){
    return 0
  }else if now == len(e){
    if i == 2{
      if estado == 1 || estado == 4 {
        fmt.Println("Estado final: ",estado)
        return 3
      }else{
        return 0
      }
    }else{
      if estado == 3 {
        fmt.Println("Estado final: ",estado)
        return 3
      }else{
        return 0
      }
    }
  }
  return 2
}

func primer_auto(cadena string)  {
  for {
    act := strings.ToLower(string(cadena[m]))
    fmt.Println("estado: ",estado, "Ingresa: ", act)
    if estado == 0 {
      if act == "a" {
        estado = 1
      }else if act == "i" {
        estado = 4
      }else {
        error()
      }
    }else if estado == 4 {
      if act == "u" {
        estado = 3
      }else{
        error()
      }
    }else if estado == 2 {
      if act == "i" || act == "o" {
        estado = 3
      }else{
        error()
      }
    }else if estado == 3 {
      if act == "o" {
        estado = 4
      }else{
        error()
      }
    }else{
      if act == "e" && estado == 1 {
        estado = 2
      }else{
        error()
      }
    }
    m = m+1
    check := check_n(m,cadena)
    if check == 0{
      error()
    }
    if check == 3{
      fmt.Println("Finalizado\n\n")
      break
    }
  }
}

/*
q0: 0 : q1* Si la cadena termina
q0: 1 : q3
q1: 1 : q2
q2: 1 : q1*
q3: 0 : q4 *
q4: 0 : q3
q2 nunca manda un 0 Si es 2 entonces no se verifica el 0
q3 nunca manda 1 Si es 3 no se verifica el 1
q4 nunca manda un 1 Si es 4 no se verifica el 1
*/



func segun_auto(cadena string)  {
  //este codigo entra a un for infinito, para que vaya recorriendo la cadena ingresada
  //por el usuario.
  //va verificando el estado actual para poder avanzar al que sigue
  //Si surge algún error, se manda a llamar a error()
  for {
    act := string(cadena[m])
    fmt.Println("estado: ",estado, "Ingresa: ", act)
    if estado == 2 || estado == 1 {
      h := revisar(1,act)
      if h == 1{
        if estado == 1{
          estado = 2
        }else{
          estado = 1
        }
      }else {
        error()
      }
    }else if (estado == 3 || estado == 4){
      h := revisar(0,act)
      if h == 1{
        if estado == 3{
          estado = 4
        }else{
          estado = 3
        }
      }else {
        error()
      }
    }else{ ///////
      h := revisar(0,act)
      if h == 1{
        estado = 1
      }else if h == 0{
        h := revisar(1,act)
        if h == 1{
          estado = 3
        }else{
          error()
        }
      }
    }
    m = m+1
    check := check_n(m,cadena)
    if check == 0{
      error()
    }
    if check == 3{
      fmt.Println("\nFinalizado\n\n")
      break
    }
  }
}

func main() {
  var cadena string
  for {
    fmt.Println("***** Seleccione una opción de automata")
    fmt.Println("Automatas disponibles: 1 2")
    fmt.Scanln(&i)
    switch i {
      case 1:
        file, err := ioutil.ReadFile("AFD1.txt") // For read access.
        if err != nil {
        	log.Fatal(err)
        }
        str := string(file)
        fmt.Println(str)
        fmt.Println("Ingrese su cadena")
        fmt.Scanln(&cadena)
        primer_auto(cadena)
      case 2:
        file, err := ioutil.ReadFile("AFD2.txt") // For read access.
        if err != nil {
        	log.Fatal(err)
        }
        str := string(file)
        fmt.Println(str)
        fmt.Println("Ingrese su cadena")
        fmt.Scanln(&cadena)
        segun_auto(cadena)
    }
    m = 0
    estado = 0
  }
}
