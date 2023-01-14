package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

var result string

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Println("we are in")
	rContent, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	json.Unmarshal(rContent, result)
	fmt.Println(result)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	// w.Write([]byte(resp))
	todo := Todo{
		ID:    1,
		Title: "First JSON",
		Done:  true,
		Body:  "You did it!!!!",
	}
	respo, err := json.Marshal(todo)
	if err != nil {
		log.Println("Couldn't marshal JSON")
		return
	}
	w.Write(respo)
}

func main() {
	// app := fiber.New()

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: "http://localhost:3000",
	// 	AllowHeaders: "Origin, Content-Type, Accept",
	// }))

	// todos := []Todo{}
	http.HandleFunc("/healthcheck", HealthCheck)
	// app.Post("/healthcheck", func(c *fiber.Ctx) error {
	// 	c.
	// 	temp := string(c.Body())
	// 	fmt.Println(temp + ": Data from our Front")

	// 	fmt.Println("Got it")
	// 	return c.JSON(todo)
	// })

	// app.Post("/api/todos", func(c *fiber.Ctx) error {
	// 	todo := &Todo{}

	// 	if err := c.BodyParser(todo); err != nil {
	// 		return err
	// 	}

	// 	todo.ID = len(todos) + 1

	// 	todos = append(todos, *todo)

	// 	return c.JSON(todos)
	// })

	// app.Patch("/api/todos/:id/done", func(c *fiber.Ctx) error {
	// 	id, err := c.ParamsInt("id")
	// 	if err != nil {
	// 		return c.Status(401).SendString("Invalid id")
	// 	}

	// 	for i, t := range todos {
	// 		if t.ID == id {
	// 			todos[i].Done = true
	// 			break
	// 		}
	// 	}

	// 	return c.JSON(todos)
	// })

	// app.Get("/api/todos", func(c *fiber.Ctx) error {
	// 	return c.JSON(todos)
	// })
	log.Println("Starting a server at http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
	// log.Fatal(app.Listen(":4000"))
}
