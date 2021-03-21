package vangogh_images

var imageTypeRequiresAuth []ImageType

func RequiresAuth(it ImageType) bool {
	for _, itra := range imageTypeRequiresAuth {
		if itra == it {
			return true
		}
	}
	return false
}
