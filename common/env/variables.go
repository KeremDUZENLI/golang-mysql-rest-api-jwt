package env

import (
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	LISTUSERS []primitive.M
)

var (
	MONGO_URL             string
	MONGO_COLLECTION_NAME string

	SECRET_KEY string

	URL string
)

func Load() {
	godotenv.Load(".env")

	MONGO_URL = os.Getenv("MONGO_URL")
	MONGO_COLLECTION_NAME = os.Getenv("MONGO_COLLECTION_NAME")

	SECRET_KEY = os.Getenv("SECRET_KEY")

	URL = os.Getenv("PORT")
}
