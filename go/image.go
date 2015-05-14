package main

import (
    "encoding/json"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "image"
    "image/png"
    _ "image/jpeg"
)

type ReqArgs struct {
    Cmd string `json:"cmd"`
    Src struct {
        Url      string `json:"url"`
        Mimetype string `json:"mimetype"`
        Fsize    int32  `json:"fsize"`
        Bucket   string `json:"bucket"`
        Key      string `json:"key"`
    } `json: "src"`
}

func grayscale(in io.Reader, out io.Writer) {
    src, _, err := image.Decode(in)
    if err != nil {
        log.Fatal(err)
    }

    bounds := src.Bounds()
    dst := image.NewGray(bounds)

    for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
        for x := bounds.Min.X; x < bounds.Max.X; x++ {
            color := src.At(x, y)
            dst.Set(x, y, color)
        }
    }

    png.Encode(out, dst)
}

func demoHandler(w http.ResponseWriter, req *http.Request) {
    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        w.WriteHeader(400)
        log.Println("read request body failed:", err)
        return
    }

    var args ReqArgs
    err = json.Unmarshal(body, &args)
    if err != nil {
        w.WriteHeader(400)
        log.Println("invalid request body:", err)
        return
    }

    resp, err := http.Get(args.Src.Url)
    if err != nil {
        w.WriteHeader(400)
        log.Println("fetch resource failed:", err)
        return
    }

    defer resp.Body.Close()
    grayscale(resp.Body, w)
//    contentType := http.DetectContentType(buf)
//    w.Write([]byte(contentType))
}

func main() {
    http.HandleFunc("/uop", demoHandler)
    err := http.ListenAndServe(":9100", nil)
    if err != nil {
        log.Fatal("Demo server failed to start:", err)
    }
}

