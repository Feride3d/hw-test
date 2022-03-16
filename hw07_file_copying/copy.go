package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

// Function "Copy" copies files.
func Copy(fromPath, toPath string, offset, limit int64) error {
	inFile, err := os.Open(fromPath)
	if err != nil {
		return ErrUnsupportedFile
	}
	defer inFile.Close()

	info, err := inFile.Stat()
	if err != nil {
		return err
	}

	switch {
	case info.Size() < offset:
		return ErrOffsetExceedsFileSize
	case info.Size() == 0 || (info.IsDir()):
		return ErrUnsupportedFile
	case (limit > info.Size()-offset) || (limit == 0):
		limit = info.Size() - offset
	}

	_, err = inFile.Seek(offset, io.SeekStart) // SeekStart is a constant to search relative to the origin of the file.
	if err != nil {
		return err
	}

	outFile, err := os.Create(toPath) // Create creates or truncates the named file (if the file already exists).
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Start new console progress bar to output to the console copy progress in percent.
	bar := pb.Full.Start64(limit)
	defer bar.Finish()
	// create proxy reader
	barReader := bar.NewProxyReader(inFile)
	_, err = io.CopyN(outFile, barReader, limit)
	if err != nil {
		return err
	}
	return nil
}
