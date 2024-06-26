package mimes

const (
	// MIME text
	TextPlain             = "text/plain"
	TextPlainUTF8         = "text/plain; charset=utf-8"
	TextPlainISO88591     = "text/plain; charset=iso-8859-1"
	TextPlainFormatFlowed = "text/plain; format=flowed"
	TextPlainDelSpaceYes  = "text/plain; delsp=yes"
	TextPlainDelSpaceNo   = "text/plain; delsp=no"
	TextHtml              = "text/html"
	TextCss               = "text/css"
	TextJavascript        = "text/javascript"
	MultipartPOSTForm     = "multipart/form-data"

	// MIME application
	ApplicationOctetStream  = "application/octet-stream"
	ApplicationFlash        = "application/x-shockwave-flash"
	ApplicationHTMLForm     = "application/x-www-form-urlencoded"
	ApplicationHTMLFormUTF8 = "application/x-www-form-urlencoded; charset=UTF-8"
	ApplicationTar          = "application/x-tar"
	ApplicationGZip         = "application/gzip"
	ApplicationXGZip        = "application/x-gzip"
	ApplicationBZip2        = "application/bzip2"
	ApplicationXBZip2       = "application/x-bzip2"
	ApplicationShell        = "application/x-sh"
	ApplicationDownload     = "application/x-msdownload"
	ApplicationJSON         = "application/json"
	ApplicationJSONUTF8     = "application/json; charset=utf-8"
	ApplicationXML          = "application/xml"
	ApplicationXMLUTF8      = "application/xml; charset=utf-8"
	ApplicationZip          = "application/zip"
	ApplicationPdf          = "application/pdf"
	ApplicationWord         = "application/msword"
	ApplicationExcel        = "application/vnd.ms-excel"
	ApplicationPPT          = "application/vnd.ms-powerpoint"
	ApplicationOpenXMLWord  = "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	ApplicationOpenXMLExcel = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	ApplicationOpenXMLPPT   = "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	PROTOBUF                = "application/x-protobuf"

	// MIME image
	ImageJPEG         = "image/jpeg"
	ImagePNG          = "image/png"
	ImageGIF          = "image/gif"
	ImageBitmap       = "image/bmp"
	ImageWebP         = "image/webp"
	ImageIco          = "image/x-icon"
	ImageMicrosoftICO = "image/vnd.microsoft.icon"
	ImageTIFF         = "image/tiff"
	ImageSVG          = "image/svg+xml"
	ImagePhotoshop    = "image/vnd.adobe.photoshop"

	// MIME audio
	AudioBasic     = "audio/basic"
	AudioL24       = "audio/L24"
	AudioMP3       = "audio/mp3"
	AudioMP4       = "audio/mp4"
	AudioMPEG      = "audio/mpeg"
	AudioOggVorbis = "audio/ogg"
	AudioWAVE      = "audio/vnd.wave"
	AudioWebM      = "audio/webm"
	AudioAAC       = "audio/x-aac"
	AudioAIFF      = "audio/x-aiff"
	AudioMIDI      = "audio/x-midi"
	AudioM3U       = "audio/x-mpegurl"
	AudioRealAudio = "audio/x-pn-realaudio"

	// MIME video
	VideoMPEG          = "video/mpeg"
	VideoOgg           = "video/ogg"
	VideoMP4           = "video/mp4"
	VideoQuickTime     = "video/quicktime"
	VideoWinMediaVideo = "video/x-ms-wmv"
	VideWebM           = "video/webm"
	VideoFlashVideo    = "video/x-flv"
	Video3GPP          = "video/3gpp"
	VideoAVI           = "video/x-msvideo"
	VideoMatroska      = "video/x-matroska"
)
