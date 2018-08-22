package support

import (
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

var driver selenium.WebDriver

//WDInit retorna uma instancia do Webdriver
func WDInit() selenium.WebDriver {
	var err error
	caps := selenium.Capabilities(map[string]interface{}{"browserName": "chrome"})
	driver, err = selenium.NewRemote(caps, "")

	if err != nil {
		fmt.Println("Erro ao instanciar o driver", err.Error())
	}
	driver.SetImplicitWaitTimeout(time.Second * 3)
	driver.ResizeWindow("note", 1280, 800)

	return driver
}

//SaveImage print of webdriver in bytes and save a png into the proyect
/*
	func SaveImage(foto []byte, name string){

		img, _, _ := image.Decode(bytes.NewReader(foto))

		out, err := os.Create("./log/screenshot/" + name + ".png")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		err = png.Encode(out, img)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}*/
