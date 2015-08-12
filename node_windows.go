package restic

import (
	"os"
	"syscall"
)

func (node *Node) OpenForReading() (*os.File, error) {
	file, err := os.OpenFile(node.path, os.O_RDONLY, 0)
	if os.IsPermission(err) {
		return os.OpenFile(node.path, os.O_RDONLY, 0)
	}
	return file, err
}

// mknod() creates a filesystem node (file, device
// special file, or named pipe) named pathname, with attributes
// specified by mode and dev.
var mknod = func(path string, mode uint32, dev int) (err error) {
	panic("mknod not implemented")
}

func (node Node) restoreSymlinkTimestamps(path string, utimes [2]syscall.Timespec) error {
	return nil
}

type statWin syscall.Win32FileAttributeData

func toStatT(i interface{}) (statT, bool) {
	if i == nil {
		return nil, false
	}
	s, ok := i.(*syscall.Win32FileAttributeData)
	if ok && s != nil {
		return statWin(*s), true
	}
	return nil, false
}

func (s statWin) dev() uint64   { return 0 }
func (s statWin) ino() uint64   { return 0 }
func (s statWin) nlink() uint64 { return 0 }
func (s statWin) uid() uint32   { return 0 }
func (s statWin) gid() uint32   { return 0 }
func (s statWin) rdev() uint64  { return 0 }

func (s statWin) size() int64 {
	return int64(s.FileSizeLow) | (int64(s.FileSizeHigh) << 32)
}

func (s statWin) atim() syscall.Timespec {
	return syscall.NsecToTimespec(s.LastAccessTime.Nanoseconds())
}

func (s statWin) mtim() syscall.Timespec {
	return syscall.NsecToTimespec(s.LastWriteTime.Nanoseconds())
}

func (s statWin) ctim() syscall.Timespec {
	return syscall.NsecToTimespec(s.CreationTime.Nanoseconds())
}