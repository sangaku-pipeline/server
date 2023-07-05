package models

import (
	"tonothan/sangaku-pipeline-server/configs"
)

// insert following from database
// SiteEN, SiteJP
// LocationEN, LocationJP
// Image URIs
// Year
// AuthorEN, AuthorJP
// SchoolEN, SchoolJP
// Dimensions
// MediumEN, MediumJP
// PermissionsEN, PermissionsJP
// Manifest label will be Site+Location concatenated

type ManifestData struct {
	UUID          string  `form:"uuid"`
	SiteEN        string  `form:"siteEN"`
	SiteJP        string  `form:"siteJP"`
	LocationEN    string  `form:"loclabelationEN"`
	LocationJP    string  `form:"locationJP"`
	Year          int     `form:"year"`
	AuthorEN      string  `form:"authorEN"`
	AuthorJP      string  `form:"authorJP"`
	SchoolEN      string  `form:"schoolEN"`
	SchoolJP      string  `form:"schoolJP"`
	Dimensions    string  `form:"dimensions"`
	MediumEN      string  `form:"mediumEN"`
	MediumJP      string  `form:"mediumJP"`
	PermissionsEN string  `form:"permissionsEN"`
	PermissionsJP string  `form:"permissionsJP"`
	Images        []Image `form:"images,omitempty"`
}

type Image struct {
	ID string `json:"id"`
}

type Label struct {
	English  []string `json:"en"`
	Japanese []string `json:"jp"`
}

type Manifest struct {
	Context   string                   `json:"@context"`
	ID        string                   `json:"id"`
	Type      string                   `json:"type"`
	Label     Label                    `json:"label"`
	Summary   Label                    `json:"summary"`
	Metadata  []map[string]interface{} `json:"metadata"`
	Thumbnail []map[string]interface{} `json:"thumbnail"`
	Items     []CanvasItem             `json:"items"`
}

type AnnotationItem struct {
	ID         string         `json:"id"`
	Type       string         `json:"type"`
	Motivation string         `json:"motivation,omitempty"`
	Target     string         `json:"target,omitempty"`
	Body       AnnotationBody `json:"body"`
}

type CanvasItem struct {
	ID     string           `json:"id"`
	Type   string           `json:"type"`
	Label  Label            `json:"label"`
	Height int              `json:"height"`
	Width  int              `json:"width"`
	Items  []AnnotationItem `json:"items"`
}

type ResponseImageAPI struct {
	Id     string `json:"id"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
}

type Service2 struct {
	Id   string `json:"@id"`
	Type string `json:"@type"`
}

type Service struct {
	Id      string     `json:"id"`
	Type    string     `json:"type"`
	Profile string     `json:"profile"`
	Service []Service2 `json:"service"`
}

type AnnotationBody struct {
	Id      string    `json:"id"`
	Type    string    `json:"type"`
	Format  string    `json:"format"`
	Service []Service `json:"service"`
	Height  int       `json:"height"`
	Width   int       `json:"width"`
}

func GenerateManifest(retrievedManifest ManifestData) Manifest {

	baseID := configs.EnvBaseURI() + ":8080/" + retrievedManifest.UUID
	manifestID := baseID + ".json"

	items := []CanvasItem{}

	// Items:       []map[string]interface{}{{"id": "someIDstring", "type": "sc:canvas", "label": map[string]interface{}{"en": []string{"Canvas with one sangaku image"}, "jp": []string{"translation of English"}}, "height": 3024, "width": 4032, "items": map[string]interface{}{"id": "someIDstring", "type": "AnnotationPage", "items": map[string]interface{}{"id": "someURI", "type": "Annotation", "motivation": "painting", "body": map[string]interface{}{"id": "someURI", "type": "Image", "format": "image/jpeg", "height": 3024, "width": 4032, "service": map[string]interface{}{"id": "someURI", "profile": "level1", "type": "ImageService3"}, "target": "targetURI"}}}}, {}},
	for _, image := range retrievedManifest.Images {
		image_id := configs.EnvBaseURI() + ":8182/iiif/3/" + image.ID

		var response ResponseImageAPI
		configs.GetJson(image_id, &response)

		items = append(items, CanvasItem{
			ID:     baseID + "/canvas/1",
			Type:   "sc:canvas",
			Height: response.Height,
			Width:  response.Width,
			Label:  Label{English: []string{"Canvas with one sangaku image"}, Japanese: []string{"translation of English"}},
			Items: []AnnotationItem{{
				ID:         baseID + "/page/p1/1",
				Type:       "AnnotationPage",
				Motivation: "painting",
				Body: AnnotationBody{
					Id:     image_id + "/full/max/0/default.jpg",
					Type:   "Image",
					Format: "image/jpeg",
					Height: response.Height,
					Width:  response.Width,
					Service: []Service{{
						Id:      "https://example.org/iiif/book1/page1",
						Type:    "ImageService3",
						Profile: "level2",
						Service: []Service2{{
							Id:   "https://example.org/iiif/book1/page1",
							Type: "ImageService3",
						}},
					}},
				},
				Target: baseID + "/canvas/1",
			}}})
	}

	// canvas > page > annotation > image

	//"This will be Site+Location"

	manifest := Manifest{
		Context: "http://iiif.io/api/presentation/3/context.json",
		ID:      manifestID,
		Type:    "sc:Manifest",
		Label:   Label{English: []string{retrievedManifest.SiteEN + retrievedManifest.LocationEN}, Japanese: []string{retrievedManifest.SiteJP + retrievedManifest.LocationJP}},
		Summary: Label{English: []string{"Sangaku in X temple, Y location, by Z author"}, Japanese: []string{"算額の説明"}},
		//Attribution: []map[string]interface{}{{"en": map[string]interface{}{"label": []string{"Attribution"}, "value": []string{"Copyright by XYZ"}}, "jp": map[string]interface{}{"label": []string{"神社名"}, "value": []string{"八幡お寺"}}}},
		Metadata:  []map[string]interface{}{{"label": map[string]interface{}{"en": []string{"Site"}, "jp": []string{"SiteJP"}}, "value": map[string]interface{}{"en": []string{"Hachiman Temple"}, "jp": []string{"八幡お寺"}}}, {"label": map[string]interface{}{"en": []string{"Location"}, "jp": []string{"住所"}}, "value": map[string]interface{}{"en": []string{"Tokyo"}, "jp": []string{"東京"}}}, {"label": map[string]interface{}{"en": []string{"Year"}, "jp": []string{"年"}}, "value": map[string]interface{}{"none": []string{"1767"}}}, {"label": map[string]interface{}{"en": []string{"Author"}, "jp": []string{"作者"}}, "value": map[string]interface{}{"en": []string{"Remo Grillo"}, "jp": []string{"Grillo-san"}}}, {"label": map[string]interface{}{"en": []string{"School"}, "jp": []string{"塾"}}, "value": map[string]interface{}{"en": []string{"Mashiko school"}, "jp": []string{"益子"}}}, {"label": map[string]interface{}{"en": []string{"Dimensions"}, "jp": []string{"サイズ"}}, "value": map[string]interface{}{"none": []string{"120x40cm"}}}, {"label": map[string]interface{}{"en": []string{"Medium"}, "jp": []string{"方法"}}, "value": map[string]interface{}{"en": []string{"Wood replica"}, "jp": []string{"レプリカント"}}}},
		Thumbnail: []map[string]interface{}{},
		Items:     items,
	}

	return manifest
}
