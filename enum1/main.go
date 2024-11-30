package main

import "fmt"

type TimeOfDay int // Enum tipi tanılama işemi

const (
	Morning TimeOfDay = iota
	Afternoon
	Evening
	Night
)

// enum int değerler alıyorud enum değerlerini stringe çevirmeliyiz...
// enum değerini String'e çeviren Strin metodu

func (t TimeOfDay) String() string {

	switch t {
	case Morning:
		return "Morning"

	case Afternoon:
		return "Afternoon"
	case Evening:
		return "Evening"

	case Night:
		return "Night"

	default:
		return "Unknow"
	}

}
func main() {
	// var currentTime TimeOfDay = Afternoon
	currentTime := Afternoon

	fmt.Println(currentTime)



	/*
	fmt.Println(currentTime) çağrılır.
currentTime bir TimeOfDay türüdür ve bu türün String() metodu vardır.
fmt.Println otomatik olarak currentTime.String() metodunu çağırır.
String() metodu ilgili zaman dilimine karşılık gelen string değeri döndürür ("Afternoon").
Bu string değer fmt.Println tarafından ekrana yazdırılır.
*/
}
