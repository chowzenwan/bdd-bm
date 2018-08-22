package bookingmotor

import (
	"mark/support"
	"time"

	"github.com/DATA-DOG/godog"
	"github.com/tebeka/selenium"
)

var Driver selenium.WebDriver

func queSeVaALaPaginaPrincipalDe(url string) (err error) {

	Driver.Get("https://" + url)
	return nil
}

func enLaPestaaDeHotelesSeSeleccionaLaCiudadDe(city string) (err error) {
	hotelTab, err := Driver.FindElement(selenium.ByCSSSelector, "i.icon-bm-hotel")
	if err != nil {
		return
	}
	hotelTab.Click()
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	destinyInput, err := Driver.FindElement(selenium.ByID, "destiny-hotel")
	if err != nil {
		return
	}
	destinyInput.SendKeys(city)
	// fmt.Println("debe decir Paris, Francia y dice : ", city)
	Driver.SetImplicitWaitTimeout(time.Second * 10)
	citySelected, err := Driver.FindElement(selenium.ByCSSSelector, "li.ui-menu-item:first-child")
	if err != nil {
		return
	}
	citySelected.Click()
	return nil
}

func laFechaDelCheckinEsPorNoche(arg1, arg2 string) error {
	return godog.ErrPending
}

func seDaClickEnElBotonDeBuscar() error {
	return godog.ErrPending
}

func deberemosVerUnListadoDeHoteles() error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^que se va a la pagina principal de "([^"]*)"$`, queSeVaALaPaginaPrincipalDe)
	s.Step(`^en la pestaña de Hoteles, se selecciona la ciudad de "([^"]*)"$`, enLaPestaaDeHotelesSeSeleccionaLaCiudadDe)
	s.Step(`^la fecha del checkin es "([^"]*)" por "([^"]*)" noche$`, laFechaDelCheckinEsPorNoche)
	s.Step(`^se da click en el boton de buscar$`, seDaClickEnElBotonDeBuscar)
	s.Step(`^deberemos ver un listado de hoteles$`, deberemosVerUnListadoDeHoteles)

	s.BeforeScenario(func(interface{}) {
		Driver = support.WDInit()
	})
}
