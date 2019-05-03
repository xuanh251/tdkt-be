package routes

import (
	"TDKTServer/controllers"
	"TDKTServer/middlewares"

	//"TDKTServer/middlewares"

	//"TDKTServer/controllers"
	//"TDKTServer/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("./resources"))))
	r.HandleFunc("/api/auth/login", controllers.Login).Methods("POST")

	////DonVi route
	r.HandleFunc("/api/donvi/create", middlewares.IsAuth(controllers.CreateDonVi)).Methods("POST")
	r.HandleFunc("/api/listdonvi", middlewares.IsAuth(controllers.GetAllDonVi)).Methods("GET")
	r.HandleFunc("/api/donvi/{id}", middlewares.IsAuth(controllers.GetDonViById)).Methods("GET")
	r.HandleFunc("/api/donvi/{id}", middlewares.IsAuth(controllers.UpdateDonVi)).Methods("PUT")
	r.HandleFunc("/api/donvi/{id}", middlewares.IsAuth(controllers.DeleteDonVi)).Methods("DELETE")

	////ChucVu route
	r.HandleFunc("/api/chucvu/create", middlewares.IsAuth(controllers.CreateChucVu)).Methods("POST")
	r.HandleFunc("/api/listchucvu", middlewares.IsAuth(controllers.GetAllChucVu)).Methods("GET")
	r.HandleFunc("/api/chucvu/{id}", middlewares.IsAuth(controllers.GetChucVuById)).Methods("GET")
	r.HandleFunc("/api/chucvu/{id}", middlewares.IsAuth(controllers.UpdateChucVu)).Methods("PUT")
	r.HandleFunc("/api/chucvu/{id}", middlewares.IsAuth(controllers.DeleteChucVu)).Methods("DELETE")

	////CanBo route
	r.HandleFunc("/api/canbo/create", controllers.CreateCanBo).Methods("POST")
	r.HandleFunc("/api/listcanbo", middlewares.IsAuth(controllers.GetAllCanBos)).Methods("GET")
	r.HandleFunc("/api/canbo/{id}", middlewares.IsAuth(controllers.GetCanBoById)).Methods("GET")
	r.HandleFunc("/api/canbo/{id}", middlewares.IsAuth(controllers.UpdateCanBo)).Methods("PUT")
	r.HandleFunc("/api/canbo/{id}", middlewares.IsAuth(controllers.DeleteCanBo)).Methods("DELETE")

	r.HandleFunc("/api/canbo/setActive/{id}", middlewares.IsAuth(controllers.SetActive)).Methods("GET")
	r.HandleFunc("/api/canbo/setDeactive/{id}", middlewares.IsAuth(controllers.SetDeactive)).Methods("GET")

	//Quyen route
	r.HandleFunc("/api/getquyenlimit", middlewares.IsAuth(controllers.GetAllQuyen)).Methods("GET")
	r.HandleFunc("/api/getquyenbyid/{id}", middlewares.IsAuth(controllers.GetQuyenById)).Methods("GET")

	////DanhHieuThiDua route
	r.HandleFunc("/api/danhhieuthidua/create", middlewares.IsAuth(controllers.CreateDanhHieuThiDua)).Methods("POST")
	r.HandleFunc("/api/listdanhhieuthidua", middlewares.IsAuth(controllers.GetAllDanhHieuThiDua)).Methods("GET")
	r.HandleFunc("/api/danhhieuthidua/{id}", middlewares.IsAuth(controllers.GetDanhHieuThiDuaById)).Methods("GET")
	r.HandleFunc("/api/danhhieuthidua/{id}", middlewares.IsAuth(controllers.UpdateDanhHieuThiDua)).Methods("PUT")
	r.HandleFunc("/api/danhhieuthidua/{id}", middlewares.IsAuth(controllers.DeleteDanhHieuThiDua)).Methods("DELETE")

	//HinhThucKhenThuong route
	r.HandleFunc("/api/hinhthuckhenthuong/create", middlewares.IsAuth(controllers.CreateHinhThucKhenThuong)).Methods("POST")
	r.HandleFunc("/api/listhinhthuckhenthuong", middlewares.IsAuth(controllers.GetAllHinhThucKhenThuong)).Methods("GET")
	r.HandleFunc("/api/hinhthuckhenthuong/{id}", middlewares.IsAuth(controllers.GetHinhThucKhenThuongById)).Methods("GET")
	r.HandleFunc("/api/hinhthuckhenthuong/{id}", middlewares.IsAuth(controllers.UpdateHinhThucKhenThuong)).Methods("PUT")
	r.HandleFunc("/api/hinhthuckhenthuong/{id}", middlewares.IsAuth(controllers.DeleteHinhThucKhenThuong)).Methods("DELETE")

	//CanCu route
	r.HandleFunc("/api/cancu/create", middlewares.IsAuth(controllers.CreateCanCu)).Methods("POST")
	r.HandleFunc("/api/listcancu", middlewares.IsAuth(controllers.GetAllCanCuChung)).Methods("GET")
	r.HandleFunc("/api/listcancurieng", middlewares.IsAuth(controllers.GetAllCanCuRieng)).Methods("GET")
	r.HandleFunc("/api/cancu/{id}", middlewares.IsAuth(controllers.GetCanCuById)).Methods("GET")
	r.HandleFunc("/api/cancu/{id}", middlewares.IsAuth(controllers.UpdateCanCu)).Methods("PUT")
	r.HandleFunc("/api/cancu/{id}", middlewares.IsAuth(controllers.DeleteCanCu)).Methods("DELETE")

	//NamHoc route
	r.HandleFunc("/api/namhoc/create", middlewares.IsAuth(controllers.CreateNamHoc)).Methods("POST")
	r.HandleFunc("/api/namhoc/list", middlewares.IsAuth(controllers.GetAllNamHoc)).Methods("GET")

	//HoiDong route
	r.HandleFunc("/api/hoidong/getlisthoidong", middlewares.IsAuth(controllers.GetAllHoiDong)).Methods("GET")
	r.HandleFunc("/api/hoidong/getlasthoidong", middlewares.IsAuth(controllers.GetLastHoiDong)).Methods("GET")
	r.HandleFunc("/api/hoidong/gethoidong/{id}", middlewares.IsAuth(controllers.GetHoiDongById)).Methods("GET")
	r.HandleFunc("/api/hoidong/getlistthanhvienbyhoidong/{id}", middlewares.IsAuth(controllers.GetListThanhVien)).Methods("GET")
	r.HandleFunc("/api/hoidong/getlistcancubyhoidong/{id}", middlewares.IsAuth(controllers.GetListCanCuByHoiDong)).Methods("GET")
	r.HandleFunc("/api/hoidong/getlistthanhvien", middlewares.IsAuth(controllers.GetListThanhVienByHoiDong)).Methods("POST")
	r.HandleFunc("/api/hoidong/create", middlewares.IsAuth(controllers.CreateHoiDong)).Methods("POST")
	r.HandleFunc("/api/hoidong/capnhatcancu", middlewares.IsAuth(controllers.CapNhatCanCuHoiDong)).Methods("POST")
	r.HandleFunc("/api/hoidong/capnhatthanhvien", middlewares.IsAuth(controllers.CapNhatThanhVienHoiDong)).Methods("POST")

	r.HandleFunc("/api/hoidong/getlistdanhhieu/{id}", middlewares.IsAuth(controllers.GetListDanhHieuThiDuaByHoiDong)).Methods("GET")
	r.HandleFunc("/api/hoidong/getlistdanhhieubyhoidong/{id}", middlewares.IsAuth(controllers.GetListDHTDByHoiDong)).Methods("GET")
	r.HandleFunc("/api/hoidong/bauchon/active/{id}", middlewares.IsAuth(controllers.ActiveDanhHieu)).Methods("GET")
	r.HandleFunc("/api/hoidong/bauchon/deactive/{id}", middlewares.IsAuth(controllers.DeactiveDanhHieu)).Methods("GET")

	r.HandleFunc("/api/hoidong/open/{id}", middlewares.IsAuth(controllers.MoHoiDong)).Methods("GET")
	r.HandleFunc("/api/hoidong/close/{id}", middlewares.IsAuth(controllers.DongHoiDong)).Methods("GET")

	//ChucDanh route
	r.HandleFunc("/api/chucdanh/create", middlewares.IsAuth(controllers.CreateChucDanh)).Methods("POST")
	r.HandleFunc("/api/listchucdanh", middlewares.IsAuth(controllers.GetAllChucDanh)).Methods("GET")
	r.HandleFunc("/api/chucdanh/{id}", middlewares.IsAuth(controllers.GetChucDanhById)).Methods("GET")
	r.HandleFunc("/api/chucdanh/{id}", middlewares.IsAuth(controllers.UpdateChucDanh)).Methods("PUT")
	r.HandleFunc("/api/chucdanh/{id}", middlewares.IsAuth(controllers.DeleteChucDanh)).Methods("DELETE")

	//XetThiDua route
	r.HandleFunc("/api/xetthidua/tapthe/capnhat", middlewares.IsAuth(controllers.CapNhatXetThiDuaTapThe)).Methods("POST")
	r.HandleFunc("/api/xetthidua/tapthe/danhsach/{id}", middlewares.IsAuth(controllers.GetDanhSach_TDTT)).Methods("GET")
	r.HandleFunc("/api/xetthidua/tapthe/getobj/{id}", middlewares.IsAuth(controllers.GetObjById_TDTT)).Methods("GET")
	r.HandleFunc("/api/xetthidua/tapthe/{id}", middlewares.IsAuth(controllers.DeleteObj_TDTT)).Methods("DELETE")

	r.HandleFunc("/api/xetthidua/canhan/capnhat", middlewares.IsAuth(controllers.CapNhatXetThiDuaCaNhan)).Methods("POST")
	r.HandleFunc("/api/xetthidua/canhan/danhsach/{id}", middlewares.IsAuth(controllers.GetDanhSach_TDCN)).Methods("GET")
	r.HandleFunc("/api/xetthidua/canhan/getobj/{id}", middlewares.IsAuth(controllers.GetObjById_TDCN)).Methods("GET")
	r.HandleFunc("/api/xetthidua/canhan/{id}", middlewares.IsAuth(controllers.DeleteObj_TDCN)).Methods("DELETE")

	//XetKhenThuong route
	r.HandleFunc("/api/xetkhenthuong/tapthe/capnhat", middlewares.IsAuth(controllers.CapNhatXetKhenThuongTapThe)).Methods("POST")
	r.HandleFunc("/api/xetkhenthuong/tapthe/danhsach/{id}", middlewares.IsAuth(controllers.GetDanhSach_KTTT)).Methods("GET")
	r.HandleFunc("/api/xetkhenthuong/tapthe/getobj/{id}", middlewares.IsAuth(controllers.GetObjById_KTTT)).Methods("GET")
	r.HandleFunc("/api/xetkhenthuong/tapthe/{id}", middlewares.IsAuth(controllers.DeleteObj_KTTT)).Methods("DELETE")

	r.HandleFunc("/api/xetkhenthuong/canhan/capnhat", middlewares.IsAuth(controllers.CapNhatXetKhenThuongCaNhan)).Methods("POST")
	r.HandleFunc("/api/xetkhenthuong/canhan/danhsach/{id}", middlewares.IsAuth(controllers.GetDanhSach_KTCN)).Methods("GET")
	r.HandleFunc("/api/xetkhenthuong/canhan/getobj/{id}", middlewares.IsAuth(controllers.GetObjById_KTCN)).Methods("GET")
	r.HandleFunc("/api/xetkhenthuong/canhan/{id}", middlewares.IsAuth(controllers.DeleteObj_KTCN)).Methods("DELETE")

	//BauChon Route
	r.HandleFunc("/api/bauchon/getlistobjttbydanhhieu", middlewares.IsAuth(controllers.GetListBauChonThiDuaTapTheByDanhHieu)).Methods("POST")
	return r
}
