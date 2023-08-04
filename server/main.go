package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	pb "todo-list-grpc/todo"

	"google.golang.org/grpc"
)

type dataTodoServer struct {
	pb.UnimplementedDataTodoServer
	todos []*pb.Todo
}

func (d *dataTodoServer) FindAllTodo(ctx context.Context, todo *pb.Todo) (*pb.Todos, error) {
	fmt.Println("fetch FindAllTodo server")

	datas := pb.Todos{Todos: d.todos}

	return &datas, nil
}

func (d *dataTodoServer) FindTodoById(ctx context.Context, todo *pb.Todo) (*pb.Todo, error) {
	fmt.Println("fetch FindTodoById server")

	for _, v := range d.todos {
		if v.Id == todo.Id {
			fmt.Println(todo.Id, " is Found")
			return v, nil
		}
	}

	return nil, errors.New("Not Found")
}

func (d *dataTodoServer) FindTodoByTitle(ctx context.Context, todo *pb.Todo) (*pb.Todo, error) {
	fmt.Println("fetch FindTodoByTitle server")

	for _, v := range d.todos {
		if v.Title == todo.Title {
			fmt.Println(todo.Title, " is Found")
			return v, nil
		}
	}

	return nil, errors.New("Not Found")
}

func (d *dataTodoServer) UpdateTodoById(ctx context.Context, todo *pb.Todo) (*pb.Todo, error) {
	fmt.Println("fetch UpdateTodoById server")

	existsTodo, err := d.FindTodoById(ctx, todo)
	if err != nil {
		return nil, err
	}

	// Update
	existsTodo.Title = todo.Title
	existsTodo.Description = todo.Description
	existsTodo.IsDone = todo.IsDone

	return existsTodo, nil
}

func (d *dataTodoServer) DeleteTodoById(ctx context.Context, todo *pb.Todo) (*pb.Todo, error) {
	fmt.Println("fetch DeleteTodoById server")

	existsTodo, err := d.FindTodoById(ctx, todo)
	if err != nil {
		return nil, err
	}

	// Delete
	var index int
	for i, v := range d.todos {
		if v.Id == existsTodo.Id {
			index = i
		}
	}

	d.todos = append(d.todos[:index], d.todos[index+1:]...)

	fmt.Println("Exists todos ", d.todos)

	return &pb.Todo{}, nil
}

func (d *dataTodoServer) loadData() {
	data, err := ioutil.ReadFile("./json/todos.json")
	if err != nil {
		log.Fatalln("error load file json :: ", err.Error())
	}

	if err := json.Unmarshal(data, &d.todos); err != nil {
		log.Fatalln("error unmarshal data :: ", err.Error())
	}
}

func newServer() *dataTodoServer {
	s := dataTodoServer{}
	s.loadData()
	return &s
}

func main() {
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("error listen :: ", err.Error())
	}

	grpcServer := grpc.NewServer()
	// newServer := newServer().UnimplementedDataTodoServer
	pb.RegisterDataTodoServer(grpcServer, newServer())

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalln("error run grpc serve :: ", err.Error())
	}

	fmt.Println("Server is running . . . ")
}
