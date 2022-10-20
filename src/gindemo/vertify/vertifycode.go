package vertify

import (
	"bytes"
	"fmt"
	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SessionConfig() sessions.Store {
	sessionMxAge := 3600
	sessionSecret := "topgoer"
	var store sessions.Store
	store = cookie.NewStore([]byte(sessionSecret))
	store.Options(sessions.Options{
		MaxAge: sessionMxAge,
		Path:   "/",
	})
	return store
}

func Session(kp string) gin.HandlerFunc {
	store := SessionConfig()
	return sessions.Sessions(kp, store)
}

func Captcha(c *gin.Context, lengths ...int) {
	l := captcha.DefaultLen
	w, h := 107, 36
	if len(lengths) == 1 {
		l = lengths[0]
	}
	if len(lengths) == 2 {
		w = lengths[1]
	}
	if len(lengths) == 3 {
		h = lengths[2]
	}
	id := captcha.NewLen(l)

	session := sessions.Default(c)

	session.Set("captcha", id)
	_ = session.Save()
	_ = Serve(c.Writer, c.Request, id, ".png", "zh", false, w, h)

}

func CaptchaVerify(c *gin.Context, code string) bool {
	session := sessions.Default(c)
	if id := session.Get("captcha"); id != nil {
		session.Delete("captcha")
		_ = session.Save()
		if captcha.VerifyString(id.(string), code) {
			return true
		} else {
			return false
		}
	}
	return false
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("./*.html")
	r.Use(Session("topgoer"))
	r.GET("/capthca", func(c *gin.Context) {
		Captcha(c, 4)
	})
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.GET("/captcha/verify/:value", func(c *gin.Context) {
		value := c.Param("value")
		if CaptchaVerify(c, value) {
			c.JSON(http.StatusOK, gin.H{"status": 0, "msg": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})
		}
	})

	if err := r.Run(":8080"); err != nil {
		fmt.Println("get name success")
	}
}
