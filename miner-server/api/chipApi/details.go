package chipApi

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// TPUCards struct of tpu card parameter
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

// BMChips struct of bmchip parameter
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

// BMChipsInfos reading data from txt written from bm-smi
func BMChipsInfos() []TPUCards {
	// open file
	cardList := make([]TPUCards, 0)
	file, err := os.Open("bm_smi_1.txt")
	if err != nil {
		fmt.Println("fail to open file: ", err)
		return cardList
	}
	//defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Println("scan file error: ", err)
	}

	// 定义正则表达式，用于匹配控制字符
	reg := regexp.MustCompile("\x1b\\[.*?[@-~]")

	lineCount := 0
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
		var readLine = lineCount - 1
		// 第一行：卡和第一张芯片信息
		if readLine%9 == 8 && readLine >= 8 {
			//fmt.Println(cleanLine)
			//fmt.Println("+++++++++++++++++++++++++")

			elements := strings.Fields(cleanLine)
			if len(elements) == 0 {
				listChipDone = true
				break
			}
			chipList = make([]BMChips, 0)
			if len(elements) <= 12 {
				listChipDone = true
				break
			}
			for i, elem := range elements {
				if i == 1 {
					card.CardID = elem
				} else if i == 2 {
					card.Name = elem
				} else if i == 3 {
					card.Mode = elem
				} else if i == 4 {
					card.SerialNum = elem
				}
				if readLine <= 35 {
					if i == 6 {
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
				} else {
					if i == 5 {
						ss := elem
						tpu.DevId = strings.Replace(ss, "|", "", -1)
					} else if i == 6 {
						tpu.BoardT = elem
					} else if i == 7 {
						tpu.ChipT = elem
					} else if i == 8 {
						tpu.TPUP = elem
					} else if i == 9 {
						tpu.TPUV = elem
					} else if i == 11 {
						ss := elem
						tpu.TPUUti = strings.Replace(ss, "N/A", "", -1)
					}
				}
			}

		}

		if listChipDone {
			break
		}

		// 第二行：卡和第一张芯片信息
		if readLine%9 == 0 && readLine >= 9 {
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
		if (readLine%9 == 2 && readLine >= 11) || (readLine%9 == 5 && readLine >= 14) {
			elements := strings.Fields(cleanLine)
			for i, elem := range elements {
				if readLine <= 32 {
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
		if (readLine%9 == 3 && readLine >= 12) || (readLine%9 == 6 && readLine >= 15) {
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

		if readLine == 16 || (readLine%9 == 7 && readLine > 16) {
			card.Chips = chipList
			cardList = append(cardList, card)
			//fmt.Println(cardList)
			//fmt.Println("+++++++++++++++++++++++++")
		}

	}
	return cardList
}
