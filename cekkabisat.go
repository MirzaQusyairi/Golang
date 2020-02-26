/*
Program untuk mengecek apakah tahun yang diinputkan termasuk tahun kabisat atau tidak,
yang menghasilkan output berupa true atau false
*/

package main 
import "fmt"

func main(){
	var tahun int
	var proses bool

	fmt.Print("Tahun: ")
	fmt.Scanln(&tahun)
	proses = tahun%400 == 0 || tahun%4 == 0 && tahun%100 > 0
	fmt.Println("Kabisat: ",proses)
	fmt.Scanln()
}
