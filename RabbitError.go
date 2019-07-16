package main

type RabbitError string

const(
	SERVER_ERROR RabbitError = "Unable to conntect to RabbitMQ server",

)