package models

import (
	"database/sql"
	"github.com/lib/pq"
)

type NamHoc struct {
	MaNamHoc uint32 `gorm:"primary_key" json:"ma_nam_hoc"`
	TuNgay   string `json:"tu_ngay"`
}

func (NamHoc) TableName() string {
	return "nam_hoc"
}

type DonVi struct {
	MaDonVi   uint32 `gorm:"primary_key" json:"ma_don_vi"`
	TenDonVi  string `json:"ten_don_vi"`
	LoaiDonVi string `json:"loai_don_vi"`
}

func (DonVi) TableName() string {
	return "don_vi"
}

type ChucVu struct {
	MaChucVu  uint32 `gorm:"primary_key" json:"ma_chuc_vu"`
	TenChucVu string `json:"ten_chuc_vu"`
}

func (ChucVu) TableName() string {
	return "chuc_vu"
}

type CanBo struct {
	MaCanBo    uint32 `gorm:"primary_key" json:"ma_can_bo"`
	HoTen      string `json:"ho_ten"`
	NgaySinh   string `json:"ngay_sinh"`
	QueQuan    string `json:"que_quan"`
	GioiTinh   bool   `json:"gioi_tinh"`
	DonVi      DonVi  `gorm:"association_foreignkey:MaDonVi"`
	MaDonVi    uint32 `json:"ma_don_vi,string,omitempty"`
	ChucVu     ChucVu
	MaChucVu   uint32 `json:"ma_chuc_vu,string,omitempty"`
	Quyen      Quyen
	MaQuyen    uint32         `json:"ma_quyen,string,omitempty"`
	Email      string         `json:"email"`
	Password   string         `json:"password"`
	AnhDaiDien sql.NullString `json:"anh_dai_dien"`
	TrangThai  bool           `json:"trang_thai"`
}

func (CanBo) TableName() string {
	return "can_bo"
}

type Quyen struct {
	MaQuyen  uint32 `gorm:"primary_key" json:"ma_quyen"`
	TenQuyen string `json:"ten_quyen"`
}

func (Quyen) TableName() string {
	return "quyen"
}

type DanhHieuThiDua struct {
	MaDanhHieu     uint32 `gorm:"primary_key" json:"ma_danh_hieu"`
	TenDanhHieu    string `json:"ten_danh_hieu"`
	DoiTuongApDung string `json:"doi_tuong_ap_dung"`
	TrangThai      bool   `json:"trang_thai"`
	DaBauChon      int32  `json:"da_bau_chon"`
	ThoiGianMo     int64  `json:"thoi_gian_mo"`
	TiLeDat        int8   `json:"ti_le_dat"`
	SoPhut         int    `json:"so_phut"`
	SlBoPhieu      int    `json:"sl_bo_phieu"`
}

func (DanhHieuThiDua) TableName() string {
	return "danh_hieu_thi_dua"
}

type HinhThucKhenThuong struct {
	MaHinhThuc  uint32 `gorm:"primary_key" json:"ma_hinh_thuc"`
	TenHinhThuc string `json:"ten_hinh_thuc"`
	TrangThai   bool   `json:"trang_thai"`
	DaBauChon   int32  `json:"da_bau_chon"`
	ThoiGianMo  int64  `json:"thoi_gian_mo"`
	TiLeDat     int8   `json:"ti_le_dat"`
	SoPhut      int    `json:"so_phut"`
	SlBoPhieu   int    `json:"sl_bo_phieu"`
}

func (HinhThucKhenThuong) TableName() string {
	return "hinh_thuc_khen_thuong"
}

type CanCu struct {
	MaCanCu     int64  `gorm:"primary_key" json:"ma_can_cu"`
	LoaiCanCu   string `json:"loai_can_cu"`
	SoHieuCanCu string `json:"so_hieu_can_cu"`
	NgayKy      string `json:"ngay_ky"`
	NoiBanHanh  string `json:"noi_ban_hanh"`
	TrichYeu    string `json:"trich_yeu"`
	TrangThai   int    `json:"trang_thai"`
}

func (CanCu) TableName() string {
	return "can_cu"
}

