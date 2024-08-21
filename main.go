package main

import (
    "fmt"
    "image"
    "image/color"
    "image/gif"
    "os"
    "time"
)

var eightRegions = [8][2]int{
    {0, 31}, {32, 63}, {64, 95}, {96, 127},
    {128, 159}, {160, 191}, {192, 223}, {224, 255},
}

var rRegionMappings = make([][]int, 8)
var gRegionMappings = make([][]int, 8)
var bRegionMappings = make([][]int, 8)

var rRepresentativeColorPerRegion = make([]int, 8)
var gRepresentativeColorPerRegion = make([]int, 8)
var bRepresentativeColorPerRegion = make([]int, 8)

func getRegionIndex(colorValue int) int {
    for index, region := range eightRegions {
        if colorValue >= region[0] && colorValue <= region[1] {
            return index
        }
    }
    return 0
}

func main() {
    start := time.Now() // Start timer

    file, err := os.Open("sd.gif")
    if err != nil {
        fmt.Println("Error opening image file:", err)
        return
    }
    defer file.Close()

    g, err := gif.DecodeAll(file)
    if err != nil {
        fmt.Println("Error decoding GIF:", err)
        return
    }

    newGif := gif.GIF{}

    for _, img := range g.Image {
        bounds := img.Bounds()
        width, height := bounds.Max.X, bounds.Max.Y

        for i := 0; i < 8; i++ {
            rRegionMappings[i] = []int{}
            gRegionMappings[i] = []int{}
            bRegionMappings[i] = []int{}
        }

        for y := 0; y < height; y++ {
            for x := 0; x < width; x++ {
                r, g, b, _ := img.At(x, y).RGBA()
                r, g, b = r>>8, g>>8, b>>8

                rRegionMappings[getRegionIndex(int(r))] = append(rRegionMappings[getRegionIndex(int(r))], int(r))
                gRegionMappings[getRegionIndex(int(g))] = append(gRegionMappings[getRegionIndex(int(g))], int(g))
                bRegionMappings[getRegionIndex(int(b))] = append(bRegionMappings[getRegionIndex(int(b))], int(b))
            }
        }

        for i := 0; i < 8; i++ {
            rRepresentativeColorPerRegion[i] = calculateMean(rRegionMappings[i])
            gRepresentativeColorPerRegion[i] = calculateMean(gRegionMappings[i])
            bRepresentativeColorPerRegion[i] = calculateMean(bRegionMappings[i])
        }

        newImg := image.NewPaletted(bounds, img.Palette)
        for y := 0; y < height; y++ {
            for x := 0; x < width; x++ {
                r, g, b, a := img.At(x, y).RGBA()
                r, g, b = r>>8, g>>8, b>>8

                newR := rRepresentativeColorPerRegion[getRegionIndex(int(r))]
                newG := gRepresentativeColorPerRegion[getRegionIndex(int(g))]
                newB := bRepresentativeColorPerRegion[getRegionIndex(int(b))]

                newImg.SetColorIndex(x, y, uint8(img.Palette.Index(color.RGBA{uint8(newR), uint8(newG), uint8(newB), uint8(a >> 8)})))
            }
        }

        newGif.Image = append(newGif.Image, newImg)
        newGif.Delay = append(newGif.Delay, g.Delay[0])
    }

    outputFile, err := os.Create("uniform_quantized.gif")
    if err != nil {
        fmt.Println("Error creating output image file:", err)
        return
    }
    defer outputFile.Close()

    err = gif.EncodeAll(outputFile, &newGif)
    if err != nil {
        fmt.Println("Error encoding GIF:", err)
        return
    }

    fmt.Println("Quantized GIF saved as uniform_quantized.gif")

    elapsed := time.Since(start) // Calculate elapsed time
    fmt.Printf("Execution time: %s\n", elapsed)
}

func calculateMean(values []int) int {
    if len(values) == 0 {
        return 0
    }
    sum := 0
    for _, v := range values {
        sum += v
    }
    return sum / len(values)
}