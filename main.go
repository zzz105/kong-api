package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/spf13/viper"
)

func main() {
    // 读取配置文件
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath(".")
    if err := viper.ReadInConfig(); err != nil {
        panic("Failed to read config file: " + err.Error())
    }

    // 获取配置的端口
    port := viper.GetString("server.port")
    if port == "" {
        port = "8080" // 默认端口
    }

    // 创建一个默认的Gin引擎
    r := gin.Default()

    // 定义健康检查API
    r.GET("/api/healthz", healthCheck)

    // 启动服务器，监听配置的端口
    r.Run(":" + port)
}

// healthCheck 处理健康检查请求
func healthCheck(c *gin.Context) {
    // 返回200状态码和JSON响应
    c.JSON(http.StatusOK, gin.H{
        "status": "OK",
    })
}