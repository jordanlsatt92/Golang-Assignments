package main

/*
Create a map that includes at least ten key:value pairs.
The pairs should make sense, but you can choose whatever you wish.
Suggestions include produce categories (e.g., apple:fruit and onion:vegetable), animal categories
(e.g., mammals, birds, fish), city populations, or word definitions.
Prompt the user to input a search term.
Display all key:value pairs that include the input search term, either in the key or in the value.
For example, if the topic is produce, the user can enter either apple or fruit and see all matching
map entries.
If the value does not exist in the dictionary, display a user-friendly error message.
Prompt the user to start over again or exit the program.
*/

import (
	"fmt"
	"strings"
)

func main(){
	m:=make(map[string]string)
	m["apple"] = "fruit"
	m["orange"] = "fruit"
	m["grape"] = "fruit"
	m["clementine"] = "fruit"
	m["mango"] = "fruit"
	m["onion"] = "vegetable"
	m["carrots"] = "vegetable"
	m["potato"] = "vegetable"
	
	var input string
	for {
		fmt.Print("Please enter the name of a fruit or vegetable or input \"fruit\" or \"vegetable\" (type \"exit\" to end the program):")
		fmt.Scan(&input)
		if strings.ToLower(input) == "exit"{
			break
		}
		var printed bool = false
		for i, j := range m{
			if i == input || j == input{
				fmt.Println("Key:",i,", Value:", j)
				printed = true
			}
		}
		if !printed{
			fmt.Printf("\nSorry, \"%s\" does not exist in the dictionary. Please try again\n\n", input)
		}
	}
}