package test

import (
	"github.com/ruanlianjun/image-tools/wimg"
	"io/ioutil"
	"os"
	"testing"
)

func TestCompressJpeg(t *testing.T) {
	file, err := ioutil.ReadFile("./file/1.jpg")
	if err != nil {
		t.Fatalf("open file err:%#v", err)
	}
	create, err := os.Create("./file/out.webp")
	if err != nil {
		t.Fatalf("create file err:%#v", err)
	}

	compress, err := os.Create("./file/Compress.jpeg")
	if err != nil {
		t.Fatalf("create file err:%#v", err)
	}

	if err = wimg.Util.ReadImage(file).WebpEncoder(10).Save(create); err != nil {
		t.Fatalf("save file err:%#v", err)
	}
	if base64Img := wimg.Util.ReadImage(file).WebpEncoder(10).ToBase64(); base64Img != "" {
		t.Logf("save file data:%#v", base64Img)
	}

	if err = wimg.Util.ReadImage(file).CompressJpeg(10).Save(compress); err != nil {
		t.Logf("save CompressJpeg file err:%s\n", err)
	}
}
