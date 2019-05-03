package controllers

//func CreateAccount(w http.ResponseWriter, r *http.Request) {
//	body, _ := ioutil.ReadAll(r.Body)
//	var user models.Account
//	err := json.Unmarshal(body, &user)
//	if err != nil {
//		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
//		return
//	}
//	user, err = validations.ValidationNewUser(user)
//	if err != nil {
//		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
//		return
//	}
//	_, err = repository.NewAccount(user)
//	if err != nil {
//		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
//		return
//	}
//	w.WriteHeader(http.StatusCreated)
//	utils.ToJson(w, "Tạo tài khoản thành công!")
//}
//func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
//	accounts, err := repository.GetAllAccounts()
//	if err != nil {
//		utils.ErrorResponse(w, err, http.StatusBadRequest)
//		return
//	}
//	utils.ToJson(w, accounts)
//}
//func GetAccountById(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	id, _ := strconv.ParseInt(params["uid"], 10, 12)
//	account, err := repository.GetAccountByID(uint32(id))
//	if err != nil {
//		utils.ErrorResponse(w, err, http.StatusBadRequest)
//		return
//	}
//	utils.ToJson(w, account)
//}
//func UpdateAccount(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	uid, _ := strconv.ParseInt(params["uid"], 10, 12)
//	body, _ := ioutil.ReadAll(r.Body)
//	var account models.Account
//	err := json.Unmarshal(body, &account)
//	if err != nil {
//		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
//		return
//	}
//	account.ID = uint32(uid)
//	_, err = repository.UpdateAccount(account)
//	if err != nil {
//		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
//		return
//	}
//	utils.ToJson(w, "Cập nhật thành công!")
//}
//func DeleteAccount(w http.ResponseWriter, r *http.Request) {
//	params := mux.Vars(r)
//	uid, _ := strconv.ParseInt(params["uid"], 10, 12)
//	_, err := repository.DeleteAccount(uint32(uid))
//	if err!=nil{
//		utils.ErrorResponse(w,err,http.StatusUnprocessableEntity)
//		return
//	}
//	utils.ToJson(w, "Xoá dữ liệu thành công!")
//}
