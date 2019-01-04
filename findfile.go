/*
 * Copyright (c) 2014-2017 MongoDB, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the license is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gowin32

import (
	"strings"
	"time"

	"github.com/winlabs/gowin32/wrappers"

	"path/filepath"
	"syscall"
	"unsafe"
)

type ReparseTag uint32

const (
	ReparseTagMountPoint   ReparseTag = wrappers.IO_REPARSE_TAG_MOUNT_POINT
	ReparseTagHSM          ReparseTag = wrappers.IO_REPARSE_TAG_HSM
	ReparseTagHSM2         ReparseTag = wrappers.IO_REPARSE_TAG_HSM2
	ReparseTagSIS          ReparseTag = wrappers.IO_REPARSE_TAG_SIS
	ReparseTagCSV          ReparseTag = wrappers.IO_REPARSE_TAG_CSV
	ReparseTagDFS          ReparseTag = wrappers.IO_REPARSE_TAG_DFS
	ReparseTagSymbolicLink ReparseTag = wrappers.IO_REPARSE_TAG_SYMLINK
	ReparseTagDFSR         ReparseTag = wrappers.IO_REPARSE_TAG_DFSR
)

type FindFileItem struct {
	FileAttributes    FileAttributes
	CreationTime      time.Time
	LastAccessTime    time.Time
	LastWriteTime     time.Time
	FileSize          uint64
	ReparseTag        ReparseTag
	FileName          string
	AlternateFileName string
}

type FindFile struct {
	handle   syscall.Handle
	fileName string
	current  wrappers.WIN32_FIND_DATA
}

func OpenFindFile(fileName string) *FindFile {
	return &FindFile{fileName: fileName}
}

func (self *FindFile) Close() error {
	if self.handle != 0 {
		if err := wrappers.FindClose(self.handle); err != nil {
			return NewWindowsError("FindClose", err)
		}
		self.handle = 0
		wrappers.RtlZeroMemory((*byte)(unsafe.Pointer(&self.current)), unsafe.Sizeof(self.current))
	}
	return nil
}

func (self *FindFile) Current() FindFileItem {
	return FindFileItem{
		FileAttributes:    FileAttributes(self.current.FileAttributes),
		CreationTime:      windowsFileTimeToTime(int64(fileTimeToUint64(self.current.CreationTime))),
		LastAccessTime:    windowsFileTimeToTime(int64(fileTimeToUint64(self.current.LastAccessTime))),
		LastWriteTime:     windowsFileTimeToTime(int64(fileTimeToUint64(self.current.LastWriteTime))),
		FileSize:          (uint64(self.current.FileSizeHigh) << 32) | uint64(self.current.FileSizeLow),
		ReparseTag:        ReparseTag(self.current.Reserved0),
		FileName:          syscall.UTF16ToString(self.current.FileName[:]),
		AlternateFileName: syscall.UTF16ToString(self.current.AlternateFileName[:]),
	}
}

func (self *FindFile) Next() (bool, error) {
	if self.handle == 0 {
		handle, err := wrappers.FindFirstFile(syscall.StringToUTF16Ptr(self.fileName), &self.current)
		if err == wrappers.ERROR_FILE_NOT_FOUND {
			return false, nil
		} else if err != nil {
			return false, NewWindowsError("FindFirstFile", err)
		}
		self.handle = handle
	} else {
		if err := wrappers.FindNextFile(self.handle, &self.current); err == wrappers.ERROR_NO_MORE_FILES {
			return false, nil
		} else if err != nil {
			return false, NewWindowsError("FindNextFile", err)
		}
	}
	return true, nil
}

func GetDirectorySize(dirName string) (uint64, error) {
	wildcard := filepath.Join(dirName, "*.*")
	ff := OpenFindFile(wildcard)
	defer ff.Close()
	var totalSize uint64
	for {
		if more, err := ff.Next(); err != nil {
			return 0, err
		} else if !more {
			break
		}
		info := ff.Current()
		if info.FileName == "." || info.FileName == ".." {
			continue
		} else if (info.FileAttributes & FileAttributeDirectory) != 0 {
			subdir := filepath.Join(dirName, info.FileName)
			subdirSize, err := GetDirectorySize(subdir)
			if err != nil {
				return 0, err
			}
			totalSize += subdirSize
		} else {
			totalSize += info.FileSize
		}
	}
	return totalSize, nil
}

func getDirectorySizeOnDisk(dirName string, clusterSize uint64) (uint64, error) {
	wildcard := filepath.Join(dirName, "*.*")
	ff := OpenFindFile(wildcard)
	defer ff.Close()
	var totalSize uint64
	for {
		if more, err := ff.Next(); err != nil {
			return 0, err
		} else if !more {
			break
		}
		info := ff.Current()
		if info.FileName == "." || info.FileName == ".." {
			continue
		} else if (info.FileAttributes & FileAttributeDirectory) != 0 {
			subdir := filepath.Join(dirName, info.FileName)
			subdirSize, err := getDirectorySizeOnDisk(subdir, clusterSize)
			if err != nil {
				return 0, err
			}
			totalSize += subdirSize
		} else if (info.FileAttributes & FileAttributeCompressed) != 0 {
			fileName := filepath.Join(dirName, info.FileName)
			compressedSize, err := GetCompressedSize(fileName)
			if err != nil {
				return 0, err
			}
			totalSize += compressedSize
		} else if (info.FileSize % clusterSize) != 0 {
			totalSize += info.FileSize - (info.FileSize % clusterSize) + clusterSize
		} else {
			totalSize += info.FileSize
		}
	}
	return totalSize, nil
}

func GetDirectorySizeOnDisk(dirName string) (uint64, error) {
	volume, err := GetVolumePath(dirName)
	if err != nil {
		return 0, err
	}
	sectorsPerCluster, bytesPerSector, _, _, err := GetSectorsAndClusters(volume)
	if err != nil {
		return 0, err
	}
	return getDirectorySizeOnDisk(dirName, uint64(sectorsPerCluster)*uint64(bytesPerSector))
}

// GetSubdirsInfo returns information about all subdirectories of the directory
func GetSubdirsInfo(dirName string) ([]FindFileItem, error) {

	ff := OpenFindFile(filepath.Join(dirName, "*"))
	defer ff.Close()

	result := make([]FindFileItem, 0)
	for {
		found, err := ff.Next()
		if err != nil {
			return nil, err
		}

		if found {
			info := ff.Current()
			if info.FileName != "." && info.FileName != ".." && (info.FileAttributes&FileAttributeDirectory) != 0 {
				result = append(result, info)
			}
		} else {
			break
		}
	}

	return result, nil
}

func getFilesInfo(dirName, mask string, recursive bool, dirsProcessed *[]string, result *[]FindFileItem) error {
	// get symlinkDirName
	symlinkDirName, _ := GetFinalPathNameAsDOSName(dirName)
	if strings.EqualFold(dirName, symlinkDirName) {
		// if dirName is normal (not symlinked) GetFinalPathNameAsDosName return dirName, so we clearing symlinkDirName
		symlinkDirName = ""
	}

	// skip already processed dirs
	for _, dir := range *dirsProcessed {
		if strings.EqualFold(dir, dirName) || strings.EqualFold(dir, symlinkDirName) {
			return nil
		}
	}

	// add dirName and its symlink target to processed dirs
	*dirsProcessed = append(*dirsProcessed, dirName)
	if symlinkDirName != "" {
		*dirsProcessed = append(*dirsProcessed, symlinkDirName)
	}

	ff := OpenFindFile(filepath.Join(dirName, mask))
	defer ff.Close()
	for {
		found, err := ff.Next()
		if err != nil {
			return err
		}
		if !found {
			break
		}

		info := ff.Current()
		if info.FileName == "." || info.FileName == ".." {
			continue
		}
		info.FileName = filepath.Join(dirName, info.FileName)
		*result = append(*result, info)

		// for "*" mask all dirs are processed in this loop
		if recursive && mask == "*" && ((info.FileAttributes & FileAttributeDirectory) != 0) {
			if err = getFilesInfo(info.FileName, mask, recursive, dirsProcessed, result); err != nil {
				return err
			}
		}
	}

	// for mask != "*" we processed dirs after files
	if recursive && mask != "*" {
		dirs, err := GetSubdirsInfo(dirName)
		if err != nil {
			return err
		}

		for _, dirInfo := range dirs {
			getFilesInfo(filepath.Join(dirName, dirInfo.FileName), mask, recursive, dirsProcessed, result)
		}
	}
	return nil
}

// GetFilesInfo returns all files from dirName that matches a specific file mask (for example *.txt)
// This function can search files recursively and process symlinked directories. If file exists in
// directory and also in symlinked directory it will be present in result array only one time
// In result array file names are returned with full path
func GetFilesInfo(dirName, mask string, recursive bool) ([]FindFileItem, error) {
	result := make([]FindFileItem, 0)
	dirsProcessed := make([]string, 0)
	err := getFilesInfo(dirName, mask, recursive, &dirsProcessed, &result)
	return result, err
}
