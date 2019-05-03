package repository

//var (
//	ErrAccountNotFound = errors.New("Account not found!")
//)
//
//func NewAccount(account models.Account) (bool, error) {
//	con := driver.Connect()
//	defer con.Close()
//	var timeNow = int64(time.Now().Unix()) * 1000
//
//	hashedPassword, err := utils.Bcrypt(account.Password)
//	if err != nil {
//		return false, err
//	}
//	account.Password= string(hashedPassword)
//	addRc:=con.NewRecord(account)
//	if addRc {
//		err = con.Create(&account).Error
//		if err != nil {
//			return false, err
//		}
//	}
//	return true, nil
//}
//func GetAllAccounts() ([]models.Account, error) {
//	con := driver.Connect()
//	defer con.Close()
//	var accounts []models.Account
//	err := con.Find(&accounts).Error
//	if err != nil {
//		return nil, err
//	}
//	return accounts, nil
//}
//
//func GetAccountByID(id uint32) (models.Account, error) {
//	con := driver.Connect()
//	defer con.Close()
//	acc := models.Account{}
//	err:=con.First(&acc,id).Error
//	if err!=nil {
//		return models.Account{}, err
//	}
//	return acc, nil
//}
//func GetAccountByEmail(email string) (models.Account, error) {
//	con := driver.Connect()
//	defer con.Close()
//	acc := models.Account{}
//	err:=con.Where("email=?",email).First(&acc).Error
//	if err!=nil {
//		return models.Account{}, err
//	}
//	return acc, nil
//}
//func UpdateAccount(account models.Account) (int64, error) {
//	con := driver.Connect()
//	defer con.Close()
//	acc2Update,err:=GetAccountByID(account.ID)
//	if err!=nil {
//		return 0, err
//	}
//
//	account.CreatedAt=acc2Update.CreatedAt
//
//	account.UpdatedAt=timeNow
//
//	err=con.Save(&account).Error
//	if err != nil {
//		return 0, err
//	}
//	return con.RowsAffected, nil
//}
//func DeleteAccount(uid uint32) (int64, error) {
//	con := driver.Connect()
//	defer con.Close()
//	acc, err:=GetAccountByID(uid)
//	if err!=nil {
//		return 0,err
//	}
//	err=con.Delete(&acc).Error
//	if err != nil {
//		return 0, err
//	}
//	return con.RowsAffected, nil
//}
//func UploadUserPhoto(file multipart.File, handle *multipart.FileHeader, uid uint32) (int64, error) {
//	con := driver.Connect()
//	defer con.Close()
//	account, err := GetAccountByID(uid)
//	if err != nil {
//		return 0, ErrAccountNotFound
//	}
//	fmt.Printf("Uploading file: %+v\n", handle.Filename);
//	fmt.Printf("File size: %+v\n", handle.Size);
//	fmt.Printf("MIME header: %+v\n", handle.Header);
//	tempFile, err := ioutil.TempFile("resources/images", "upload-*.png")
//	if err != nil {
//		fmt.Println(err)
//		return 0, ErrAccountNotFound
//		return 0, ErrAccountNotFound
//	}
//	defer tempFile.Close()
//
//	fileBytes, err := ioutil.ReadAll(file)
//	if err != nil {
//		fmt.Println(err)
//		return 0, ErrAccountNotFound
//	}
//	tempFile.Write(fileBytes)
//	sql := "update accounts set photo_url=$1 where uid=$2"
//	stmt, err := con.Prepare(sql)
//	if err != nil {
//		return 0, err
//	}
//	defer stmt.Close()
//	rs, err := stmt.Exec(handle.Filename, account.UID)
//	if err != nil {
//		return 0, err
//	}
//	return rs.RowsAffected()
//}
