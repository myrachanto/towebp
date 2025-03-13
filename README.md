# towebp

A simple Go package to convert PNG and JPEG/JPG images to WebP format.

## Overview

The `towebp` package provides a single function `Towebp` that converts PNG and JPEG/JPG images to WebP format using the `github.com/chai2010/webp` library. It maintains good image quality while reducing file size through WebP compression.

## Installation

To install the package, use:

```bash
go get github.com/myrachanto/towebp
```

### supported formats
- jpg, jpeg
- png
