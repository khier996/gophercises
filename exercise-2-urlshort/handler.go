package urlshort

import (
	"net/http"
  "gopkg.in/yaml.v2"
)

type pathUrl struct {
  Path string `yaml:"path"`
  Url string `yaml:"url"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
  return func(w http.ResponseWriter, r *http.Request) {
    for url, path := range pathsToUrls {
      if r.URL.Path == url {
        http.Redirect(w, r, path, 301)
        return
      }
    }
    fallback.ServeHTTP(w, r)
  }
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
  var pathUrls []pathUrl
  if err := yaml.Unmarshal(yml, &pathUrls); err != nil {
    return nil, err
  }
	pathsMap := make(map[string]string)
  for _, pathUrlObj := range pathUrls {
    pathsMap[pathUrlObj.Path] = pathUrlObj.Url
  }
  return MapHandler(pathsMap, fallback), nil
  // return nil, nil
}


