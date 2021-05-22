package main

import "cmd/main.go/internal/consumer"

func main() {
	//config.NewConfig()
	//fmt.Println("Access Key:", config.GetConfiguration().AwsConfig.AccessKey)
	//fmt.Println("Secret Key:", config.GetConfiguration().AwsConfig.SecretKey)
	//fmt.Println("Region:", config.GetConfiguration().AwsConfig.Region)

	consumer.PoolMessages()
}
