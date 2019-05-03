create database dqu.tdkt;
create database dqu.tdkt;

create table if not exists don_vi
(
    ma_don_vi   bigserial primary key,
    ten_don_vi  varchar(255) not null,
    loai_don_vi varchar(255) not null
);

create table if not exists chuc_vu
(
    ma_chuc_vu  bigserial primary key,
    ten_chuc_vu varchar(255) not null
);

create table if not exists nam_hoc
(
    ma_nam_hoc bigserial primary key,
    tu_ngay    date not null
);

create table if not exists can_bo
(
    ma_can_bo     bigserial primary key,
    ho_va_chu_lot varchar(255) not null,
    ten           varchar(255) not null,
    ngay_sinh     date         not null,
    que_quan      varchar(255),
    gioi_tinh     boolean      not null,
    don_vi        bigint       not null,
    chuc_vu       bigint       not null,
    last_updated  bigint       not null,
    constraint canbo_donvi_fk foreign key (don_vi)
        references don_vi (ma_don_vi)
        on delete cascade
        on update cascade,
    constraint canbo_chucvu_fk foreign key (chuc_vu)
        references chuc_vu (ma_chuc_vu)
        on delete cascade
        on update cascade,
);
create table if not exists danh_hieu_thi_dua
(
    ma_danh_hieu      bigserial primary key,
    ten_danh_hieu     varchar(255) not null,
    doi_tuong_ap_dung varchar(255) not null
);

create table if not exists hinh_thuc_khen_thuong
(
    ma_hinh_thuc  bigserial primary key,
    ten_hinh_thuc varchar(255) not null,
);


create table if not exists chuc_danh_hoi_dong
(
    ma_chuc_danh      bigserial primary key,
    ten_chuc_danh     varchar(255),
    tham_gia_bo_phieu bool default false
);
create table if not exists can_cu
(
    ma_can_cu      bigserial primary key,
    loai_can_cu    integer not null,
    so_hieu_can_cu varchar(255),
    ngay_ky        date,
    noi_ban_hanh   varchar(255),
    trich_yeu      varchar(255),
    trangthai      integer not null
);
create table if not exists hoi_dong
(
    ma_hoi_dong   varchar(255) primary key unique,
    ma_nam_hoc    bigint not null,
    nguoi_tao     bigint not null,
    ngay_tao      bigint not null,
    so_quyet_dinh varchar(255),
    ngay_ky       date,
    thoi_gian_mo  bigint not null,
    trang_thai    int,
    constraint hoidong_namhoc_fk foreign key (ma_nam_hoc)
        references nam_hoc (ma_nam_hoc)
        on delete cascade
        on update cascade,
    constraint hoidong_namhoc_fk foreign key (ma_nam_hoc)
        references nam_hoc (ma_nam_hoc)
        on delete cascade
        on update cascade,
    constraint hoidong_canbo_fk foreign key (nguoi_tao)
        references accounts (id)
        on delete cascade
        on update cascade
);

create table if not exists hoi_dong_can_cu
(
    ma_hoi_dong varchar(255),
    ma_can_cu   varchar(255),
    primary key (ma_can_cu, ma_hoi_dong),
    constraint hoidong_cancu_fk foreign key (ma_hoi_dong)
        references hoi_dong (ma_hoi_dong)
        on delete cascade
        on update cascade,
    constraint cancu_fk foreign key (ma_can_cu)
        references can_cu (ma_can_cu)
        on delete cascade
        on update cascade
);

create table if not exists thanh_phan_hoi_dong
(
    ma_thanh_phan bigserial primary key,
    ma_hoi_dong   varchar(255) not null,
    ma_can_bo     bigint       not null,
    ma_chuc_danh  bigint       not null,
    constraint hoidong_thanhphan_fk foreign key (ma_hoi_dong)
        references hoi_dong (ma_hoi_dong)
        on delete cascade
        on update cascade,
    constraint canbo_thanhphanhoidong_fk foreign key (ma_can_bo)
        references can_bo (ma_can_bo)
        on delete cascade
        on update cascade,
    constraint chucdanh_thanhphanhoidong_fk foreign key (ma_chuc_danh)
        references chuc_danh_hoi_dong (ma_chuc_danh)
        on delete cascade
        on update cascade
);

