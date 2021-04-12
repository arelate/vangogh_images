package vangogh_images

type ImageType int

const (
	Unknown ImageType = iota
	Image
	BoxArt
	Logo
	Icon
	Background
	GalaxyBackground
	Screenshots
	Video
)

var imageTypeStrings = map[ImageType]string{
	Unknown:          "unknown-image-type",
	Image:            "image",
	BoxArt:           "box-art",
	Logo:             "logo",
	Icon:             "icon",
	Background:       "background",
	GalaxyBackground: "galaxy-background",
	Screenshots:      "screenshots",
	Video:            "video",
}

func (it ImageType) String() string {
	str, ok := imageTypeStrings[it]
	if ok {
		return str
	}

	return imageTypeStrings[Unknown]
}

func Parse(imageType string) ImageType {
	for it, str := range imageTypeStrings {
		if str == imageType {
			return it
		}
	}
	return Unknown
}

func Valid(it ImageType) bool {
	_, ok := imageTypeStrings[it]
	return ok && it != Unknown
}
