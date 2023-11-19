package app

import (
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTMLTemplates(t *testing.T) {
	templateRoot := os.DirFS("templates")
	textTemplates, err := compileHTMLTemplates(templateRoot, "")
	require.NoError(t, err, "compileHTMLTemplates")
	var buf bytes.Buffer
	wi := welcomeInfo{Host: "http://localhost:8888", Version: "1.2.3",
		Scouts: []string{"Kalle", "Britta"}}
	err = textTemplates.ExecuteTemplate(&buf, "welcome.html", wi)
	require.NoError(t, err)
	welcomeStr := buf.String()
	require.Greater(t, strings.Index(welcomeStr, "Kalle"), 0)
}
