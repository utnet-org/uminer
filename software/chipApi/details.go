package chipApi

//#cgo CXXFLAGS: -std=c++11
//#cgo CFLAGS: -I/Users/mac/Desktop/UtilityChain/ut_miner/src
//#cgo LDFLAGS: -L/Users/mac/Desktop/UtilityChain/ut_miner/src -lchip -lstdc++
//#cgo LDFLAGS: -L/usr/local/opt/openssl/lib -lcrypto
//#include <chip.h>
//#include <openssl/ec.h>
//#include <openssl/ecdsa.h>
//#include <openssl/obj_mac.h>
//
import "C"
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unsafe"
)

// GetChipsDetails 从c++端读取参数返回
func GetChipsDetails() {

	res := C.queryChipDetails()

	// Accessing the returned array
	num := 10

	// Convert the C array to a Go slice for easier handling
	chips := (*[1 << 30]C.struct_bmsmi)(unsafe.Pointer(res))[:num]

	// Print the content of each ChipSignature
	for _, chip := range chips {
		fmt.Printf("chip id %d: %d\n", chip.dev_id, chip.dev_id)
		fmt.Printf("chip status %d: %d\n", chip.dev_id, chip.status)
	}
}

type TPUCards struct {
	CardID    string
	Name      string
	Mode      string
	SerialNum string
	ATX       string
	MaxP      string
	BoardP    string
	Minclk    string
	Maxclk    string
	// tpu chips information
	Chips []BMChips
}
type BMChips struct {
	DevId string
	BusId string
	// forms like 178MB/10694MB
	Memory string
	//percentage of usage(%)
	TPUUti string
	// temperature
	BoardT string
	ChipT  string
	// power
	TPUP string
	// voltage
	TPUV string
	// current
	TPUC    string
	Currclk string
	Status  string
}

// BMChipsInfos 直接读取bm-smi的txt文件参数
func BMChipsInfos() {
	file, err := os.Open("bm_smi.txt")
	if err != nil {
		fmt.Println("无法打开芯片参数文件:", err)
		return
	}
	//defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Println("文件扫描错误:", err)
	}

	// 定义正则表达式，用于匹配控制字符
	reg := regexp.MustCompile("\x1b\\[.*?[@-~]")

	// 计数器用于追踪当前行数
	lineCount := 0

	cardList := make([]TPUCards, 0)
	chipList := make([]BMChips, 0)
	card := TPUCards{
		CardID:    "",
		Name:      "",
		Mode:      "",
		SerialNum: "",
		ATX:       "",
		MaxP:      "",
		BoardP:    "",
		Minclk:    "",
		Maxclk:    "",
		Chips:     chipList,
	}
	tpu := BMChips{
		DevId:   "",
		BusId:   "",
		Memory:  "",
		TPUUti:  "",
		BoardT:  "",
		ChipT:   "",
		TPUP:    "",
		TPUV:    "",
		TPUC:    "",
		Currclk: "",
		Status:  "",
	}
	listChipDone := false

	for scanner.Scan() {
		line := scanner.Text()
		cleanLine := reg.ReplaceAllString(line, "")

		lineCount++
		// 第一行：卡和第一张芯片信息
		if lineCount%9 == 8 && lineCount >= 8 {
			//fmt.Println(cleanLine)
			//fmt.Println("+++++++++++++++++++++++++")

			elements := strings.Fields(cleanLine)
			if len(elements) == 0 {
				listChipDone = true
				break
			}
			chipList = make([]BMChips, 0)
			for i, elem := range elements {
				if i == 1 {
					card.CardID = elem
				} else if i == 2 {
					card.Name = elem
				} else if i == 3 {
					card.Mode = elem
				} else if i == 4 {
					card.SerialNum = elem
				} else if i == 6 {
					tpu.DevId = elem
				} else if i == 7 {
					tpu.BoardT = elem
				} else if i == 8 {
					tpu.ChipT = elem
				} else if i == 9 {
					tpu.TPUP = elem
				} else if i == 10 {
					tpu.TPUV = elem
				} else if i == 12 {
					ss := elem
					tpu.TPUUti = strings.Replace(ss, "N/A", "", -1)
				}
			}

		}

		if listChipDone {
			break
		}

		// 第二行：卡和第一张芯片信息
		if lineCount%9 == 0 && lineCount >= 9 {
			elements := strings.Fields(cleanLine)
			for i, elem := range elements {
				if i == 0 {
					ss := elem
					card.ATX = strings.Replace(ss, "|", "", -1)
				} else if i == 1 {
					card.MaxP = elem
				} else if i == 2 {
					card.BoardP = elem
				} else if i == 3 {
					card.Minclk = elem
				} else if i == 4 {
					card.Maxclk = elem
				} else if i == 5 {
					ss := elem
					tpu.BusId = strings.Replace(ss, "N/A|", "", -1)
				} else if i == 6 {
					tpu.Status = elem
				} else if i == 7 {
					re := regexp.MustCompile(`(\d+M)(\d+\.\d+A)`)
					matches := re.FindStringSubmatch(elem)
					tpu.Currclk = matches[1]
					tpu.TPUC = matches[2]
				} else if i == 8 {
					ss := elem
					tpu.Memory = strings.TrimRight(ss, "|")
				}
			}
			chipList = append(chipList, tpu)
		}

		// 第三行：第二张/第三张 第一行 芯片信息
		if (lineCount%9 == 2 && lineCount >= 11) || (lineCount%9 == 5 && lineCount >= 14) {
			elements := strings.Fields(cleanLine)
			for i, elem := range elements {
				if lineCount <= 32 {
					if i == 1 {
						tpu.DevId = elem
					} else if i == 2 {
						tpu.BoardT = elem
					} else if i == 3 {
						tpu.ChipT = elem
					} else if i == 4 {
						tpu.TPUP = elem
					} else if i == 5 {
						tpu.TPUV = elem
					} else if i == 7 {
						ss := elem
						tpu.TPUUti = strings.Replace(ss, "N/A", "", -1)
					}
				} else {
					if i == 0 {
						ss := elem
						tpu.DevId = strings.Replace(ss, "||", "", -1)
					} else if i == 1 {
						tpu.BoardT = elem
					} else if i == 2 {
						tpu.ChipT = elem
					} else if i == 3 {
						tpu.TPUP = elem
					} else if i == 4 {
						tpu.TPUV = elem
					} else if i == 6 {
						ss := elem
						tpu.TPUUti = strings.Replace(ss, "N/A", "", -1)
					}

				}

			}

		}
		// 第三行：第二张/第三张 第二行 芯片信息
		if (lineCount%9 == 3 && lineCount >= 12) || (lineCount%9 == 6 && lineCount >= 15) {
			elements := strings.Fields(cleanLine)
			for i, elem := range elements {
				if i == 0 {
					ss := elem
					tpu.BusId = strings.Replace(ss, "||", "", -1)
				} else if i == 1 {
					tpu.Status = elem
				} else if i == 2 {
					re := regexp.MustCompile(`(\d+M)(\d+\.\d+A)`)
					matches := re.FindStringSubmatch(elem)
					tpu.Currclk = matches[1]
					tpu.TPUC = matches[2]
				} else if i == 3 {
					ss := elem
					tpu.Memory = strings.TrimRight(ss, "|")
				}
			}
			chipList = append(chipList, tpu)
		}

		if lineCount == 16 || (lineCount%9 == 7 && lineCount > 16) {
			card.Chips = chipList
			cardList = append(cardList, card)
			fmt.Println(cardList)
			fmt.Println("+++++++++++++++++++++++++")
		}

	}

}
