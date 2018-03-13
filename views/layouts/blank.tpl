<!DOCTYPE html>
<html>
<head>
  <title>{{.title}}</title>
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  {{LoadCSSs .CSSs | str2html}}
</head>
<body class="container">
  <h1>Header</h1>
  {{.LayoutContent}}
  <h3>Footer</h3>
  {{LoadJSs .JSs | str2html}}
  <h5>XD</h5>
</body>
</html>
