package callapi

import (
	"bytes"
	"encoding/json"
	"image/jpeg"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"

	"github.com/hatobus/Teikyo/models"
)

func DetectFace(fstream multipart.File) (*models.FaceParts, error) {

	exe, _ := os.Getwd()
	savedir := filepath.Join(exe, "picture", "output")

	// ディレクトリの確認
	// ディレクトリがなかった場合は作成
	// それでもディレクトリを作成できなかったらエラー
	if f, err := os.Stat(savedir); os.IsNotExist(err) || !f.IsDir() {
		if err := os.MkdirAll(savedir, 0777); err != nil {
			return nil, err
		}
	}

	// 画像をデコード
	imgBin, err := jpegToBin(fstream)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// 叩くURLは以下のようにするのでこれを満たすように作成
	// https://[location].api.cognitive.microsoft.com/face/v1.0/detect[?returnFaceId][&returnFaceLandmarks][&returnFaceAttributes]

	URL := os.Getenv("URL")
	subscriptionKey := os.Getenv("KEY1")

	u, err := url.Parse(URL)
	if err != nil {
		return nil, err
	}

	// クエリパラメータはreturnFaceLandmarkdが必要なのでtrueにしておく
	query := url.Values{
		"returnFaceLandmarks": {"true"},
	}

	// URLの生成
	u.Path = path.Join(u.Path, "?", query.Encode())
	log.Println(u.Path)

	req, err := http.NewRequest("POST", u.String(), bytes.NewBuffer(imgBin))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("Ocp-Apim-Subscription-Key", subscriptionKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	facelandmark := &models.FaceParts{}
	err = json.Unmarshal(body, facelandmark)
	if err != nil {
		return nil, err
	}

	return facelandmark, nil
}

func jpegToBin(fstream multipart.File) ([]byte, error) {
	buf := new(bytes.Buffer)
	// どうやらファイルの先頭までシークをしなければいけなかったっぽい
	// https://stackoverflow.com/questions/32193395/golang-io-reader-issue-with-jpeg-decode-returning-eof
	fstream.Seek(0, 0)

	img, err := jpeg.Decode(fstream)
	if err != nil {
		log.Println("line96: ", err)
		// ここで unexpected EOF
		return buf.Bytes(), err
	}

	if err = jpeg.Encode(buf, img, nil); err != nil {
		return buf.Bytes(), err
	}

	return buf.Bytes(), nil
}
