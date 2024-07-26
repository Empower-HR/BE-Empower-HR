package cloudinary

import (
	"context"
	"io"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadCloudinary(file io.Reader, filename string) (string, error){

	keyCld := os.Getenv("CLOUDINARY_URL");

	cld, err := cloudinary.NewFromURL(keyCld);

	if err != nil {
		return "", err
	};


	// upload file to cloudinary
	uploadParams := uploader.UploadParams{
		Folder: "images_empowerHR",
		PublicID: filename,
	};

	uploadResult, err := cld.Upload.Upload(context.Background(), file, uploadParams);

	if err != nil {
		return "", err
	};

	publicURL := uploadResult.URL

	return publicURL, nil;
}