package main

import (
	"fmt"
	"unicode/utf8"
)


func main() {
var i int  = 21
j := true
// var hexaAwal = 0xf
var hasilHex  = 15
var k float64 = 123.456;
rune,_:=utf8.DecodeRuneInString("Я")


fmt.Printf("%v \n", i) //nilai dari i
fmt.Printf("%T \n", i) // tipe data dari i
fmt.Println("%") // print % ????
fmt.Printf("%t \n", j) // print nilai j
fmt.Printf("%b \n", i) // base 2 dari i
fmt.Printf("%c \n", rune) // base 2 dari i

fmt.Printf("%d \n", i) //base 10 dari i 
fmt.Printf("%o \n", i) //base 8 dari i
// fmt.Printf("%d \n", hexaAwal) // cari angka untuk dipakai di variabel hasilHex
fmt.Printf("%x \n", hasilHex) // base 16 f
fmt.Printf("%X \n", hasilHex) //base 16 F

fmt.Printf("%U \n", rune) //unicode character

fmt.Printf("%f \n", k) //float k
fmt.Printf("%E \n", k) //float science k 

}