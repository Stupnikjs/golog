package database

const InsertUser = "INSERT INTO users (name, email, password) VALUES( $1, $2, $3)"
const SelectAllUsers = " SELECT * FROM users "
const GetPasswordAndMailFromName = "SELECT email, password FROM users WHERE name = $1"
const SelectUserWithEmail = "SELECT name FROM users WHERE email = $1"

/*
const GetUserFromId = "SELECT username,PP,score FROM users WHERE id= $1"
const CheckUserEmail = "SELECT username, email FROM users WHERE username = $1 or email = $2"
const InsertScoreIntoUserWithId = "UPDATE users set score = $1 WHERE id = $2"
*/
