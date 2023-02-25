package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

func InitConfig() {
	files, err := ioutil.ReadDir(CONFIG_PATH)
	if err != nil {
		fmt.Println("config dir load error: ", err)
	}
	viper.SetConfigType("yaml")
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, err := os.Open(fmt.Sprintf("%s/%s", CONFIG_PATH, file.Name()))
		if err != nil {
			fmt.Printf("open %s error", file)
		}
		err = viper.MergeConfig(f)
		if err != nil {
			fmt.Println("viper merge config error: ", err)
		}
	}
}
