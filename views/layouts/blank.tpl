<!DOCTYPE html>
<html>
<head>
    <title>Lin Li</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    {{ block "css" . }}{{ end }}
</head>
<body class="container">
  <h1>Header</h1>
  {{ block "yield" . }}{{ end }}
  <h3>Footer</h3>
  {{ block "js" . }}{{ end }}
</body>
</html>
