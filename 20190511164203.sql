/*
PostgreSQL Backup
Database: dqu_tdkt/public
Backup Time: 2019-05-11 16:42:06
*/

DROP SEQUENCE IF EXISTS "public"."ca_nhan_xet_khen_thuong_id_seq";
DROP SEQUENCE IF EXISTS "public"."ca_nhan_xet_thi_dua_id_seq";
DROP SEQUENCE IF EXISTS "public"."can_bo_ma_can_bo_seq";
DROP SEQUENCE IF EXISTS "public"."can_cu_ma_can_cu_seq";
DROP SEQUENCE IF EXISTS "public"."chuc_danh_hoi_dong_ma_chuc_danh_seq";
DROP SEQUENCE IF EXISTS "public"."chuc_vu_ma_chuc_vu_seq";
DROP SEQUENCE IF EXISTS "public"."danh_hieu_thi_dua_ma_danh_hieu_seq";
DROP SEQUENCE IF EXISTS "public"."don_vi_ma_don_vi_seq";
DROP SEQUENCE IF EXISTS "public"."hinh_thuc_khen_thuong_ma_hinh_thuc_seq";
DROP SEQUENCE IF EXISTS "public"."nam_hoc_ma_nam_hoc_seq";
DROP SEQUENCE IF EXISTS "public"."quyen_ma_quyen_seq";
DROP SEQUENCE IF EXISTS "public"."tap_the_xet_khen_thuong_id_seq";
DROP SEQUENCE IF EXISTS "public"."tap_the_xet_thi_dua_id_seq";
DROP SEQUENCE IF EXISTS "public"."thanh_phan_hoi_dong_ma_thanh_phan_seq";
DROP TABLE IF EXISTS "public"."bau_chon_khen_thuong_ca_nhan";
DROP TABLE IF EXISTS "public"."bau_chon_khen_thuong_tap_the";
DROP TABLE IF EXISTS "public"."bau_chon_thi_dua_ca_nhan";
DROP TABLE IF EXISTS "public"."bau_chon_thi_dua_tap_the";
DROP TABLE IF EXISTS "public"."ca_nhan_xet_khen_thuong";
DROP TABLE IF EXISTS "public"."ca_nhan_xet_thi_dua";
DROP TABLE IF EXISTS "public"."can_bo";
DROP TABLE IF EXISTS "public"."can_cu";
DROP TABLE IF EXISTS "public"."canhan_quatrinh";
DROP TABLE IF EXISTS "public"."chuc_danh_hoi_dong";
DROP TABLE IF EXISTS "public"."chuc_vu";
DROP TABLE IF EXISTS "public"."danh_hieu_thi_dua";
DROP TABLE IF EXISTS "public"."don_vi";
DROP TABLE IF EXISTS "public"."hinh_thuc_khen_thuong";
DROP TABLE IF EXISTS "public"."hoi_dong";
DROP TABLE IF EXISTS "public"."hoi_dong_can_cu";
DROP TABLE IF EXISTS "public"."nam_hoc";
DROP TABLE IF EXISTS "public"."quyen";
DROP TABLE IF EXISTS "public"."tap_the_xet_khen_thuong";
DROP TABLE IF EXISTS "public"."tap_the_xet_thi_dua";
DROP TABLE IF EXISTS "public"."tapthe_quatrinh";
DROP TABLE IF EXISTS "public"."thanh_phan_hoi_dong";
CREATE SEQUENCE "ca_nhan_xet_khen_thuong_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "ca_nhan_xet_thi_dua_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "can_bo_ma_can_bo_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "can_cu_ma_can_cu_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "chuc_danh_hoi_dong_ma_chuc_danh_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "chuc_vu_ma_chuc_vu_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "danh_hieu_thi_dua_ma_danh_hieu_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "don_vi_ma_don_vi_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "hinh_thuc_khen_thuong_ma_hinh_thuc_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "nam_hoc_ma_nam_hoc_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "quyen_ma_quyen_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "tap_the_xet_khen_thuong_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "tap_the_xet_thi_dua_id_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE SEQUENCE "thanh_phan_hoi_dong_ma_thanh_phan_seq" 
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
CREATE TABLE "bau_chon_khen_thuong_ca_nhan" (
  "ma_thanh_phan" int8 NOT NULL,
  "ma_xet" int8 NOT NULL,
  "trang_thai_bau_chon" bool NOT NULL
)
;
ALTER TABLE "bau_chon_khen_thuong_ca_nhan" OWNER TO "postgres";
CREATE TABLE "bau_chon_khen_thuong_tap_the" (
  "ma_thanh_phan" int8 NOT NULL,
  "ma_xet" int8 NOT NULL,
  "trang_thai_bau_chon" bool NOT NULL
)
;
ALTER TABLE "bau_chon_khen_thuong_tap_the" OWNER TO "postgres";
CREATE TABLE "bau_chon_thi_dua_ca_nhan" (
  "ma_thanh_phan" int8 NOT NULL,
  "ma_xet" int8 NOT NULL,
  "trang_thai_bau_chon" bool NOT NULL
)
;
ALTER TABLE "bau_chon_thi_dua_ca_nhan" OWNER TO "postgres";
CREATE TABLE "bau_chon_thi_dua_tap_the" (
  "ma_thanh_phan" int8 NOT NULL,
  "ma_xet" int8 NOT NULL,
  "trang_thai_bau_chon" bool NOT NULL
)
;
ALTER TABLE "bau_chon_thi_dua_tap_the" OWNER TO "postgres";
CREATE TABLE "ca_nhan_xet_khen_thuong" (
  "id" int8 NOT NULL DEFAULT nextval('ca_nhan_xet_khen_thuong_id_seq'::regclass),
  "ma_can_bo" int8 NOT NULL,
  "ma_hoi_dong" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ma_hinh_thuc" int8 NOT NULL,
  "noi_dung" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "ca_nhan_xet_khen_thuong" OWNER TO "postgres";
CREATE TABLE "ca_nhan_xet_thi_dua" (
  "id" int8 NOT NULL DEFAULT nextval('ca_nhan_xet_thi_dua_id_seq'::regclass),
  "ma_can_bo" int8 NOT NULL,
  "ma_hoi_dong" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ma_danh_hieu" int8 NOT NULL
)
;
ALTER TABLE "ca_nhan_xet_thi_dua" OWNER TO "postgres";
CREATE TABLE "can_bo" (
  "ma_can_bo" int8 NOT NULL DEFAULT nextval('can_bo_ma_can_bo_seq'::regclass),
  "ho_ten" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ngay_sinh" date NOT NULL,
  "que_quan" varchar(255) COLLATE "pg_catalog"."default",
  "gioi_tinh" bool,
  "ma_don_vi" int8 NOT NULL,
  "ma_chuc_vu" int8 NOT NULL,
  "email" varchar(255) COLLATE "pg_catalog"."default",
  "anh_dai_dien" varchar(255) COLLATE "pg_catalog"."default",
  "password" varchar(255) COLLATE "pg_catalog"."default",
  "ma_quyen" int8,
  "trang_thai" bool
)
;
ALTER TABLE "can_bo" OWNER TO "postgres";
CREATE TABLE "can_cu" (
  "ma_can_cu" int8 NOT NULL DEFAULT nextval('can_cu_ma_can_cu_seq'::regclass),
  "loai_can_cu" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "so_hieu_can_cu" varchar(255) COLLATE "pg_catalog"."default",
  "ngay_ky" date,
  "noi_ban_hanh" varchar(255) COLLATE "pg_catalog"."default",
  "trich_yeu" varchar(255) COLLATE "pg_catalog"."default",
  "trang_thai" int4 NOT NULL DEFAULT 0
)
;
ALTER TABLE "can_cu" OWNER TO "postgres";
CREATE TABLE "canhan_quatrinh" (
  "ma_nam_hoc" int8 NOT NULL,
  "ma_can_bo" int8 NOT NULL,
  "list_danh_hieu" int4[],
  "list_khen_thuong" int4[],
  "ma_hoi_dong" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "canhan_quatrinh" OWNER TO "postgres";
CREATE TABLE "chuc_danh_hoi_dong" (
  "ma_chuc_danh" int8 NOT NULL DEFAULT nextval('chuc_danh_hoi_dong_ma_chuc_danh_seq'::regclass),
  "ten_chuc_danh" varchar(255) COLLATE "pg_catalog"."default",
  "tham_gia_bo_phieu" bool DEFAULT false
)
;
ALTER TABLE "chuc_danh_hoi_dong" OWNER TO "postgres";
CREATE TABLE "chuc_vu" (
  "ma_chuc_vu" int8 NOT NULL DEFAULT nextval('chuc_vu_ma_chuc_vu_seq'::regclass),
  "ten_chuc_vu" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "chuc_vu" OWNER TO "postgres";
CREATE TABLE "danh_hieu_thi_dua" (
  "ma_danh_hieu" int8 NOT NULL DEFAULT nextval('danh_hieu_thi_dua_ma_danh_hieu_seq'::regclass),
  "ten_danh_hieu" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "doi_tuong_ap_dung" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "trang_thai" bool DEFAULT false,
  "da_bau_chon" int2 DEFAULT 0,
  "thoi_gian_mo" int8 DEFAULT 0,
  "ti_le_dat" int8,
  "so_phut" int2 DEFAULT 0,
  "sl_bo_phieu" int2 DEFAULT 0
)
;
ALTER TABLE "danh_hieu_thi_dua" OWNER TO "postgres";
CREATE TABLE "don_vi" (
  "ma_don_vi" int8 NOT NULL DEFAULT nextval('don_vi_ma_don_vi_seq'::regclass),
  "ten_don_vi" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "loai_don_vi" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "don_vi" OWNER TO "postgres";
CREATE TABLE "hinh_thuc_khen_thuong" (
  "ma_hinh_thuc" int8 NOT NULL DEFAULT nextval('hinh_thuc_khen_thuong_ma_hinh_thuc_seq'::regclass),
  "ten_hinh_thuc" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "trang_thai" bool DEFAULT false,
  "da_bau_chon" int4 DEFAULT 0,
  "thoi_gian_mo" int8 DEFAULT 0,
  "ti_le_dat" int8,
  "so_phut" int8 DEFAULT 0,
  "sl_bo_phieu" int2 DEFAULT 0
)
;
ALTER TABLE "hinh_thuc_khen_thuong" OWNER TO "postgres";
CREATE TABLE "hoi_dong" (
  "ma_hoi_dong" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ma_nam_hoc" int8 NOT NULL,
  "nguoi_tao" int8 NOT NULL,
  "ngay_tao" int8 NOT NULL,
  "so_quyet_dinh" varchar(255) COLLATE "pg_catalog"."default",
  "ngay_ky" date,
  "thoi_gian_mo" int8,
  "trang_thai" int4
)
;
ALTER TABLE "hoi_dong" OWNER TO "postgres";
CREATE TABLE "hoi_dong_can_cu" (
  "ma_hoi_dong" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ma_can_cu" int8 NOT NULL
)
;
ALTER TABLE "hoi_dong_can_cu" OWNER TO "postgres";
CREATE TABLE "nam_hoc" (
  "ma_nam_hoc" int8 NOT NULL DEFAULT nextval('nam_hoc_ma_nam_hoc_seq'::regclass),
  "tu_ngay" date NOT NULL
)
;
ALTER TABLE "nam_hoc" OWNER TO "postgres";
CREATE TABLE "quyen" (
  "ma_quyen" int8 NOT NULL DEFAULT nextval('quyen_ma_quyen_seq'::regclass),
  "ten_quyen" varchar(255) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "quyen" OWNER TO "postgres";
CREATE TABLE "tap_the_xet_khen_thuong" (
  "id" int8 NOT NULL DEFAULT nextval('tap_the_xet_khen_thuong_id_seq'::regclass),
  "ma_don_vi" int8 NOT NULL,
  "ma_hoi_dong" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ma_hinh_thuc" int8 NOT NULL,
  "noi_dung" varchar(255) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "tap_the_xet_khen_thuong" OWNER TO "postgres";
CREATE TABLE "tap_the_xet_thi_dua" (
  "id" int8 NOT NULL DEFAULT nextval('tap_the_xet_thi_dua_id_seq'::regclass),
  "ma_don_vi" int8 NOT NULL,
  "ma_hoi_dong" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ma_danh_hieu" int8 NOT NULL
)
;
ALTER TABLE "tap_the_xet_thi_dua" OWNER TO "postgres";
CREATE TABLE "tapthe_quatrinh" (
  "ma_nam_hoc" int8 NOT NULL,
  "ma_don_vi" int8 NOT NULL,
  "ma_hoi_dong" varchar(255) COLLATE "pg_catalog"."default",
  "list_danh_hieu" int4[],
  "list_khen_thuong" int4[]
)
;
ALTER TABLE "tapthe_quatrinh" OWNER TO "postgres";
CREATE TABLE "thanh_phan_hoi_dong" (
  "ma_thanh_phan" int8 NOT NULL DEFAULT nextval('thanh_phan_hoi_dong_ma_thanh_phan_seq'::regclass),
  "ma_hoi_dong" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
  "ma_can_bo" int8 NOT NULL,
  "ma_chuc_danh" int8 NOT NULL,
  "co_mat" bool DEFAULT true,
  "bau_chon" bool DEFAULT false
)
;
ALTER TABLE "thanh_phan_hoi_dong" OWNER TO "postgres";
BEGIN;
LOCK TABLE "public"."bau_chon_khen_thuong_ca_nhan" IN SHARE MODE;
DELETE FROM "public"."bau_chon_khen_thuong_ca_nhan";
COMMIT;
BEGIN;
LOCK TABLE "public"."bau_chon_khen_thuong_tap_the" IN SHARE MODE;
DELETE FROM "public"."bau_chon_khen_thuong_tap_the";
COMMIT;
BEGIN;
LOCK TABLE "public"."bau_chon_thi_dua_ca_nhan" IN SHARE MODE;
DELETE FROM "public"."bau_chon_thi_dua_ca_nhan";
COMMIT;
BEGIN;
LOCK TABLE "public"."bau_chon_thi_dua_tap_the" IN SHARE MODE;
DELETE FROM "public"."bau_chon_thi_dua_tap_the";
INSERT INTO "public"."bau_chon_thi_dua_tap_the" ("ma_thanh_phan","ma_xet","trang_thai_bau_chon") VALUES (55, 27, 't'),(55, 28, 't'),(62, 35, 't'),(63, 35, 't'),(63, 33, 't'),(63, 34, 'f'),(62, 33, 't'),(62, 34, 't');
COMMIT;
BEGIN;
LOCK TABLE "public"."ca_nhan_xet_khen_thuong" IN SHARE MODE;
DELETE FROM "public"."ca_nhan_xet_khen_thuong";
INSERT INTO "public"."ca_nhan_xet_khen_thuong" ("id","ma_can_bo","ma_hoi_dong","ma_hinh_thuc","noi_dung") VALUES (5, 6, '1556180234000', 9, 'hoàn thành nhiệm vụ năm học'),(6, 14, '1556180234000', 10, 'hoàn thành nhiệm vụ năm học'),(9, 12, '1557137883000', 11, 'Đã hoàn thành xuất sắc nhiệm vụ trong năm học'),(10, 14, '1557137883000', 11, 'Đã hoàn thành xuất sắc nhiệm vụ trong năm học'),(11, 11, '1557137883000', 10, 'Đã hoàn thành xuất sắc nhiệm vụ trong năm học');
COMMIT;
BEGIN;
LOCK TABLE "public"."ca_nhan_xet_thi_dua" IN SHARE MODE;
DELETE FROM "public"."ca_nhan_xet_thi_dua";
INSERT INTO "public"."ca_nhan_xet_thi_dua" ("id","ma_can_bo","ma_hoi_dong","ma_danh_hieu") VALUES (11, 6, '1556180234000', 1),(12, 10, '1556180234000', 2),(16, 14, '1557137883000', 1),(17, 11, '1557137883000', 1),(18, 13, '1557137883000', 3);
COMMIT;
BEGIN;
LOCK TABLE "public"."can_bo" IN SHARE MODE;
DELETE FROM "public"."can_bo";
INSERT INTO "public"."can_bo" ("ma_can_bo","ho_ten","ngay_sinh","que_quan","gioi_tinh","ma_don_vi","ma_chuc_vu","email","anh_dai_dien","password","ma_quyen","trang_thai") VALUES (12, 'Trần Thị Diệu Hiền ', '1980-01-01', 'Quảng Nam', 't', 5, 18, 'ttdh@gmail.com', 'hoang.png', '$2a$10$48Tlbcjs7br9SK..htgxoeuouC/KJxQNvOt3AmfWhqOAK62HIIIee', 1, 'f'),(14, 'Đỗ Quang Khôi   ', '1978-01-05', 'Quảng Nam', 'f', 4, 7, 'doquangkhoi@gmail.com', 'hoang.png', '$2a$10$r0TRT99wFSI0YqwKhWdzw.FANBaJbvkWz4xBiyJCOXbkzsPgENsSO', 3, 'f'),(6, 'Huỳnh Trọng Dương', '1970-04-02', 'Quảng Nam', 'f', 9, 1, 'htduong@gmail.com', 'hoang.png', '$2a$10$Lz3wmCM8G/WZZNE2/AZDceKrZYisUYgklVtNYdzVlus5QIDwXu/8q', 3, 'f'),(9, 'Admin', '1997-12-21', 'Điẹn Bàn', 'f', 27, 17, 'lexuanhoang1997@gmail.com', 'hoang.png', '$2a$10$jAN2rDC6BHwcQVLyb9FqkeSk.kApgtuRG2MMuClGzZhV9vtb.y54u', 2, 'f'),(11, 'Dương Phương Hùng ', '1971-04-20', 'Quảng Nam', 'f', 5, 10, 'dph@gmail.com', 'hoang.png', '$2a$10$Ldn9/4h4m7OWB7N0t0mPeOSrCiEAkoriOVNUSyScsodVNDQ4A7pSS', 1, 'f'),(13, 'Nguyễn Thị Minh Châu', '1980-01-01', 'Quảng Nam', 't', 5, 11, 'nmchau_dhqnam@gmail.com', 'hoang.png', '$2a$10$KKZky1I2aq0s8tZBp9kerujdRqCPi9yTwT6DtsFNr9BirLTqf.K.a', 1, 'f'),(10, 'Vũ Thị Phương Anh', '1980-01-01', 'Quảng Nam', 't', 9, 9, 'vtpa@gmail.com', 'hoang.png', '$2a$10$Lz3wmCM8G/WZZNE2/AZDceKrZYisUYgklVtNYdzVlus5QIDwXu/8q', 1, 'f');
COMMIT;
BEGIN;
LOCK TABLE "public"."can_cu" IN SHARE MODE;
DELETE FROM "public"."can_cu";
INSERT INTO "public"."can_cu" ("ma_can_cu","loai_can_cu","so_hieu_can_cu","ngay_ky","noi_ban_hanh","trich_yeu","trang_thai") VALUES (1, 'Quyết định', '722/QĐ-TTg', '2017-07-08', 'Thủ tướng Chính phủ', 'về việc thành lập Trường Đại học Quảng Nam', 0),(3, 'Nghị định', '91/2017/NĐ-CP', '2017-07-31', 'Thủ tướng Chính phủ', 'về Quy định chi tiết thi hành một số điều của Luật thi đua, khen thưởng', 0),(31, 'Quyết định', '1234', '2019-05-06', 'Trường Đại học Quảng Nam', 'V/v ban hành quyết định thi đua khen thưởng', 1),(32, 'Quyết định', '43', '2019-05-07', 'ĐHQN', 'v/v thành lập HĐ', 1),(33, 'Quyết định', '54', '2019-05-07', 'ĐHQN', 'V/v thành lập HĐ', 1),(34, 'Quyết định', '546', '2019-05-07', 'ĐHQN', 'V/v thành lập HĐ', 1),(35, 'Quyết định', '13', '2019-05-07', 'trường Đại học Quảng Nam', 'vv ban hành quyết định thi đua khen thưởng', 1),(36, 'Quyết định', '1132', '2019-05-07', 'trương đại học Quảng Nam', 'vv thành lập hội đồng thi đua khen thưởng', 1);
COMMIT;
BEGIN;
LOCK TABLE "public"."canhan_quatrinh" IN SHARE MODE;
DELETE FROM "public"."canhan_quatrinh";
COMMIT;
BEGIN;
LOCK TABLE "public"."chuc_danh_hoi_dong" IN SHARE MODE;
DELETE FROM "public"."chuc_danh_hoi_dong";
INSERT INTO "public"."chuc_danh_hoi_dong" ("ma_chuc_danh","ten_chuc_danh","tham_gia_bo_phieu") VALUES (1, 'Chủ tịch', 't'),(2, 'Phó chủ tịch', 't'),(4, 'Uỷ viên thường trực', 't'),(6, 'Thư ký', 'f'),(5, 'Uỷ viên', 't');
COMMIT;
BEGIN;
LOCK TABLE "public"."chuc_vu" IN SHARE MODE;
DELETE FROM "public"."chuc_vu";
INSERT INTO "public"."chuc_vu" ("ma_chuc_vu","ten_chuc_vu") VALUES (2, 'Trưởng phòng'),(1, 'Hiệu trưởng'),(8, 'Phó phòng'),(9, 'Hiệu phó'),(10, 'Trưởng khoa'),(11, 'Giảng viên'),(12, 'Giáo viên'),(13, 'Chuyên viên'),(14, 'Văn thư'),(15, 'Nhân viên'),(16, 'Nhân viên BV'),(17, ''),(18, 'Phó trưởng khoa'),(7, 'Giám đốc');
COMMIT;
BEGIN;
LOCK TABLE "public"."danh_hieu_thi_dua" IN SHARE MODE;
DELETE FROM "public"."danh_hieu_thi_dua";
INSERT INTO "public"."danh_hieu_thi_dua" ("ma_danh_hieu","ten_danh_hieu","doi_tuong_ap_dung","trang_thai","da_bau_chon","thoi_gian_mo","ti_le_dat","so_phut","sl_bo_phieu") VALUES (3, 'Chiến sĩ thi đua cấp tỉnh', 'Cá nhân', 'f', 0, 0, 50, 0, 0),(4, 'Chiến sĩ thi đua toàn quốc', 'Cá nhân', 'f', 0, 0, 50, 0, 0),(2, 'Chiến sĩ thi đua cấp cơ sở', 'Cá nhân', 'f', 0, 0, 50, 0, 0),(1, 'Lao động tiên tiến', 'Cá nhân', 'f', 0, 0, 70, 0, 0),(6, 'Tập thể lao động xuất sắc', 'Tập thể', 'f', 2, 1557284508000, 70, 5, 0),(5, 'Tập thể lao động tiên tiến', 'Tập thể', 'f', 2, 1557284801000, 70, 5, 0);
COMMIT;
BEGIN;
LOCK TABLE "public"."don_vi" IN SHARE MODE;
DELETE FROM "public"."don_vi";
INSERT INTO "public"."don_vi" ("ma_don_vi","ten_don_vi","loai_don_vi") VALUES (5, 'Công nghệ thông tin', 'Khoa'),(4, 'Đào tạo - Bồi dưỡng', 'Trung tâm'),(1, 'Kế hoạch - Tài chính', 'Phòng'),(9, '', ''),(10, 'Hành Chính - Tổng hợp', 'Phòng'),(12, 'Đào tạo', 'Phòng'),(14, 'Khảo thí & Đảm bảo chất lượng', 'Phòng'),(15, 'Công tác sinh viên', 'Phòng'),(16, 'Học liệu & Công nghệ thông tin', 'Trung tâm'),(17, 'Ngoại ngữ tin học', 'Trung tâm'),(11, 'Tổ chức - Thanh tra', 'Phòng'),(13, 'Quản lí khoa học & hợp tác quốc tế', 'Phòng'),(18, 'Toán', 'Khoa'),(19, 'Tiểu học mầm non & Nghệ thuật', 'Trung tâm'),(20, 'Lý Hoá Sinh', 'Khoa'),(21, 'Các môn chung', 'Khoa'),(22, 'Ngoại ngữ', 'Khoa'),(23, 'Ngữ văn & Công tác xã hội', 'Khoa'),(24, 'Kinh tế - Du lịch', 'Khoa'),(27, 'Quản trị', '');
COMMIT;
BEGIN;
LOCK TABLE "public"."hinh_thuc_khen_thuong" IN SHARE MODE;
DELETE FROM "public"."hinh_thuc_khen_thuong";
INSERT INTO "public"."hinh_thuc_khen_thuong" ("ma_hinh_thuc","ten_hinh_thuc","trang_thai","da_bau_chon","thoi_gian_mo","ti_le_dat","so_phut","sl_bo_phieu") VALUES (2, 'Huy chương', 'f', 0, 0, 50, 0, 0),(4, 'Giải thưởng Hồ Chí Minh', 'f', 0, 0, 50, 0, 0),(8, 'Kỷ niệm chương', 'f', 0, 0, 50, 0, 0),(10, 'Bằng khen', 'f', 0, 0, 50, 0, 0),(11, 'Giấy khen', 'f', 0, 0, 50, 0, 0),(9, 'Huy hiệu', 'f', 0, 0, 50, 0, 0),(5, 'Giải thưởng nhà nước', 'f', 0, 0, 50, 0, 0),(3, 'Danh hiệu vinh dự nhà nước', 'f', 0, 0, 50, 0, 0),(1, 'Huân chương', 'f', 0, 0, 50, 0, 0);
COMMIT;
BEGIN;
LOCK TABLE "public"."hoi_dong" IN SHARE MODE;
DELETE FROM "public"."hoi_dong";
INSERT INTO "public"."hoi_dong" ("ma_hoi_dong","ma_nam_hoc","nguoi_tao","ngay_tao","so_quyet_dinh","ngay_ky","thoi_gian_mo","trang_thai") VALUES ('1556095581000', 3, 9, 1556095581000, '2423', '2019-04-25', 1556096407000, 1),('1556180234000', 4, 9, 1556180234000, '1321', '2019-04-04', 1556180763000, 1),('1557137883000', 9, 9, 1557137883000, '123/QĐ', '2019-05-06', 1557155833000, 1),('1557211233000', 10, 9, 1557211233000, '123', '2019-05-07', 1557211409000, 2);
COMMIT;
BEGIN;
LOCK TABLE "public"."hoi_dong_can_cu" IN SHARE MODE;
DELETE FROM "public"."hoi_dong_can_cu";
INSERT INTO "public"."hoi_dong_can_cu" ("ma_hoi_dong","ma_can_cu") VALUES ('1556180234000', 1),('1556180234000', 3),('1557137883000', 1),('1557137883000', 3),('1557137883000', 31),('1557211233000', 1),('1557211233000', 3),('1557211233000', 36);
COMMIT;
BEGIN;
LOCK TABLE "public"."nam_hoc" IN SHARE MODE;
DELETE FROM "public"."nam_hoc";
INSERT INTO "public"."nam_hoc" ("ma_nam_hoc","tu_ngay") VALUES (3, '2012-04-06'),(4, '2013-04-18'),(5, '2014-04-04'),(6, '2015-04-02'),(7, '2016-04-07'),(8, '2017-04-07'),(9, '2018-04-06'),(10, '2019-08-15');
COMMIT;
BEGIN;
LOCK TABLE "public"."quyen" IN SHARE MODE;
DELETE FROM "public"."quyen";
INSERT INTO "public"."quyen" ("ma_quyen","ten_quyen") VALUES (1, 'Người dùng thông thường'),(3, 'Thành viên hội đồng'),(2, 'Quản trị');
COMMIT;
BEGIN;
LOCK TABLE "public"."tap_the_xet_khen_thuong" IN SHARE MODE;
DELETE FROM "public"."tap_the_xet_khen_thuong";
INSERT INTO "public"."tap_the_xet_khen_thuong" ("id","ma_don_vi","ma_hoi_dong","ma_hinh_thuc","noi_dung") VALUES (9, 5, '1556180234000', 9, 'hoàn thành nhiệm vụ năm học'),(10, 4, '1556180234000', 10, 'hoàn thành nhiệm vụ năm học'),(14, 5, '1557137883000', 11, 'Đã hoàn thành xuất sắc nhiệm vụ trong năm học'),(15, 4, '1557137883000', 10, 'Đã hoàn thành xuất sắc nhiệm vụ trong năm học'),(16, 1, '1557137883000', 11, 'Đã hoàn thành xuất sắc nhiệm vụ trong năm học');
COMMIT;
BEGIN;
LOCK TABLE "public"."tap_the_xet_thi_dua" IN SHARE MODE;
DELETE FROM "public"."tap_the_xet_thi_dua";
INSERT INTO "public"."tap_the_xet_thi_dua" ("id","ma_don_vi","ma_hoi_dong","ma_danh_hieu") VALUES (18, 5, '1556180234000', 5),(19, 4, '1556180234000', 6),(20, 1, '1556180234000', 5),(27, 5, '1557137883000', 5),(28, 1, '1557137883000', 5),(29, 4, '1557137883000', 6),(33, 5, '1557211233000', 5),(34, 4, '1557211233000', 5),(35, 1, '1557211233000', 6);
COMMIT;
BEGIN;
LOCK TABLE "public"."tapthe_quatrinh" IN SHARE MODE;
DELETE FROM "public"."tapthe_quatrinh";
INSERT INTO "public"."tapthe_quatrinh" ("ma_nam_hoc","ma_don_vi","ma_hoi_dong","list_danh_hieu","list_khen_thuong") VALUES (9, 4, '1557137883000', '{6,5}', NULL),(9, 5, '1557137883000', '{5}', NULL),(9, 1, '1557137883000', '{5}', NULL),(10, 1, '1557211233000', '{6}', NULL),(10, 5, '1557211233000', '{5}', NULL),(10, 4, '1557211233000', '{5}', NULL);
COMMIT;
BEGIN;
LOCK TABLE "public"."thanh_phan_hoi_dong" IN SHARE MODE;
DELETE FROM "public"."thanh_phan_hoi_dong";
INSERT INTO "public"."thanh_phan_hoi_dong" ("ma_thanh_phan","ma_hoi_dong","ma_can_bo","ma_chuc_danh","co_mat","bau_chon") VALUES (56, '1557137883000', 10, 2, 't', 'f'),(55, '1557137883000', 6, 1, 't', 'f'),(63, '1557211233000', 14, 5, 't', 'f'),(62, '1557211233000', 6, 1, 't', 'f'),(38, '1556180234000', 6, 1, 't', 'f'),(39, '1556180234000', 10, 2, 't', 'f'),(42, '1556180234000', 12, 6, 't', 'f'),(40, '1556180234000', 11, 4, 't', 'f');
COMMIT;
ALTER TABLE "bau_chon_khen_thuong_ca_nhan" ADD CONSTRAINT "bau_chon_khen_thuong_ca_nhan_pkey" PRIMARY KEY ("ma_thanh_phan", "ma_xet");
ALTER TABLE "bau_chon_khen_thuong_tap_the" ADD CONSTRAINT "bau_chon_khen_thuong_tap_the_pkey" PRIMARY KEY ("ma_thanh_phan", "ma_xet");
ALTER TABLE "bau_chon_thi_dua_ca_nhan" ADD CONSTRAINT "bau_chon_thi_dua_ca_nhan_pkey" PRIMARY KEY ("ma_thanh_phan", "ma_xet");
ALTER TABLE "bau_chon_thi_dua_tap_the" ADD CONSTRAINT "bau_chon_thi_dua_tap_the_pkey" PRIMARY KEY ("ma_thanh_phan", "ma_xet");
ALTER TABLE "ca_nhan_xet_khen_thuong" ADD CONSTRAINT "ca_nhan_xet_khen_thuong_pkey" PRIMARY KEY ("id");
ALTER TABLE "ca_nhan_xet_thi_dua" ADD CONSTRAINT "ca_nhan_xet_thi_dua_pkey" PRIMARY KEY ("id");
ALTER TABLE "can_bo" ADD CONSTRAINT "can_bo_pkey" PRIMARY KEY ("ma_can_bo");
ALTER TABLE "can_cu" ADD CONSTRAINT "can_cu_pkey" PRIMARY KEY ("ma_can_cu");
ALTER TABLE "canhan_quatrinh" ADD CONSTRAINT "canhan_quatrinh_pkey" PRIMARY KEY ("ma_nam_hoc", "ma_can_bo");
ALTER TABLE "chuc_danh_hoi_dong" ADD CONSTRAINT "chuc_danh_hoi_dong_pkey" PRIMARY KEY ("ma_chuc_danh");
ALTER TABLE "chuc_vu" ADD CONSTRAINT "chuc_vu_pkey" PRIMARY KEY ("ma_chuc_vu");
ALTER TABLE "danh_hieu_thi_dua" ADD CONSTRAINT "danh_hieu_thi_dua_pkey" PRIMARY KEY ("ma_danh_hieu");
ALTER TABLE "don_vi" ADD CONSTRAINT "don_vi_pkey" PRIMARY KEY ("ma_don_vi");
ALTER TABLE "hinh_thuc_khen_thuong" ADD CONSTRAINT "hinh_thuc_khen_thuong_pkey" PRIMARY KEY ("ma_hinh_thuc");
ALTER TABLE "hoi_dong" ADD CONSTRAINT "hoi_dong_pkey" PRIMARY KEY ("ma_hoi_dong");
ALTER TABLE "hoi_dong_can_cu" ADD CONSTRAINT "hoi_dong_can_cu_pkey" PRIMARY KEY ("ma_can_cu", "ma_hoi_dong");
ALTER TABLE "nam_hoc" ADD CONSTRAINT "nam_hoc_pkey" PRIMARY KEY ("ma_nam_hoc");
ALTER TABLE "quyen" ADD CONSTRAINT "quyen_pkey" PRIMARY KEY ("ma_quyen");
ALTER TABLE "tap_the_xet_khen_thuong" ADD CONSTRAINT "tap_the_xet_khen_thuong_pkey" PRIMARY KEY ("id");
ALTER TABLE "tap_the_xet_thi_dua" ADD CONSTRAINT "tap_the_xet_thi_dua_pkey" PRIMARY KEY ("id");
ALTER TABLE "tapthe_quatrinh" ADD CONSTRAINT "tapthe_quatrinh_pkey" PRIMARY KEY ("ma_nam_hoc", "ma_don_vi");
ALTER TABLE "thanh_phan_hoi_dong" ADD CONSTRAINT "thanh_phan_hoi_dong_pkey" PRIMARY KEY ("ma_thanh_phan");
ALTER TABLE "bau_chon_khen_thuong_ca_nhan" ADD CONSTRAINT "thanhphan_bauchoncn_fk" FOREIGN KEY ("ma_thanh_phan") REFERENCES "public"."thanh_phan_hoi_dong" ("ma_thanh_phan") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "bau_chon_khen_thuong_ca_nhan" ADD CONSTRAINT "xethinhthuccn_bauchon_fk" FOREIGN KEY ("ma_xet") REFERENCES "public"."ca_nhan_xet_khen_thuong" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "bau_chon_khen_thuong_tap_the" ADD CONSTRAINT "thanhphan_bauchontt_fk" FOREIGN KEY ("ma_thanh_phan") REFERENCES "public"."thanh_phan_hoi_dong" ("ma_thanh_phan") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "bau_chon_khen_thuong_tap_the" ADD CONSTRAINT "xethinhthuctt_bauchon_fk" FOREIGN KEY ("ma_xet") REFERENCES "public"."tap_the_xet_khen_thuong" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "bau_chon_thi_dua_ca_nhan" ADD CONSTRAINT "thanhphan_bauchoncn_fk" FOREIGN KEY ("ma_thanh_phan") REFERENCES "public"."thanh_phan_hoi_dong" ("ma_thanh_phan") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "bau_chon_thi_dua_ca_nhan" ADD CONSTRAINT "xetdanhhieucn_bauchon_fk" FOREIGN KEY ("ma_xet") REFERENCES "public"."ca_nhan_xet_thi_dua" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "bau_chon_thi_dua_tap_the" ADD CONSTRAINT "thanhphan_bauchontt_fk" FOREIGN KEY ("ma_thanh_phan") REFERENCES "public"."thanh_phan_hoi_dong" ("ma_thanh_phan") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "bau_chon_thi_dua_tap_the" ADD CONSTRAINT "xetdanhhieutt_bauchon_fk" FOREIGN KEY ("ma_xet") REFERENCES "public"."tap_the_xet_thi_dua" ("id") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "ca_nhan_xet_khen_thuong" ADD CONSTRAINT "xetkhenthuongcn_canbo_fk" FOREIGN KEY ("ma_can_bo") REFERENCES "public"."can_bo" ("ma_can_bo") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "ca_nhan_xet_khen_thuong" ADD CONSTRAINT "xetkhenthuongcn_hinhthuc_fk" FOREIGN KEY ("ma_hinh_thuc") REFERENCES "public"."hinh_thuc_khen_thuong" ("ma_hinh_thuc") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "ca_nhan_xet_khen_thuong" ADD CONSTRAINT "xetkhenthuongcn_hoidong_fk" FOREIGN KEY ("ma_hoi_dong") REFERENCES "public"."hoi_dong" ("ma_hoi_dong") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "ca_nhan_xet_thi_dua" ADD CONSTRAINT "xetthiduacn_canbo_fk" FOREIGN KEY ("ma_can_bo") REFERENCES "public"."can_bo" ("ma_can_bo") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "ca_nhan_xet_thi_dua" ADD CONSTRAINT "xetthiduacn_danhhieu_fk" FOREIGN KEY ("ma_danh_hieu") REFERENCES "public"."danh_hieu_thi_dua" ("ma_danh_hieu") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "ca_nhan_xet_thi_dua" ADD CONSTRAINT "xetthiduacn_hoidong_fk" FOREIGN KEY ("ma_hoi_dong") REFERENCES "public"."hoi_dong" ("ma_hoi_dong") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "can_bo" ADD CONSTRAINT "can_bo_email_key" UNIQUE ("email");
COMMENT ON CONSTRAINT "can_bo_email_key" ON "can_bo" IS 'Email này đã tồn tại';
ALTER TABLE "can_bo" ADD CONSTRAINT "canbo_chucvu_fk" FOREIGN KEY ("ma_chuc_vu") REFERENCES "public"."chuc_vu" ("ma_chuc_vu") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "can_bo" ADD CONSTRAINT "canbo_donvi_fk" FOREIGN KEY ("ma_don_vi") REFERENCES "public"."don_vi" ("ma_don_vi") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "can_bo" ADD CONSTRAINT "quyen_canbo_fk" FOREIGN KEY ("ma_quyen") REFERENCES "public"."quyen" ("ma_quyen") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "canhan_quatrinh" ADD CONSTRAINT "canhan_quatrinhfk" FOREIGN KEY ("ma_can_bo") REFERENCES "public"."can_bo" ("ma_can_bo") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "canhan_quatrinh" ADD CONSTRAINT "namhoc_quatrinhfk" FOREIGN KEY ("ma_nam_hoc") REFERENCES "public"."nam_hoc" ("ma_nam_hoc") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "hoi_dong" ADD CONSTRAINT "hoidong_namhoc_fk" FOREIGN KEY ("ma_nam_hoc") REFERENCES "public"."nam_hoc" ("ma_nam_hoc") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "hoi_dong_can_cu" ADD CONSTRAINT "cancu_fk" FOREIGN KEY ("ma_can_cu") REFERENCES "public"."can_cu" ("ma_can_cu") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "hoi_dong_can_cu" ADD CONSTRAINT "hoidong_fk" FOREIGN KEY ("ma_hoi_dong") REFERENCES "public"."hoi_dong" ("ma_hoi_dong") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "quyen" ADD CONSTRAINT "quyen_ten_quyen_key" UNIQUE ("ten_quyen");
ALTER TABLE "tap_the_xet_khen_thuong" ADD CONSTRAINT "xetkhenthuongtt_donvi_fk" FOREIGN KEY ("ma_don_vi") REFERENCES "public"."don_vi" ("ma_don_vi") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "tap_the_xet_khen_thuong" ADD CONSTRAINT "xetkhenthuongtt_hinhthuc_fk" FOREIGN KEY ("ma_hinh_thuc") REFERENCES "public"."hinh_thuc_khen_thuong" ("ma_hinh_thuc") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "tap_the_xet_khen_thuong" ADD CONSTRAINT "xetkhenthuongtt_hoidong_fk" FOREIGN KEY ("ma_hoi_dong") REFERENCES "public"."hoi_dong" ("ma_hoi_dong") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "tap_the_xet_thi_dua" ADD CONSTRAINT "xetthiduatt_danhhieu_fk" FOREIGN KEY ("ma_danh_hieu") REFERENCES "public"."danh_hieu_thi_dua" ("ma_danh_hieu") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "tap_the_xet_thi_dua" ADD CONSTRAINT "xetthiduatt_donvi_fk" FOREIGN KEY ("ma_don_vi") REFERENCES "public"."don_vi" ("ma_don_vi") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "tap_the_xet_thi_dua" ADD CONSTRAINT "xetthiduatt_hoidong_fk" FOREIGN KEY ("ma_hoi_dong") REFERENCES "public"."hoi_dong" ("ma_hoi_dong") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "tapthe_quatrinh" ADD CONSTRAINT "donvi_quatrinhfk" FOREIGN KEY ("ma_don_vi") REFERENCES "public"."don_vi" ("ma_don_vi") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "tapthe_quatrinh" ADD CONSTRAINT "namhoc_quatrinhfk" FOREIGN KEY ("ma_nam_hoc") REFERENCES "public"."nam_hoc" ("ma_nam_hoc") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "thanh_phan_hoi_dong" ADD CONSTRAINT "canbo_thanhphanhoidong_fk" FOREIGN KEY ("ma_can_bo") REFERENCES "public"."can_bo" ("ma_can_bo") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "thanh_phan_hoi_dong" ADD CONSTRAINT "chucdanh_thanhphanhoidong_fk" FOREIGN KEY ("ma_chuc_danh") REFERENCES "public"."chuc_danh_hoi_dong" ("ma_chuc_danh") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER TABLE "thanh_phan_hoi_dong" ADD CONSTRAINT "hoidong_thanhphan_fk" FOREIGN KEY ("ma_hoi_dong") REFERENCES "public"."hoi_dong" ("ma_hoi_dong") ON DELETE CASCADE ON UPDATE CASCADE;
ALTER SEQUENCE "ca_nhan_xet_khen_thuong_id_seq"
OWNED BY "ca_nhan_xet_khen_thuong"."id";
SELECT setval('"ca_nhan_xet_khen_thuong_id_seq"', 12, true);
ALTER SEQUENCE "ca_nhan_xet_khen_thuong_id_seq" OWNER TO "postgres";
ALTER SEQUENCE "ca_nhan_xet_thi_dua_id_seq"
OWNED BY "ca_nhan_xet_thi_dua"."id";
SELECT setval('"ca_nhan_xet_thi_dua_id_seq"', 19, true);
ALTER SEQUENCE "ca_nhan_xet_thi_dua_id_seq" OWNER TO "postgres";
ALTER SEQUENCE "can_bo_ma_can_bo_seq"
OWNED BY "can_bo"."ma_can_bo";
SELECT setval('"can_bo_ma_can_bo_seq"', 15, true);
ALTER SEQUENCE "can_bo_ma_can_bo_seq" OWNER TO "postgres";
ALTER SEQUENCE "can_cu_ma_can_cu_seq"
OWNED BY "can_cu"."ma_can_cu";
SELECT setval('"can_cu_ma_can_cu_seq"', 37, true);
ALTER SEQUENCE "can_cu_ma_can_cu_seq" OWNER TO "postgres";
ALTER SEQUENCE "chuc_danh_hoi_dong_ma_chuc_danh_seq"
OWNED BY "chuc_danh_hoi_dong"."ma_chuc_danh";
SELECT setval('"chuc_danh_hoi_dong_ma_chuc_danh_seq"', 7, true);
ALTER SEQUENCE "chuc_danh_hoi_dong_ma_chuc_danh_seq" OWNER TO "postgres";
ALTER SEQUENCE "chuc_vu_ma_chuc_vu_seq"
OWNED BY "chuc_vu"."ma_chuc_vu";
SELECT setval('"chuc_vu_ma_chuc_vu_seq"', 19, true);
ALTER SEQUENCE "chuc_vu_ma_chuc_vu_seq" OWNER TO "postgres";
ALTER SEQUENCE "danh_hieu_thi_dua_ma_danh_hieu_seq"
OWNED BY "danh_hieu_thi_dua"."ma_danh_hieu";
SELECT setval('"danh_hieu_thi_dua_ma_danh_hieu_seq"', 11, true);
ALTER SEQUENCE "danh_hieu_thi_dua_ma_danh_hieu_seq" OWNER TO "postgres";
ALTER SEQUENCE "don_vi_ma_don_vi_seq"
OWNED BY "don_vi"."ma_don_vi";
SELECT setval('"don_vi_ma_don_vi_seq"', 28, true);
ALTER SEQUENCE "don_vi_ma_don_vi_seq" OWNER TO "postgres";
ALTER SEQUENCE "hinh_thuc_khen_thuong_ma_hinh_thuc_seq"
OWNED BY "hinh_thuc_khen_thuong"."ma_hinh_thuc";
SELECT setval('"hinh_thuc_khen_thuong_ma_hinh_thuc_seq"', 12, true);
ALTER SEQUENCE "hinh_thuc_khen_thuong_ma_hinh_thuc_seq" OWNER TO "postgres";
ALTER SEQUENCE "nam_hoc_ma_nam_hoc_seq"
OWNED BY "nam_hoc"."ma_nam_hoc";
SELECT setval('"nam_hoc_ma_nam_hoc_seq"', 11, true);
ALTER SEQUENCE "nam_hoc_ma_nam_hoc_seq" OWNER TO "postgres";
ALTER SEQUENCE "quyen_ma_quyen_seq"
OWNED BY "quyen"."ma_quyen";
SELECT setval('"quyen_ma_quyen_seq"', 4, true);
ALTER SEQUENCE "quyen_ma_quyen_seq" OWNER TO "postgres";
ALTER SEQUENCE "tap_the_xet_khen_thuong_id_seq"
OWNED BY "tap_the_xet_khen_thuong"."id";
SELECT setval('"tap_the_xet_khen_thuong_id_seq"', 17, true);
ALTER SEQUENCE "tap_the_xet_khen_thuong_id_seq" OWNER TO "postgres";
ALTER SEQUENCE "tap_the_xet_thi_dua_id_seq"
OWNED BY "tap_the_xet_thi_dua"."id";
SELECT setval('"tap_the_xet_thi_dua_id_seq"', 36, true);
ALTER SEQUENCE "tap_the_xet_thi_dua_id_seq" OWNER TO "postgres";
ALTER SEQUENCE "thanh_phan_hoi_dong_ma_thanh_phan_seq"
OWNED BY "thanh_phan_hoi_dong"."ma_thanh_phan";
SELECT setval('"thanh_phan_hoi_dong_ma_thanh_phan_seq"', 64, true);
ALTER SEQUENCE "thanh_phan_hoi_dong_ma_thanh_phan_seq" OWNER TO "postgres";
