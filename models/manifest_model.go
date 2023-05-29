package models

type ManifestData struct {
	UUID   string  `form:"uuid"`
	Label  string  `form:"label"`
	Images []Image `form:"images,omitempty"`
}

type Image struct {
	ID string `json:"id"`
}

type Language struct {
	English  []string `jsonld:"en"`
	Japanese []string `jsonld:"jp"`
}

type Manifest struct {
	Context     string                   `jsonld:"@context"`
	ID          string                   `jsonld:"@id"`
	Type        string                   `jsonld:"@type"`
	Label       Language                 `jsonld:"label"`
	Summary     Language                 `jsonld:"summary"`
	Attribution []map[string]interface{} `jsonld:"requiredStatement"`
	Metadata    []map[string]interface{} `jsonld:"metadata"`
	Thumbnail   []map[string]interface{} `jsonld:"thumbnail"`
	Items       []map[string]interface{} `jsonld:"items"`
}

func GenerateManifest(manifestID string) Manifest {

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

	manifest := Manifest{
		Context:     "http://iiif.io/api/presentation/3/context.json",
		ID:          manifestID,
		Type:        "sc:Manifest",
		Label:       Language{English: []string{"This will be Site+Location"}, Japanese: []string{"マニフェスト"}},
		Summary:     Language{English: []string{"Sangaku in X temple, Y location, by Z author"}, Japanese: []string{"算額の説明"}},
		Attribution: []map[string]interface{}{{"en": map[string]interface{}{"label": []string{"Attribution"}, "value": []string{"Copyright by XYZ"}}, "jp": map[string]interface{}{"label": []string{"神社名"}, "value": []string{"八幡お寺"}}}},
		Metadata:    []map[string]interface{}{{"label": map[string]interface{}{"en": []string{"Site"}, "jp": []string{"SiteJP"}}, "value": map[string]interface{}{"en": []string{"Hachiman Temple"}, "jp": []string{"八幡お寺"}}}, {"label": map[string]interface{}{"en": []string{"Location"}, "jp": []string{"住所"}}, "value": map[string]interface{}{"en": []string{"Tokyo"}, "jp": []string{"東京"}}}, {"label": map[string]interface{}{"en": []string{"Year"}, "jp": []string{"年"}}, "value": map[string]interface{}{"none": []int{1767}}}, {"label": map[string]interface{}{"en": []string{"Author"}, "jp": []string{"作者"}}, "value": map[string]interface{}{"en": []string{"Remo Grillo"}, "jp": []string{"Grillo-san"}}}, {"label": map[string]interface{}{"en": []string{"School"}, "jp": []string{"塾"}}, "value": map[string]interface{}{"en": []string{"Mashiko school"}, "jp": []string{"益子"}}}, {"label": map[string]interface{}{"en": []string{"Dimensions"}, "jp": []string{"サイズ"}}, "value": map[string]interface{}{"none": "120x40cm"}}, {"label": map[string]interface{}{"en": []string{"Medium"}, "jp": []string{"方法"}}, "value": map[string]interface{}{"en": []string{"Wood replica"}, "jp": []string{"レプリカント"}}}},
		Thumbnail:   []map[string]interface{}{},
		Items:       []map[string]interface{}{{"id": "someIDstring", "type": "sc:canvas", "label": map[string]interface{}{"en": []string{"Canvas with one sangaku image"}, "jp": []string{"translation of English"}}, "height": 3024, "width": 4032, "items": map[string]interface{}{"id": "someIDstring", "type": "AnnotationPage", "items": map[string]interface{}{"id": "someURI", "type": "Annotation", "motivation": "painting", "body": map[string]interface{}{"id": "someURI", "type": "Image", "format": "image/jpeg", "height": 3024, "width": 4032, "service": map[string]interface{}{"id": "someURI", "profile": "level1", "type": "ImageService3"}, "target": "targetURI"}}}}, {}},
	}

	return manifest
	/*
		content, err := jsonld.Marshal(manifest)
		if err != nil {
			fmt.Println(err)
		}
		return string(content)
	*/
}
