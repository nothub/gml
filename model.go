package gml

import (
	"bytes"
	"encoding/xml"
)

type GML struct {
	XMLName xml.Name `xml:"gml"`
	Spec    string   `xml:"spec,attr"`
	Tag     struct {
		Header  Header `xml:"header"`
		Drawing struct {
			Strokes []Stroke `xml:"stroke"`
		} `xml:"drawing"`
	} `xml:"tag"`
}

type Header struct {
	Client struct {
		Name      string `xml:"name"`
		Version   string `xml:"version"`
		Username  string `xml:"username"`
		Permalink string `xml:"permalink"`
		Keywords  string `xml:"keywords"`
		UniqueKey string `xml:"uniquekey"`
		Ip        string `xml:"ip"`
		Time      string `xml:"time"`
		Location  struct {
			Lat string `xml:"lat"`
			Lon string `xml:"lon"`
		} `xml:"location"`
	} `xml:"client"`
	Environment struct {
		Offset struct {
			X float32 `xml:"x"`
			Y float32 `xml:"y"`
			Z float32 `xml:"z"`
		} `xml:"offset"`
		Rotation struct {
			X float32 `xml:"x"`
			Y float32 `xml:"y"`
			Z float32 `xml:"z"`
		} `xml:"rotation"`
		Up struct {
			X float32 `xml:"x"`
			Y float32 `xml:"y"`
			Z float32 `xml:"z"`
		} `xml:"up"`
		ScreenBounds struct {
			X float32 `xml:"x"`
			Y float32 `xml:"y"`
			Z float32 `xml:"z"`
		} `xml:"screenbounds"`
		Origin struct {
			X float32 `xml:"x"`
			Y float32 `xml:"y"`
			Z float32 `xml:"z"`
		} `xml:"origin"`
		RealScale struct {
			X    float32 `xml:"x"`
			Y    float32 `xml:"y"`
			Z    float32 `xml:"z"`
			Unit string  `xml:"unit"`
		} `xml:"realscale"`
		Audio      string `xml:"audio"`
		Background string `xml:"background"`
	} `xml:"environment"`
}

type Stroke struct {
	IsDrawing bool `xml:"isdrawing,attr"`
	Pt        []struct {
		X    float32 `xml:"x"`
		Y    float32 `xml:"y"`
		Z    float32 `xml:"z"`
		T    float32 `xml:"t"`
		Pres float32 `xml:"pres"`
		Rot  float32 `xml:"rot"`
		Dir  struct {
			X float32 `xml:"x"`
			Y float32 `xml:"y"`
			Z float32 `xml:"z"`
		} `xml:"dir"`
	} `xml:"pt"`
	Brush Brush `xml:"brush"`
	Info  struct {
		Curved string `xml:"curved"`
	} `xml:"info"`
}

type Brush struct {
	Mode              string  `xml:"mode"`
	UniqueStyleId     string  `xml:"uniquestyleid"`
	Spec              string  `xml:"spec"`
	Width             int     `xml:"width"`
	SpeedToWidthRatio float32 `xml:"speedtowidthratio"`
	DripAmnt          float32 `xml:"dripamnt"`
	DripSpeed         float32 `xml:"dripspeed"`
	LayerAbsolute     int     `xml:"layerabsolute"`
	Color             struct {
		R uint8 `xml:"r"`
		G uint8 `xml:"g"`
		B uint8 `xml:"b"`
		A uint8 `xml:"a"`
	} `xml:"color"`
	DripVecRelativeToUp struct {
		X float32 `xml:"x"`
		Y float32 `xml:"y"`
		Z float32 `xml:"z"`
	} `xml:"dripvecrelativetoup"`
	LayerRelative int `xml:"layerrelative"`
}

func (g *GML) Serialize() ([]byte, error) {
	return nil, nil
}

func Deserialize(raw []byte) (*GML, error) {
	gml := &GML{}
	err := xml.NewDecoder(bytes.NewReader(raw)).Decode(gml)
	if err != nil {
		return nil, err
	}
	return gml, nil
}
