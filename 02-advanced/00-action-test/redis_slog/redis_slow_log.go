package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

func Execute() {
	deleteTime := time.Now().Add(-time.Duration(10) * time.Second)
	infos, err := ioutil.ReadDir("Dir")
	if err != nil {
		return
	}
	for _, info := range infos {
		if !strings.HasSuffix(info.Name(), ".tar.gz") {
			continue
		}
		ts := strings.Split(info.Name(), "_")
		if len(ts) != 2 {
			continue
		}
		t := strings.TrimSuffix(ts[1], ".tar.gz")
		fileTime, err := time.Parse(time.RFC3339, t)
		if err != nil {
			continue
		}
		if deleteTime.After(fileTime) {
			err = os.Remove(path.Join("Dir", info.Name()))
			if err != nil {
				continue
			}
		}
	}
}

var rmExp = "^\\d{4}_(\\d{2}_){5}.*\\.slow_log.tar.gz$"
var rmReg = regexp.MustCompile(rmExp)

var zipExp = "^\\d{4}_(\\d{2}_){5}.*\\.log$"
var zipReg = regexp.MustCompile(zipExp)

var l sync.Mutex

const LogTimeStamp = "2006-01-02 15:04:05.000"

func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return nil == err
}

func Recycle44() {
	//toRemoves := []string{}
	//toZips := []string{}
	var size int64
	files, err := ioutil.ReadDir("./redis_slow_log")
	if nil != err {
		return
	}
	fmt.Println(files)
	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	file, err := os.Stat("./redis_slow_log/slow.log")
	if err != nil {
		fmt.Println("文件不存在")
		return
	}
	//size += file.Size()
	// 是都大于2M
	if file.Size() >= 2*1024*1024 {
		if err = os.Rename(fmt.Sprintf("%s/redis_slow_log/%s", pwd, file.Name()), fmt.Sprintf("%s/redis_slow_log/slow.log", pwd)); err != nil {
			panic(err)
		}
	}

	//for key, file := range files {
	//	fmt.Println(key, file.Name())
	//	if !file.IsDir() {
	//		//fmt.Println("is_not_dir", file)
	//		//		rmMatched := rmReg.Match([]byte(file.Name()))
	//		//		zipMatched := zipReg.Match([]byte(file.Name()))
	//		//		if rmMatched || zipMatched {
	//		//			size += file.Size()
	//		//		}
	//		if file.Name() == "slow.log" {
	//			//size += file.Size()
	//			// 是都大于2M
	//			if file.Size() >= 2*1024*1024 {
	//				if err = os.Rename(fmt.Sprintf("%s/redis_slow_log/%s", pwd, file.Name()), fmt.Sprintf("%s/redis_slow_log/slow_%d.log", pwd, key)); err != nil {
	//					panic(err)
	//				}
	//			}
	//		}
	//		//		if rmMatched {
	//		//			toRemoves = append(toRemoves, file.Name())
	//		//		}
	//		//		if zipMatched {
	//		//			toZips = append(toZips, file.Name())
	//		//		}
	//	}
	//}
	fmt.Println(size)
	//
	////remove first file in toRemoves-toZip, and loop
	//l.Lock()
	//totalLimit := 100
	//l.Unlock()
	//
	//if size >= int64(totalLimit)*1024*1024 {
	//	for _, toRemove := range toRemoves {
	//		isRemove := true
	//		for _, toZip := range toZips {
	//			if strings.HasPrefix(toRemove, toZip) {
	//				isRemove = false
	//				break
	//			}
	//		}
	//		if !isRemove {
	//			continue
	//		}
	//		removeFilePath := fmt.Sprintf("%v/%v", "DIR", toRemove)
	//
	//		fmt.Fprintf(os.Stderr, "[%v][LOG][REMOVE]%v\n", time.Now().Format(LogTimeStamp), removeFilePath)
	//		err := os.Remove(removeFilePath)
	//		if nil != err {
	//			fmt.Fprintf(os.Stderr, "[%v][LOG][ERROR] remove file err:%v\n", time.Now().Format(LogTimeStamp), err)
	//		}
	//		break
	//
	//	}
	//}
}

const (
	redisSlowLogFilename = "slow.log"
	// redisSlowLogPath /ustats/redis_slow_log/{instance_id}
	redisSlowLogPath = "./redis_slow_log/%s/"
	redisLogContext  = "%s [ID] %d [ClientAddr] %s [ClientName] %s [Args] %s [ExecutionTime] %s"
)

var (
	Instance   *RedisSlowLogRecycle
	instanceMu sync.Mutex
)

type RedisSlowLogRecycle struct {
	removeQueue  []string
	slowLogCount int
}

func NewRedisSlowLogRecycle() *RedisSlowLogRecycle {
	// 初始化
	ins := &RedisSlowLogRecycle{
		removeQueue:  make([]string, 0),
		slowLogCount: 0,
	}
	files, err := ioutil.ReadDir("./redis_slow_log/")
	if err != nil {
		return ins
	}

	for key, file := range files {
		if file.Name() != "slow.log" {
			ins.removeQueue = append(ins.removeQueue, "./redis_slow_log/"+file.Name())
			ins.slowLogCount = key
		}
	}

	return ins
}

