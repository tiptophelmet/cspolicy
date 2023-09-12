# cspolicy
[![codecov](https://codecov.io/gh/tiptophelmet/cspolicy/graph/badge.svg?token=i1BrUFenkV)](https://codecov.io/gh/tiptophelmet/cspolicy)

User-friendly API to build your Content-Security-Policy HTTP header.

## ⬇️ Overview

CSPolicy is designed to simplify the process of creating and managing CSP headers in Go applications. With a structured approach to defining sources and directives, this library ensures that you maintain a robust and secure Content-Security-Policy HTTP header for your Go applications.

## 💻 Getting Started

### ⚙️ Installation
```
go get github.com/tiptophelmet/cspolicy
```

### 💡 Usage
Here's a basic example to get you started:
```
package main

import (
    "github.com/tiptophelmet/cspolicy"
    "github.com/tiptophelmet/cspolicy/src"
    "github.com/tiptophelmet/cspolicy/directives"
)

func main() {
	csp := cspolicy.Build(
		directives.DefaultSrc(src.None()),
		directives.BaseURI(src.Self(), src.Host("*.example.com")),
		directives.ChildSrc(
			src.Host("cdn.example.com/assets"),
			src.Host("resources.example.com/artifacts"),
		),
		directives.ConnectSrc(
			src.Host("uploads.example.com"),
			src.Host("status.example.com"),
			src.Host("api.example.com"),
		),
		directives.FrameSrc(
			src.Host("notes.example.com"),
			src.Host("viewbox.example.com"),
		),
		directives.ImgSrc(
			src.Self(),
			src.Scheme("data:"),
			src.Host("media.example.com"),
			src.Host("avatars.example.com"),
		),
		directives.UpgradeInsecureRequests(),
	)

    fmt.Println(csp)
}
```

## 📖 Documentation
Visit pkg.go.dev/github.com/tiptophelmet/cspolicy.

## 🙋 Contributing
Your contribution is valuable! If you find a bug or have a feature request, please open an issue. If you'd like to contribute code, please open a pull request.

## 🏷️ License
This project is licensed under the MIT License. See the LICENSE file for details.