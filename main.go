package main

import "log"

func taskInTasks(task string, tasks []string) bool {
	found := false
	for _, t := range task {
		if t == task {
			found = true
		}
	}
	return found
}

func main() {

	tasks := []string{"Parler des Tableaux", "Donner un cours sur go"}

	for _, task := range tasks {

		log.Print(task)
	}

	log.Printf("notre slice : %#v", tasks)

	task := "Donner un cours sur go"
	log.Printf("Est ce que la task %v est presente : %v", task, taskInTasks(task, tasks))
}
