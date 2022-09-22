package model

import (
	"context"
	"fmt"
	"github.com/nickalie/go-webpbin"
	"golang.org/x/image/draw"
	"google.golang.org/api/drive/v3"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
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

func ProcessGD(service *drive.Service, e *EventData) {
	path := download(service, *e)
	if path == "" {
		return
	}

	filename := e.eventOrgName
	filename = strings.Replace(filename, "<", "-", -1)
	filename = strings.Replace(filename, ">", "-", -1)
	filename = strings.Replace(filename, " ", "_", -1)
	filename = strings.Replace(filename, "?", "_", -1)
	filename = strings.Replace(filename, "\"", "-", -1)
	filename = strings.Replace(filename, "*", "-", -1)
	filename = strings.Replace(filename, "|", "-", -1)
	filename = strings.Replace(filename, "/", "-", -1)
	filename = strings.Replace(filename, ":", "_", -1)
	path0 := "icons/" + filename + "." + "png"
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
	img, err = checkImageSize(img)
	if err != nil {
		if img == nil {
			fmt.Printf("CRITICAL: error checking size, %s\n", err)
			return
		}
		e.ImgStatus = err.Error()
	}

	output0, err := os.Create(path0)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer output0.Close()

	output1, err := os.Create(path1)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer output1.Close()

	if err := png.Encode(output0, img); err != nil {
		fmt.Printf("writing png: %s\n", err)
		return
	}
	if err := webpbin.Encode(output1, img); err != nil {
		fmt.Printf("writing webp: %s\n", err)
		return
	}
	if e.ImgStatus == "" {
		e.ImgStatus = "正常です。"
	}
	return
}

func download(service *drive.Service, e EventData) string {
	if e.iconDataId == "" {
		return ""
	}
	f, err := service.Files.Get(e.iconDataId).Do()
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
	case "image/heif":
		extension = "heif"
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
	path := "icons-original/" + string(e.eventIdMD5) + "." + extension
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
	bounds := target.Bounds()
	w := bounds.Dx()
	h := bounds.Dy()
	if w == h && w <= 500 {
		return target, nil
	}
	if w != h {
		err = fmt.Errorf("画像が正方形でなかったため画像サイズの変更を行いました。(%d,%d)->(500,500)", w, h)
		fmt.Println(err)
	}
	imgData := image.NewRGBA(image.Rect(0, 0, 500, 500))
	draw.CatmullRom.Scale(imgData, imgData.Bounds(), target, target.Bounds(), draw.Over, nil)
	return imgData, err
}
