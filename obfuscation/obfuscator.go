package obfuscation

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/codersrank-org/repo_info_extractor/v2/commit"
)

// Obfuscate private info, like filename, username and emails
func Obfuscate(c *commit.OptimizedCommitForExport) {
	c.AuthorEmail = toMD5(c.AuthorEmail)
}

func toMD5(text string) string {
	algorithm := md5.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
