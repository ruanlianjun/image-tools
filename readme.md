#### 图片处理工具
> 支持图片格式转化为webp格式，图片缩放与压缩
```go
file, err := ioutil.ReadFile("./file/1.jpg")
if err != nil {
    t.Fatalf("open file err:%#v", err)
}
create, err := os.Create("./file/out.webp")
if err != nil {
    t.Fatalf("create file err:%#v", err)
}

if err = utils.Util.ReadImage(file).WebpEncoder(10).Save(create); err != nil {
    t.Fatalf("save file err:%#v", err)
}
if base64Img := utils.Util.ReadImage(file).WebpEncoder(10).ToBase64(); base64Img != "" {
    t.Logf("save file data:%#v", base64Img)
}
```