package models

type FaceParts struct {
	FaceID        string `json:"faceId"`
	FaceRectangle struct {
		Top    int `json:"top"`
		Left   int `json:"left"`
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"faceRectangle"`
	FaceLandmarks struct {
		PupilLeft struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"pupilLeft"`
		PupilRight struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"pupilRight"`
		NoseTip struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseTip"`
		MouthLeft struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"mouthLeft"`
		MouthRight struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"mouthRight"`
		EyebrowLeftOuter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowLeftOuter"`
		EyebrowLeftInner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowLeftInner"`
		EyeLeftOuter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftOuter"`
		EyeLeftTop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftTop"`
		EyeLeftBottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftBottom"`
		EyeLeftInner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeLeftInner"`
		EyebrowRightInner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowRightInner"`
		EyebrowRightOuter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyebrowRightOuter"`
		EyeRightInner struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightInner"`
		EyeRightTop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightTop"`
		EyeRightBottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightBottom"`
		EyeRightOuter struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"eyeRightOuter"`
		NoseRootLeft struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRootLeft"`
		NoseRootRight struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRootRight"`
		NoseLeftAlarTop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseLeftAlarTop"`
		NoseRightAlarTop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRightAlarTop"`
		NoseLeftAlarOutTip struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseLeftAlarOutTip"`
		NoseRightAlarOutTip struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"noseRightAlarOutTip"`
		UpperLipTop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"upperLipTop"`
		UpperLipBottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"upperLipBottom"`
		UnderLipTop struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"underLipTop"`
		UnderLipBottom struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"underLipBottom"`
	} `json:"faceLandmarks"`
}
