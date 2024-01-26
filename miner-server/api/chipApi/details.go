package chipApi

import (
	"bufio"
	"fmt"
	"github.com/prometheus/common/expfmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
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
	BoardT    string
	BoardP    string
	Minclk    string
	Maxclk    string
	// tpu chips information
	Chips []BMChips
}

// BMChips struct of bmchip parameter
type BMChips struct {
	DevId       string
	BusId       string
	Memory      string // forms like 178MB/10694MB
	UsedMemory  string
	TotalMemory string
	TPUUti      string //percentage of usage(%)
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

// RemoteGetChipInfo information from "curl 10.0.3.178:9100"
func RemoteGetChipInfo(url string) []TPUCards {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP请求失败:", err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
	}
	parser := &expfmt.TextParser{}
	metricFamilies, err := parser.TextToMetricFamilies(strings.NewReader(string(body)))
	if err != nil {
		fmt.Println("解析失败:", err)
	}

	// card and chip structure
	cardLists := make([]TPUCards, 0)
	for _, mf := range metricFamilies {
		if mf.GetName() == "bitmain_board_chip_start_index" {
			for i, m := range mf.Metric {
				chip := make([]BMChips, 0)
				coreNum, _ := strconv.ParseInt(m.Label[0].GetValue(), 10, 64)
				sliceLength := len(m.Label[3].GetValue())
				// busid for every chip, form like 000:0a:00.0
				for n := 0; n < int(coreNum); n++ {
					chip = append(chip, BMChips{
						BusId:  m.Label[3].GetValue()[:sliceLength-1] + strconv.Itoa(n),
						Status: "Active",
					})
				}
				cardLists = append(cardLists, TPUCards{
					CardID:    strconv.Itoa(i),
					Name:      m.Label[7].GetValue(),
					Mode:      "PCIE",
					SerialNum: m.Label[8].GetValue(),
					ATX:       "ATX",
					MaxP:      "",
					BoardP:    "",
					Minclk:    "",
					Maxclk:    "",
					Chips:     chip,
				})
			}
		}
		//for _, m := range mf.Metric {
		//	fmt.Printf("  Metric: %v\n", m)
		//	for _, label := range m.Label {
		//		fmt.Printf("    Label: %s=%s\n", label.GetName(), label.GetValue())
		//	}
		//	fmt.Printf("    Value: %s\n", m.GetGauge().GetValue())
		//}
	}
	// card params
	for _, mf := range metricFamilies {
		//fmt.Printf("Metric Family: %s\n", mf.GetName())
		if mf.GetName() == "bitmain_board_max_power" {
			for i, m := range mf.Metric {
				cardLists[i].MaxP = strconv.FormatFloat(m.GetGauge().GetValue(), 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_board_current_power" {
			for i, m := range mf.Metric {
				cardLists[i].BoardP = strconv.FormatFloat(m.GetGauge().GetValue(), 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_board_temperature_celsius" {
			for i, m := range mf.Metric {
				cardLists[i].BoardT = strconv.FormatFloat(m.GetGauge().GetValue(), 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_board_tpu_max_clock" {
			for i, m := range mf.Metric {
				cardLists[i].Maxclk = strconv.FormatFloat(m.GetGauge().GetValue(), 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_board_tpu_min_clock" {
			for i, m := range mf.Metric {
				cardLists[i].Minclk = strconv.FormatFloat(m.GetGauge().GetValue(), 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_board_current_atx12v" {
			for i, m := range mf.Metric {
				cardLists[i].ATX = strconv.FormatFloat(m.GetGauge().GetValue(), 'f', -1, 64)
			}
		}

	}
	// chip params
	for _, mf := range metricFamilies {
		if mf.GetName() == "bitmain_chip_memory_used_bytes" {
			for i, m := range mf.Metric {
				cardId, _ := strconv.ParseInt(m.Label[0].GetValue(), 10, 64)
				coreNum := len(cardLists[cardId].Chips)
				cardLists[cardId].Chips[i%coreNum].UsedMemory = strconv.FormatFloat(m.GetGauge().GetValue()/(1), 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_chip_memory_total_bytes" {
			for i, m := range mf.Metric {
				cardId, _ := strconv.ParseInt(m.Label[0].GetValue(), 10, 64)
				coreNum := len(cardLists[cardId].Chips)
				cardLists[cardId].Chips[i%coreNum].TotalMemory = strconv.FormatFloat(m.GetGauge().GetValue()/(1), 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_chip_tpu_utilization" {
			for i, m := range mf.Metric {
				cardId, _ := strconv.ParseInt(m.Label[0].GetValue(), 10, 64)
				coreNum := len(cardLists[cardId].Chips)
				cardLists[cardId].Chips[i%coreNum].TPUUti = strconv.FormatFloat(m.GetGauge().GetValue(), 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_chip_tpu_power" {
			for i, m := range mf.Metric {
				cardId, _ := strconv.ParseInt(m.Label[0].GetValue(), 10, 64)
				coreNum := len(cardLists[cardId].Chips)
				cardLists[cardId].Chips[i%coreNum].TPUP = strconv.FormatFloat(m.GetGauge().GetValue(), 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_chip_tpu_voltage" {
			for i, m := range mf.Metric {
				cardId, _ := strconv.ParseInt(m.Label[0].GetValue(), 10, 64)
				coreNum := len(cardLists[cardId].Chips)
				cardLists[cardId].Chips[i%coreNum].TPUV = strconv.FormatFloat(m.GetGauge().GetValue(), 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_chip_tpu_current" {
			for i, m := range mf.Metric {
				cardId, _ := strconv.ParseInt(m.Label[0].GetValue(), 10, 64)
				coreNum := len(cardLists[cardId].Chips)
				cardLists[cardId].Chips[i%coreNum].TPUC = strconv.FormatFloat(m.GetGauge().GetValue()/1000, 'f', -1, 64)
			}
		}
		if mf.GetName() == "bitmain_chip_tpu_curr_clock" {
			for i, m := range mf.Metric {
				cardId, _ := strconv.ParseInt(m.Label[0].GetValue(), 10, 64)
				coreNum := len(cardLists[cardId].Chips)
				cardLists[cardId].Chips[i%coreNum].Currclk = strconv.FormatFloat(m.GetGauge().GetValue(), 'f', -1, 64)
			}
		}
	}

	return cardLists

}

// BMChipsInfos reading data from txt written from bm-smi
func BMChipsInfos(directory string) []TPUCards {
	// open file
	cardList := make([]TPUCards, 0)
	file, err := os.Open(directory)
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
