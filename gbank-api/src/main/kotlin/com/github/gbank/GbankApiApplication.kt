package com.github.gbank

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class GbankApiApplication

fun main(args: Array<String>) {
	runApplication<GbankApiApplication>(*args)
}
