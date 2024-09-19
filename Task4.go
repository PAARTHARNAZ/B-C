package main

import (
	"crypto/sha256"
	"fmt"
	"strings"
)

// Student struct holds the student's name and the subjects they are studying.
type Student struct {
	Name     string
	Subjects []string
}

// CalculateHash takes a string input and returns its SHA256 hash as a hexadecimal string.
func CalculateHash(data string) string {
	sum := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", sum)
}

func main() {
	// Create a list of students with their subjects.
	students := []Student{
		{Name: "Alice", Subjects: []string{"Math", "Science", "History"}},
		{Name: "Bob", Subjects: []string{"English", "Biology", "Physics"}},
		{Name: "Charlie", Subjects: []string{"Chemistry", "Geography", "Art"}},
	}

	// Iterate over each student and calculate the hash for their data.
	for _, student := range students {
		// Concatenate the student's name and subjects into a single string.
		dataToHash := student.Name + ": " + strings.Join(student.Subjects, ", ")
		// Calculate the hash of the concatenated data.
		hash := CalculateHash(dataToHash)
		// Print the student data and its hash.
		fmt.Printf("Student: %s\nSubjects: %s\nHash: %s\n\n", student.Name, strings.Join(student.Subjects, ", "), hash)
	}
}
