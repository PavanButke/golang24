import{
	"flag"
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
}

type templateHandler struct{
	once sync.once
	filename string
	templ *template.Template
}

func 