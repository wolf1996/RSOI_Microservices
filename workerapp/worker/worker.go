package worker


type Config struct {
	Rsconfig ResponserConfig
	Dtconfig DatabaseWorkerConfig
	Hconfig	HandlerConfig
}

type ReciverConfig struct {
	RabConf 	    RabbitConfig
	NumberOfWorkers int
}


type RabbitConfig struct {
	Addres string
	User string
	Pass string
}


func StartWorkers(config Config)(err error){
	respInput, err := startResponserWorkers(config.Rsconfig)
	if err != nil {
		return
	}
	dbIn, err := StartDatabaseWorker(config.Dtconfig, respInput)
	if err != nil {
		return
	}
	err = startHandlerWorkers(config.Hconfig, dbIn)
	if err != nil {
		return
	}
	return
}

