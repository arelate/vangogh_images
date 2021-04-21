package vangogh_images

//starting with empty collection and no image types require auth at the moment
var imageTypeRequiresAuth []ImageType

func RequiresAuth(it ImageType) bool {
	for _, itra := range imageTypeRequiresAuth {
		if itra == it {
			return true
		}
	}
	return false
}
