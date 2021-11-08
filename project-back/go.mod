module main

go 1.17

require (
	connection v0.0.0-00010101000000-000000000000
	customhttp v0.0.0-00010101000000-000000000000
	dataacces v0.0.0-00010101000000-000000000000 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/mattn/go-sqlite3 v2.0.3+incompatible
	github.com/rs/cors v1.8.0
)

replace (
	connection => ./connection
	customhttp => ./customhttp
	dataacces => ./dataacces
)
