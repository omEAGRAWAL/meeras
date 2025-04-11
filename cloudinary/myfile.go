package cloudinary

import (
	"context"
	"log"
	"net/http"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

var (
	cld *cloudinary.Cloudinary
	ctx context.Context
)

func init() {
	var err error
	cld, err = cloudinary.NewFromURL("cloudinary://896883517971133:eQTok1iSw5eRW4f48vemQO89AhU@drzgqp3pc")
	if err != nil {
		log.Fatalf("❌ Cloudinary init error: %v", err)
	}
	cld.Config.URL.Secure = true
	ctx = context.Background()
}

func UploadHandler(c *gin.Context) {
	// Get uploaded file
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image not provided"})
		return
	}

	// Open file stream
	src, err := file.Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
		return
	}
	defer src.Close()

	// Upload to Cloudinary
	resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{
		PublicID: file.Filename,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Cloudinary upload failed"})
		return
	}

	// Return Cloudinary URL
	c.JSON(http.StatusOK, gin.H{"url": resp.SecureURL})
}
