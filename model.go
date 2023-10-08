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
			X string `xml:"x"`
			Y string `xml:"y"`
			Z string `xml:"z"`
		} `xml:"offset"`
		Rotation struct {
			X string `xml:"x"`
			Y string `xml:"y"`
			Z string `xml:"z"`
		} `xml:"rotation"`
		Up struct {
			X string `xml:"x"`
			Y string `xml:"y"`
			Z string `xml:"z"`
		} `xml:"up"`
		ScreenBounds struct {
			X string `xml:"x"`
			Y string `xml:"y"`
			Z string `xml:"z"`
		} `xml:"screenbounds"`
		Origin struct {
			X string `xml:"x"`
			Y string `xml:"y"`
			Z string `xml:"z"`
		} `xml:"origin"`
		RealScale struct {
			X    string `xml:"x"`
			Y    string `xml:"y"`
			Z    string `xml:"z"`
			Unit string `xml:"unit"`
		} `xml:"realscale"`
		Audio      string `xml:"audio"`
		Background string `xml:"background"`
	} `xml:"environment"`
}

type Stroke struct {
	IsDrawing string `xml:"isdrawing,attr"`
	Pt        []struct {
		X    string `xml:"x"`
		Y    string `xml:"y"`
		Z    string `xml:"z"`
		T    string `xml:"t"`
		Pres string `xml:"pres"`
		Rot  string `xml:"rot"`
		Dir  struct {
			X string `xml:"x"`
			Y string `xml:"y"`
			Z string `xml:"z"`
		} `xml:"dir"`
	} `xml:"pt"`
	Brush Brush `xml:"brush"`
	Info  struct {
		Curved string `xml:"curved"`
	} `xml:"info"`
}

type Brush struct {
	Mode              string `xml:"mode"`
	UniqueStyleId     string `xml:"uniquestyleid"`
	Spec              string `xml:"spec"`
	Width             string `xml:"width"`
	SpeedToWidthRatio string `xml:"speedtowidthratio"`
	DripAmnt          string `xml:"dripamnt"`
	DripSpeed         string `xml:"dripspeed"`
	LayerAbsolute     string `xml:"layerabsolute"`
	Color             struct {
		R string `xml:"r"`
		G string `xml:"g"`
		B string `xml:"b"`
		A string `xml:"a"`
	} `xml:"color"`
	DripVecRelativeToUp struct {
		X string `xml:"x"`
		Y string `xml:"y"`
		Z string `xml:"z"`
	} `xml:"dripvecrelativetoup"`
	LayerRelative string `xml:"layerrelative"`
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
