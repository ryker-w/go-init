package static

import (
	"embed"
)

/**
-pkg 指定生成文件的包名
-o 指定生成文件名
--prefix 替换文件路径
*/

//go:embed *
var Static embed.FS
