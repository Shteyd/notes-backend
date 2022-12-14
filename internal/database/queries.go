package database

const CreateNewUser = `
INSERT INTO Users (email, password_hash)
VALUES ($1, $2, $3);
`

const GetUserById = `
SELECT * FROM Users WHERE id=$1 RETURNING id
`

const GetUserID = `
SELECT id IN Users WHERE email = $1 AND password_hash = $2`

const CreateNote = `
INSERT INTO Notes (user_id, title, content)
VALUES ($1, $2, $3) RETURNING id;
`

const CreateNoteWithoutContent = `
INSERT INTO Notes (user_id, title)
VALUES($1, $2) RETURNING id;
`

const GetNoteByID = `
SELECT * FROM Notes WHERE id=$1 AND user_id=$2
`

const GetAllNotesByUserID = `
SELECT * FROM Notes WHERE user_id=$1
`

const UpdateNote = `
UPDATE Notes
SET title = $1, context = $2
WHERE id=$3 AND user_id=$4
`

const UpdateNoteWithoutContent = `
UPDATE Notes
SET title = $1
WHERE id=$2 AND user_id=$3
`

const UpdateNoteWithoutTitle = `
UPDATE Notes
SET context = $1
WHERE id=$2 AND user_id=$3
`

const DeleteNoteByID = `
DELETE FROM Notes WHERE id=$1 AND user_id=$2
`
