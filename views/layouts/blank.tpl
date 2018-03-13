<!DOCTYPE html>
<html>
<head>
  <title>{{.title}}</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  {{range $key, $val := .CSSs}}
  <link href="{{htmlquote $val}}.css" rel="stylesheet" type="text/css"/>
  {{end}}
</head>
<body class="container">
  <h1>Header</h1>
  {{.LayoutContent}}
  <h3>Footer</h3>

  {{range $key, $val := .JSs}}
  <script type="text/javascript" src="{{htmlquote $val}}.js"></script>
  {{end}}
</body>
</html>
