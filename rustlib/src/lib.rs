use std::ffi::{CStr, CString};
use std::os::raw::c_char;

use blake3::Hasher;

#[unsafe(no_mangle)]
pub extern "C" fn hash_identity(email: *const c_char, timestamp: u64) -> *mut c_char {
    if email.is_null() {
        return std::ptr::null_mut();
    }

    let c_str = unsafe { CStr::from_ptr(email) };
    let email_str = match c_str.to_str() {
        Ok(s) => s,
        Err(_) => return std::ptr::null_mut(),
    };

    let input = format!("{}:{}", email_str, timestamp);
    let mut hasher = Hasher::new();
    hasher.update(input.as_bytes());
    let hash = hasher.finalize();

    let hex_str = hash.to_hex().to_string();
    CString::new(hex_str).unwrap().into_raw()
}

#[unsafe(no_mangle)]
pub extern "C" fn free_str(s: *mut c_char) {
    if s.is_null() {
        return;
    }
    unsafe {
        drop(CString::from_raw(s));
    }
}