var wg sync.WaitGroup

func (r *RedisSlowLogRecycle) RedisSlowLogRecycle() {
	instanceMu.Lock()
	Instance = NewRedisSlowLogRecycle()
	instanceMu.Unlock()

	wg.Add(1)
	go func() {
		for {
			r.recycle()
			time.Sleep(10 * time.Second)
		}
	}()
	wg.Wait()
}

func (r *RedisSlowLogRecycle) recycle() {
	fmt.Println(111)

	// Step 1: 查找指定目录下的慢日志文件 slow.log；
	// Step 2: 通过获取除slow.log外的文件数量
	// Step 3: 判断慢日志文件大小，如果大小超过元数据配置指定大小，则对文件重命名
	// Step 4: 判断目录下的文件数量是否达到设置的最大允许保留数量，如果超过数量则删除
	//files, err := ioutil.ReadDir("./redis_slow_log/")
	file, err := os.Stat("./redis_slow_log/slow.log")
	if nil != err {
		fmt.Println(err)
		return
	}

	// Step1
	if file.Name() == "slow.log" {
		// Step3
		// 是否大于2M
		if file.Size() >= 2*1024*1024 {
			fmt.Println(Instance.slowLogCount)
			oldLogFile := fmt.Sprintf("./redis_slow_log/%s", redisSlowLogFilename)
			//fmt.Println(oldLogFile)
			newFile := fmt.Sprintf("./redis_slow_log/slow_%d.log", Instance.slowLogCount)
			//fmt.Println(newFile)
			if err = os.Rename(oldLogFile, newFile); err != nil {
				return
			}
			Instance.removeQueue = append(Instance.removeQueue, newFile)
			Instance.slowLogCount++
		}
	} else {
		// Step2
		Instance.slowLogCount++
	}

	fmt.Println(Instance.removeQueue)
	// Step4
	if Instance.slowLogCount >= 5 && len(Instance.removeQueue) > 0 {
		fmt.Println("remove operation")
		removedFile := Instance.removeQueue[0]
		Instance.removeQueue = Instance.removeQueue[1:]
		if err := os.Remove(removedFile); err != nil {
			return
		}
		Instance.slowLogCount = 0
	}
	fmt.Println(Instance.removeQueue)
}

type FileInfo struct {
	Name string // 文件名
	Size int64  // 文件大小
	Path string // 文件路径
}

// GetFileList 递归获取指定目录下的所有文件
func GetFileList(path string, fileList *[]FileInfo) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			GetFileList(path+file.Name()+`\`, fileList) // 递归调用
		} else {
			*fileList = append(*fileList, FileInfo{
				Name: file.Name(),
				Size: file.Size(),
				Path: path + "/" + file.Name(),
			})
		}
	}
}

func main2() {
	//Instance.RedisSlowLogRecycle()
	//defer wg.Done()

	//err := filepath.Walk("redis_slow_log",
	//	func(path string, info os.FileInfo, err error) error {
	//		if err != nil {
	//			return err
	//		}
	//		fmt.Println(path, info.Size())
	//
	//		return nil
	//	})
	//if err != nil {
	//	log.Println(err)
	//}

	//var fileList []FileInfo
	//GetFileList("./redis_slow_log", &fileList)
	//fmt.Println("文件数量：", len(fileList))
	//// 打印文件信息
	//for _, file := range fileList {
	//	fmt.Println("file：", file.Name, file.Size, file.Path)
	//}

	//var s = "MemTotal: 1001332 kB"
	//var valid = regexp.MustCompile("[0-9]")
	//fmt.Println(valid.FindAllStringSubmatch(s, -1))

	re := regexp.MustCompile("[0-9]+")
	fmt.Println(re.FindAllString("abc123def987asdf", -1))

}

//if err := filepath.Walk(redisSlowLogDir, func(path string, info os.FileInfo, err error) error {
//	if err != nil {
//		log.Detail(stage, "filepath.Walk error: ", err)
//		return err
//	}
//	if !strings.Contains(path, redisSlowLogFilename) {
//		ins.removeQueue = append(ins.removeQueue, fmt.Sprintf("./%s", path))
//	}
//	return nil
//}); err != nil {
//	log.Detail(stage, "filepath.Walk error: ", err)
//}

var redisSlowLogDir = "./redis_slow_log"

func main() {
	//slowLogDir, err := ioutil.ReadDir(redisSlowLogDir)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(slowLogDir)
	//}
	//for _, dir := range slowLogDir {
	//	fmt.Println(dir.Name())
	//}

	slogLogFilePath := filepath.Join(redisSlowLogDir, "redis-1", redisSlowLogFilename)
	file, err := os.Stat(slogLogFilePath)
	if nil != err {
		fmt.Println(err)
	} else {
		fmt.Println(file.Name(), file.ModTime().Unix())
	}
}
