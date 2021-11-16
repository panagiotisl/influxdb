package tsm1_test

import (
	"bufio"
	"compress/gzip"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/influxdata/influxdb/v2/tsdb/engine/tsm1"
)


func TestCompress_FloatBlock_SlopeFloats(t *testing.T) {
	rand.Seed(23)
	var firstTimestamp int64 = 1444238178437870000
	var iterations = 1000
	var size = 1000
	values := make([]tsm1.Value, size)
	var totalSize = int(0)
	for iteration:= 0; iteration < iterations; iteration++ {
		for i := 0; i < size; i++ {
			var value float64 = 300 * float64(i) + 20 + float64(rand.Int() % 10) * 0.1
			values[i] = tsm1.NewValue(firstTimestamp, value)
		}
		b, err := tsm1.Values(values).Encode(nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		totalSize += binary.Size(b)
	}
	fmt.Printf("Total bits: %v\n", totalSize)
}

func TestCompress_Stocks_Germany(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../Stocks_Germany_TKAG_XETRA_NoExpiry.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}
	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)
}

func TestCompress_Stocks_UK(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../Stocks_United_Kingdom_BLND.LSE_NoExpiry.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}

func TestCompress_Stocks_USA(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../Stocks_USA_BAX_NYSE_NoExpiry.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}


func TestCompress_Stocks_Germany_All(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../Stocks-Germany.txt.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}


func TestCompress_Stocks_UK_All(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../Stocks-UK.txt.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}


func TestCompress_Stocks_USA_All(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../Stocks-USA.txt.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}

func TestCompress_Temp_Air_FNLT(t *testing.T) {
	size := 1000
	layout := "2006-01-02T15:04:05Z"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_temp-air-buoy-FNLT.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, row[0])
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[1], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}

func TestCompress_Wind_ABBY(t *testing.T) {
	size := 1000
	layout := "2006-01-02T15:04:05Z"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../ABBY.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, row[0])
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[1], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}

func TestCompress_Rel_Humidity_DewTemp(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_rel-humidity-buoy-dewTempMean.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}


func TestCompress_Rel_Humidity_RHMean(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_rel-humidity-buoy-RHMean.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}


func TestCompress_Rel_Humidity_TempRHMean(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_rel-humidity-buoy-tempRHMean.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}


func TestCompress_Pressure_Air_StaPresMean(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_pressure-air_staPresMean.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}


func TestCompress_Temp_BioMean(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_temp-bio-bioTempMean.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}

func TestCompress_Size_Dust_Particulate_PM10Median(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_size-dust-particulate-PM10Median.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}



func TestCompress_Size_Dust_Particulate_PM10sub50RHMedian(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_size-dust-particulate-PM10sub50RHMedian.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}

func TestCompress_Size_Dust_Particulate_PM15Median(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_size-dust-particulate-PM15Median.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}



func TestCompress_Size_Dust_Particulate_PM15sub50RHMedian(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_size-dust-particulate-PM15sub50RHMedian.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}

func TestCompress_Size_Dust_Particulate_PM1Median(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_size-dust-particulate-PM1Median.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}



func TestCompress_Size_Dust_Particulate_PM1sub50RHMedian(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_size-dust-particulate-PM1sub50RHMedian.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}

func TestCompress_Size_Dust_Particulate_PM25Median(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_size-dust-particulate-PM2.5Median.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}



func TestCompress_Size_Dust_Particulate_PM25sub50RHMedian(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_size-dust-particulate-PM2.5sub50RHMedian.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}

func TestCompress_Size_Dust_Particulate_PM4Median(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_size-dust-particulate-PM4Median.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}



func TestCompress_Size_Dust_Particulate_PM4sub50RHMedian(t *testing.T) {
	size := 1000
	layout := "01/02/2006 15:04:05"
	values := make([]tsm1.Value, size)

	f, err := os.Open("../../../NEON_size-dust-particulate-PM4sub50RHMedian.csv.gz")
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
	}
	defer gz.Close()
	scanner := bufio.NewScanner(gz)
	currentRow := 0
	totalSize := 0
	totalTime := time.Duration(0)
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), ",")
		t, err := time.Parse(layout, fmt.Sprintf("%s %s", row[0], row[1]))
		if err != nil {
			fmt.Println(err)
		}
		if value, err := strconv.ParseFloat(row[2], 64); err == nil {
			values[currentRow] = tsm1.NewValue(t.UnixNano(), value)
			//fmt.Printf("%d: %v\n", t.UnixNano(), value)
		}
		currentRow += 1
		if currentRow == size {
			currentRow = 0
			start := time.Now()
			if b, err := tsm1.Values(values).Encode(nil); err == nil {
				//fmt.Println(len(b))
				totalSize += len(b)
			}
			elapsed := time.Since(start)
			totalTime += elapsed
		}
	}

	fmt.Printf("Total size: %v, Execution took %s\n", totalSize, totalTime)

}