create table if not exists tap_the_xet_thi_dua
(
    id                bigserial primary key,
    ma_don_vi         bigint       not null,
    ma_hoi_dong       varchar(255) not null,
    danh_hieu_dang_ky bigint       not null,
    ti_le_dat         int          not null,
    constraint xetthiduatt_donvi_fk foreign key (ma_don_vi)
        references don_vi (ma_don_vi)
        on delete cascade
        on update cascade,
    constraint xetthiduatt_hoidong_fk foreign key (ma_hoi_dong)
        references hoi_dong (ma_hoi_dong)
        on delete cascade
        on update cascade,
    constraint xetthiduatt_danhhieu_fk foreign key (danh_hieu_dang_ky)
        references danh_hieu_thi_dua (ma_danh_hieu)
        on delete cascade
        on update cascade
);

create table if not exists ca_nhan_xet_thi_dua
(
    id                bigserial primary key,
    ma_can_bo         bigint       not null,
    ma_hoi_dong       varchar(255) not null,
    danh_hieu_dang_ky bigint       not null,
    ti_le_dat         int          not null,
    constraint xetthiduacn_canbo_fk foreign key (ma_can_bo)
        references can_bo (ma_can_bo)
        on delete cascade
        on update cascade,
    constraint xetthiduacn_hoidong_fk foreign key (ma_hoi_dong)
        references hoi_dong (ma_hoi_dong)
        on delete cascade
        on update cascade,
    constraint xetthiduacn_danhhieu_fk foreign key (danh_hieu_dang_ky)
        references danh_hieu_thi_dua (ma_danh_hieu)
        on delete cascade
        on update cascade
);
create table if not exists tap_the_xet_khen_thuong
(
    id                bigserial primary key,
    ma_don_vi         bigint       not null,
    ma_hoi_dong       varchar(255) not null,
    hinh_thuc_dang_ky bigint       not null,
    ti_le_dat         int          not null,
    constraint xetkhenthuongtt_donvi_fk foreign key (ma_don_vi)
        references don_vi (ma_don_vi)
        on delete cascade
        on update cascade,
    constraint xetkhenthuongtt_hoidong_fk foreign key (ma_hoi_dong)
        references hoi_dong (ma_hoi_dong)
        on delete cascade
        on update cascade,
    constraint xetkhenthuongtt_hinhthuc_fk foreign key (hinh_thuc_dang_ky)
        references hinh_thuc_khen_thuong (ma_hinh_thuc)
        on delete cascade
        on update cascade
);
create table if not exists ca_nhan_xet_khen_thuong
(
    id                bigserial primary key,
    ma_can_bo         bigint       not null,
    ma_hoi_dong       varchar(255) not null,
    hinh_thuc_dang_ky bigint       not null,
    ti_le_dat         int          not null,
    constraint xetkhenthuongcn_canbo_fk foreign key (ma_can_bo)
        references can_bo (ma_can_bo)
        on delete cascade
        on update cascade,
    constraint xetkhenthuongcn_hoidong_fk foreign key (ma_hoi_dong)
        references hoi_dong (ma_hoi_dong)
        on delete cascade
        on update cascade,
    constraint xetkhenthuongcn_hinhthuc_fk foreign key (hinh_thuc_dang_ky)
        references hinh_thuc_khen_thuong (ma_hinh_thuc)
        on delete cascade
        on update cascade
)
create table bau_chon_thi_dua_tap_the
(
    ma_thanh_phan       varchar(255) not null,
    ma_xet              bigint       not null,
    trang_thai_bau_chon int          not null,
    constraint thanhphan_bauchon_fk foreign key (ma_thanh_phan)
        references thanh_phan_hoi_dong (ma_thanh_phan)
        on delete cascade
        on update cascade,
    constraint xetdanhhieutt_bauchon_fk foreign key (ma_xet)
        references tap_the_xet_thi_dua (id)
        on delete cascade
        on update cascade
);
create table bau_chon_thi_dua_ca_nhan
(
    ma_thanh_phan       varchar(255) not null,
    ma_xet              bigint       not null,
    trang_thai_bau_chon int          not null,
    constraint thanhphan_bauchon_fk foreign key (ma_thanh_phan)
        references thanh_phan_hoi_dong (ma_thanh_phan)
        on delete cascade
        on update cascade,
    constraint xetdanhhieucn_bauchon_fk foreign key (ma_xet)
        references ca_nhan_xet_thi_dua (id)
        on delete cascade
        on update cascade
);




