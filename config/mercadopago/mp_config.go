package mercadopago

import (
	"log"
	"errors"
	"os"
	
	"github.com/mercadopago/sdk-go/pkg/config"
	
	"github.com/joho/godotenv"

	"chopipay/internal/models/dto/mp"
)

func Initialize() (*mp.MPConfig, error){
	log.Println("Initializing MercadoPago...")

	err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }

    // Obtener variables de entorno
    access_token := os.Getenv("ACCESS_TOKEN")
	if access_token == "" {
		log.Fatalf("Error al obtener ACCESS_TOKEN")
	}
	
	cfg, err := config.New(access_token)
	if err != nil {
		log.Println("Error initializing MercadoPago")
		return nil, errors.New("error initializing MercadoPago: Cause: " + err.Error())
	}

	return &mp.MPConfig{
		Cfg: cfg,
	}, nil
}
