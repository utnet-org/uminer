#include "../include/bm_smi_recovery.hpp"

bm_smi_recovery::bm_smi_recovery(bm_smi_cmdline &cmdline) : bm_smi_test(cmdline) {
  start_dev = 0;
#ifdef __linux__
  int driver_version = 0;
  fd = bm_smi_open_bmctl(&driver_version);
  dev_cnt = bm_smi_get_dev_cnt(fd);
#endif
}

bm_smi_recovery::~bm_smi_recovery() {}

int bm_smi_recovery::validate_input_para() {
#ifndef SOC_MODE
#ifndef __linux__
  printf("windows does not support recovery!\n");
  return -ENOSPC;
#else
  if ((g_cmdline.m_dev != 0xff) && ((g_cmdline.m_dev < 0) || (g_cmdline.m_dev >= dev_cnt))) {
    printf("error dev = %d\n", g_cmdline.m_dev);
    return -EINVAL;
  }

  if (g_cmdline.m_dev == 0xff) {
    start_dev = 0;
  } else {
    start_dev = g_cmdline.m_dev;
    dev_cnt = 1;
  }
#endif
#endif
  return 0;
}
#ifdef __linux__
#ifndef SOC_MODE
static std::string install_file_path((getenv("START_A53_PATH")));
#endif
int bm_smi_recovery::restart_a53(int start_dev)
{
#ifndef SOC_MODE
  bm_handle_t handle;
  bm_handle_t handle_tmp;
  bm_status_t ret;
  unsigned int card_id = 0;
  unsigned int chip_num = 0;
  unsigned int start_dev_index = 0;
  int i = 0;


  std::string spi_path = install_file_path + "/fip.bin";
  std::string itb_path = install_file_path + "/ramboot_rootfs.itb";

  if (access(spi_path.c_str(), F_OK) != 0) {
	return -1;
  }

  if (access(itb_path.c_str(), F_OK) != 0) {
	return -1;
  }

  ret = bm_dev_request(&handle, start_dev);
  if (ret != BM_SUCCESS) {
    return -1;
  }

  ret = bm_get_card_id(handle, &card_id);
  if (ret != BM_SUCCESS) {
    return -1;
  }

  ret = bm_get_chip_num_from_card(card_id, &chip_num, &start_dev_index);
  if (ret != BM_SUCCESS) {
    return -1;
  }

  if (chip_num == 0x3) {
    for (i = 0; i < 3; i++) {
      ret = bm_dev_request(&handle_tmp, start_dev_index + i);
      if (ret != BM_SUCCESS) {
        return -1;
      }
      ret = bmcpu_start_cpu(handle_tmp, (char *)spi_path.c_str(),  (char *)itb_path.c_str());
      if (ret != BM_SUCCESS) {
        bm_dev_free(handle_tmp);
        return -1;
      }
      bm_dev_free(handle_tmp);
    }
  } else {
    ret = bmcpu_start_cpu(handle, (char *)spi_path.c_str(),  (char *)itb_path.c_str());
    if (ret != BM_SUCCESS) {
      bm_dev_free(handle);
      return -1;
    } else {
      printf("Start cpu success!\n");
    }

    bm_dev_free(handle);
  }
#endif
  return 0;
}
#else
int bm_smi_recovery::restart_a53(int start_dev)
{
  printf("windows does not support recovery!\n");
  return 0;
}
#endif
int bm_smi_recovery::run_opmode() {
#ifdef __linux__
#ifndef SOC_MODE
  char conf[10];
  bm_handle_t handle;
  int ret = 0x0;
  struct bm_misc_info misc_info;

  ret = bm_dev_request(&handle, start_dev);
  if (ret != BM_SUCCESS)
    return -EINVAL;

  bm_get_misc_info(handle, &misc_info);

  bm_dev_free(handle);

  /* set operations: recovery*/
  if (misc_info.pcie_soc_mode == 1) {
    printf("recovery failed!\n");
    printf("Recovery is not support on SOC mode.\n");
    return -1;
  } else {
    printf("Please migrate the upper application business away from the physical board"
                   " where dev%d is located.\n", start_dev);
    printf("In our local experiments, we found that not all servers support recovery"
                    ",and some servers will restart.\n"
                    "For more information, please refer to the documents provided.\n");
    printf("Are you sure to perform recovery option?(yes/no)\n");
    scanf("%s", conf);
    if (strcmp(conf, "yes") == 0 || strcmp(conf, "Y") == 0 || strcmp(conf, "y") == 0 || strcmp(conf, "YES") == 0) {
      if (dev_cnt == 1) {
        ret = ioctl(fd, BMCTL_DEV_RECOVERY, start_dev);
        if (ret < 0) {
          if (errno == EPERM) {
            printf("recovery on dev %d not support\n", start_dev);
            close(fd);
            return -1;
          }
          if (errno == EBUSY) {
            printf("recovery on dev%d failed because of this devcie is busy\n", start_dev);
            close(fd);
            return -1;
          }
          printf("recovery dev %d fail\n", start_dev);
          close(fd);
          return -1;
        } else {
          restart_a53(start_dev);
          printf("recovery dev %d success\n", start_dev);
          close(fd);
          return 0;
        }
      } else {
        printf("Recover one device a time!\n");
        printf("Input correct dev id to recovery.\n");
        close(fd);
        return -EINVAL;
      }
    } else if (strcmp(conf, "no") == 0 || strcmp(conf, "N") == 0 || strcmp(conf, "n") == 0 || strcmp(conf, "NO") == 0) {
      printf("Cancel operation, recovery operation will not be performed.\n");
      close(fd);
      return 0;
    } else {
      printf("wrong key! please input correct key to perform.\n");
      close(fd);
      return -EINVAL;
    }
  }
#else
  close(fd);
  return 0;
#endif

#else
  return 0;
#endif
}

int bm_smi_recovery::check_result() { return 0; }

void bm_smi_recovery::case_help() {}
