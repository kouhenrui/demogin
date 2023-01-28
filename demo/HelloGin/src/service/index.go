package service

//var user pojo.UserInterface
//var admin pojo.AdminInterface
//
//type Login struct {
//}
//
//func LoginService() Login {
//	return Login{}
//}
//func (l *Login) LoginByAccount(account string, password string) (bool, string) {
//	u, ue := user.FindByAccount(account)
//	if ue != nil {
//		log.Println("查询数据库", u)
//		return false, util.ACCOUT_NOT_EXIST_ERROR
//	}
//	log.Println("查询数据库,weichucuoi", u)
//
//	enpwd := u.Password
//	salt := u.Salt
//	pwd, deerr := util.DePwdCode(enpwd, salt)
//	if deerr != nil {
//		return false, util.PASSWORD_RESOLUTION_ERROR
//	}
//	if pwd != password {
//		return false, util.AUTH_LOGIN_PASSWORD_ERROR
//	}
//	return true, util.SUCCESS
//}

//func validateExistWhere(account string) {
//	switch expr {
//
//	}
//}

//type CurrencyService struct {
//}
//
//func NewCommonService() CurrencyService {
//	return CurrencyService{}
//}

//func (c *CurrencyService) Login(loginParam ) bool {
//	go func() {
//		res:=validateExist(loginParam)
//	}()
//	return true
//}
//
//func validateExist(param struct{}) interface{} {
//	var (
//		name = param.name
//		password=param.password
//		account=param.account
//		type=param.type
//	)
//	if type=="account" {
//
//	}
//
//}
