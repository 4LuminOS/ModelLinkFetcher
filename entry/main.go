package entry

import (
	"context"
	"github.com/4LuminOS/ModelLinkFetcher/api"
	"fmt"
	"os"
)
func GetLink(modelName string) {
	fmt.Println("Fetching direct download link for model:", modelName)

	parsedModelPath := api.ParseModelPath(modelName)

	manifest, manifestLink, err := api.GetManifest(context.TODO(), parsedModelPath)
	if err != nil {
		fmt.Println("Error getting manifest:", err)
		os.Exit(1)
	}

	var layers []*api.Layer
	layers = append(layers, manifest.Layers...)
	layers = append(layers, manifest.Config)

	var downloadLinks []string

	for _, layer := range layers {
		config := api.DownloadLinkConfig{
			ModelPath: parsedModelPath,
			Digest:    layer.Digest,
		}
		link := config.GetDownloadLink()
		downloadLinks = append(downloadLinks, link)
	}

	fmt.Printf("Manifest download link: %s\n", manifestLink)
	fmt.Println("Download links for layers:")
	for i, link := range downloadLinks {
		fmt.Printf("%d - %s\n", i+1, link)
	}
	fmt.Println("Generated download links for model:", modelName)
	fmt.Println("Finished successfully!")
	os.Exit(0)
}

func Install(modelName string, blobsPath string, ) {
	if modelName == "" || blobsPath == "" {
		fmt.Println("model name and blobs path are required")
		os.Exit(1)
	}

	hasPermission, err := api.HasElevatedPermissions()
	if err != nil {
		fmt.Println("error checking permissions: ", err)
		os.Exit(1)
	}
	if !hasPermission {
		fmt.Println("please run the command with elevated permissions")
		os.Exit(1)
	}

	fmt.Println("installing model :", modelName)

	err = api.VerifyDownloadedModel(modelName, blobsPath)
	if err != nil {
		fmt.Println("error verifying model: ", err)
		os.Exit(1)
	}

	err = api.InstallModel(modelName, blobsPath)
	if err != nil {
		fmt.Println("error installing model: ", err)
		os.Exit(1)
	}

}
