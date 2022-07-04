package controllers

import (
	"crud/services"
	"crud/utils"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

type FileC interface {
	Upload(c *fiber.Ctx) error
	Download(c *fiber.Ctx) error
}

type fileC struct {
	firebase services.Storage
}

func NewFileC(firebase services.Storage) FileC {
	return &fileC{
		firebase: firebase,
	}
}

func (f *fileC) Upload(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err != nil {
		return utils.Response(c, 400, nil, err.Error(), false)
	}

	files := form.File["image"]
	timeNow := strconv.Itoa(int(time.Now().UnixMilli()))
	for _, file := range files {
		img := fmt.Sprintf("./temp/%s%s", timeNow, file.Filename)
		if err := c.SaveFile(file, img); err != nil {
			log.Println(err)
		}

		f.firebase.UploadToStorage(c.Context(), img)
	}

	return utils.Response(c, 200, nil, "Image uploaded!", true)
}

func (f *fileC) Download(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	if err == nil {
		files := form.File["image"]
		timeNow := strconv.Itoa(int(time.Now().UnixMilli()))
		for _, file := range files {
			if err := c.SaveFile(file, fmt.Sprintf("./temp/%s%s", timeNow, file.Filename)); err != nil {
				return err
			}
		}
	}

	return err
}
