package main

import (
	"net/http/httputil"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// // Setup proxy
	// url1, err := url.Parse("http://localhost:8081")
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// url2, err := url.Parse("http://localhost:8082")
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// url3, err := url.Parse("http://localhost-intools:9002/")
	// if err != nil {
	// 	e.Logger.Fatal(err)
	// }
	// targets := []*middleware.ProxyTarget{
	// 	{
	// 		URL: url1,
	// 	},
	// 	{
	// 		URL: url2,
	// 	},
	// 	{
	// 		URL: url3,
	// 	},
	// }
	// e.Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:9002",
	})

	proxy_appAuth_code := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:  "http",
		Host:    "localhost:9002",
		RawPath: "/appauth/code",
	})

	proxy_common_auth := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:  "http",
		Host:    "localhost:9002",
		RawPath: "/common/auth",
	})

	proxy_common_v2_auth := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:  "http",
		Host:    "localhost:9002",
		RawPath: "/common/v2/auth",
	})

	proxy_common_wsauth := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:  "http",
		Host:    "localhost:9002",
		RawPath: "/common/wsauth",
	})

	proxy_common_token_upload := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:  "http",
		Host:    "localhost:9002",
		RawPath: "/common/token/upload",
	})

	proxy_login := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:  "http",
		Host:    "localhost:9002",
		RawPath: "/login",
	})

	proxy_logout := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:  "http",
		Host:    "localhost:9002",
		RawPath: "/logout",
	})

	e.GET("/", echo.WrapHandler(proxy))

	e.GET("/common/auth", echo.WrapHandler(proxy_common_auth))
	e.GET("/common/v2/auth", echo.WrapHandler(proxy_common_v2_auth))
	e.POST("/common/v2/auth", echo.WrapHandler(proxy_common_v2_auth))
	e.GET("/common/wsauth", echo.WrapHandler(proxy_common_wsauth))
	e.GET("/common/token/upload", echo.WrapHandler(proxy_common_token_upload))

	e.GET("/appauth/code", echo.WrapHandler(proxy_appAuth_code))
	e.GET("/login", echo.WrapHandler(proxy_login))
	e.GET("/logout", echo.WrapHandler(proxy_logout))

	proxy_mitraApp_saldoMitra := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:  "http",
		Host:    "localhost:9002",
		RawPath: "/mitraapp/saldomitra",
	})

	proxy_mitraApp_saldoMitra_ajax := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme:  "http",
		Host:    "localhost:9002",
		RawPath: "/mitraapp/saldomitra/ajax",
	})

	e.GET("/mitraapp/saldomitra", echo.WrapHandler(proxy_mitraApp_saldoMitra))
	e.GET("/mitraapp/saldomitra/ajax", echo.WrapHandler(proxy_mitraApp_saldoMitra_ajax))

	e.Logger.Fatal(e.Start(":1323"))
}
