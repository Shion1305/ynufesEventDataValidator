package model

import (
	"context"
	"fmt"
	"github.com/nickalie/go-webpbin"
	"golang.org/x/image/draw"
	"google.golang.org/api/drive/v3"
	"image"
	_ "image/jpeg"
	_ "image/png"
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
	path := download(service, e)
	if path == "" {
		return
	}

	path1 := "webps/" + string(e.eventIdMD5) + "." + "webp"
	source, err := os.Open(path)
	if err != nil {
		fmt.Printf("failed to reopen source: %s\n", err)
		fmt.Println(path)
		return
	}
	defer source.Close()

	img, _, err := image.Decode(source)
	if err != nil {
		fmt.Printf("failed to decode image: %s\n", err)
		return
	}
	if img, err = checkImageSize(img); err != nil {
		fmt.Printf("image requirement not met: %s\n", err)
		return
	}
	output1, err := os.Create(path1)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer output1.Close()

	if err := webpbin.Encode(output1, img); err != nil {
		fmt.Printf("writing webp: %s\n", err)
		return
	}
	return
}

func download(service *drive.Service, e EventData) string {
	if e.iconDataId == "" {
		return ""
	}
	f, err := service.Files.Get(e.iconDataId).Do()
	if err == nil {
		fmt.Println(f.MimeType)
	}
	var extension string
	switch f.MimeType {
	case "image/png":
		extension = "png"
		break
	case "image/jpeg":
		extension = "jpeg"
		break
	case "image/heic":
		extension = "heic"
		break
	default:
		fmt.Printf("MIME ERROR: %s\n", f.MimeType)
		return ""
	}
	resp, err := service.Files.Get(e.iconDataId).Download()
	if err != nil {
		fmt.Println("get drive file: %w", err)
		fmt.Println(e.iconDataId)
		return ""
	}
	defer resp.Body.Close()
	path := "icons/" + string(e.eventIdMD5) + "." + extension
	output, err := os.Create(path)
	defer output.Close()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return ""
	}

	if _, err := io.Copy(output, resp.Body); err != nil {
		fmt.Println("copying downloaded data: %w", err)
		return ""
	}
	return path
}

func checkImageSize(target image.Image) (image.Image, error) {
	var err error
	resizeTo := 500
	bounds := target.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	if w == h && w <= 500 {
		return target, nil
	}
	if w != h {
		err = fmt.Errorf("画像が正方形ではありません。(%d,%d)", w, h)
		fmt.Println(err)
		if w > h {
			resizeTo = w
		} else {
			resizeTo = h
		}
	}
	imgData := image.NewRGBA(image.Rect(0, 0, resizeTo, resizeTo))
	draw.CatmullRom.Scale(imgData, imgData.Bounds(), target, target.Bounds(), draw.Over, nil)
	return imgData, err
}
