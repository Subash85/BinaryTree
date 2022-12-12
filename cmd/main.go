package main

import (
	"log"
)
import "BinaryTree/internal"

func main() {
	data := []string{"delta", "bravo", "charlie", "echo", "alpha"}

	tree := &BinaryTree.Tree{}
	for i := 0; i < len(data); i++ {
		err := tree.Insert(data[i])
		if err != nil {
			log.Fatal("Error inserting value '", data[i], "': ", err)
		}
	}
}
