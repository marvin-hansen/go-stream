package v1

import . "go-stream/api/types"

func (s SDKImpl) handle(message DataMessage, invokeFunction InvokeFunction) (err error) {

	if invokeFunction == nil {
		return
	} else {
		return invokeFunction(message)
	}
}
