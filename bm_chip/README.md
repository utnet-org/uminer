# bmChips_handler

1. install_driver for 1684 (2.7.0)\
   sudo rmmod bmsophon (clear previous version if not compatible)\
   sudo bmnnsdk2_bm1684_v2.7.0_20220810patched/bmnnsdk2-bm1684_v2.7.0/scripts/remove_driver_pcie.sh\
   sudo bmnnsdk2_bm1684_v2.7.0_20220810patched/bmnnsdk2-bm1684_v2.7.0/scripts/install_driver_pcie.sh
2. install_driver for 1684x (0.5.0) (the install package is provided)\
   sudo rmmod bmsophon, sudo apt purge sophon-rpc(clear previous version if not compatible)\
   sudo apt install dkms libncurses5\
   sudo dpkg -i sc7_driver/libsophon/sophon-libsophon-dev_0.5.0_amd64.deb
   sudo dpkg -i sc7_driver/rpc/sophon-rpc_3.3.0_amd64.deb

3. compile the libraries under the src directory:\
   g++ -shared -o libchip.so *.cpp ../bmlib/src/*.cpp ../bmlib/src/linux/*.cpp ../common/bm1684/src/common.c -I../bmlib/include -I../bmlib/src/linux -I../common/bm1684/include -I../config -I/usr/local/opt/openssl/include -L/usr/local/opt/openssl/lib -lssl -lcrypto -fPIC
4. make a copy to lib:\
cp ./libchip.so /usr/local/lib/ (cp ./libchip.so /usr/lib/)