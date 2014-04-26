package MimeTypes

var allTypes [32]MimeType

func init() {
	allTypes[0] = TypeWebHTML
	allTypes[1] = TypeWebCSS
	allTypes[2] = TypeWebJavaScript
	allTypes[3] = TypeXML
	allTypes[4] = TypeArchiveZIP
	allTypes[5] = TypeArchiveGZ
	allTypes[6] = TypeWebOCTET
	allTypes[7] = TypeDocumentPDF
	allTypes[8] = TypeDocumentLaTeX
	allTypes[9] = TypeShockwave
	allTypes[10] = TypeArchiveTAR
	allTypes[11] = TypeAudioWAV
	allTypes[12] = TypeAudioMP3
	allTypes[13] = TypeAudioAAC
	allTypes[14] = TypeAudioOGG
	allTypes[15] = TypeAudioWMA
	allTypes[16] = TypeImageGIF
	allTypes[17] = TypeImageCommon
	allTypes[18] = TypeUnknown
	allTypes[19] = TypeImageJPEG
	allTypes[20] = TypeImagePNG
	allTypes[21] = TypePlainText
	allTypes[22] = TypeVideoMPEG
	allTypes[23] = TypeVideoMOV
	allTypes[24] = TypeVideoAVI
	allTypes[25] = TypeVideoMP4
	allTypes[26] = TypeFontEOT
	allTypes[27] = TypeFontOTF
	allTypes[28] = TypeImageSVG
	allTypes[29] = TypeFontTTF
	allTypes[30] = TypeFontWOFF
	allTypes[31] = TypeWebJSON
}
