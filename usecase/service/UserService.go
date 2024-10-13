package service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/safe_msvc_user/clients/rol"
	"github.com/safe_msvc_user/clients/statesstruct"
	"github.com/safe_msvc_user/core"
	"github.com/safe_msvc_user/insfractruture/entities"
	"github.com/safe_msvc_user/insfractruture/ui/global"
	"github.com/safe_msvc_user/insfractruture/ui/uicore"
	"github.com/safe_msvc_user/insfractruture/utils"
	"github.com/safe_msvc_user/usecase/dto"
	"github.com/ulule/deepcopier"
)

type userService struct {
	uiUser uicore.UIUserCore
}

func NewUserService() global.UIUserGlobal {
	return &userService{uiUser: core.NewUserRepository()}
}

func (s *userService) GetUserFindAll(c *fiber.Ctx) error {
	var userDto []dto.UserResponseDTO
	limit := 5
	page, begin := Pagination(c, int(limit))
	results, countUser, err := s.uiUser.GetUserFindAll(begin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	for _, item := range results {
		var userDtoItem dto.UserResponseDTO
		dataRol, _ := rol.MsvcRolFindById(item.RolId)
		dataState, _ := statesstruct.MsvcStateFindById(item.StateId)

		if dataRol.Name == "" {
			userDtoItem.RolName = "Not Rol"
		} else {
			userDtoItem.RolName = dataRol.Name
		}
		userDtoItem.StateName = dataState.Name

		deepcopier.Copy(&item).To(&userDtoItem)
		userDto = append(userDto, userDtoItem)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS:  fiber.StatusOK,
		utils.DATA:    userDto,
		"totalCount":  countUser,
		"pageCount":   page,
		"begin":       begin,
	})
}

func (s *userService) GetStudentsFindAll(c *fiber.Ctx) error {
	var userDto []dto.UserResponseDTO
	limit := 5
	page, begin := Pagination(c, int(limit))
	results, countUser, err := s.uiUser.GetStudentsFindAll(begin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	for _, item := range results {
		var userDtoItem dto.UserResponseDTO
		dataRol, _ := rol.MsvcRolFindById(item.RolId)
		dataState, _ := statesstruct.MsvcStateFindById(item.StateId)

		if dataRol.Name == "" {
			userDtoItem.RolName = "Not Rol"
		} else {
			userDtoItem.RolName = dataRol.Name
		}
		userDtoItem.StateName = dataState.Name

		deepcopier.Copy(&item).To(&userDtoItem)
		userDto = append(userDto, userDtoItem)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS:  fiber.StatusOK,
		utils.DATA:    userDto,
		"totalCount":  countUser,
		"pageCount":   page,
		"begin":       begin,
	})
}


func (s *userService) GetInstructorFindAll(c *fiber.Ctx) error {
	var userDto []dto.UserResponseDTO
	limit := 5
	page, begin := Pagination(c, int(limit))
	results, countUser, err := s.uiUser.GetInstructorFindAll(begin)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	for _, item := range results {
		var userDtoItem dto.UserResponseDTO
		dataRol, _ := rol.MsvcRolFindById(item.RolId)
		dataState, _ := statesstruct.MsvcStateFindById(item.StateId)

		if dataRol.Name == "" {
			userDtoItem.RolName = "Not Rol"
		} else {
			userDtoItem.RolName = dataRol.Name
		}
		userDtoItem.StateName = dataState.Name

		deepcopier.Copy(&item).To(&userDtoItem)
		userDto = append(userDto, userDtoItem)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS:  fiber.StatusOK,
		utils.DATA:    userDto,
		"totalCount":  countUser,
		"pageCount":   page,
		"begin":       begin,
	})
}

func (s *userService) GetUserFindById(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.uiUser.GetUserFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS: fiber.StatusOK,
		utils.DATA:   result,
	})
}

func (s *userService) CreateUser(c *fiber.Ctx) error {
	var userCreate entities.User

	userDto, msgError := validateUser(0, s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(userDto).To(&userCreate)
	dataRol, _ := rol.MsvcRolFindById(userDto.RolId)

	if dataRol.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ROL_NOT_FOUND,
		})
	}
	dataState, _ := statesstruct.MsvcStateFindById(userDto.StateId)
	if dataState.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.STATE_NOT_FOUND,
		})
	}
	userCreate.Password = utils.HashAndSalt([]byte(userCreate.Password))
	_, err := s.uiUser.CreateUser(userCreate)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_CREATE,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   utils.CREATED,
	})
}