type HoiDong struct {
	MaHoiDong   string `gorm:"primary_key" json:"ma_hoi_dong"`
	MaNamHoc    uint32 `json:"ma_nam_hoc"`
	NamHoc      NamHoc
	NguoiTao    int `json:"nguoi_tao"`
	CanBo       CanBo
	NgayTao     int64  `json:"ngay_tao"`
	SoQuyetDinh string `json:"so_quyet_dinh"`
	NgayKy      string `json:"ngay_ky"`
	ThoiGianMo  int64  `json:"thoi_gian_mo"`
	TrangThai   int    `json:"trang_thai"`
}

func (HoiDong) TableName() string {
	return "hoi_dong"
}

type HoiDongCanCu struct {
	MaHoiDong string `json:"ma_hoi_dong"`
	MaCanCu   int64  `json:"ma_can_cu"`
	CanCu     CanCu
}

func (HoiDongCanCu) TableName() string {
	return "hoi_dong_can_cu"
}

type ThanhPhanHoiDong struct {
	MaThanhPhan uint32 `gorm:"primary_key" json:"ma_thanh_phan"`
	MaHoiDong   string `json:"ma_hoi_dong"`
	MaCanBo     uint32 `json:"ma_can_bo"`
	CanBo       CanBo
	MaChucDanh  uint32 `json:"ma_chuc_danh"`
	ChucDanh    ChucDanh
	CoMat       bool `json:"co_mat"`
	BauChon     bool `json:"bau_chon"`
}

func (ThanhPhanHoiDong) TableName() string {
	return "thanh_phan_hoi_dong"
}

type ChucDanh struct {
	MaChucDanh     int8   `gorm:"primary_key" json:"ma_chuc_danh"`
	TenChucDanh    string `json:"ten_chuc_danh"`
	ThamGiaBoPhieu bool   `json:"tham_gia_bo_phieu"`
}

func (ChucDanh) TableName() string {
	return "chuc_danh_hoi_dong"
}

type XetThiDuaTapThe struct {
	ID                  int8           `gorm:"primary_key" json:"id"`
	MaDonVi             uint32         `json:"ma_don_vi"`
	DonVi               DonVi          `json:"don_vi"`
	MaDanhHieu          uint32         `json:"ma_danh_hieu"`
	DanhHieuThiDua      DanhHieuThiDua `json:"danh_hieu"`
	MaHoiDong           string         `json:"ma_hoi_dong"`
	HoiDong             HoiDong        `json:"hoi_dong"`
	BauChonThiDuaTapThe []BauChonThiDuaTapThe
}

func (XetThiDuaTapThe) TableName() string {
	return "tap_the_xet_thi_dua"
}

type XetThiDuaCaNhan struct {
	ID                  int8           `gorm:"primary_key" json:"id"`
	MaCanBo             uint32         `json:"ma_can_bo"`
	CanBo               CanBo          `json:"can_bo"`
	MaDanhHieu          uint32         `json:"ma_danh_hieu"`
	DanhHieuThiDua      DanhHieuThiDua `json:"danh_hieu"`
	MaHoiDong           string         `json:"ma_hoi_dong"`
	HoiDong             HoiDong        `json:"hoi_dong"`
	BauChonThiDuaCaNhan []BauChonThiDuaCaNhan
}

func (XetThiDuaCaNhan) TableName() string {
	return "ca_nhan_xet_thi_dua"
}

type XetKhenThuongTapThe struct {
	ID                      int8               `gorm:"primary_key" json:"id"`
	MaDonVi                 uint32             `json:"ma_don_vi"`
	DonVi                   DonVi              `json:"don_vi"`
	MaHinhThuc              uint32             `json:"ma_hinh_thuc"`
	HinhThucKhenThuong      HinhThucKhenThuong `json:"hinh_thuc"`
	MaHoiDong               string             `json:"ma_hoi_dong"`
	HoiDong                 HoiDong            `json:"hoi_dong"`
	NoiDung                 string             `json:"noi_dung"`
	BauChonKhenThuongTapThe []BauChonKhenThuongTapThe
}

func (XetKhenThuongTapThe) TableName() string {
	return "tap_the_xet_khen_thuong"
}

