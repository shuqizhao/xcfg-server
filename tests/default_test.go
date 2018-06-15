package test

import (
	"fmt"
	"testing"
	//models "xcfg-server/models"
	"regexp"
	"strings"
)

func init() {
}

// TestMain is a sample to run an endpoint test
func TestMain(m *testing.M) {
	CfgFile := `<?xml version="1.0" encoding="UTF-8"?>
<beans minorVersion="3" majorVersion="2" xmlns="http://www.springframework.org/schema/beans"
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:p="http://www.springframework.org/schema/p"
    xmlns:context="http://www.springframework.org/schema/context"
    xmlns:mvc="http://www.springframework.org/schema/mvc"
    xmlns:cache="http://www.springframework.org/schema/cache"
    xsi:schemaLocation="http://www.springframework.org/schema/beans
                        http://www.springframework.org/schema/beans/spring-beans-4.2.xsd
                        http://www.springframework.org/schema/context
                        http://www.springframework.org/schema/context/spring-context-4.2.xsd
                        http://www.springframework.org/schema/mvc
                        http://www.springframework.org/schema/mvc/spring-mvc-4.2.xsd
                        http://www.springframework.org/schema/cache
                        http://www.springframework.org/schema/cache/spring-cache-4.2.xsd">

　　<!-- 配置文件加载 -->

　　<context:property-placeholder location="classpath:*.properties"/>

<cache:annotation-driven cache-manager="cacheManager"/>
    <!-- redis连接池 -->
    <bean id="poolConfig" class="redis.clients.jedis.JedisPoolConfig">
        <property name="maxIdle" value="${redis.maxIdle}" />
         <property name="maxWaitMillis" value="${redis.maxWait}" />
         <property name="testOnBorrow" value="${redis.testOnBorrow}" />
    </bean>
    <!-- 连接工厂 -->
    <bean id="JedisConnectionFactory" class="org.springframework.data.redis.connection.jedis.JedisConnectionFactory"
       p:host-name="${redis.host}" p:port="${redis.port}" p:password="${redis.pass}" p:pool-config-ref="poolConfig"/>
    <!-- redis模板 -->
    <bean id="redisTemplate" class="org.springframework.data.redis.core.RedisTemplate">
         <property name="connectionFactory" ref="JedisConnectionFactory" />
    </bean>

    <bean id="cacheManager" class="org.springframework.cache.support.SimpleCacheManager">
         <property name="caches">
            <set>
                <!-- 这里可以配置多个redis -->
                <bean class="com.cky.rest.utils.RedisCache">
                     <property name="redisTemplate" ref="redisTemplate" />
                     <property name="name" value="content"/>
                     <!-- name对应的名称要在类或方法的注解中使用 -->
                </bean>
            </set>
         </property>
     </bean>
</beans>`
	CfgFile = dealVersion(CfgFile,"6","7")
	fmt.Println(CfgFile)
}

func dealVersion(cfgFile string,majorVersion string,minorVersion string) string {
	b := []byte(cfgFile)
	reg := regexp.MustCompile(`\<\?.*?\?\>`)
	rep := []byte("${1}")
	cfgFile = string(reg.ReplaceAll(b, rep))

	//regDocType := regexp.MustCompile(`\<!DOCTYPE[\w\W]*?\>`)
	//dotTypeString := regDocType.FindString(cfgFile)
	////fmt.Println(dotTypeString)
	//if dotTypeString != "" {
	//	cfgFile = strings.Replace(cfgFile, dotTypeString, "", 1)
	//}

	//dotTypeString = strings.Trim(dotTypeString, "")

	minorVersionType := regexp.MustCompile(`minorVersion.*?=.*?"\d+"`)
	minorVersionStr := minorVersionType.FindString(cfgFile)
	//fmt.Println(minorVersionStr)
	if minorVersionStr !="" {
		cfgFile = strings.Replace(cfgFile, minorVersionStr, `minorVersion="`+minorVersion+`"`, 1)
	}else {
		rootType := regexp.MustCompile(`\<\w+?\s`)
		rootStr := rootType.FindString(cfgFile)
		cfgFile = strings.Replace(cfgFile, rootStr, rootStr+` minorVersion="`+minorVersion+`" `, 1)
	}

	majorVersionType := regexp.MustCompile(`majorVersion.*?=.*?"\d+"`)
	majorVersionStr := majorVersionType.FindString(cfgFile)
	//fmt.Println(majorVersionStr)
	if majorVersionStr !="" {
		cfgFile = strings.Replace(cfgFile, majorVersionStr, `majorVersion="`+majorVersion+`"`, 1)
	}else {
		rootType := regexp.MustCompile(`\<\w+?\s`)
		rootStr := rootType.FindString(cfgFile)
		cfgFile = strings.Replace(cfgFile, rootStr, rootStr+` majorVersion="`+majorVersion+`" `, 1)
	}
	cfgFile = "<?xml version='1.0' encoding='utf-8' ?>"  + cfgFile
	return cfgFile
}
