package wimg

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/chai2010/webp"
	"github.com/ruanlianjun/graphics"
	"github.com/smallnest/rpcx/log"
	"golang.org/x/image/draw"
	"image"
	"image/color"
	"image/jpeg"
	"io"
)

func (u *util) WebpEncoder(quality ...float32) *util {
	contentType := u.imgType

	if contentType != "image/jpeg" && contentType != "image/bmp" &&
		contentType != "image/png" && contentType != "image/gif" {
		log.Fatalf("webp conv to image type not supper")
	}
	var buf bytes.Buffer
	var q float32 = 40
	if len(quality) > 0 {
		q = quality[0]
	}
	if err := webp.Encode(&buf, u.img, &webp.Options{
		Lossless: false,
		Quality:  q,
		Exact:    false,
	}); err != nil {
		return nil
	}
	u.imgType = "image/webp"
	u.saveBuf = buf
	return u
}

func (u *util) CompressJpeg(quality ...int) *util {
	if u.img == nil {
		log.Fatalf("compress jpeg u.img err:%#v", u.img)
	}

	var buf bytes.Buffer
	q := 40
	if len(quality) > 0 {
		q = quality[0]
	}
	if err := jpeg.Encode(&buf, u.img, &jpeg.Options{Quality: q}); err != nil {
		return u
	}
	u.saveBuf = buf
	u.imgType = "image/jpeg"
	return u
}

func (u *util) ReadImage(content []byte) *util {
	decode, _, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		log.Fatalf("decode image err:%#v", err)
	}
	contentType := u.FileContentType(content[:512])
	log.Infof("file content type  %#v\n", contentType)
	return &util{img: decode, imgType: contentType}
}

func (u *util) CompressPng(quality ...int) *util {
	if u.img == nil {
		log.Fatalf("compress png data err:%#v", u.img)
	}

	newRGBA := image.NewRGBA(u.img.Bounds())
	draw.Draw(newRGBA, newRGBA.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
	draw.Draw(newRGBA, newRGBA.Bounds(), u.img, image.Point{}, draw.Over)

	var buf bytes.Buffer

	q := 40
	if len(quality) > 0 {
		q = quality[0]
	}

	if err := jpeg.Encode(&buf, newRGBA, &jpeg.Options{Quality: q}); err != nil {
		return u
	}
	u.saveBuf = buf
	u.imgType = "image/jpeg"
	return u
}

func (u *util) ScaleImages(scalePercentage int) *util {
	if u.img == nil {
		log.Fatalf("scale image data is nil:%#v", u.img)
	}

	scaleDx := u.img.Bounds().Dx() * scalePercentage
	scaleDy := u.img.Bounds().Dy() * scalePercentage

	newRGBA := image.NewRGBA(image.Rect(0, 0, scaleDx, scaleDy))

	if err := graphics.Scale(newRGBA, u.img); err != nil {
		log.Fatalf("graphics scale image error:$#v", err)
	}
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, newRGBA, &jpeg.Options{Quality: 40}); err != nil {
		log.Fatalf("scale jpeg encode image err:%#v", err)
	}
	u.saveBuf = buf
	u.imgType = "image/jpeg"
	return u
}

func (u *util) ToBase64() string {
	if u.saveBuf.Len() <= 0 {
		log.Fatalf("save to base64 data is nil")
	}
	var bf bytes.Buffer

	bf.WriteString(fmt.Sprintf("data:%s;base64,", u.imgType))
	bf.WriteString(base64.StdEncoding.EncodeToString(u.saveBuf.Bytes()))
	return bf.String()
}

func (u *util) Save(w io.Writer) error {

	if u.saveBuf.Len() <= 0 {
		log.Fatalf("save image  data:%#v", u.saveBuf)
	}

	if _, err := u.saveBuf.WriteTo(w); err != nil {
		log.Fatalf("save image to write:%#v", err)
	}

	return nil
}
