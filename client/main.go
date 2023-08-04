package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	pb "todo-list-grpc/todo"
)

func getDataAllTodo(client pb.DataTodoClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	fmt.Println("fetch getDataAllTodo")

	defer cancel()

	todos, err := client.FindAllTodo(ctx, &pb.Todo{})
	if err != nil {
		log.Fatalln("error when getDataAllTodo :: ", err.Error())
	}

	fmt.Println(todos)
}

func getDataTodoById(client pb.DataTodoClient, id int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	fmt.Println("fetch getDataTodoById :: ", string(id))

	defer cancel()

	todoSearch := pb.Todo{Id: id}
	todo, err := client.FindTodoById(ctx, &todoSearch)
	if err != nil {
		log.Fatalln("error when getTodoByTitle :: ", err.Error())
	}

	fmt.Println(todo)
}

func getDataTodoByTitle(client pb.DataTodoClient, title string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	fmt.Println("fetch getDataTodoByTitle :: ", string(title))

	defer cancel()

	todoSearch := pb.Todo{Title: title}
	todo, err := client.FindTodoByTitle(ctx, &todoSearch)
	if err != nil {
		log.Fatalln("error when getTodoByTitle :: ", err.Error())
	}

	fmt.Println(todo)
}

func updateDataTodoById(client pb.DataTodoClient, todo *pb.Todo) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	fmt.Println("fetch updateDataTodoById")

	defer cancel()

	updatedTodo, err := client.UpdateTodoById(ctx, todo)
	if err != nil {
		log.Fatalln("error when updateDataTodoById :: ", err.Error())
	}

	fmt.Println(updatedTodo)
}

func deleteDataTodoById(client pb.DataTodoClient, id int32) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	fmt.Println("fetch deleteDataTodoById")

	defer cancel()

	removeTodo := pb.Todo{Id: id}
	_, err := client.DeleteTodoById(ctx, &removeTodo)
	if err != nil {
		log.Fatalln("error when deleteDataTodoById :: ", err.Error())
	}

	fmt.Println("Deleted ", id, " is successfully")
}

func main() {
	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":1200", opts...)
	if err != nil {
		log.Fatalln("error in dial :: ", err.Error())
	}

	defer conn.Close()

	client := pb.NewDataTodoClient(conn)

	// Get All
	getDataAllTodo(client)

	// Get by id
	getDataTodoById(client, 1)

	// Get by title
	getDataTodoByTitle(client, "Todo 2")

	// Update by id
	dataUpdate := pb.Todo{
		Id:          3,
		Title:       "Todo 3 updated",
		Description: "Description 3 updated",
		IsDone:      true,
	}
	updateDataTodoById(client, &dataUpdate)

	// Delete by id
	deleteDataTodoById(client, 2)
}
