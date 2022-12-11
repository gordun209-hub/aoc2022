package day11

//
// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )
//
// const MonkeyCount = 7
//
// type Monke struct {
// 	items     string
// 	operation func(x int) int
// 	test      func(x int) bool
// }
//
// type Monkaes []Monke
//
// func ReadFile() []string {
// 	file, _ := os.ReadFile("./day11/sample")
//
// 	lines := strings.Split(string(file), "\n\n")
// 	return lines
// }
//
// func Maim() {
// 	monkas := Monkaes{}
// 	input := ReadFile()
// 	for _, monke := range input {
// 		indivitualMonke := strings.Split(monke, "\n")
// 		monkeitem := getMonkeItem(indivitualMonke[2]) // this returns monkeitems
// 		getMonkeOperation(indivitualMonke[1])
// 		monkas = append(monkas, Monke{
// 			items: monkeitem,
// 		})
//
// 		// for _, monkePerson := range indivitualMonke {
// 		// 	getMonkeItem(monkePerson)
// 		// }
// 	}
// }
//
// func getMonkeItem(s string) string {
// 	Item := ""
// 	itemDesc := strings.Fields(s)
// 	items := itemDesc[2:]
//
// 	for _, item := range items {
// 		format := strings.ReplaceAll(item, ",", " ")
// 		Item += format
//
// 	}
// 	return Item
// }
//
// func getMonkeOperation(s string) {
// 	fmt.Println(s)
// }
//
// func getMonkeTest(s []string) {
// 	fmt.Println(s)
// }
