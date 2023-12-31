package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"github.com/lanetb/CMDTodo"
)

const (
	todoFile = "C:/Users/tommy/Desktop/CMDTodo/.todo.json"
)

func main(){

	add := flag.Bool("a", false, "Add a new task to the list")
	complete := flag.Int("c", 0, "Complete the task with the provided index")
	del := flag.Int("d", 0, "Delete the task with the provided index")
	list := flag.Bool("l", false, "List all tasks")
	delCompleted := flag.Bool("dc", false, "Delete all completed tasks")
	clear := flag.Bool("clear", false, "Clear all tasks")

	flag.Parse()

	todos := &todo.Todos{}
	
	if err := todos.Load(todoFile); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	switch{
		case *add:
			task, err := getInput(os.Stdin, flag.Args()...)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			todos.Add(task)

			err = todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		case *complete > 0:
			err := todos.Complete(*complete)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			err = todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		case *del > 0:
			err := todos.Delete(*del)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}

			err = todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		case *list:
			todos.Print()
		case *delCompleted:
			todos.DeleteCompleted()

			err := todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		case *clear:
			todos.Clear()

			err := todos.Store(todoFile)
			if err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
				os.Exit(1)
			}
		default:
			fmt.Fprintln(os.Stdout, "Invalid command")
			os.Exit(1)
	}
}



func getInput(r io.Reader, args ...string) (string, error){
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()

	if len(scanner.Text()) == 0{
		return "", errors.New("empty todo is not allowed")
	}

	return text, nil
}