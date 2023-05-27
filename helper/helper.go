package helper

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

var txtExcludeMutex sync.Mutex

func TxtExclude(srcpath string, excludepath string, escape bool) (int, []string) {
	fileB, err := os.Open(excludepath)
	if err != nil {
		return CountFileLines(srcpath), nil
	}
	defer fileB.Close()

	fileA, err := os.Open(srcpath)
	if err != nil {
		return 0, nil
	}
	defer fileA.Close()

	txtExcludeMutex.Lock()
	defer txtExcludeMutex.Unlock()

	bLines := make(map[string]bool)
	scannerB := bufio.NewScanner(fileB)
	for scannerB.Scan() {
		line := scannerB.Text()
		if escape {
			line = Escape2(line)
		}
		bLines[line] = true
	}

	var result []string
	scannerA := bufio.NewScanner(fileA)
	for scannerA.Scan() {
		line := scannerA.Text()
		if escape {
			line = Escape2(line)
		}
		if !bLines[line] {
			result = append(result, line)
		}
	}

	return len(result), result
}
func Escape(id string) string {
	InvalidChars := []rune{
		'\u0022', '\u003c', '\u003e', '\u003a', '\u002a', '\u003f', '\u005c',
		'\u002f', '\u007c',
	}
	WhiteChars := []rune{
		'\u0000', '\u0001', '\u0002', '\u0003', '\u0004', '\u0005', '\u0006', '\u0007',
		'\u0008', '\u0009', '\u000a', '\u000b', '\u000c', '\u000d', '\u000e', '\u000f',
		'\u0010', '\u0011', '\u0012', '\u0013', '\u0014', '\u0015', '\u0016', '\u0017',
		'\u0018', '\u0019', '\u001a', '\u001b', '\u001c', '\u001d', '\u001e', '\u001f',
		'\u0020', '\u0085', '\u00c2', '\u00a0', '\u007f',
		'\u2000', '\u2001', '\u2002', '\u2003', '\u2004', '\u2005',
		'\u2006', '\u2007', '\u2008', '\u2009', '\u200a',
		'\u2028', '\u2029', '\u205f', '\u3000',
	}
	var result strings.Builder
	for _, c := range id {
		filter := false
		for _, w := range WhiteChars {
			if c == w {
				filter = true
			}
		}
		if filter {
			continue
		}
		for _, w := range InvalidChars {
			if c == w {
				filter = true
			}
		}
		if filter {
			result.WriteRune('_')
		} else {
			result.WriteRune(c)
		}
	}
	return result.String()
}
func Escape2(id string) string {
	WhiteChars := []rune{
		'\u0000', '\u0001', '\u0002', '\u0003', '\u0004', '\u0005', '\u0006', '\u0007',
		'\u0008', '\u0009', '\u000a', '\u000b', '\u000c', '\u000d', '\u000e', '\u000f',
		'\u0010', '\u0011', '\u0012', '\u0013', '\u0014', '\u0015', '\u0016', '\u0017',
		'\u0018', '\u0019', '\u001a', '\u001b', '\u001c', '\u001d', '\u001e', '\u001f',
		'\u0020', '\u0085', '\u00c2', '\u00a0', '\u007f',
		'\u2000', '\u2001', '\u2002', '\u2003', '\u2004', '\u2005',
		'\u2006', '\u2007', '\u2008', '\u2009', '\u200a',
		'\u2028', '\u2029', '\u205f', '\u3000',
	}
	var result strings.Builder
	for _, c := range id {
		filter := false
		for _, w := range WhiteChars {
			if c == w {
				filter = true
			}
		}
		if filter {
			continue
		} else {
			result.WriteRune(c)
		}
	}
	return result.String()
}
func Escape3(id string) string {
	whiteChars := map[rune]bool{
		'\u0000': true, '\u0001': true, '\u0002': true, '\u0003': true,
		'\u0004': true, '\u0005': true, '\u0006': true, '\u0007': true,
		'\u0008': true, '\u0009': true, '\u000a': true, '\u000b': true,
		'\u000c': true, '\u000d': true, '\u000e': true, '\u000f': true,
		'\u0010': true, '\u0011': true, '\u0012': true, '\u0013': true,
		'\u0014': true, '\u0015': true, '\u0016': true, '\u0017': true,
		'\u0018': true, '\u0019': true, '\u001a': true, '\u001b': true,
		'\u001c': true, '\u001d': true, '\u001e': true, '\u001f': true,
		'\u0020': true, '\u0085': true, '\u00c2': true, '\u00a0': true,
		'\u007f': true,
		'\u2000': true, '\u2001': true, '\u2002': true, '\u2003': true,
		'\u2004': true, '\u2005': true, '\u2006': true, '\u2007': true,
		'\u2008': true, '\u2009': true, '\u200a': true, '\u2028': true,
		'\u2029': true, '\u205f': true, '\u3000': true,
	}
	invalidChars := map[rune]bool{
		'\u0022': true, '\u003c': true, '\u003e': true, '\u003a': true,
		'\u002a': true, '\u003f': true, '\u005c': true, '\u002f': true,
		'\u007c': true,
	}

	var result strings.Builder
	for _, c := range id {
		if whiteChars[c] {
			continue
		}
		if invalidChars[c] {
			result.WriteString("_")
		} else {
			result.WriteString(string(c))
		}
	}
	return result.String()
}
func Escape4(id string) string {
	whiteChars := map[rune]bool{
		'\u0000': true, '\u0001': true, '\u0002': true, '\u0003': true,
		'\u0004': true, '\u0005': true, '\u0006': true, '\u0007': true,
		'\u0008': true, '\u0009': true, '\u000a': true, '\u000b': true,
		'\u000c': true, '\u000d': true, '\u000e': true, '\u000f': true,
		'\u0010': true, '\u0011': true, '\u0012': true, '\u0013': true,
		'\u0014': true, '\u0015': true, '\u0016': true, '\u0017': true,
		'\u0018': true, '\u0019': true, '\u001a': true, '\u001b': true,
		'\u001c': true, '\u001d': true, '\u001e': true, '\u001f': true,
		'\u0020': true, '\u0085': true, '\u00c2': true, '\u00a0': true,
		'\u007f': true,
		'\u2000': true, '\u2001': true, '\u2002': true, '\u2003': true,
		'\u2004': true, '\u2005': true, '\u2006': true, '\u2007': true,
		'\u2008': true, '\u2009': true, '\u200a': true, '\u2028': true,
		'\u2029': true, '\u205f': true, '\u3000': true,
	}

	var result strings.Builder
	for _, c := range id {
		if whiteChars[c] {
			continue
		}
		result.WriteString(string(c))
	}
	return result.String()
}

