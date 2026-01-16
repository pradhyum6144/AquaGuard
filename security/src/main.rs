use actix_web::{web, App, HttpResponse, HttpServer};
use anyhow::Result;
use base64::{engine::general_purpose::STANDARD, Engine};
use serde::{Deserialize, Serialize};

/// Request payload for encryption
#[derive(Deserialize)]
struct EncryptRequest {
    data: String,
}

/// Response payload with encrypted data
#[derive(Serialize)]
struct EncryptResponse {
    encrypted: String,
}

/// Encrypts a string using Base64 encoding.
/// This is a placeholder for real encryption - we'll upgrade to AES-GCM later.
fn encrypt_data(input: &str) -> Result<String> {
    // For now, just Base64 encode the input
    // This isn't real encryption, but it's a starting point
    let encoded = STANDARD.encode(input.as_bytes());
    Ok(encoded)
}

/// Decrypts a Base64 encoded string back to plaintext.
fn decrypt_data(input: &str) -> Result<String> {
    let decoded_bytes = STANDARD.decode(input)?;
    let decoded = String::from_utf8(decoded_bytes)?;
    Ok(decoded)
}

/// HTTP handler for encryption endpoint
async fn encrypt_handler(req: web::Json<EncryptRequest>) -> HttpResponse {
    match encrypt_data(&req.data) {
        Ok(encrypted) => HttpResponse::Ok().json(EncryptResponse { encrypted }),
        Err(e) => HttpResponse::InternalServerError().body(format!("Encryption failed: {}", e)),
    }
}

/// HTTP handler for decryption endpoint
#[derive(Deserialize)]
struct DecryptRequest {
    data: String,
}

#[derive(Serialize)]
struct DecryptResponse {
    decrypted: String,
}

async fn decrypt_handler(req: web::Json<DecryptRequest>) -> HttpResponse {
    match decrypt_data(&req.data) {
        Ok(decrypted) => HttpResponse::Ok().json(DecryptResponse { decrypted }),
        Err(e) => HttpResponse::InternalServerError().body(format!("Decryption failed: {}", e)),
    }
}

/// Health check endpoint
async fn health() -> HttpResponse {
    HttpResponse::Ok().json(serde_json::json!({
        "status": "healthy",
        "service": "aquaguard-security"
    }))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("AquaGuard Security Service starting on port 8081");
    
    HttpServer::new(|| {
        App::new()
            .route("/health", web::get().to(health))
            .route("/encrypt", web::post().to(encrypt_handler))
            .route("/decrypt", web::post().to(decrypt_handler))
    })
    .bind("0.0.0.0:8081")?
    .run()
    .await
}
