package mod_media

import "mime/multipart"

func UploadMedia(file *multipart.FileHeader) error {
	bucket, ctx := InitFirebaseStorage()

	err := bucket.Create(ctx, "assets", nil)
	return err
}
