package statesstruct

import (
	"net/http"
	"strconv"

	"github.com/safe_msvc_user/insfractruture/utils"
)

func MsvcStateFindById(id uint) (StateResponse, string) {
	data, err := http.NewRequest(utils.GET, utils.MSVC_STATES_URL+"/"+strconv.Itoa(int(id)), nil)
	resp, msg := DataStateClient(data, err)
	return resp, msg
}
