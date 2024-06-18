package barcode

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/qr"
)

// func main() {

// 	lContent := "THANKS FOR SCANNING"
// 	lContent = "https://www.instagram.com/swetha_appadurai/?hl=en"
// 	lErr := GenerateQRCode(lContent)
// 	if lErr != nil {
// 		log.Println("ERROR WHILE CREATING QR CODE ", lErr)
// 		return
// 	} else {
// 		GenerateBarCode(lContent)
// 		if lErr != nil {
// 			log.Println("ERROR WHILE CREATING BAR CODE ", lErr)
// 			return
// 		}
// 	}
// }

func GenerateQRCode(pContent string) (lErr error) {
	log.Println("GenerateQRCode (+)")
	// Generate QR code
	lQRCode, lErr := qr.Encode(pContent, qr.M, qr.Auto)
	if lErr != nil {
		log.Println("ERROR WHILE ENCODE ", lErr)
		return
	} else {
		lQRCode, lErr = barcode.Scale(lQRCode, 200, 200) // Scale to 200x200 pixels
		if lErr != nil {
			log.Println("ERROR WHILE CREATING BARCODE ", lErr)
			return
		} else {
			var lQRFile *os.File
			// Create a PNG file for QR code
			lQRFile, lErr = os.Create("qrcode.png")
			if lErr != nil {
				log.Println("ERROR WHILE CREATING BARCODE ", lErr)
				return
			} else {
				defer lQRFile.Close()
				png.Encode(lQRFile, lQRCode)
			}
		}
	}
	log.Println("GenerateQRCode (-)")
	return
}

func GenerateBarCode(pContent string) (lErr error) {
	log.Println("GenerateBarCode (+)")
	// Generate Code 128 barcode
	lCode128Code, lErr := code128.Encode(pContent)
	if lErr != nil {
		log.Println("ERROR WHILE CREATING BARCODE ", lErr)
		return
	} else {

		// Calculate the width and height of the barcode image
		lBarcodeWidth := 600
		lBarcodeHeight := 200

		// Create a new image with white background
		lIimg := image.NewRGBA(image.Rect(0, 0, lBarcodeWidth, lBarcodeHeight))
		lBackground := color.White
		lDrawColor := color.Black

		// Calculate the scaling factors
		lScaleX := float64(lBarcodeWidth) / float64(lCode128Code.Bounds().Dx())
		lScaleY := float64(lBarcodeHeight) / 100.0 // The height of Code 128 is always 100

		// Draw the barcode on the image
		for x := 0; x < lBarcodeWidth; x++ {
			for y := 0; y < lBarcodeHeight; y++ {
				if lCode128Code.At(int(float64(x)/lScaleX), int(float64(y)/lScaleY)) == color.Black {
					lIimg.Set(x, y, lDrawColor)
				} else {
					lIimg.Set(x, y, lBackground)
				}
			}
		}
		var lBarCodeFile *os.File
		// Create a PNG lBarCodeFile for the Code 128 barcode
		lBarCodeFile, lErr = os.Create("code128.png")
		if lErr != nil {
			log.Println("ERROR WHILE CREATING BARCODE ", lErr.Error())
			return
		} else {
			defer lBarCodeFile.Close()
			png.Encode(lBarCodeFile, lIimg)
		}
	}
	log.Println("GenerateBarCode (-)")
	return
}
