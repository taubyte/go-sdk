package event

import "io"

/************************ BODY READCLOSER ****************************/

var (
	_HttpEventBody HttpEventBody
	_              io.Reader = _HttpEventBody
	_              io.Closer = _HttpEventBody
)

/************************** EVENT TYPES ****************************/
