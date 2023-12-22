# bmChips_handler

1. compile the libraries under the src directory:\
   g++ -shared -o libchip.so *.cpp ../bmlib/src/*.cpp ../bmlib/src/linux/*.cpp ../common/bm1684/src/common.c -I../bmlib/include -I../bmlib/src/linux -I../common/bm1684/include -I../config -I/usr/local/opt/openssl/include -L/usr/local/opt/openssl/lib -lssl -lcrypto -fPIC
2. make a copy to lib:\
cp ./libchip.so /usr/local/lib/ (cp ./libchip.so /usr/lib/)