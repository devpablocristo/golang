package renderjet

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

// views is the jet view set
var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./chat/infrastructure/client"),
	jet.InDevelopmentMode(), // stop the need to reload after a change
)

// renderPage renders a jet template
// data can be nil
func RenderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	// view is the template
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
