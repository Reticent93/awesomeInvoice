package main

import (
	"html/template"
	"os"
)

type Player struct {
	Name     string
	Position string
	Team     string
	Age      int
	Jersey   int
	Retired  bool
}

func BestPlayer(players []Player) Player {
	players = []Player{
		{
			Name:     "Michael Jordan",
			Position: "Shooting Guard",
			Team:     "Chicago Bulls",
			Age:      57,
			Jersey:   23,
			Retired:  true,
		},
		{
			Name:     "Kobe Bryant",
			Position: "Shooting Guard",
			Team:     "Los Angeles Lakers",
			Age:      41,
			Jersey:   24,
			Retired:  true,
		},
		{
			Name:     "Lebron James",
			Position: "Small Forward",
			Team:     "Los Angeles Lakers",
			Age:      36,
			Jersey:   23,
			Retired:  false,
		},
		{
			Name:     "Stephen Curry",
			Position: "Point Guard",
			Team:     "Golden State Warriors",
			Age:      33,
			Jersey:   30,
			Retired:  false,
		},
		{
			Name:     "Karim Abdul-Jabbar",
			Position: "Center",
			Team:     "Los Angeles Lakers",
			Age:      71,
			Jersey:   33,
			Retired:  true,
		},
		{
			Name:     "Magic Johnson",
			Position: "Point Guard",
			Team:     "Los Angeles Lakers",
			Age:      60,
			Jersey:   32,
			Retired:  true,
		},
		{
			Name:     "Larry Bird",
			Position: "Small Forward",
			Team:     "Boston Celtics",
			Age:      64,
			Jersey:   33,
			Retired:  true,
		},
		{
			Name:     "Wilt Chamberlain",
			Position: "Center",
			Team:     "Los Angeles Lakers",
			Age:      87,
			Jersey:   13,
			Retired:  true,
		},
		{
			Name:     "Bill Russell",
			Position: "Center",
			Team:     "Boston Celtics",
			Age:      86,
			Jersey:   6,
			Retired:  true,
		},
	}
	tmplFile := "playersHtml.tmpl"
	tmpl, err := template.New(tmplFile).ParseFiles("./templates/" + tmplFile)
	if err != nil {
		panic(err)
	}
	var file *os.File
	file, err = os.Create("players.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(file, players)
	if err != nil {
		panic(err)
	}

	return players[0]
}

func main() {
	BestPlayer(nil)

}
