package main

import (
	"fmt"

	"github.com/kungbaman/practice/mapping"
)

type movie struct {
	name        string
	year        int
	rating      float64
	genre       []string
	isSuperHero bool
}

type movie_2 struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

func (m movie) info() {
	fmt.Printf("\n%s (%d) - %.2f\n", m.name, m.year, m.rating)

	for _, g := range m.genre {
		fmt.Printf("\t%s\n", g)
	}
}

func (m *movie_2) addVote(rating float64) {
	m.votes = append(m.votes, rating)
}

type voter interface {
	addVote(float64)
}

func vote(v voter, rating float64) {
	v.addVote(rating)
}

func main() {
	// variables

	/*
		msg := "Hello Natchapol!!!"
		age := 28
		price := 22.32
		check := true

		fmt.Println("name:", msg)
		fmt.Println("age:", age)
		fmt.Println("price:", price)
		fmt.Println("check:", check)
	*/

	// if else
	num := 35

	/*
		if num == 34 {
			fmt.Println("Yes, it is Thirty Four!!")
		}
	*/

	/*
		if num == 34 && (num > 30 || num < 39) {
			fmt.Println("Yes, it is Thirty Four!!")
		}
	*/

	if num%2 == 0 {
		fmt.Println("เลขคู่")
	} else if num%2 != 0 {
		fmt.Println("เลขคี่")
	} else {
		fmt.Println("__")
	}

	// workshop If else

	// rating := 8.4

	if rating := 14.4; rating < 5.0 {
		fmt.Printf("Disappoint \n")
	} else if rating >= 5.0 && rating < 7.0 {
		fmt.Println("Normal")
	} else if rating >= 7.0 && rating < 10.0 {
		fmt.Println("Good")
	} else {
		fmt.Println("_____")
	}

	// switch case
	today := "Saturday"

	switch today {
	case "Saturday":
		fmt.Println("Today is Saturday")
		fallthrough
	case "Monday":
		fmt.Println("Today is Monday")
		fallthrough
	default:
		fmt.Printf("สวัสดีวัน %s\n", today)
	}

	// Array
	/*
		genres := [...]string{"Action", "Adventure", "Fantasy"}
		fmt.Println(len(genres))
		fmt.Printf("genres[0]: %v\n", genres[0])
		fmt.Printf("genres[1]: %v\n", genres[1])
		fmt.Printf("genres[2]: %v\n", genres[2])

		genres[1] = "Sci-Fi"
		fmt.Println("genres[1]:", genres[1])

		// (for) loop
		fmt.Println("Before for loop: ", genres)
		for i := 0; i < len(genres); i++ {
			genres[i] = "Movie: " + genres[i]
		}

		fmt.Println("After for loop: ", genres)

		for i, genre := range genres {
			fmt.Printf("genres[%d]: %s\n", i, genre)
		}
	*/

	// Slices
	// เราได้เก็บสะสมคะแนนโหวตผู้ชมไว้เป็น 2 ชุด ที่เก็บอยู่ในตัวแปร xs และ ys เรียบร้อยแล้ว
	// ให้ทำการรวมคะแนนโหวตที่อยู่ในตัวแปร xs และ ys ไปเก็บไว้ในตัวแปร votes ตามลำดับ (ค่าใน xs ทั้งหมดแล้วต่อด้วย ys).
	// ทำการแสดงผลคะแนนโหวตของผู้ชมที่อยู่ในตำแหน่ง(index)ที่ 5 ถึง 8 ของ votes ออกมาทางหน้าจอ

	var xs []float64 = []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}

	var votes []float64
	votes = append(xs, ys...)
	fmt.Printf("votes: %.0f\n", votes)
	fmt.Printf("votes 5 - 8: %.0f\n", votes[5:9])

	// Structures
	var m1 = movie{
		name:        "Avenger: Endgame",
		year:        2019,
		rating:      8.4,
		genre:       []string{"Action", "Drama"},
		isSuperHero: true,
	}

	m2 := movie{
		name:        "Infinity War",
		year:        2018,
		rating:      8.4,
		genre:       []string{"Action", "Sci-Fi"},
		isSuperHero: true,
	}

	// ประกาศ struct ได้หลายแบบ

	var mvs []movie
	mvs = append(mvs, m1, m2)

	/*
		mvs := []movie{m1, m2}
	*/

	/*
		var mvs []movie = []movie{m1, m2}
	*/

	for _, mv := range mvs {
		fmt.Printf("%#v\n", mv)
	}

	// Workshop: method
	// กำหนด: 1. ให้สร้างmethod info สำหรับ movie เพื่อเก็บแสดงผลรายละเอียด โดยประกอบด้วย ชื่อเรื่อง(string) ปี(ตัวเลข) เรตติ้ง(ตัวเลขทศนิยม) ประเภท(slice ของ string) และ isSuperHero(bool).
	//
	// Output:
	// Avengers: Endgame (2019) - 8.40
	// Genres:
	// 				Action
	// 				Drama

	m1.info()

	// pointer
	eg := &movie_2{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	// eg.addVote(8)
	vote(eg, 8)
	fmt.Println("votes: ", eg.votes)

	// Workshop: maps
	// กำหนด: 1. ให้สร้างเขียนฟังก์ชัน WordCount เพื่อนำคำซ้ำในประโยค
	//           strings.Fields น่าจะเป็นตัวช่วยได้ https://pkg.go.dev/strings#Fields
	//
	// Output:
	// map[string]int{"If":1, "a":4, "and":1, "duck":4, "is":1, "it":2, "like":3, "looks":1, "probably":1, "quacks":1, "swims":1, "then":1}

	s := "If it looks like a duck swims like a duck and quacks like a duck then it probably is a duck"
	w := mapping.WordCount(s)
	fmt.Printf("%#v\n", w)

}
