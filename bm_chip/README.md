# bmChips_handler

1. compile the libraries under the src directory:\
g++ -shared -o libchip.so chip.cpp chipStart.cpp chipBurning.cpp chipSignature.cpp chipVerify.cpp main.cpp -I/usr/local/opt/openssl/include -L/usr/local/opt/openssl/lib -lssl -lcrypto -fPIC
2. make a copy to lib:\
cp ./libchip.so /usr/local/lib/ (cp ./libchip.so /usr/lib/)