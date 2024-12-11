package url

import (
	"net/http"

	"gopkg.in/yaml.v3"
)
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path 
		if value, ok := pathsToUrls[key]; ok {
			http.Redirect(w, r, value, http.StatusPermanentRedirect) 
		}
		fallback.ServeHTTP(w, r) 
	}

}


func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	var myData []Element
	err := yaml.Unmarshal(yml, &myData) 
	if err != nil {
		return fallback.ServeHTTP, err
	}
	m := make(map[string]string, len(myData))
	for _, element := range myData {
		m[element.Path] = element.URL
		
	}
	return MapHandler(m, fallback), nil
}                                                                      

type Element struct {
	Path string
	URL  string
}