type XetKhenThuongCaNhan struct {
	ID                      int8               `gorm:"primary_key" json:"id"`
	MaCanBo                 uint32             `json:"ma_can_bo"`
	CanBo                   CanBo              `json:"can_bo"`
	MaHinhThuc              uint32             `json:"ma_hinh_thuc"`
	HinhThucKhenThuong      HinhThucKhenThuong `json:"hinh_thuc"`
	MaHoiDong               string             `json:"ma_hoi_dong"`
	HoiDong                 HoiDong            `json:"hoi_dong"`
	NoiDung                 string             `json:"noi_dung"`
	BauChonKhenThuongCaNhan []BauChonKhenThuongCaNhan
}

func (XetKhenThuongCaNhan) TableName() string {
	return "ca_nhan_xet_khen_thuong"
}

type BauChonThiDuaTapThe struct {
	MaThanhPhan      int8 `json:"ma_thanh_phan"`
	CanBo            CanBo
	MaXet            int8 `json:"ma_xet"`
	XetThiDuaTapThe  XetThiDuaTapThe
	TrangThaiBauChon bool `json:"trang_thai_bau_chon"`
}

func (BauChonThiDuaTapThe) TableName() string {
	return "bau_chon_thi_dua_tap_the"
}

type BauChonThiDuaCaNhan struct {
	MaThanhPhan      int8 `json:"ma_thanh_phan"`
	CanBo            CanBo
	MaXet            int8 `json:"ma_xet"`
	XetThiDuaCaNhan  XetThiDuaCaNhan
	TrangThaiBauChon bool `json:"trang_thai_bau_chon"`
}

func (BauChonThiDuaCaNhan) TableName() string {
	return "bau_chon_thi_dua_ca_nhan"
}

type BauChonKhenThuongTapThe struct {
	MaThanhPhan         int8 `json:"ma_thanh_phan"`
	CanBo               CanBo
	MaXet               int8 `json:"ma_xet"`
	XetKhenThuongTapThe XetKhenThuongTapThe
	TrangThaiBauChon    bool `json:"trang_thai_bau_chon"`
}

func (BauChonKhenThuongTapThe) TableName() string {
	return "bau_chon_khen_thuong_tap_the"
}

type BauChonKhenThuongCaNhan struct {
	MaThanhPhan         int8 `json:"ma_thanh_phan"`
	CanBo               CanBo
	MaXet               int8 `json:"ma_xet"`
	XetKhenThuongCaNhan XetKhenThuongCaNhan
	TrangThaiBauChon    bool `json:"trang_thai_bau_chon"`
}

func (BauChonKhenThuongCaNhan) TableName() string {
	return "bau_chon_khen_thuong_ca_nhan"
}

type TapTheQuaTrinh struct {
	MaNamHoc           int8 `json:"ma_nam_hoc"`
	NamHoc             NamHoc
	MaDonVi            int8 `json:"ma_don_vi"`
	DonVi              DonVi
	MaHoiDong          string `json:"ma_hoi_dong"`
	HoiDong            HoiDong
	ListDanhHieu       pq.Int64Array `gorm:"type:int8[]" json:"list_danh_hieu"`
	DanhHieuThiDua     []DanhHieuThiDua
	ListKhenThuong     pq.Int64Array `gorm:"type:int8[]" json:"list_khen_thuong"`
	HinhThucKhenThuong []HinhThucKhenThuong
}

func (TapTheQuaTrinh) TableName() string {
	return "tapthe_quatrinh"
}

type CaNhanQuaTrinh struct {
	MaNamHoc           int8 `json:"ma_nam_hoc"`
	NamHoc             NamHoc
	MaCanBo            int8 `json:"ma_can_bo"`
	CanBo              CanBo
	MaHoiDong          string `json:"ma_hoi_dong"`
	HoiDong            HoiDong
	ListDanhHieu       pq.Int64Array `gorm:"type:int8[]" json:"list_danh_hieu"`
	DanhHieuThiDua     []DanhHieuThiDua
	ListKhenThuong     pq.Int64Array `gorm:"type:int8[]" json:"list_khen_thuong"`
	HinhThucKhenThuong []HinhThucKhenThuong
}

func (CaNhanQuaTrinh) TableName() string {
	return "canhan_quatrinh"
}
