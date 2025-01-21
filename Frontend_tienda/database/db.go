package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func ConnectDB(mongoURI string) {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Printf("Archivo .env no encontrado")
	}

	if mongoURI == "" {
		mongoURI = os.Getenv("MONGODB_URI")
		if mongoURI == "" {
			log.Fatal("MONGODB_URI no está configurada en el ambiente")
		}
	}

	log.Printf("Conectando a MongoDB en: %s", mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// Ping a la base de datos
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error al verificar conexión con la base de datos:", err)
	}

	DB = client.Database("Tienda")
	
	// Verificar las colecciones
	collections := []string{"products", "orders", "admins"}
	for _, collName := range collections {
		collection := DB.Collection(collName)
		count, err := collection.CountDocuments(ctx, bson.M{})
		if err != nil {
			log.Printf("Error al contar documentos en %s: %v", collName, err)
		} else {
			log.Printf("Se encontraron %d documentos en la colección %s", count, collName)
		}
	}

	log.Printf("Conectado a la base de datos MongoDB: %s", DB.Name())
}