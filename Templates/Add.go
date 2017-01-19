package Templates

//AddTemplate adds a template to the template cache so it can be used by ProcessHTML
func AddTemplate(src string) error {
	_, err := templates.Parse(src)
	return err
}
