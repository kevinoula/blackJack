module github.com/kevinoula/blackjack

go 1.17

require (
	cards v0.0.1
	player v0.0.1
)

replace cards v0.0.1 => ./cards

replace player v0.0.1 => ./player