create table bau_chon_khen_thuong_tap_the
(
    ma_thanh_phan       varchar(255) not null,
    ma_xet        bigint       not null,
    trang_thai_bau_chon int          not null,
    constraint thanhphan_bauchon_fk foreign key (ma_thanh_phan)
        references thanh_phan_hoi_dong (ma_thanh_phan)
        on delete cascade
        on update cascade,
    constraint xethinhthuctt_bauchon_fk foreign key (ma_xet)
        references tap_the_xet_khen_thuong (id)
        on delete cascade
        on update cascade
)
create table bau_chon_khen_thuong_ca_nhan
(
    ma_thanh_phan       varchar(255) not null,
    ma_xet        bigint       not null,
    trang_thai_bau_chon int          not null,
    constraint thanhphan_bauchon_fk foreign key (ma_thanh_phan)
        references thanh_phan_hoi_dong (ma_thanh_phan)
        on delete cascade
        on update cascade,
    constraint xethinhthuccn_bauchon_fk foreign key (ma_xet)
        references ca_nhan_xet_khen_thuong (id)
        on delete cascade
        on update cascade
)


CREATE TABLE if not EXISTS tapthe_quatrinh
(
    ma_nam_hoc       bigint not null,
    ma_don_vi        bigint not null,
    ma_hoi_dong      varchar(255),
    list_danh_hieu   integer[],
    list_khen_thuong integer[],
    primary key (ma_nam_hoc, ma_don_vi),
    constraint namhoc_quatrinhfk foreign key (ma_nam_hoc)
        references nam_hoc (ma_nam_hoc)
        on delete cascade
        on update cascade,
    constraint donvi_quatrinhfk foreign key (ma_don_vi)
        references don_vi (ma_don_vi)
        on delete cascade
        on update cascade
)
CREATE TABLE if not EXISTS canhan_quatrinh
(
    id               bigserial primary key,
    ma_nam_hoc       bigint not null,
    ma_can_bo        bigint not null,
    list_danh_hieu   integer[],
    list_khen_thuong integer[],
    ma_hoi_dong      varchar(255),
    primary key (ma_nam_hoc, ma_can_bo),
    constraint namhoc_quatrinhfk foreign key (ma_nam_hoc)
        references nam_hoc (ma_nam_hoc)
        on delete cascade
        on update cascade,
    constraint canhan_quatrinhfk foreign key (ma_can_bo)
        references don_vi (ma_can_bo)
        on delete cascade
        on update cascade
)
CREATE TABLE if not EXISTS phan_quyen
(
    ma_can_bo int8 not null,
    ma_quyen  int8 not null,
    constraint canbo_phanquyen_fk foreign key (ma_can_bo)
        references can_bo (ma_can_bo)
        on delete cascade
        on update cascade,
    constraint quyen_phanquyen_fk foreign key (ma_quyen)
        references quyen (ma_quyen)
        on delete cascade
        on update cascade
)

CREATE table if not EXISTS quyen
(
    ma_quyen  bigserial primary key,
    ten_quyen varchar(255) not null UNIQUE
)
