这是一些关于Go的小脚本或项目，记录一下
=
0721,16
Template

它的文档写的真的很不人性化(应该是我的英语烂)
如果要使用到xml的一些判断循环就要按照它的语法来做。

比如range .obj(for i in obj)
if eq .sth ""(if sth == "")

eg：
	{{println .List}}
	{{if eq .Result "success" }}<Device >b</Device>
	{{ else}}<Device >a</Device>
	{{ end}}
