package copy

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb"
)

const bufferSize uint32 = 1024 // 1kb

// Copy one file to another with given limit and offset.
// from - origin file
// to - destination file
// limit - maximum transferred bytes
// offset - input file begining offset
func Copy(from string, to string, limit int, offset int) error {
	inFile, err := os.Open(from)
	var limit64 int64
	var offset64 int64 = int64(offset)
	if err != nil {
		return err
	}

	defer inFile.Close()

	inFileInfo, err := inFile.Stat()
	if err != nil {
		return err
	}

	inFileSize := inFileInfo.Size()

	if offset64 > inFileSize {
		return errors.New("Offset is bigger than actual file size")
	}

	inFile.Seek(offset64, 0)

	outFile, err := os.Create(to)
	if err != nil {
		return err
	}

	defer outFile.Close()

	if limit == 0 {
		limit64 = inFileSize - offset64
	} else {
		limit64 = int64(limit)
	}

	sectionReader := io.NewSectionReader(inFile, offset64, limit64)

	bar := pb.StartNew(int(limit64))
	bar.SetUnits(pb.U_BYTES)

	buf := make([]byte, bufferSize)

	for {
		bytesRead, err := sectionReader.Read(buf)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		_, err = outFile.Write(buf[:bytesRead])
		if err != nil {
			return err
		}
		bar.Add(bytesRead)
	}

	return nil
}
