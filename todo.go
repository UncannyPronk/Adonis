package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	var command string
	running := true
	fmt.Println("Adonis - Your personal CLI Progress Tracking Assistant\n(type help for commands list)\n")

	current_tasks := []string{}
	completed_tasks := []string{}

	for running == true{
		fmt.Println("\n\nv")
		fmt.Scan(&command)
		fmt.Print("\n")
		switch command {
		case "showPending":
			if len(current_tasks) > 0 {
				for i, v := range current_tasks {
					fmt.Println(i+1, " - ", v)
				}
			} else {
				if len(completed_tasks) > 0 {fmt.Println("Congratulations!")}
				fmt.Println("You have no pending tasks.")
			}
		case "showCompleted":
			if len(completed_tasks) > 0 {
				for i, v := range completed_tasks {
					fmt.Println(i+1, " - ", v)
				}
			} else {
				fmt.Println("There are no completed tasks yet.")
			}
		case "add":
			// fmt.Println("New Task :-")
			newTask, _ := reader.ReadString('\n')
			newTask = strings.TrimSpace(newTask)
			// fmt.Println(newTask)
			current_tasks = append(current_tasks, newTask)
			fmt.Println("Task Added!")
		case "delete":
			var delno int
			// fmt.Println("Enter Task Number :-")
			fmt.Scan(&delno)
			current_tasks = append(current_tasks[:delno-1], current_tasks[delno:]...)
			fmt.Println("Task Deleted!")
		case "flushCompleted":
			completed_tasks = completed_tasks[:0]
			fmt.Println("Completed Tasks Flushed.")
		case "completed":
			var delno int
			// fmt.Println("Enter Task Number :-")
			fmt.Scan(&delno)
			completed_tasks = append(completed_tasks, current_tasks[delno-1])
			current_tasks = append(current_tasks[:delno-1], current_tasks[delno:]...)
			fmt.Println("Task Marked as Completed!")
		case "help":
			fmt.Print("List of commands:\nadd [task name] - Add a new task\ndelete [task index/number] - Delete a pending task\ncompleted [task index/number] - Mark a task as completed\nshowPending - Show pending tasks\nshowCompleted - Show completed tasks\nflushCompleted - Delete all completed tasks\nexit - Exit this tool")
		case "exit": 
			fmt.Println("Goodbye!  ;)")
			running = false
			break
		default:
			fmt.Println("Invalid Command - enter 'help' for more info")
		}
		fmt.Print("\n")
	}
}