func (s *userService) UpdateUser(c *fiber.Ctx) error {
	var updatedUser entities.User
	id, _ := strconv.Atoi(c.Params(utils.ID))
	result, err := s.uiUser.GetUserFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if result.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	deepcopier.Copy(result).To(&updatedUser)
	dataRol, _ := rol.MsvcRolFindById(result.RolId)

	if dataRol.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ROL_NOT_FOUND,
		})
	}
	dataState, _ := statesstruct.MsvcStateFindById(result.StateId)
	if dataState.Id == 0 {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.STATE_NOT_FOUND,
		})
	}
	userDto, msgError := validateUser(uint(id), s, c)
	if msgError != utils.EMPTY {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: msgError,
		})
	}
	deepcopier.Copy(userDto).To(&updatedUser)
	user, err := s.uiUser.UpdateUser(uint(id), updatedUser)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_UPDATE,
		})
	}
	return c.Status(http.StatusOK).JSON(fiber.Map{
		utils.STATUS: http.StatusOK,
		utils.DATA:   user,
	})
}
func (s *userService) DeleteUser(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params(utils.ID))
	userFindById, err := s.uiUser.GetUserFindById(uint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS: fiber.StatusBadRequest,
			utils.DATA:   utils.ERROR_QUERY,
		})
	}
	if userFindById.Id == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			utils.STATUS:  fiber.StatusNotFound,
			utils.MESSAGE: utils.ID_NO_EXIST,
		})
	}
	result, err := s.uiUser.DeleteUser(uint(id))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			utils.STATUS:  http.StatusBadRequest,
			utils.MESSAGE: utils.ERROR_DELETE,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		utils.STATUS:  http.StatusOK,
		utils.MESSAGE: utils.REMOVED,
		utils.DATA:    result,
	})
}

func validateUser(id uint, s *userService, c *fiber.Ctx) (dto.UserDTO, string) {
	var userDto dto.UserDTO
	var msg string = utils.EMPTY
	body := c.Body()

	var dataMap map[string]interface{}
	err := json.Unmarshal([]byte(body), &dataMap)
	if err != nil {
		msg = err.Error()
	}

	msgValid := validateField(dataMap)
	if msgValid != utils.EMPTY {
		return dto.UserDTO{}, msgValid
	}

	MapToStructUser(&userDto, dataMap)
	msgRequired := validateRequired(userDto)
	if msgRequired != utils.EMPTY {
		return dto.UserDTO{}, msgRequired
	}
	existEmail, _ := s.uiUser.GetUserFindByEmailAndId(id, userDto.Email)

	if existEmail {
		msg = utils.EMAIL_ALREADY_EXIST
	}
	return userDto, msg
}
func MapToStructUser(userDto *dto.UserDTO, dataMap map[string]interface{}) {
	user := dto.UserDTO{
		FirstName:    dataMap[utils.FIRST_NAME].(string),
		FirstSurName: dataMap[utils.FIRST_SUR_NAME].(string),
		SeconSurName: dataMap[utils.SECOND_SUR_NAME].(string),
		Address:      dataMap[utils.ADDRESS].(string),
		Phone:        dataMap[utils.PHONE].(string),
		ZipCode:      dataMap[utils.ZIP_CODE].(string),
		StateId:      uint(dataMap[utils.STATE_ID].(float64)),
		RolId:        uint(dataMap[utils.ROL_ID].(float64)),
		Email:        dataMap[utils.EMAIL].(string),
		Password:     dataMap[utils.PASSWORD].(string),
	}
	*userDto = user
}
func validateField(value map[string]interface{}) string {
	var msg string = utils.EMPTY
	if value[utils.FIRST_NAME] == nil {
		msg = utils.FIRST_NAME_IS_FIELD_REQUIRED
	}
	if value[utils.FIRST_SUR_NAME] == nil {
		msg = utils.FIRST_SUR_NAME_IS_FIELD_REQUIRED
	}
	if value[utils.ADDRESS] == nil {
		msg = utils.ADDRESS_IS_FIELD_REQUIRED
	}
	if value[utils.PHONE] == nil {
		msg = utils.PHONE_IS_FIELD_REQUIRED
	}
	if value[utils.ZIP_CODE] == nil {
		msg = utils.ZIP_CODE_IS_FIELD_REQUIRED
	}
	if value[utils.STATE_ID] == nil {
		msg = utils.STATE_ID_IS_FIELD_REQUIRED
	}
	if value[utils.ROL_ID] == nil {
		msg = utils.ROL_ID_IS_FIELD_REQUIRED
	}
	return msg
}

func validateRequired(user dto.UserDTO) string {
	var msg string = utils.EMPTY
	if user.FirstName == utils.EMPTY {
		msg = utils.FIRST_NAME_IS_REQUIRED
	}
	if user.FirstSurName == utils.EMPTY {
		msg = utils.FIRST_SUR_NAME_IS_REQUIRED
	}
	if user.Address == utils.EMPTY {
		msg = utils.ADDRESS_IS_REQUIRED
	}
	if user.Phone == utils.EMPTY {
		msg = utils.PHONE_IS_REQUIRED
	}
	if user.ZipCode == utils.EMPTY {
		msg = utils.ZIP_CODE_IS_REQUIRED
	}
	if user.StateId == 0 {
		msg = utils.STATE_ID_IS_REQUIRED
	}
	if user.RolId == 0 {
		msg = utils.ROL_ID_IS_REQUIRED
	}
	if user.Email == utils.EMPTY {
		msg = utils.EMAIL_IS_REQUIRED
	}

	return msg
}
func Pagination(c *fiber.Ctx, limit int) (int, int) {
	log.Println(c)
	pageParam := c.Query("page")

	if pageParam == "" {
		return 1, 0
	}
	page, _ := strconv.Atoi(c.Query("page"))
	if page < 1 {
		return 1, 0
	}

	begin := (limit * page) - limit
	return page, begin
}
