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

## Code Overview

### Main Functions

- `forGif()`: Handles the quantization process for GIF images.
- `forImage()`: Handles the quantization process for PNG images.

### Helper Functions

- `getRegionIndex(colorValue int) int`: Determines the region index for a given color value.
- `calculateMean(values []int) int`: Calculates the mean of a slice of integers.
