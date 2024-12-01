Задание
Реализовать сервисную архитектуру (три четыре сервиса, которые общаются между собой)
Реализовать общение между сервисами с использованием Rabbit MQ 
Реализовать паттерн Inbox или Outbox
Реализовать паттерн  “хореография”

UserService — создаёт заказ и публикует событие о создании заказа в очередь "созданные заказы" RabbitMQ. слушает очередь  "доставка" и получает уведомление о доставке заказ
OrderService — слушает очередь событий от OrderService "заказы", принимает событие о создании заказа, оформляет заказ и публикует событие о создании заказа в очередь RabbitMQ "заказы".
DeliveryService — слушает очередь событий от OrderService "заказы", принимает событие о создании заказа и отправляет уведомление в очередь RabbitMQ "доставка".
 
 



https://go.dev/dl/


go mod init main

go get https://github.com/rabbitmq/amqp091-go

docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:management


go run user-service/user_service.go
go run order-service/order_service.go
go run notification-service/notification_service.go