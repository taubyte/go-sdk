package event

import "io"

/************************ BODY READCLOSER ****************************/

var (
	_EventBody EventBody
	_          io.Reader = _EventBody
	_          io.Closer = _EventBody
)

/************************** EVENT TYPES ****************************/
