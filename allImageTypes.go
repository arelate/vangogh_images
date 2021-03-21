package vangogh_images

func All() []ImageType {
	imageTypes := make([]ImageType, 0, len(imageTypeStrings))
	for it, _ := range imageTypeStrings {
		if it == Unknown {
			continue
		}
		imageTypes = append(imageTypes, it)
	}
	return imageTypes
}
