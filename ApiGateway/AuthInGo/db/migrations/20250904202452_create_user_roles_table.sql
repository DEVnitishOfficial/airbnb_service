-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_roles(
    id INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    user_id BIGINT UNSIGNED NOT NULL,
    role_id INT NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose StatementEnd
INSERT INTO user_roles(user_id, role_id)
SELECT id, 2 FROM users;

INSERT INTO user_roles(user_id, role_id)
SELECT id, 1 FROM users WHERE email = 'nkumar52@r1rcm.com';

INSERT INTO user_roles(user_id, role_id)
SELECT id, 3 FROM users WHERE email = 'test@gmail.com';

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_roles;
-- +goose StatementEnd
