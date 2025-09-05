-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS permissions(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    resource VARCHAR(100) NOT NULL,
    action VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
-- +goose StatementEnd
INSERT INTO permissions(name, description, resource, action)VALUES
('user:read', 'permission to read user data', 'user', 'read'),
('user:write', 'permission to modify user data', 'user', 'write'),
('user:delete', 'permission to delete user data', 'user', 'delete'),
('role:read', 'permission to read role data', 'role', 'read'),
('role:write', 'permission to modify role data', 'role', 'write'),
('role:delete', 'permission to delete role data', 'role', 'delete'),
('role:manage', 'permission to assign or remove roles from users', 'role', 'manage'),
('permission:read', 'permission to read permission data', 'permission', 'read'),
('permission:write', 'permission to modify permission data', 'permission', 'write'),
('permission:delete', 'permission to delete permission data', 'permission', 'delete'),
('permission:manage', 'permission to assign or remove permissions from roles', 'permission', 'manage');
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS permissions;
-- +goose StatementEnd
