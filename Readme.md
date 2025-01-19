
# Configs

Pastikan Config Env berikut ditambahkan pada env ori-shared-config:

VERSION_SOKAS_SERVICE=1.0.0
APP_LOKASI_SOKAS_SERVICE = '/home/donny/project_rnd/ORI/microservice/ori-so-dana-kas'
PORT_SOKAS_SERVICE=7132
PREFORK= true
DB_NAME_SOKAS= so_dana_kas
CEKSTORE_SECRET_KEY= IDMC0mmandMustbeSetFor5ecr3t@3DP
CEKSTORE_KEY= R4h451A_3DP@4MaN

URL_CEKSTORE_LOGIN= "http://172.24.52.30:7321/login"
URL_CEKSTORE= "http://172.24.52.30:7321/ReportFromListener/v1/CekStore"

URL_TRACER= "192.168.131.18"
PORT_TRACER= 4318

# Build
Silahkan custom ./build.sh jika diperlukan.

Jalankan ./build.sh pada terminal Anda.