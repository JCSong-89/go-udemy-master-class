package utill

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DBdDiver   string `mapstructure:"DB_DRIVER"`
	DBDNS      string `mapstructure:"DB_DNS"`
	ServerHost string `mapstructure:"SERVER_HOST"`
}

func LoadConfig(path string) (config Config, err error) {
	fmt.Println("시작 컨피그 파일 리딩")

	viper.AddConfigPath(path)  // config 파일이 있는 경로를 지정
	viper.SetConfigName("app") // 읽어들일 파일의 이름을 지정
	viper.SetConfigType("env") // 일어들일 파일의 형식을 지정

	viper.AutomaticEnv() // 환경변수를 자동으로 읽어들이도록 설정

	err = viper.ReadInConfig() // 설정파일을 읽어들임
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config) // 읽어들인 설정파일을 Config 구조체에 매핑
	fmt.Println(config)
	return
}
