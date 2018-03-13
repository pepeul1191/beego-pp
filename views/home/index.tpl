{{ template "../layouts/blank.tpl" }}
{{ define "css" }}
<link href="static/bower_components/bootstrap/dist/css/bootstrap.min.css" rel="stylesheet" type="text/css"/>
<link href="static/bower_components/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css"/>
{{ end }}
{{ define "yield" }}
hola mundo
<br/>
  Contact me: {{ .Email }}
</p>
{{ end }}
<br/>
Website : {{.Website}}
{{ define "js" }}
<script type="text/javascript" src="static/bower_components/jquery/dist/jquery.min.js"></script>
<script type="text/javascript" src="static/bower_components/bootstrap/dist/js/bootstrap.min.js"></script>
{{ end }}