use md5::{Digest, Md5};

// CWE-328: MD5 password hash.
pub fn hash_password(password: &str) -> String {
    let mut hasher = Md5::new();
    hasher.update(password.as_bytes());
    format!("{:x}", hasher.finalize())
}

// CWE-208: direct comparison.
pub fn check_password(password: &str, stored: &str) -> bool {
    hash_password(password) == stored
}

// CWE-295: TLS verification disabled.
pub fn insecure_client() -> reqwest::Client {
    reqwest::Client::builder()
        .danger_accept_invalid_certs(true)
        .build()
        .unwrap()
}

// CWE-338: time-seeded pseudo-random token.
pub fn weak_token() -> u64 {
    let now = time::now().to_timespec();
    (now.sec as u64) ^ (now.nsec as u64)
}

// CWE-798: secret helper.
pub fn jwt_secret() -> &'static str {
    "4Qm9pLxV2nT7wZ1cB6sH3jF5gA0eY8uI9oP2qR4tV7x"
}
