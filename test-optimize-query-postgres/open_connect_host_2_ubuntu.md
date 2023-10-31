# Hướng dẫn kết nối từ máy host Windows tới PostgreSQL trên WSL Ubuntu

## Bước 1: Xác định IP của WSL

Đầu tiên, bạn cần biết địa chỉ IP của WSL để có thể kết nối. Mở Terminal trong WSL và sử dụng lệnh sau:

```bash
ip addr show eth0
```
***IP WSL Ubuntu: 172.23.102.182.***

Kết quả sẽ hiển thị thông tin về giao diện mạng eth0. Tìm dòng bắt đầu bằng "inet" để tìm địa chỉ IP của WSL.

## Bước 2: Cấu hình PostgreSQL trên WSL

Đảm bảo PostgreSQL trên WSL đang lắng nghe trên địa chỉ IP của WSL (không phải localhost). Bạn cần chỉnh sửa tệp cấu hình PostgreSQL `postgresql.conf` để đặt giá trị `listen_addresses` là địa chỉ IP của WSL. Mở tệp cấu hình bằng lệnh sau:

```bash
sudo nano /etc/postgresql/13/main/postgresql.conf
```

Tìm dòng sau:

```plaintext
#listen_addresses = 'localhost'
```

Và sửa thành:

```plaintext
listen_addresses = 'địa_chỉ_IP_của_WSL'
```

Sau khi sửa, lưu và thoát khỏi tệp.

## Bước 3: Cấu hình phân quyền trong PostgreSQL

Bạn cần cấu hình PostgreSQL để cho phép kết nối từ xa. Mở tệp `pg_hba.conf` bằng lệnh sau:

```bash
sudo nano /etc/postgresql/13/main/pg_hba.conf
```

Thêm dòng sau vào tệp để cho phép kết nối từ bất kỳ địa chỉ IP nào:

```plaintext
host    all             all             0.0.0.0/0               md5
```

Lưu và thoát khỏi tệp.

## Bước 4: Khởi động lại dịch vụ PostgreSQL

Khởi động lại dịch vụ PostgreSQL để áp dụng các thay đổi:

```bash
sudo service postgresql restart
```

## Bước 5: Kết nối từ máy host Windows

Sử dụng một PostgreSQL client trên máy host Windows (như pgAdmin, psql, hoặc một ứng dụng PostgreSQL khác) để kết nối vào địa chỉ IP của WSL và cổng PostgreSQL (mặc định là 5432).

Ví dụ sử dụng `psql` trên máy host Windows:

```bash
psql -h địa_chỉ_IP_của_WSL -U tên_người_dùng_postgres -d tên_cơ_sở_dữ_liệu
```

- `địa_chỉ_IP_của_WSL`: Điền địa chỉ IP bạn đã xác định trong Bước 1.
- `tên_người_dùng_postgres`: Thay thế bằng tên người dùng PostgreSQL (mặc định là "postgres").
- `tên_cơ_sở_dữ_liệu`: Thay thế bằng tên cơ sở dữ liệu bạn muốn kết nối.

Sau khi bạn đã nhập các thông tin, bạn sẽ được yêu cầu nhập mật khẩu để kết nối vào PostgreSQL trên WSL.
```