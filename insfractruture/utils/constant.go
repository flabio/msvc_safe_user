package utils

const (
	ID                           = "id"
	MESSAGE                      = "message"
	STATUS                       = "status"
	DATA                         = "data"
	USER                         = "user"
	CREATED                      = "was successfully created"
	UPDATED                      = "was updated successfully"
	REMOVED                      = "was successfully removed"
	ERROR_QUERY                  = "error query, please try again later"
	ERROR_CREATE                 = "error creating"
	ERROR_PARSING_BODY           = "error parsing body"
	ERROR_UPDATE                 = "error updating"
	ERROR_DELETE                 = "error deleting"
	ERROR_REQUIRED_FIELDS        = " is required."
	EMPTY                        = ""
	USER_NOT_FOUND        string = "username o password incorrect."
)
const (
	DB_DIFF_ID        = "id <>?"
	DB_EQUAL_ID       = "id=?"
	DB_EQUAL_NAME     = "name =?"
	DB_ORDER_DESC     = "id desc"
	DB_EQUAL_EMAIL_ID = "email=?"
)

const (
	ID_NO_EXIST         string = "The id not exists"
	NAME_ALREADY_EXIST  string = "Name already exists"
	EMAIL_ALREADY_EXIST string = "Email already exists"
	ROL_NOT_FOUND       string = "Rol not found"
	STATE_NOT_FOUND     string = "State not found"
)
const (
	FIRST_NAME      = "first_name"
	FIRST_SUR_NAME  = "first_sur_name"
	SECOND_SUR_NAME = "second_sur_name"
	ADDRESS         = "address"
	PHONE           = "phone"
	ZIP_CODE        = "zip_code"
	STATE_ID        = "state_id"
	ROL_ID          = "rol_id"
	EMAIL           = "email"
	PASSWORD        = "password"
)
const (
	FIRST_NAME_IS_FIELD_REQUIRED      = "The field first name is required."
	FIRST_SUR_NAME_IS_FIELD_REQUIRED  = "The field first sur name is required."
	SECOND_SUR_NAME_IS_FIELD_REQUIRED = "The field second sur name is required."
	ADDRESS_IS_FIELD_REQUIRED         = "The field address is required."
	PHONE_IS_FIELD_REQUIRED           = "The field phone is required."
	ZIP_CODE_IS_FIELD_REQUIRED        = "The field zip code is required."
	STATE_ID_IS_FIELD_REQUIRED        = "The field state id is required."
	ROL_ID_IS_FIELD_REQUIRED          = "The field rol id is required."
	EMAIL_IS_FIELD_REQUIRED           = "The field email is required."
	PASSWORD_IS_FIELD_REQUIRED        = "The field password is required."
)
const (
	FIRST_NAME_IS_MIN_LENGTH_REQUIRED      = "The field first name must have a minimum length of 3 characters."
	FIRST_SUR_NAME_IS_MIN_LENGTH_REQUIRED  = "The field first sur name must have a minimum length of 3 characters."
	SECOND_SUR_NAME_IS_MIN_LENGTH_REQUIRED = "The field second sur name must have a minimum length of 3 characters."
	ADDRESS_IS_MIN_LENGTH_REQUIRED         = "The field address must have a minimum length of 10 characters."
	PHONE_IS_MIN_LENGTH_REQUIRED           = "The field phone must have a minimum length of 10 characters."
	ZIP_CODE_IS_MIN_LENGTH_REQUIRED        = "The field zip code must have a minimum length of 5 characters."
	STATE_ID_IS_MIN_LENGTH_REQUIRED        = "The field state id "
)
const (
	FIRST_NAME_IS_REQUIRED      = "The first name is required."
	FIRST_SUR_NAME_IS_REQUIRED  = "The first sur name is required."
	SECOND_SUR_NAME_IS_REQUIRED = "The second sur name is required."
	ADDRESS_IS_REQUIRED         = "The address is required."
	PHONE_IS_REQUIRED           = "The phone is required."
	ZIP_CODE_IS_REQUIRED        = "The zip code is required."
	STATE_ID_IS_REQUIRED        = "The state id is required."
	ROL_ID_IS_REQUIRED          = "The rol id is required."
	EMAIL_IS_REQUIRED           = "The email is required."
	PASSWORD_IS_REQUIRED        = "The password is required."
	EMAIL_IS_INVALID            = "The email is invalid."
	PASSWORD_IS_INVALID_LENGTH  = "The password must have a minimum length of 8 characters."
)

const (
	MSVC_ROL_URL               string = "http://localhost:3001/api/rol"
	MSVC_EMERGENCY_CONTACT_URL string = "http://localhost:3003/api/emergency"
	MSVC_STATES_URL            string = "http://localhost:3004/api/states"
	MSVC_SCHOOL_URL            string = "http://localhost:3006/api/school"
	MSVC_COURSE_URL            string = "http://localhost:3007/api/course"
)
const (
	GET           string = "GET"
	POST          string = "POST"
	PUT           string = "PUT"
	DELETE        string = "DELETE"
	AUTHORIZATION string = "Authorization"
	BEARER        string = "Bearer "
)
const (
	TOKEN_INVALID string = "Token not provided or invalid"
)
