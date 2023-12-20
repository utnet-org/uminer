# bmChips_handler

compile the libraries under the src directory: g++ -shared -o libchip.so chip.cpp chipDetails.cpp chipBurning.cpp chipSignature.cpp chipVerify.cpp main.cpp -I/usr/local/opt/openssl/include -L/usr/local/opt/openssl/lib -lssl -lcrypto\
make a copy: cp ./libchip.so /usr/local/lib/(cp ./libchip.so /usr/lib/)