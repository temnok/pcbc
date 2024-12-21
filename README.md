<img src="doc/gallery/logo.png" width="200"/>

## PCBC: fast PCB prototyping with a fiber laser

### Usage Example

```go
func main() {
	err := eda.GeneratePCB(&eda.Component{
		Components: eda.Components{
			pcbc.Board35x45,
			eda.ComponentGrid(3, 11, 5,
				x2.X2("LED+", "R2V-"),
				x2.X2("LED+", "G3V-"),
				x2.X2("LED+", "B3V-"),
				x2.X2("LED+", "Y2V-"),
				x2.X2("LED+", "W3V-"),
				x2.X2("R ", "50R"),
				x2.X2("R ", "50R"),
				x2.X2("R ", "K10"),
				x2.X2("R ", "K10"),
				x2.X2("R ", "K15"),
				x2.X2("R ", "K15"),
				x2.X2("R ", "K20"),
				x2.X2("R ", "K20"),
				x2.X2("R ", "K25"),
				x2.X2("R ", "K25"),
			).Arrange(transform.Rotate(90)),
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}
```

### Generated LightBurn files

* Etch
* Mask
* Stencil

### Generated Overview

* Dark green: FR4
* Bright Green: FR4 cuts
* Orange: copper
* Blue: soldermask cuts
* Pale Green/White: soldermask marks (silkscreen)
* Bright White: stencil cuts

![PY32](doc/gallery/py32/overview.png)
![E73](doc/gallery/e73/overview.png)
![Micro boards](doc/gallery/resistors.png)

### Gallery

![Finished](doc/gallery/e73/finished.jpg)
![Baked](doc/gallery/e73/baked.jpg)
![Stencil](doc/gallery/e73/stencil.jpg)
![Masked](doc/gallery/e73/masked.jpg)
![Etched](doc/gallery/e73/etched.jpg)
![Pre-etch](doc/gallery/e73/pre-etch.jpg)

