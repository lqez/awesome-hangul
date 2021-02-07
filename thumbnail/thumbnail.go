package main

import (
    "github.com/fogleman/gg"
		"bufio"
		"os"
    "fmt"
    "log"
    "strings"
)

func countLanguages() int {
    file, err := os.Open("../README.md")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    mode := 0
    languages := 0

    for scanner.Scan() {
        text := scanner.Text()
        switch mode {
        case 0:
            if strings.HasPrefix(text, "## Programming Languages") {
                mode = 1
            }
        case 1:
            if strings.HasPrefix(text, "### ") {
                languages += 1
            } else if strings.HasPrefix(text, "## ") {
                break
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return languages
}

func main() {
		languages := countLanguages()
    fmt.Println(languages, "programming language(s) found.")

    img, err := gg.LoadImage("thumbnail-base.png")
    if err != nil {
        log.Fatal(err)
    }

		bounds := img.Bounds()
    dc := gg.NewContext(bounds.Max.X, bounds.Max.Y)

    if err := dc.LoadFontFace("./RobotoSlab-Regular.ttf", 42); err != nil {
        panic(err)
    }

    dc.DrawImage(img, 0, 0)
		dc.SetRGB(0, 0, 0)

    s := fmt.Sprintf("Includes %d+ programming languages", languages)
    dc.DrawString(s, 324, 316)

    dc.SavePNG("thumbnail.png")
}
