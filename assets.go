package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets747a811bc5536d86ab3f34108f10053e2fffc266 = "<!DOCTYPE html>\n<html>\n\n<head>\n    <title>Twitch bot - {{ .MainChannel }}</title>\n    <!-- CSS -->\n    <link rel=\"stylesheet\" href=\"{{ .Context.BaseURL }}/static/css/bootstrap.min.css\">\n\n    <!-- jQuery and JS bundle w/ Popper.js -->\n    <script src=\"{{ .Context.BaseURL }}/static/js/jquery.min.js\"></script>\n    <script src=\"{{ .Context.BaseURL }}/static/js/bootstrap.bundle.min.js\"></script>\n</head>\n\n<body>\n    <nav class=\"navbar navbar-expand-lg navbar-light bg-light\">\n        <a class=\"navbar-brand\" href=\"#\">Fonctions du Bot</a>\n        <button class=\"navbar-toggler\" type=\"button\" data-toggle=\"collapse\" data-target=\"#navbarNav\" aria-controls=\"navbarNav\" aria-expanded=\"false\" aria-label=\"Toggle navigation\">\n            <span class=\"navbar-toggler-icon\"></span>\n        </button>\n        <div class=\"collapse navbar-collapse\" id=\"navbarNav\">\n            <ul class=\"navbar-nav navbar-right\">\n            {{ range $route, $routeDetails := .Context.Navigation }}\n            <li class=\"nav-item\"><a class=\"nav-link\" href=\"{{ $routeDetails.Route }}\">{{ $routeDetails.Desc }}</a></li>\n            {{ end }}\n            </ul>\n        </div>\n    </nav>\n    {{template \"content\" .}}\n</body>\n</html>"
var _Assetsd03dfd4ffc41356e90e04d3791e679b0adfb978b = "{{define \"content\"}}\n<div class=\"form-bottom\">\n    <form role=\"form\" action=\"\" method=\"post\" class=\"login-form\">\n        <div class=\"form-group\">\n            <label for=\"form-username\">Nom d'utilisateur</label>\n            <input name=\"username\" placeholder=\"Username...\" class=\"form-username form-control\" id=\"form-username\" type=\"text\" required>\n        </div>\n        <div class=\"form-group\">\n            <label for=\"form-email\">email</label>\n            <input name=\"email\" placeholder=\"email...\" class=\"form-email form-control\" id=\"form-email\" type=\"text\" required>\n        </div>\n        <button type=\"submit\" class=\"btn btn-primary\">Se connecter</button>\n        <input type=\"hidden\" name=\"_csrf\" value=\"{{ .Context.CSRFToken }}\" />\n    </form>\n</div>\n{{end}}"
var _Assets55a37844ac8c3474cb9319c374bdc05b04d76fb5 = "{{define \"content\"}}\n    <h1>Welcome sur la gestion du bot Twitch pour : {{ .Context.MainChannel }}</h1>\n{{end}}"
var _Assetsee7a22cac18b698ae6c74dc1bfe05f4a92a61e87 = "Error {{ .error }}"
var _Assets792d102622f6e891107803cd74fc7253e1d8f61f = "{{define \"content\"}}\n    <h1>Statistiques pour {{ .Context.MainChannel }}</h1>\n    <table class=\"table table-bordered\">\n        <thead>\n            <tr>\n                <th scope=\"col\">User</th>\n                <th scope=\"col\">Score</th>\n            </tr>\n        </thead>\n        <tbody>\n    {{range $user, $score := .Data.Statistiques}}\n            <tr>\n                <td>{{ $user }}</td>\n                <td>{{ $score }}</td>\n            </tr>\n    {{end}}\n        </tbody>\n    </table>\n{{end}}\n"
var _Assets72baf2332493c53d1630406a465d95d6e729a9b1 = "{{define \"content\"}}\n<div class=\"form-bottom\">\n    <form role=\"form\" action=\"\" method=\"post\" class=\"login-form\">\n        <div class=\"form-group\">\n            <label for=\"form-username\">Nom d'utilisateur</label>\n            <input name=\"username\" placeholder=\"Username...\" class=\"form-username form-control\" id=\"form-username\" type=\"text\" required>\n        </div>\n        <div class=\"form-group\">\n            <label for=\"form-password\">Mot de passe</label>\n            <input name=\"password\" placeholder=\"Password...\" class=\"form-password form-control\" id=\"form-password\" type=\"password\" required>\n        </div>\n        <button type=\"submit\" class=\"btn btn-primary\">Se connecter</button>\n        <input type=\"hidden\" name=\"_csrf\" value=\"{{ .Context.CSRFToken }}\" />\n    </form>\n</div>\n{{end}}"
var _Assetsa9531d392a275cf19c734adc26dcc55447e7419e = "{{define \"content\"}}\n    <h1>\n    {{range $index, $channel := .Channels}}\n        {{ $channel }}\n    {{end}}\n    {{ .Context.MainChannel }}\n    </h1>\n        <table id=\"messages\" data-url=\"{{ .Context.BaseURL }}/json/messages\" data-auto-refresh=\"true\" data-auto-refresh-interval=\"5\">\n        <thead>\n            <tr>\n                <th data-field=\"Channel\">Channel</th>\n                <th data-field=\"Time\" data-formatter=\"timeFormater\">Time</th>\n                <th data-field=\"User.DisplayName\">User</th>\n                <th data-field=\"Message\">Message</th>\n            </tr>\n        </thead>\n    </table>\n    <script src=\"{{ .Context.BaseURL }}/static/js/bootstrap-table.min.js\"></script>\n    <script src=\"{{ .Context.BaseURL }}/static/js/bootstrap-table-auto-refresh.min.js\"></script>\n    <script>\n    $(document).ready(function(){\n        $('#messages').bootstrapTable();\n    });\n\n    /*function chanFormater(value) {\n        return '#' + value\n    }\n\n    function userFormater(value) {\n        return '&lt;' + value + '&gt;'\n    }*/\n\n    function timeFormater(value) {\n        var time = new Date(value);\n        h = time.getHours();\n        m = time.getMinutes();\n        s = time.getSeconds();\n\n        if (h < 10) {\n            h = '0' + h\n        }\n        if (m < 10) {\n            m = '0' + m\n        }\n        if (s < 10) {\n            s = '0' + s\n        }\n        return h + ':' + m + ':' + s\n    }\n\n    </script>\n{{end}}\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"includes", "layouts"}, "/includes": []string{"create_user_form.html", "home.html", "stats.html", "error.html", "login_form.html", "aggregator.html"}, "/layouts": []string{"base.html"}}, map[string]*assets.File{
	"/includes/stats.html": &assets.File{
		Path:     "/includes/stats.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1605714206, 1605714206929190544),
		Data:     []byte(_Assets792d102622f6e891107803cd74fc7253e1d8f61f),
	}, "/includes/login_form.html": &assets.File{
		Path:     "/includes/login_form.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1605714206, 1605714206929190544),
		Data:     []byte(_Assets72baf2332493c53d1630406a465d95d6e729a9b1),
	}, "/includes/aggregator.html": &assets.File{
		Path:     "/includes/aggregator.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1605714206, 1605714206929190544),
		Data:     []byte(_Assetsa9531d392a275cf19c734adc26dcc55447e7419e),
	}, "/layouts": &assets.File{
		Path:     "/layouts",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1605718006, 1605718006076734845),
		Data:     nil,
	}, "/layouts/base.html": &assets.File{
		Path:     "/layouts/base.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1605718006, 1605718006076734845),
		Data:     []byte(_Assets747a811bc5536d86ab3f34108f10053e2fffc266),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1605719127, 1605719127491315307),
		Data:     nil,
	}, "/includes": &assets.File{
		Path:     "/includes",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1605718667, 1605718667794103069),
		Data:     nil,
	}, "/includes/create_user_form.html": &assets.File{
		Path:     "/includes/create_user_form.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1605719124, 1605719124383360464),
		Data:     []byte(_Assetsd03dfd4ffc41356e90e04d3791e679b0adfb978b),
	}, "/includes/home.html": &assets.File{
		Path:     "/includes/home.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1605714206, 1605714206929190544),
		Data:     []byte(_Assets55a37844ac8c3474cb9319c374bdc05b04d76fb5),
	}, "/includes/error.html": &assets.File{
		Path:     "/includes/error.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1605714206, 1605714206929190544),
		Data:     []byte(_Assetsee7a22cac18b698ae6c74dc1bfe05f4a92a61e87),
	}}, "")
