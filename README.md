# Uniform Image Quantization

This project provides a Go implementation for uniform quantization of images. It supports both GIF and PNG formats.

The main difference with Uniform Image Quantization is that we no longer need to generate palettes for the images.

## Features

- Quantizes GIF images and saves the result as `uniform_quantized.gif`.
- Quantizes PNG images and saves the result as `uniform_quantized.png`.
- Calculates representative colors for each region in the RGB color space.

## Installation

1. Ensure you have Go installed on your machine.
2. Clone this repository.
3. Navigate to the project directory.

## Usage

### Quantizing a GIF Image

1. Place the GIF image you want to quantize in the project directory and name it `sd.gif`.
2. Run the following command:

   ```sh
   go run main.go
   ```

3. The quantized GIF will be saved as `uniform_quantized.gif`.

### Quantizing a PNG Image

1. Place the PNG image you want to quantize in the project directory and name it `rem.png`.
2. Run the following command:

   ```sh
   go run main.go
   ```

3. The quantized PNG will be saved as `uniform_quantized.png`.

## Code Overview

### Main Functions

- `forGif()`: Handles the quantization process for GIF images.
- `forImage()`: Handles the quantization process for PNG images.

### Helper Functions

- `getRegionIndex(colorValue int) int`: Determines the region index for a given color value.
- `calculateMean(values []int) int`: Calculates the mean of a slice of integers.
