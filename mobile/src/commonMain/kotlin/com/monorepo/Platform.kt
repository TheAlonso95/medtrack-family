package com.monorepo

interface Platform {
    val name: String
}

expect fun getPlatform(): Platform