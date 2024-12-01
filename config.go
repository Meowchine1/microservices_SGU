package config

const{

	RabbitMQURL      = "amqp://guest:guest@localhost:5672/"
	order_create_queue_name = "order_create__queue"
	order_queue_name = "order_queue"
	delivery_queue_name = "delivery_queue"
}

type OutboxEvent struct {
	ID        int
	Aggregate string
	EventType string
	Payload   string
	CreatedAt time.Time
}

type InboxEvent struct {
	ID        int
	Aggregate string
	EventType string
	Payload   string
	CreatedAt time.Time
}

func insertOutboxEvent(db *sql.DB, event OutboxEvent) error {
	query := `INSERT INTO outbox (aggregate, event_type, payload, created_at) 
				VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(query, event.Aggregate, event.EventType, event.Payload, event.CreatedAt).Scan(&event.ID)
	return err
}

func insertInboxEvent(db *sql.DB, event OutboxEvent) error {
	query := `INSERT INTO inbox (aggregate, event_type, payload, created_at) 
				VALUES ($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(query, event.Aggregate, event.EventType, event.Payload, event.CreatedAt).Scan(&event.ID)
	return err
}


func declareQueue(ch *amqp.Channel, queueName string) (amqp.Queue, error) {
	// Объявление очереди в RabbitMQ
	q, err := ch.QueueDeclare(
		queueName, // имя очереди
		false,     // durable (не будет сохраняться после перезапуска RabbitMQ)
		false,     // delete when unused
		false,     // exclusive (доступна только для этого соединения)
		false,     // no-wait
		nil,       // дополнительные аргументы
	)
	if err != nil {
		log.Printf("Failed to declare queue %s: %s", queueName, err)
		return q, err
	}
	log.Printf("Queue %s declared", queueName)
	return q, nil
}