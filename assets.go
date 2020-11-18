package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets792d102622f6e891107803cd74fc7253e1d8f61f = "{{define \"content\"}}\r\n    <h1>Statistiques pour {{ .Context.MainChannel }}</h1>\r\n    <table class=\"table table-bordered\">\r\n        <thead>\r\n            <tr>\r\n                <th scope=\"col\">User</th>\r\n                <th scope=\"col\">Score</th>\r\n            </tr>\r\n        </thead>\r\n        <tbody>\r\n    {{range $user, $score := .Data.Statistiques}}\r\n            <tr>\r\n                <td>{{ $user }}</td>\r\n                <td>{{ $score }}</td>\r\n            </tr>\r\n    {{end}}\r\n        </tbody>\r\n    </table>\r\n{{end}}\r\n"
var _Assetsa9531d392a275cf19c734adc26dcc55447e7419e = "{{define \"content\"}}\r\n    <h1>\r\n    {{range $index, $channel := .Channels}}\r\n        {{ $channel }}\r\n    {{end}}\r\n    {{ .Context.MainChannel }}\r\n    </h1>\r\n        <table id=\"messages\" data-url=\"{{ .Context.BaseURL }}/json/messages\" data-auto-refresh=\"true\" data-auto-refresh-interval=\"5\">\r\n        <thead>\r\n            <tr>\r\n                <th data-field=\"Channel\">Channel</th>\r\n                <th data-field=\"Time\" data-formatter=\"timeFormater\">Time</th>\r\n                <th data-field=\"User.DisplayName\">User</th>\r\n                <th data-field=\"Message\">Message</th>\r\n            </tr>\r\n        </thead>\r\n    </table>\r\n    <script src=\"{{ .Context.BaseURL }}/static/js/bootstrap-table.min.js\"></script>\r\n    <script src=\"{{ .Context.BaseURL }}/static/js/bootstrap-table-auto-refresh.min.js\"></script>\r\n    <script>\r\n    $(document).ready(function(){\r\n        $('#messages').bootstrapTable();\r\n    });\r\n\r\n    /*function chanFormater(value) {\r\n        return '#' + value\r\n    }\r\n\r\n    function userFormater(value) {\r\n        return '&lt;' + value + '&gt;'\r\n    }*/\r\n\r\n    function timeFormater(value) {\r\n        var time = new Date(value);\r\n        h = time.getHours();\r\n        m = time.getMinutes();\r\n        s = time.getSeconds();\r\n\r\n        if (h < 10) {\r\n            h = '0' + h\r\n        }\r\n        if (m < 10) {\r\n            m = '0' + m\r\n        }\r\n        if (s < 10) {\r\n            s = '0' + s\r\n        }\r\n        return h + ':' + m + ':' + s\r\n    }\r\n\r\n    </script>\r\n{{end}}\r\n"
var _Assetsee7a22cac18b698ae6c74dc1bfe05f4a92a61e87 = "Error {{ .error }}"
var _Assets55a37844ac8c3474cb9319c374bdc05b04d76fb5 = "{{define \"content\"}}\r\n    <h1>Welcome sur la gestion du bot Twitch pour : {{ .Context.MainChannel }}</h1>\r\n{{end}}"
var _Assets72baf2332493c53d1630406a465d95d6e729a9b1 = "{{define \"content\"}}\r\n<div class=\"form-bottom\">\r\n    <form role=\"form\" action=\"\" method=\"post\" class=\"login-form\">\r\n        <div class=\"form-group\">\r\n            <label for=\"form-username\">Nom d'utilisateur</label>\r\n            <input name=\"username\" placeholder=\"Username...\" class=\"form-username form-control\" id=\"form-username\" type=\"text\" required>\r\n        </div>\r\n        <div class=\"form-group\">\r\n            <label for=\"form-password\">Mot de passe</label>\r\n            <input name=\"password\" placeholder=\"Password...\" class=\"form-password form-control\" id=\"form-password\" type=\"password\" required>\r\n        </div>\r\n        <button type=\"submit\" class=\"btn btn-primary\">Se connecter</button>\r\n        <input type=\"hidden\" name=\"_csrf\" value=\"{{ .Context.CSRFToken }}\" />\r\n    </form>\r\n</div>\r\n{{end}}"
var _Assets747a811bc5536d86ab3f34108f10053e2fffc266 = "<!DOCTYPE html>\r\n<html>\r\n\r\n<head>\r\n    <title>Twitch bot - {{ .MainChannel }}</title>\r\n    <!-- CSS -->\r\n    <link rel=\"stylesheet\" href=\"{{ .Context.BaseURL }}/static/css/bootstrap.min.css\">\r\n\r\n    <!-- jQuery and JS bundle w/ Popper.js -->\r\n    <script src=\"{{ .Context.BaseURL }}/static/js/jquery.min.js\"></script>\r\n    <script src=\"{{ .Context.BaseURL }}/static/js/bootstrap.bundle.min.js\"></script>\r\n</head>\r\n\r\n<body>\r\n    <nav class=\"navbar navbar-expand-lg navbar-light bg-light\">\r\n        <a class=\"navbar-brand\" href=\"#\">Fonctions du Bot</a>\r\n        <button class=\"navbar-toggler\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarNav\" aria-controls=\"navbarNav\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\r\n            <span class=\"navbar-toggler-icon\"></span>\r\n        </button>\r\n        <div class=\"collapse navbar-collapse\" id=\"navbarNav\">\r\n            <ul class=\"navbar-nav navbar-right\">\r\n            {{ range $route, $routeDetails := .Context.Navigation }}\r\n            <li class=\"nav-item\"><a class=\"nav-link\" href=\"{{ $routeDetails.Route }}\">{{ $routeDetails.Desc }}</a></li>\r\n            {{ end }}\r\n            </ul>\r\n        </div>\r\n    </nav>\r\n    {{template \"content\" .}}\r\n</body>\r\n</html>"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"includes", "layouts"}, "/includes": []string{"aggregator.html", "error.html", "home.html", "login_form.html", "stats.html"}, "/layouts": []string{"base.html"}}, map[string]*assets.File{
	"/includes/aggregator.html": &assets.File{
		Path:     "/includes/aggregator.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1605647592, 1605647592787399300),
		Data:     []byte(_Assetsa9531d392a275cf19c734adc26dcc55447e7419e),
	}, "/includes/stats.html": &assets.File{
		Path:     "/includes/stats.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1605647981, 1605647981419464800),
		Data:     []byte(_Assets792d102622f6e891107803cd74fc7253e1d8f61f),
	}, "/layouts": &assets.File{
		Path:     "/layouts",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1605644251, 1605644251971837500),
		Data:     nil,
	}, "/includes/login_form.html": &assets.File{
		Path:     "/includes/login_form.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1605647967, 1605647967791393100),
		Data:     []byte(_Assets72baf2332493c53d1630406a465d95d6e729a9b1),
	}, "/layouts/base.html": &assets.File{
		Path:     "/layouts/base.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1605713586, 1605713586507321300),
		Data:     []byte(_Assets747a811bc5536d86ab3f34108f10053e2fffc266),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1605713592, 1605713592533678900),
		Data:     nil,
	}, "/includes": &assets.File{
		Path:     "/includes",
		FileMode: 0x800001ff,
		Mtime:    time.Unix(1605647615, 1605647615945280900),
		Data:     nil,
	}, "/includes/error.html": &assets.File{
		Path:     "/includes/error.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1605481193, 1605481193521701700),
		Data:     []byte(_Assetsee7a22cac18b698ae6c74dc1bfe05f4a92a61e87),
	}, "/includes/home.html": &assets.File{
		Path:     "/includes/home.html",
		FileMode: 0x1b6,
		Mtime:    time.Unix(1605648308, 1605648308394155900),
		Data:     []byte(_Assets55a37844ac8c3474cb9319c374bdc05b04d76fb5),
	}}, "")
