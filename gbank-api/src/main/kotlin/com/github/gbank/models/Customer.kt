package com.github.gbank.models

import jakarta.persistence.*
import java.time.LocalDateTime
import java.util.UUID

@Entity
@Table(name = "customers")
data class Customer(
        @Id @GeneratedValue(strategy = GenerationType.AUTO)
        val id: UUID?,

        @Column(name = "name")
        var name: String?,

        @Column(name = "created_at")
        var createdAt: LocalDateTime?,

        @Column(name = "updated_at")
        var updatedAt: LocalDateTime?,
)