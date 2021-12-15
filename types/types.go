package main

type userID int

func main() {
	idx := 1
	var uid userID = 42

	// Даже если тип одинаковый, разные типы несовместимы
	// cannot use uid (type UserID) as type int64 in assignment
	// myID := idx

	myID := userID(idx)
	println(uid, myID)
}
