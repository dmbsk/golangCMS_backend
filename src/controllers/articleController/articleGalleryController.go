package articleController

import (
	"bytes"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"../../models/articleModel"
	. "../../respondFormating"
	"github.com/disintegration/imaging"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func CreateGalleryEndPoint(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error parsing form")
		fmt.Println(err)
		return
	}

	formData := r.MultipartForm
	files := formData.File["images"]

	thumbnailName := formData.Value["thumbnailImageName"][0]

	fmt.Println(len(files))
	fmt.Println("Uploading...")
	for i := range files {
		imageFile, err := files[i].Open()
		if err != nil {
			RespondWithError(w, http.StatusBadRequest, "Error reading image")
			fmt.Println(err)
			return
		}

		if files[i].Filename == thumbnailName {
			handleThumbnailFile(imageFile, w, r)
			imageFile, err = files[i].Open()
			if err != nil {
				RespondWithError(w, http.StatusBadRequest, "Error reading image again after creating thumbnail")
				fmt.Println(err)
				return
			}
		}
		fmt.Println(imageFile)
		handleFile(imageFile, w, r)
	}

	articleId := formData.Value["articleId"][0]
	if respond, err := updateGallery(articleId); err != nil {
		RespondWithError(w, http.StatusInternalServerError, respond)
		fmt.Println(err)
		return
	}
	RespondWithJson(w, http.StatusCreated, "image uploaded")
}

func updateGallery(articleId string) (respond string, err error){
	selectedArticle, err := dao.FindById(articleId)
	if err != nil {
		respond = "Cant find article"
		return
	}
	imageLinks, thumbnailLink := getImageLinks(selectedArticle.Gallery)
	selectedArticle.Gallery.ImagesLinks, selectedArticle.Gallery.Thumbnail = imageLinks, thumbnailLink
	if err := dao.Update(selectedArticle); err != nil {
		respond = "Cant update article gallery"
		return respond, err
	}

	respond = "Article gallery updated"
	return

}

func getImageLinks(gallery articleModel.Gallery) (fileArr []string, thumbnailImageLink string){
	galleryIdString := bson.ObjectId(gallery.Id).Hex()
	path := "temp/images/gallery." + galleryIdString
	files, err := ioutil.ReadDir(path)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error reading images dir")
	}
	fmt.Println(len(files))
	href := "/article/gallery/" + galleryIdString + "/"

	for _, file := range files {
		if strings.Split(file.Name(), "-")[0] != "thumbnail" {
			fileArr = append(fileArr, href + file.Name())
		} else {
			thumbnailImageLink = href + file.Name()
		}
	}

	return
}

func ImageEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	path := "temp/images/gallery." + params["id"] + "/" + params["imgName"]

	img, err := os.Open(path)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error reading image")
		fmt.Println(err)
		return
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/jpeg")
	if _, err := io.Copy(w, img); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error coping image")
		fmt.Println(err)
		return
	}

}
func handleFile(imageFile multipart.File, w http.ResponseWriter, r *http.Request) {
	defer imageFile.Close()
	galleryId := r.FormValue("galleryId")
	galleryDirectory := "temp/images/gallery." + galleryId

	if _, err := os.Stat(galleryDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(galleryDirectory, 0766); err != nil {
			RespondWithError(w, http.StatusBadRequest, "Error creating gallery directory")
			fmt.Println(err)
			return
		}
	}

	tempFileName := "upload-*.png"
	tempFile, err := ioutil.TempFile(galleryDirectory, tempFileName)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error creating temp image")
		fmt.Println(err)
		return
	}


	imageBytes, err := ioutil.ReadAll(imageFile)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error reading image file")
		fmt.Println(err)
	}

	if _, err := tempFile.Write(imageBytes); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error writing image")
		fmt.Println(err)
	}

	if err := tempFile.Close(); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error closing temp image")
		fmt.Println(err)
		return
	}
}

func handleThumbnailFile(imageFile multipart.File, w http.ResponseWriter, r *http.Request) {
	defer imageFile.Close()
	galleryId := r.FormValue("galleryId")
	galleryDirectory := "temp/images/gallery." + galleryId

	if _, err := os.Stat(galleryDirectory); os.IsNotExist(err) {
		if err := os.Mkdir(galleryDirectory, 0766); err != nil {
			RespondWithError(w, http.StatusBadRequest, "Error creating gallery directory")
			fmt.Println(err)
			return
		}
	}

	tempFileName := "thumbnail-upload.png"
	imageBytes, err := ioutil.ReadAll(imageFile)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error reading image file")
		fmt.Println(err)
		return
	}

	bigImage, _ , err := image.Decode(bytes.NewReader(imageBytes))
	imageSizes, _, err := image.DecodeConfig(bytes.NewReader(imageBytes))
	height := 400
	width := (imageSizes.Height / imageSizes.Width) * height
	smallImage := imaging.Resize(bigImage, width, height, imaging.Lanczos)

	imagingSavePath := galleryDirectory + "/" + tempFileName
	if err := imaging.Save(smallImage, imagingSavePath); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Error reading image file")
		fmt.Println(err)
		return
	}
}