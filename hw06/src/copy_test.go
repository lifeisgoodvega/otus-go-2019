package copy

import (
	"bufio"
	"github.com/alecthomas/units"
	"github.com/stretchr/testify/assert"
	"io"
	"os"
	"path/filepath"
	"testing"
)

const humanReadableText = "This is the house that Jack built." +
	"This is the malt" +
	"That lay in the house that Jack built." +
	"This is the rat," +
	"That ate the malt" +
	"That lay in the house that Jack built." +
	"This is the cat," +
	"That kill'd the rat," +
	"That ate the malt" +
	"That lay in the house that Jack built." +
	"This is the dog," +
	"That worried the cat," +
	"That kill'd the rat," +
	"That ate the malt" +
	"That lay in the house that Jack built." +
	"This is the cow with the crumpled horn," +
	"That toss'd the dog," +
	"That worried the cat," +
	"That kill'd the rat," +
	"That ate the malt" +
	"That lay in the house that Jack built." +
	"This is the maiden all forlorn," +
	"That milk'd the cow with the crumpled horn," +
	"That tossed the dog," +
	"That worried the cat," +
	"That kill'd the rat," +
	"That ate the malt" +
	"That lay in the house that Jack built." +
	"This is the man all tatter'd and torn," +
	"That kissed the maiden all forlorn," +
	"That milk'd the cow with the crumpled horn," +
	"That tossed the dog," +
	"That worried the cat," +
	"That kill'd the rat," +
	"That ate the malt" +
	"That lay in the house that Jack built." +
	"This is the priest all shaven and shorn," +
	"That married the man all tatter'd and torn," +
	"That kissed the maiden all forlorn," +
	"That milked the cow with the crumpled horn," +
	"That tossed the dog," +
	"That worried the cat," +
	"That kill'd the rat," +
	"That ate the malt" +
	"That lay in the house that Jack built." +
	"This is the cock that crow'd in the morn," +
	"That waked the priest all shaven and shorn," +
	"That married the man all tatter'd and torn," +
	"That kissed the maiden all forlorn," +
	"That milk'd the cow with the crumpled horn," +
	"That tossed the dog," +
	"That worried the cat," +
	"That kill'd the rat," +
	"That ate the malt" +
	"That lay in the house that Jack built." +
	"This is the farmer sowing his corn," +
	"That kept the cock that crow'd in the morn," +
	"That waked the priest all shaven and shorn," +
	"That married the man all tatter'd and torn," +
	"That kissed the maiden all forlorn," +
	"That milk'd the cow with the crumpled horn," +
	"That tossed the dog," +
	"That worried the cat," +
	"That killed the rat," +
	"That ate the malt" +
	"That lay in the house that Jack built."

const testDir = "test_data"

func ensureDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}
}

func createUnbearableBigFile() string {
	ensureDir(testDir)
	fileName := filepath.Join(testDir, "unbearable_big_file.txt")

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	err = file.Truncate(int64(units.Gigabyte))
	if err != nil {
		panic(err)
	}

	return fileName
}

func createTestFile() string {
	ensureDir(testDir)
	fileName := filepath.Join(testDir, "human_readable_file.txt")

	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	w := bufio.NewWriter(file)
	_, err = w.WriteString(humanReadableText)
	if err != nil {
		panic(err)
	}

	w.Flush()
	return fileName
}

func cleanUp() {
	err := os.RemoveAll(testDir + "/")
	if err != nil {
		panic(err)
	}
}

func TestTooBigOffset(t *testing.T) {
	fileName := createTestFile()
	fileNameCopy := fileName + "_copy"
	defer cleanUp()

	err := Copy(fileName, fileNameCopy, 1, 16384)
	assert.Error(t, err)
}

func TestFileNotExists(t *testing.T) {
	err := Copy("blabla", "blabla_copy", 0, 0)
	assert.Error(t, err)
}

func TestFirstCharacter(t *testing.T) {
	fileName := createTestFile()
	fileNameCopy := fileName + "_copy"
	defer cleanUp()

	err := Copy(fileName, fileNameCopy, 1, 0)
	assert.NoError(t, err)

	file, err := os.Open(fileNameCopy)
	assert.NoError(t, err)
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := io.ReadFull(file, buf)

	assert.Equal(t, 1, n)
	slice := buf[:n]
	assert.Equal(t, "T", string(slice)) // First letter
}

func TestLastCharacter(t *testing.T) {
	fileName := createTestFile()
	fileNameCopy := fileName + "_copy"
	defer cleanUp()

	stat, err := os.Stat(fileName)
	assert.NoError(t, err)
	offset := stat.Size() - 1

	err = Copy(fileName, fileNameCopy, 0, int(offset))
	assert.NoError(t, err)

	file, err := os.Open(fileNameCopy)
	assert.NoError(t, err)
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := io.ReadFull(file, buf)

	assert.Equal(t, 1, n)
	slice := buf[:n]
	assert.Equal(t, ".", string(slice)) // First letter
}

func TestFullText(t *testing.T) {
	fileName := createTestFile()
	fileNameCopy := fileName + "_copy"
	defer cleanUp()

	stat, err := os.Stat(fileName)
	assert.NoError(t, err)
	size := int(stat.Size())

	err = Copy(fileName, fileNameCopy, 0, 0)
	assert.NoError(t, err)

	file, err := os.Open(fileNameCopy)
	assert.NoError(t, err)
	defer file.Close()

	buf := make([]byte, 2048)
	n, err := io.ReadFull(file, buf)

	assert.Equal(t, size, n)
	slice := buf[:n]
	assert.Equal(t, humanReadableText, string(slice)) // First letter
}

func TestSecondSentence(t *testing.T) {
	fileName := createTestFile()
	fileNameCopy := fileName + "_copy"
	defer cleanUp()

	err := Copy(fileName, fileNameCopy, 16, 34)
	assert.NoError(t, err)

	file, err := os.Open(fileNameCopy)
	assert.NoError(t, err)
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := io.ReadFull(file, buf)

	assert.Equal(t, 16, n)
	slice := buf[:n]
	assert.Equal(t, "This is the malt", string(slice)) // First letter
}

func TestBigFile(t *testing.T) {
	fileName := createUnbearableBigFile()
	fileNameCopy := fileName + "_copy"
	defer cleanUp()

	err := Copy(fileName, fileNameCopy, 0, 0)
	assert.NoError(t, err)

	stat, err := os.Stat(fileName)
	assert.NoError(t, err)
	sizeOrigin := stat.Size()

	stat, err = os.Stat(fileName)
	assert.NoError(t, err)
	sizeCopy := stat.Size()

	assert.Equal(t, sizeOrigin, sizeCopy)
}
