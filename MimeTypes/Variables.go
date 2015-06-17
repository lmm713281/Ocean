package MimeTypes

var (
	allTypes [34]MimeType // Array for all known types

	// Create instances for each known types:
	TypeWebHTML       = MimeType{MimeType: "text/html", FileExtension: []string{".html", ".htm"}}
	TypeWebCSS        = MimeType{MimeType: "text/css", FileExtension: []string{".css"}}
	TypeWebJavaScript = MimeType{MimeType: "text/javascript", FileExtension: []string{".js"}}
	TypeWebDart       = MimeType{MimeType: "application/dart", FileExtension: []string{".dart"}}
	TypeXML           = MimeType{MimeType: "text/xml", FileExtension: []string{".xml"}}
	TypeArchiveZIP    = MimeType{MimeType: "application/zip", FileExtension: []string{".zip"}}
	TypeArchiveGZ     = MimeType{MimeType: "application/gzip", FileExtension: []string{".gz"}}
	TypeWebOCTET      = MimeType{MimeType: "application/octet-stream", FileExtension: []string{".bin", ".exe", ".dll", ".class"}}
	TypeDocumentPDF   = MimeType{MimeType: "application/pdf", FileExtension: []string{".pdf"}}
	TypeDocumentLaTeX = MimeType{MimeType: "application/x-latex", FileExtension: []string{".tex", ".latex"}}
	TypeShockwave     = MimeType{MimeType: "application/x-shockwave-flash", FileExtension: []string{".swf"}}
	TypeArchiveTAR    = MimeType{MimeType: "application/x-tar", FileExtension: []string{".tar"}}
	TypeAudioWAV      = MimeType{MimeType: "application/x-wav", FileExtension: []string{".wav"}}
	TypeAudioMP3      = MimeType{MimeType: "audio/mpeg", FileExtension: []string{".mp3"}}
	TypeAudioAAC      = MimeType{MimeType: "audio/aac", FileExtension: []string{".aac", ".m4a"}}
	TypeAudioOGG      = MimeType{MimeType: "audio/ogg", FileExtension: []string{"vogg", ".oga"}}
	TypeAudioWMA      = MimeType{MimeType: "audio/x-ms-wma", FileExtension: []string{".wma"}}
	TypeImageGIF      = MimeType{MimeType: "image/gif", FileExtension: []string{".gif"}}
	TypeImageCommon   = MimeType{MimeType: "image", FileExtension: []string{}}
	TypeUnknown       = MimeType{MimeType: "application/octet-stream", FileExtension: []string{}}
	TypeImageJPEG     = MimeType{MimeType: "image/jpeg", FileExtension: []string{".jpg", ".jpeg"}}
	TypeImagePNG      = MimeType{MimeType: "image/png", FileExtension: []string{".png"}}
	TypePlainText     = MimeType{MimeType: "text/plain", FileExtension: []string{".txt"}}
	TypeVideoMPEG     = MimeType{MimeType: "video/mpeg", FileExtension: []string{".mpeg", ".mpg"}}
	TypeVideoMOV      = MimeType{MimeType: "video/quicktime", FileExtension: []string{".mov", ".qt"}}
	TypeVideoAVI      = MimeType{MimeType: "video/x-msvideo", FileExtension: []string{".avi"}}
	TypeVideoMP4      = MimeType{MimeType: "video/mp4", FileExtension: []string{".mp4"}}
	TypeFontEOT       = MimeType{MimeType: "application/vnd.ms-fontobject", FileExtension: []string{".eot"}}
	TypeFontOTF       = MimeType{MimeType: "application/font-sfnt", FileExtension: []string{".otf"}}
	TypeImageSVG      = MimeType{MimeType: "image/svg+xml", FileExtension: []string{".svg"}}
	TypeFontTTF       = MimeType{MimeType: "application/font-sfnt", FileExtension: []string{".ttf"}}
	TypeFontWOFF      = MimeType{MimeType: "application/font-woff", FileExtension: []string{".woff"}}
	TypeWebJSON       = MimeType{MimeType: "application/json", FileExtension: []string{".json"}}
	TypeCSV           = MimeType{MimeType: "text/csv", FileExtension: []string{".csv"}}
)
