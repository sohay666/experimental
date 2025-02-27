plugins {
    kotlin("jvm") version "1.6.10" // Use the latest stable Kotlin version
    kotlin("plugin.spring") version "1.6.10"
    id("org.springframework.boot") version "2.6.5" // Ensure compatibility with Spring Boot
    id("io.spring.dependency-management") version "1.0.11.RELEASE"
}


java.toolchain.languageVersion.set(JavaLanguageVersion.of(17))

repositories {
    mavenCentral()
}


dependencies {
    // redis
    implementation("org.springframework.boot:spring-boot-starter-data-redis:2.4.5")


    //json
    implementation("com.fasterxml.jackson.core:jackson-databind:2.15.0")
    implementation("com.fasterxml.jackson.module:jackson-module-kotlin")


    implementation("org.springframework.boot:spring-boot-starter-web")
    testImplementation("org.springframework.boot:spring-boot-starter-test")
}

// Set Kotlin to target JVM 11
tasks.withType<org.jetbrains.kotlin.gradle.tasks.KotlinCompile>().configureEach {
    kotlinOptions {
        freeCompilerArgs = listOf("-Xjsr305=strict", "-Xemit-jvm-type-annotations")
        jvmTarget = "17"
    }
}
