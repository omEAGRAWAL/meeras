package config

import (
	"log"

	"github.com/cloudinary/cloudinary-go"
)

var Cld *cloudinary.Cloudinary

func InitCloudinary() *cloudinary.Cloudinary {
	Cld, err := cloudinary.NewFromURL("cloudinary://856923756819212:6W83yMejCehNPuoLN4xwzeVf1Ss@dk4rxt4x1")
	if err != nil {
		log.Fatalf("Cloudinary init failed: %s", err)
	}
	return Cld
}
