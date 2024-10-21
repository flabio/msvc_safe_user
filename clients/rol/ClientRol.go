package rol

import (
	"net/http"
	"strconv"

	"github.com/safe_msvc_user/insfractruture/utils"
)

func MsvcRolFindById(id uint) (Rol, string) {
	data, err := http.NewRequest(utils.GET, utils.MSVC_ROL_URL+"/"+strconv.Itoa(int(id)), nil)
	rol, msg := DataRolClient(data, err)
	return rol, msg
}