type lineCounter struct {
	linecount int
	mu        sync.Mutex
}

func (lc *lineCounter) countFileLinesIncrement() {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	lc.linecount++
}

func CountFileLines(path string) int {
	lineCounter := lineCounter{}

	file, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineCounter.countFileLinesIncrement()
	}

	if err := scanner.Err(); err != nil {
		return 0
	}

	return lineCounter.linecount
}
func GenerateURL(baseURL string, unknownPath string) (string, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	if strings.HasPrefix(unknownPath, "//") {
		unknownPath = base.Scheme + ":" + unknownPath
	}
	unknown, err := url.Parse(unknownPath)
	if err != nil {
		return "", err
	}
	if unknown.Scheme != "" {
		return unknown.String(), nil
	}
	if strings.HasPrefix(unknownPath, "/") {
		base.Path = unknownPath
	} else {
		base.Path += "/" + unknownPath
	}
	return base.String(), nil
}
func RandomUA() string {
	seed := rand.NewSource(time.Now().UnixNano())
	r := rand.New(seed)

	chromeMajorVersion := r.Intn(27) + 80

	chromeMinorVersion := r.Intn(1000)

	chromePatchVersion := r.Intn(10000)

	userAgent := fmt.Sprintf("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/%d.%d.%d.0 Safari/537.36", chromeMajorVersion, chromeMinorVersion, chromePatchVersion)
	return userAgent
}
