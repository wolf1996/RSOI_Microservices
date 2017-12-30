package workerapp

type RabbitConfig struct {
	Addres string
	User string
	Pass string
}

type Config struct {
	Rabbit       RabbitConfig
}

func applyConfig(config Config)  {
	
}

func StartApplication(config Config){
	applyConfig(config)
}
