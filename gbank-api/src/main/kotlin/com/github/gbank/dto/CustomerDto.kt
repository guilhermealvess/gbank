package com.github.gbank.dto

import java.io.Serializable
import java.time.LocalDateTime
import java.util.*

data class CustomerDto(
        var id: UUID?,
        var name: String?,
): Serializable
