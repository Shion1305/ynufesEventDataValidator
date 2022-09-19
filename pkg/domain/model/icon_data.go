package model

import (
	"context"
	"fmt"
	"google.golang.org/api/drive/v3"
	"io"
	"log"
	"os"
	"regexp"
)

func getIconId(url string) string {
	if url == "" {
		return ""
	}
	re := regexp.MustCompile("https://drive\\.google\\.com/open\\?id=([\\w-]+)$")
	res := re.FindStringSubmatch(url)
	if res == nil {
		fmt.Println("Regex Error: ", url)
		return ""
	}
	fmt.Println(res)
	return res[1]
}

func InitGD() *drive.Service {
	ctx := context.Background()

	d, err := drive.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to retrieve Drive client: %v", err)
	}
	return d
}

func TestGD(service *drive.Service, e EventData) {
	if e.iconDataId == "" {
		return
	}
	f, err := service.Files.Get(e.iconDataId).Do()
	if err == nil {
		fmt.Println(f.MimeType)
	}
	resp, err := service.Files.Get(e.iconDataId).Download()
	if err != nil {
		fmt.Println("get drive file: %w", err)
		fmt.Println(e.iconDataId)
		return
	}
	defer resp.Body.Close()
	path := "icons/" + string(e.eventIdMD5) + ".jpeg"
	fmt.Println(path)
	output, err := os.Create(path)
	defer output.Close()

	if _, err := io.Copy(output, resp.Body); err != nil {
		fmt.Println("write file: %w", err)
		return
	}

	return
}

func checkImageSize(target image.Image) (image.Image, error) {
	bounds := target.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	if w != h {
		return nil, errors.New("画像が正方形ではありません。")
	}
	if w <= 500 {
		return target, nil
	}
	imgData := image.NewRGBA(image.Rect(0, 0, w, w))
	draw.CatmullRom.Scale(imgData, imgData.Bounds(), target, target.Bounds(), draw.Over, nil)
	return imgData, nil
}
