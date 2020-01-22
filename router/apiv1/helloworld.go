package apiv1

import (
	"github.com/labstack/echo/v4"
)

func HelloWorld(c echo.Context) error {
	return c.HTML(200, `
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Single file upload</title>
</head>
<body>
<h1>Upload single file with fields</h1>

<form action="/api/v1/files" method="post" enctype="multipart/form-data">
    Files: <input type="file" name="file"><br><br>
    <input type="submit" value="Submit">
</form>
</body>
</html>`)
}
