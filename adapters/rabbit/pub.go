package rabbit

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fmoral2/mux-api/adapters/repository"
	application "github.com/fmoral2/mux-api/application/employee"
	"github.com/fmoral2/mux-api/application/model"

	"github.com/streadway/amqp"
)

var (
	emp model.Employee
)

func Publish(app *application.App) {
	// connect rabbit via amqp
	rabbitmqHost := os.Getenv("RABBITMQ_HOST")
	host := fmt.Sprintf("amqp://guest:guest@%s/", rabbitmqHost)
	conn, err := amqp.Dial(host)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	emps, err := app.GetEmployees(emp)
	if err != nil {
		log.Fatalf("Failed to get employees: %v", err)
	}

	sendEmps, err := json.Marshal(emps)
	if err != nil {
		log.Fatalf("Failed to marshal employees: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open channel: %v", err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"TestQueue",
		true,
		false,
		false,
		false,
		nil,
	)
	fmt.Println(q)

	if err != nil {
		log.Fatalf("Failed to declare queue: %v", err)
	}

	err = ch.Publish(
		"",
		"TestQueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        sendEmps,
		},
	)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	fmt.Println("Message published")

}

func MakeAppRb() {
	db := repository.CreateConnection()
	repository := repository.MakeRepository(db)

	app := application.MakeApplication(repository)
	Publish(app)
}
