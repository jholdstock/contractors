# ******************************************************************************
# Settings
# ******************************************************************************
SET foreign_key_checks = 1;
SET time_zone = '+00:00';

# ******************************************************************************
# Create new tables
# ******************************************************************************
CREATE DATABASE IF NOT EXISTS dcr_contractors DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;
USE dcr_contractors;

# ******************************************************************************
# Create tables
# ******************************************************************************
CREATE TABLE user_status (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    status VARCHAR(25) NOT NULL,
    
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    PRIMARY KEY (id)
);

CREATE TABLE user (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password CHAR(60) NOT NULL,
    
    admin BOOLEAN NOT NULL DEFAULT 0,

    status_id INT(10) UNSIGNED NOT NULL DEFAULT 1,
    
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    UNIQUE KEY (email),
    CONSTRAINT `f_user_status` FOREIGN KEY (`status_id`) REFERENCES `user_status` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    
    PRIMARY KEY (id)
);

INSERT INTO `user_status` (`id`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'active',   CURRENT_TIMESTAMP,  NULL,  NULL),
(2, 'inactive', CURRENT_TIMESTAMP,  NULL,  NULL);

CREATE TABLE invoice (
    id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    
    content TEXT NOT NULL,
    
    user_id INT(10) UNSIGNED NOT NULL,
    
    created_at TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,
    
    CONSTRAINT `f_invoice_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    
    PRIMARY KEY (id)
);