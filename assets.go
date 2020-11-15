package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assetsa9531d392a275cf19c734adc26dcc55447e7419e = "{{define \"content\"}}\r\n    <script type=\"text/javascript\" language=\"javascript\">\r\n    setTimeout(function(){\r\n        window.location.reload(1);\r\n    }, 5000);\r\n    </script>\r\n    <h1>\r\n    {{range $channel := .Data.Channels}}\r\n        {{ . }}\r\n    {{end}}\r\n    {{ .MainChannel }}\r\n    </h1><ul>\r\n    {{range $message := .Data.Messages}}\r\n        <li>{{ . }}</li>\r\n    {{end}}\r\n    </ul>\r\n{{end}}\r\n"
var _Assets9c21febaee572857d6d7f084f3f6b603538796a3 = "<!DOCTYPE html>\r\n<html>\r\n\r\n<head>\r\n    <title>Twitch bot - {{ .MainChannel }}</title>\r\n    <!-- CSS -->\r\n    <link rel=\"stylesheet\" href=\"https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/css/bootstrap.min.css\" integrity=\"sha384-TX8t27EcRE3e/ihU7zmQxVncDAy5uIKz4rEkgIXeMed4M0jlfIDPvg6uqKI2xXr2\" crossorigin=\"anonymous\">\r\n\r\n    <!-- jQuery and JS bundle w/ Popper.js -->\r\n    <script src=\"https://code.jquery.com/jquery-3.5.1.slim.min.js\" integrity=\"sha384-DfXdz2htPH0lsSSs5nCTpuj/zy4C+OGpamoFVy38MVBnE+IbbVYUew+OrCXaRkfj\" crossorigin=\"anonymous\"></script>\r\n    <script src=\"https://cdn.jsdelivr.net/npm/bootstrap@4.5.3/dist/js/bootstrap.bundle.min.js\" integrity=\"sha384-ho+j7jyWK8fNQe+A12Hb8AhRq26LrZ/JpcUGGOn+Y7RsweNrtN/tE3MoK7ZeZDyx\" crossorigin=\"anonymous\"></script>\r\n</head>\r\n\r\n<body>\r\n    <nav class=\"navbar navbar-expand-lg navbar-light bg-light\">\r\n        <a class=\"navbar-brand\" href=\"#\">Fonctions du Bot</a>\r\n        <button class=\"navbar-toggler\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarNav\" aria-controls=\"navbarNav\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\r\n            <span class=\"navbar-toggler-icon\"></span>\r\n        </button>\r\n        <div class=\"collapse navbar-collapse\" id=\"navbarNav\">\r\n            <ul class=\"navbar-nav\">\r\n            {{ range $route, $routeDetails := .WebRoutes }}\r\n            <li class=\"nav-item\"><a class=\"nav-link\" href=\"{{ $route }}\">{{ $routeDetails.RouteDesc }}</a></li>\r\n            {{ end }}\r\n            </ul>\r\n        </div>\r\n    </nav>\r\n    {{template \"content\" .}}\r\n</body>\r\n\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"includes", "layouts"}, "/includes": []string{"aggregator.html"}, "/layouts": []string{"index.html"}}, map[string]*assets.File{
	"/includes": &assets.File{
		Path:     "/includes",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1605392068, 1605392068175111600),
		Data:     nil,
	}, "/includes/aggregator.html": &assets.File{
		Path:     "/includes/aggregator.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1605395581, 1605395581791536800),
		Data:     []byte(_Assetsa9531d392a275cf19c734adc26dcc55447e7419e),
	}, "/layouts": &assets.File{
		Path:     "/layouts",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1605392048, 1605392048179565000),
		Data:     nil,
	}, "/layouts/index.html": &assets.File{
		Path:     "/layouts/index.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1605396127, 1605396127151181000),
		Data:     []byte(_Assets9c21febaee572857d6d7f084f3f6b603538796a3),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1605399067, 1605399067776600300),
		Data:     nil,
	}}, "")
