use actix_web::{web, App, HttpResponse, HttpServer};
use aes_gcm::{
    aead::{Aead, KeyInit},
    Aes256Gcm, Nonce,
};
use anyhow::{Context, Result};
use base64::{engine::general_purpose::STANDARD, Engine};
use rand::Rng;
use serde::{Deserialize, Serialize};

/// The encryption key - in production, this would be loaded from a secure vault
/// For demo purposes, we use a fixed 256-bit key
const ENCRYPTION_KEY: &[u8; 32] = b"aquaguard-security-key-32bytes!";

/// Request payload for encryption
#[derive(Deserialize)]
struct EncryptRequest {
    data: String,
}

/// Response payload with encrypted data
#[derive(Serialize)]
struct EncryptResponse {
    encrypted: String,
    nonce: String,
}

/// Encrypts data using AES-256-GCM.
/// This is industrial-grade encryption used to secure bot commands
/// and prevent unauthorized hijacking of the cleaning bots.
fn encrypt_data(plaintext: &str) -> Result<(String, String)> {
    // Create the cipher using our 256-bit key
    let cipher = Aes256Gcm::new_from_slice(ENCRYPTION_KEY)
        .context("Failed to create cipher from key")?;

    // Generate a random 96-bit nonce (12 bytes)
    // Each encryption must use a unique nonce for security
    let mut nonce_bytes = [0u8; 12];
    rand::thread_rng().fill(&mut nonce_bytes);
    let nonce = Nonce::from_slice(&nonce_bytes);

    // Encrypt the plaintext
    let ciphertext = cipher
        .encrypt(nonce, plaintext.as_bytes())
        .map_err(|e| anyhow::anyhow!("Encryption failed: {}", e))?;

    // Return Base64-encoded ciphertext and nonce for transport
    let encrypted_b64 = STANDARD.encode(&ciphertext);
    let nonce_b64 = STANDARD.encode(&nonce_bytes);

    Ok((encrypted_b64, nonce_b64))
}

/// Decrypts AES-256-GCM encrypted data back to plaintext.
fn decrypt_data(encrypted_b64: &str, nonce_b64: &str) -> Result<String> {
    // Decode the Base64 inputs
    let ciphertext = STANDARD.decode(encrypted_b64)
        .context("Failed to decode encrypted data from Base64")?;
    let nonce_bytes = STANDARD.decode(nonce_b64)
        .context("Failed to decode nonce from Base64")?;

    // Validate nonce length
    if nonce_bytes.len() != 12 {
        anyhow::bail!("Invalid nonce length: expected 12 bytes");
    }

    // Create cipher and nonce
    let cipher = Aes256Gcm::new_from_slice(ENCRYPTION_KEY)
        .context("Failed to create cipher from key")?;
    let nonce = Nonce::from_slice(&nonce_bytes);

    // Decrypt the ciphertext
    let plaintext_bytes = cipher
        .decrypt(nonce, ciphertext.as_ref())
        .map_err(|e| anyhow::anyhow!("Decryption failed: {}", e))?;

    let plaintext = String::from_utf8(plaintext_bytes)
        .context("Decrypted data is not valid UTF-8")?;

    Ok(plaintext)
}

/// HTTP handler for encryption endpoint
async fn encrypt_handler(req: web::Json<EncryptRequest>) -> HttpResponse {
    match encrypt_data(&req.data) {
        Ok((encrypted, nonce)) => HttpResponse::Ok().json(EncryptResponse { encrypted, nonce }),
        Err(e) => HttpResponse::InternalServerError().body(format!("Encryption failed: {}", e)),
    }
}

/// HTTP handler for decryption endpoint
#[derive(Deserialize)]
struct DecryptRequest {
    data: String,
    nonce: String,
}

#[derive(Serialize)]
struct DecryptResponse {
    decrypted: String,
}

async fn decrypt_handler(req: web::Json<DecryptRequest>) -> HttpResponse {
    match decrypt_data(&req.data, &req.nonce) {
        Ok(decrypted) => HttpResponse::Ok().json(DecryptResponse { decrypted }),
        Err(e) => HttpResponse::InternalServerError().body(format!("Decryption failed: {}", e)),
    }
}

/// Health check endpoint
async fn health() -> HttpResponse {
    HttpResponse::Ok().json(serde_json::json!({
        "status": "healthy",
        "service": "aquaguard-security",
        "encryption": "AES-256-GCM"
    }))
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    println!("🔐 AquaGuard Security Service starting on port 8081");
    println!("   Using AES-256-GCM encryption for bot command security");
    
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
