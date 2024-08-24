package main

import "fmt"

type User struct {
	Id   int64
	Name string
}

func main() {
	var defaultMap map[int64]string

	fmt.Printf("Type: %T Value: %#v \n", defaultMap, defaultMap)
	fmt.Printf("Length: %d \n\n", len(defaultMap))

	mapByMake := make(map[string]string)
	fmt.Printf("Type: %T Value: %#v \n", mapByMake, mapByMake)

	mapByMakeWithCap := make(map[string]string, 3)
	fmt.Printf("Type: %T Value: %#v \n\n", mapByMakeWithCap, mapByMakeWithCap)

	mapByLiteral := map[string]int{"Vasya": 18, "Fima": 14}
	fmt.Printf("Type: %T Value: %#v \n", mapByMakeWithCap, mapByMakeWithCap)
	fmt.Printf("Length: %d \n\n", len(mapByLiteral))

	mapWithNew := *new(map[string]string)
	fmt.Printf("Type: %T Value: %#v \n", mapWithNew, mapWithNew)

	// insert
	mapByMake["Vasya"] = "Test"
	fmt.Printf("Type: %T Value: %#v \n", mapByMake, mapByMake)
	fmt.Printf("Length: %d \n\n", len(mapByMake))

	// update
	mapByMake["Vasya"] = "Test2"
	fmt.Printf("Type: %T Value: %#v \n", mapByMake, mapByMake)
	fmt.Printf("Length: %d \n\n", len(mapByMake))

	// get
	fmt.Println(mapByLiteral["Vasya"])

	// get map default value
	fmt.Println(mapByLiteral["Fima"])

	// check value existence
	value, ok := mapByLiteral["Fima"]
	fmt.Printf("Value: %d IsExist: %t \n\n", value, ok)

	// delete value
	delete(mapByMake, "Vasya")
	fmt.Printf("Type: %T Value: %#v \n\n", mapByMake, mapByMake)

	mapForIteration := map[string]int{"Vasya": 18, "Fima": 14}
	for key, value := range mapForIteration {
		fmt.Printf("Key: %#v Value: %#v \n\n", key, value)
	}

	users := []User{
		{
			Id:   1,
			Name: "Tom",
		},
		{
			Id:   2,
			Name: "John",
		},
		{
			Id:   3,
			Name: "Test",
		},
		{
			Id:   2,
			Name: "John",
		},
	}

	uniqueUsers := make(map[int64]struct{}, len(users))

	for _, user := range users {
		if _, ok := uniqueUsers[user.Id]; !ok {
			uniqueUsers[user.Id] = struct{}{}
		}
	}
	fmt.Printf("Type: %T Value: %#v \n\n", uniqueUsers, uniqueUsers)

	fmt.Println(findInSlice(2, users))

	usersMap := make(map[int64]User, len(users))
	for _, user := range users {
		if _, ok := usersMap[user.Id]; !ok {
			usersMap[user.Id] = user
		}
	}
	fmt.Println(findInMap(2, usersMap))
}

func findInSlice(id int64, users []User) *User {
	for _, user := range users {
		if user.Id == id {
			return &user
		}
	}

	return nil
}

func findInMap(id int64, usersMap map[int64]User) *User {
	if user, ok := usersMap[id]; ok {
		return &user
	}

	return nil
}
