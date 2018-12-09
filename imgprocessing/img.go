package imageprocessing

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"math"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/hatobus/Teikyo/models"
	"github.com/nfnt/resize"
)

func GenTeikyo(fstream multipart.File, parts *models.Landmark) error {

	goPath := os.Getenv("GOPATH")
	imgPath := filepath.Join(goPath, "src", "github.com", "hatobus", "Teikyo", "picture", "material")
	outputfile := filepath.Join(goPath, "src", "github.com", "hatobus", "Teikyo", "picture", "output", "output.png")

	t, _ := os.Open(filepath.Join(imgPath, "teikyo-tei.png"))
	log.Println(filepath.Join(imgPath, "teikyo-tei.png"))
	defer t.Close()

	k, _ := os.Open(filepath.Join(imgPath, "teikyo-kyo.png"))
	defer k.Close()

	if _, err := fstream.Seek(0, 0); err != nil {
		return err
	}

	Tei, _, err := image.Decode(t)
	if err != nil {
		return err
	}

	Kyo, _, err := image.Decode(k)
	if err != nil {
		return err
	}

	dstimg, _, err := image.Decode(fstream)
	if err != nil {
		return err
	}
	RightEyeSizeHeight := parts.EyeRight.BottomX - parts.EyeRight.TopX
	RightEyeSizeWidth := parts.EyeRight.BottomY - parts.EyeRight.TopY

	LeftEyeSizeHeight := parts.EyeLeft.BottomX - parts.EyeLeft.TopX
	LeftEyeSizeWidth := parts.EyeLeft.BottomY - parts.EyeLeft.TopY

	LTei := resize.Resize(
		uint(math.Abs(RightEyeSizeWidth)),
		uint(math.Abs(RightEyeSizeHeight)),
		Tei,
		resize.Lanczos3)

	RKyo := resize.Resize(
		uint(math.Abs(LeftEyeSizeWidth)),
		uint(math.Abs(LeftEyeSizeHeight)),
		Kyo,
		resize.Lanczos3)

	// References this
	// https://blog.golang.org/go-imagedraw-package

	/*

		// 書き出し用のイメージの準備
		outRect := image.Rectangle{image.Pt(, 0), dstimg.Bounds().Size()}
		out := image.NewRGBA(outRect)

		// 描画する処理
		dstRectangle := image.Rectangle{image.Pt(0, 0), dstimg.Bounds().Size()}
		draw.Draw(out, dstRectangle, dstimg, image.Pt(0, 0), draw.Src)

		// 自分から見て左目に提供の提の字
		TeiRectangle := image.Rectangle{image.Pt(int(parts.EyeLeft.TopX), int(parts.EyeLeft.TopY)), LTei.Bounds().Size()}
		draw.Draw(
			out,
			TeiRectangle,
			LTei,
			TeiRectangle.Min,
			draw.Src)

		// 自分から見て右目に提供の供の字
		KyoRectangle := image.Rectangle{image.Pt(0, 0), RKyo.Bounds().Size()}
		draw.DrawMask(
			out,
			dstimg.Bounds(),
			RKyo,
			image.ZP,
			KyoRectangle,
			image.ZP,
			draw.Over)

		// 書き出し用のファイルを作る
		outfile, err := os.Create(outputfile)
		defer outfile.Close()
		if err != nil {
			return err
		}

		png.Encode(outfile, out)
	*/
	TeistartPoint := image.Point{int(parts.EyeLeft.TopX), int(parts.EyeLeft.TopY)}
	KyostartPoint := image.Point{int(parts.EyeRight.TopX), int(parts.EyeRight.TopY)}

	TeiRectangle := image.Rectangle{TeistartPoint, TeistartPoint.Add(Tei.Bounds().Size())}
	KyoRectangle := image.Rectangle{KyostartPoint, KyostartPoint.Add(Kyo.Bounds().Size())}
	DstRectangle := image.Rectangle{image.Point{0, 0}, dstimg.Bounds().Size()}

	rgba := image.NewRGBA(DstRectangle)
	draw.Draw(rgba, DstRectangle, dstimg, image.Pt(0, 0), draw.Src)
	draw.Draw(rgba, TeiRectangle, LTei, image.Pt(0, 0), draw.Src)
	draw.Draw(rgba, KyoRectangle, RKyo, image.Pt(0, 0), draw.Src)

	outfile, err := os.Create(outputfile)
	defer outfile.Close()
	if err != nil {
		return err
	}

	png.Encode(outfile, rgba)

	return nil
}
