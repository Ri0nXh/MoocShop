package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

func initConfig() (err error) {
	files, err := ioutil.ReadDir(CONFIG_PATH)
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
	return err
}
